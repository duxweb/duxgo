package admin

import (
	event2 "dux-project/app/tools/event"
	model "dux-project/app/tools/models"
	"encoding/json"
	"fmt"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	coreModel "github.com/duxweb/go-fast/models"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/validator"
	"github.com/go-errors/errors"
	"github.com/go-resty/resty/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/event"
	"github.com/gookit/goutil/fsutil"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BackupRes @Resource(app="admin", name = "tools.backup", route = "/tools/backup")
func BackupRes() action.Result {
	res := action.New[model.ToolsBackup](model.ToolsBackup{})

	res.ActionStore = false
	res.ActionEdit = false
	res.ActionCreate = false

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		tx = tx.Order("id desc")
		return tx
	})

	res.Transform(func(item *model.ToolsBackup, index int) map[string]any {
		return map[string]any{
			"id":         item.ID,
			"name":       item.Name,
			"created_at": item.CreatedAt,
		}
	})

	return res.Result()
}

// BackupDownloadList @Action(method = "POST", name = "download", route = "/download/:id")
func BackupDownloadList(c echo.Context) error {

	id := c.Param("id")
	info := model.ToolsBackup{}
	err := database.Gorm().Model(model.ToolsBackup{}).Find(&info, id).Error
	if err != nil {
		return err
	}

	if !fsutil.FileExist(info.Url) {
		return response.BusinessError("备份文件不存在")
	}

	return c.File(info.Url)
}

// BackupImport @Action(method = "POST", name = "import", route = "/import")
func BackupImport(c echo.Context) error {

	dataJson, err := helper.Body(c)
	if err != nil {
		return err
	}

	if !dataJson.Get("file.0.url").Exists() {
		return response.BusinessError("上传文件不存在")
	}
	url := dataJson.Get("file.0.url").String()

	resp, err := resty.New().SetTimeout(10 * time.Second).R().Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}

	md5 := helper.Md5(url)

	file, err := os.CreateTemp("", "backup_"+md5+".zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	file.Write(resp.Body())

	backupTmp := os.TempDir() + "/backup_tmp"

	err = os.MkdirAll(backupTmp, 0777)
	if err != nil {
		return err
	}

	err = helper.UnzipFiles(file.Name(), backupTmp)
	if err != nil {
		return err
	}

	jsonFiles, err := filepath.Glob(backupTmp + "/*.json")
	if err != nil {
		return err
	}

	if len(jsonFiles) == 0 {
		return response.BusinessError("备份文件不存在")
	}

	e := &event2.BackupEvent{BackupData: []map[string]any{}}
	e.SetName("tools.backup")
	err = event.FireEvent(e)
	if err != nil {
		return err
	}

	backupData := lo.KeyBy[string, map[string]any](e.GetBackupData(), func(item map[string]any) string {
		return cast.ToString(item["name"])
	})

	// 遍历每个JSON文件
	models := []map[string]any{}
	for _, filePath := range jsonFiles {
		fileContent := fsutil.ReadFile(filePath)
		fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
		item, ok := backupData[fileName]
		if !ok {
			continue
		}

		fileMaps := []map[string]any{}
		err = json.Unmarshal(fileContent, &fileMaps)
		if err != nil {
			return err
		}
		models = append(models, map[string]any{
			"name":  item["name"],
			"model": item["model"],
			"data":  fileMaps,
		})
	}

	if len(jsonFiles) == 0 {
		return response.BusinessError("找不到可恢复数据")
	}

	for _, item := range models {
		table := coreModel.GetTableName(database.Gorm(), item["model"])
		err = database.Gorm().Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)).Error
		if err != nil {
			return err
		}
	}
	err = database.Gorm().Transaction(func(tx *gorm.DB) error {
		for _, item := range models {
			err = database.Gorm().Model(item["model"]).CreateInBatches(item["data"], 1000).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return response.Send(c, response.Data{
		Message: "导入成功",
	}, 200)
}

// BackupExportList @Action(method = "GET", name = "export", route = "/export")
func BackupExportList(c echo.Context) error {

	// backup event
	e := &event2.BackupEvent{BackupData: []map[string]any{}}
	e.SetName("tools.backup")
	err := event.FireEvent(e)
	if err != nil {
		return err
	}

	return response.Send(c, response.Data{
		Data: lo.Map[map[string]any, map[string]string](e.GetBackupData(), func(item map[string]any, index int) map[string]string {
			name := cast.ToString(item["name"])
			return map[string]string{
				"label": name,
				"value": name,
			}
		}),
	}, 200)
}

// BackupExport @Action(method = "POST", name = "export", route = "/export")
func BackupExport(c echo.Context) error {

	var params struct {
		Name string   `json:"name" validate:"required" validateMsg:"请输入描述"`
		Data []string `json:"data" validate:"required" validateMsg:"请选择导出数据"`
	}
	if err := validator.RequestParser(c, &params); err != nil {
		return err
	}

	e := &event2.BackupEvent{BackupData: []map[string]any{}}
	e.SetName("tools.backup")
	err := event.FireEvent(e)
	if err != nil {
		return err
	}

	files := []string{}
	backupTmp := os.TempDir() + "/backup_tmp"
	defer os.RemoveAll(backupTmp)
	for _, item := range e.GetBackupData() {
		if lo.IndexOf[string](params.Data, cast.ToString(item["name"])) == -1 {
			continue
		}

		data := []map[string]any{}
		err = database.Gorm().Model(item["model"]).Find(&data).Error
		if err != nil {
			return err
		}

		marshal, err := json.Marshal(data)
		if err != nil {
			return err
		}

		file, err := saveJsonFile(backupTmp, cast.ToString(item["name"]), marshal)
		if err != nil {
			return err
		}
		files = append(files, file)

	}
	if len(files) == 0 {
		return response.BusinessError("备份数据不存在")
	}

	backupDir := "./data/backup"
	backupName := carbon.Now().Format("20060102150405")

	err = helper.ZipFiles(backupDir, backupName, files, backupTmp)
	if err != nil {
		return err
	}

	database.Gorm().Model(model.ToolsBackup{}).Create(&model.ToolsBackup{
		Name: params.Name,
		Url:  fmt.Sprintf("%s/%s.zip", backupDir, backupName),
	})

	return response.Send(c, response.Data{}, 200)
}

func saveJsonFile(dir, name string, data []byte) (string, error) {
	filePath := fmt.Sprintf("%s/%s.json", dir, name)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

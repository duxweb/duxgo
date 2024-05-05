package handlers

import (
	"context"
	"crypto/rand"
	"dux-project/app/tools/models"
	"encoding/hex"
	"fmt"
	"github.com/duxweb/go-fast/config"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/storage"
	"github.com/gabriel-vasile/mimetype"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"mime/multipart"
	"path/filepath"
)

func UploadHandler(hasType string, c echo.Context) error {

	dirID := c.FormValue("dir_id")

	// 获取表单文件
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// 读取上传文件
	openFile, err := file.Open()
	if err != nil {
		return err
	}
	defer func(openFile multipart.File) {
		err := openFile.Close()
		if err != nil {

		}
	}(openFile)

	name := file.Filename
	size := file.Size
	mime := file.Header.Get("Content-Type")
	extension := filepath.Ext(name)
	if extension == "" {
		mtype, err := mimetype.DetectReader(openFile)
		if err != nil {
			return err
		}
		extension = mtype.Extension()
		fmt.Println("extension:", extension)
		mime = mtype.String()
	}

	filename := getFilename(extension)
	path := filepath.Join(carbon.Now().ToDateString(), filename)

	ctx := context.Background()

	err = storage.Storage().WriteStream(ctx, path, openFile, map[string]any{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	url, err := storage.Storage().PublicUrl(ctx, path)
	if err != nil {
		fmt.Println("x", err)
		return err
	}

	types := config.Load("storage").GetString("type")
	database.Gorm().Model(models.ToolsFile{}).Create(&models.ToolsFile{
		DirID:   cast.ToUint(dirID),
		HasType: hasType,
		Driver:  types,
		Url:     url,
		Path:    path,
		Name:    name,
		Ext:     extension,
		Size:    cast.ToUint(size),
		Mime:    mime,
	})

	list := []map[string]any{
		{
			"url":  url,
			"name": name,
			"ext":  extension,
			"size": size,
			"mime": mime,
		},
	}

	return response.Send(c, response.Data{
		Data: list,
	}, 200)
}

func getFilename(extension string) string {
	randomBytes := make([]byte, 10)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	randomHex := hex.EncodeToString(randomBytes)

	return fmt.Sprintf("%s%s", randomHex, extension)
}

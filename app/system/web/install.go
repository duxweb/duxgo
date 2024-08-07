package web

import (
	"fmt"
	"github.com/duxweb/go-fast/config"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/global"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/response"
	"github.com/gookit/goutil/fsutil"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

// @RouteGroup(app="web", name = "install", route = "/install")

// InstallLocation @Route(method = "GET", name = "install.location", route = "")
func InstallLocation(ctx echo.Context) error {
	return ctx.Redirect(302, "/install/")
}

// InstallIndex @Route(method = "GET", name = "install.index", route = "/")
func InstallIndex(ctx echo.Context) error {

	file, err := global.StaticFs.ReadFile("static/web/.vite/manifest.json")
	isManifest := true
	if err != nil {
		isManifest = false
	}

	json := gjson.ParseBytes(file)

	vite := config.Load("use").GetStringMap("vite")
	port := lo.Ternary[any](vite["port"] != nil, vite["port"], 5173)
	dev := PortCheck(5173)

	if !isManifest && !dev {
		return echo.NewHTTPError(500, "未编译前端或者未开启前端服务")
	}
	vite["port"] = port
	vite["dev"] = dev

	data := map[string]any{
		"title": config.Load("use").GetString("app.name"),
		"vite":  vite,
		"manifest": map[string]any{
			"js":  json.Get("src\\/install\\.tsx").Get("file").String(),
			"css": json.Get("style\\.css").Get("file").String(),
		},
	}

	ctx.Set("tpl", "manage")
	return ctx.Render(200, "install.gohtml", data)
}

// InstallDetection @Route(method = "GET", name = "install.detection", route = "/detection")
func InstallDetection(ctx echo.Context) error {

	return response.Send(ctx, response.Data{
		Data: map[string]any{
			"ext": []map[string]any{
				{
					"name":   "sqlite",
					"status": true,
				},
			},
			"packages": []map[string]any{
				{
					"name": "go-fast",
					"ver":  global.Version,
				},
			},
			"status": true,
		},
	})
}

// InstallConfig @Route(method = "POST", name = "install.config", route = "/config")
func InstallConfig(ctx echo.Context) error {

	data, err := helper.Body(ctx)
	if err != nil {
		return err
	}

	dbConfig := data.Get("database").Map()

	dbType := dbConfig["type"].String()
	if dbType == "" {
		return response.BusinessError("Please select database type")
	}

	var connect gorm.Dialector
	switch dbType {
	case "mysql":
		connect = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig["username"].String(),
			dbConfig["password"].String(),
			dbConfig["host"].String(),
			dbConfig["port"].String(),
			dbConfig["database"].String(),
		))
		break
	case "postgresql":
		connect = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbConfig["host"].String(),
			dbConfig["username"].String(),
			dbConfig["password"].String(),
			dbConfig["database"].String(),
			dbConfig["port"].String(),
		))
		break
	}

	var message string
	db, err := gorm.Open(connect, &gorm.Config{})
	if err != nil {
		message = err.Error()
	}

	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	}()

	return response.Send(ctx, response.Data{
		Data: map[string]any{
			"error":   err != nil,
			"message": message,
		},
	})
}

// InstallConfig @Route(method = "POST", name = "install.complete", route = "/complete")
func InstallComplete(ctx echo.Context) error {

	data, err := helper.Body(ctx)
	if err != nil {
		return err
	}

	dbConfig := data.Get("database").Map()
	useConfig := data.Get("use").Map()

	dbType := dbConfig["type"].String()
	if dbType == "" {
		return response.BusinessError("Please select database type")
	}

	messageData := []string{}

	// 处理数据库
	dbViper := viper.New()
	dbViper.SetConfigFile("./config/database.yaml")
	dbViper.SetConfigType("yaml")
	err = dbViper.ReadInConfig()
	if err != nil {
		return err
	}

	databaseConfig := map[string]any{
		"type":         dbType,
		"maxIdleConns": 10,
		"maxOpenConns": 100,
	}

	switch dbType {
	case "sqlite":
		databaseConfig["file"] = "./database/app.db"
	case "mysql":
		databaseConfig["host"] = dbConfig["host"].String()
		databaseConfig["port"] = dbConfig["port"].String()
		databaseConfig["username"] = dbConfig["username"].String()
		databaseConfig["password"] = dbConfig["password"].String()
		databaseConfig["database"] = dbConfig["database"].String()
		break
	case "postgresql":
		databaseConfig["host"] = dbConfig["host"].String()
		databaseConfig["port"] = dbConfig["port"].String()
		databaseConfig["username"] = dbConfig["username"].String()
		databaseConfig["password"] = dbConfig["password"].String()
		databaseConfig["database"] = dbConfig["database"].String()
		break
	}
	dbViper.Set("db.drivers.default", databaseConfig)

	err = dbViper.WriteConfig()
	if err != nil {
		return err
	}

	messageData = append(messageData, "config database success")

	// 处理用户文件
	useViper := viper.New()
	useViper.SetConfigFile("./config/use.yaml")
	useViper.SetConfigType("yaml")
	err = useViper.ReadInConfig()
	if err != nil {
		return err
	}

	useViper.Set("app.name", useConfig["name"].String())
	useViper.Set("app.domain", useConfig["domain"].String())
	useViper.Set("app.secret", helper.RandString(32))
	useViper.Set("app.lang", useConfig["lang"].String())

	err = useViper.WriteConfig()
	if err != nil {
		return err
	}
	messageData = append(messageData, "config use success")

	// 处理存储文件
	storageViper := viper.New()
	storageViper.SetConfigFile("./config/storage.yaml")
	storageViper.SetConfigType("yaml")
	err = storageViper.ReadInConfig()
	if err != nil {
		return err
	}
	storageViper.Set("drivers.local.domain", useConfig["domain"].String())

	err = storageViper.WriteConfig()
	if err != nil {
		return err
	}
	messageData = append(messageData, "config storage success")

	// 重载配置文件
	err = config.Load("database").ReadInConfig()
	if err != nil {
		return err
	}
	err = config.Load("use").ReadInConfig()
	if err != nil {
		return err
	}
	err = config.Load("storage").ReadInConfig()
	if err != nil {
		return err
	}

	err = database.SwitchGorm("default")
	if err != nil {
		return err
	}

	// 同步数据库
	err = database.SyncDatabase()
	if err != nil {
		return err
	}
	messageData = append(messageData, "sync database success")

	// 安装锁定
	fsutil.MustCreateFile("./data/install.lock", 0777, 0777)

	return response.Send(ctx, response.Data{
		Data: map[string]any{
			"error":   false,
			"message": "",
			"logs":    strings.Join(messageData, "\n"),
		},
	})
}

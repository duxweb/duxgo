package runtime

import (
	appSystemAdmin "dux-project/app/system/admin"
	appSystemModels "dux-project/app/system/models"
	appSystemWeb "dux-project/app/system/web"
	appToolsAdmin "dux-project/app/tools/admin"
	appToolsListener "dux-project/app/tools/listener"
	appToolsModels "dux-project/app/tools/models"
	"github.com/duxweb/go-fast/annotation"
)

var Annotations = []*annotation.File{
	{
		Name: "dux-project/app/system/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "system.api",
					"route": "/system/api",
				},
				Func: appSystemAdmin.ApiRes,
			},
		},
	},
	{
		Name: "dux-project/app/system/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"app":   "web",
					"name":  "auth",
					"route": "/admin",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "POST",
					"name":   "login",
					"route":  "/login",
				},
				Func: appSystemAdmin.Login,
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "POST",
					"name":   "check",
					"route":  "/check",
				},
				Func: appSystemAdmin.Check,
			},
		},
	},
	{
		Name: "dux-project/app/system/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "system.operate",
					"route": "/system/operate",
				},
				Func: appSystemAdmin.OperateRes,
			},
		},
	},
	{
		Name: "dux-project/app/system/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "system.role",
					"route": "/system/role",
				},
				Func: appSystemAdmin.RoleRes,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "GET",
					"name":   "permission",
					"route":  "/permission",
				},
				Func: appSystemAdmin.RolePermission,
			},
		},
	},
	{
		Name: "dux-project/app/system/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"app":   "admin",
					"name":  "auth",
					"route": "",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "GET",
					"name":   "menu",
					"route":  "/menu",
				},
				Func: appSystemAdmin.Menu,
			},
		},
	},
	{
		Name: "dux-project/app/system/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "system.user",
					"route": "/system/user",
				},
				Func: appSystemAdmin.UserRes,
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.Config{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogApi{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.SystemApi{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.SystemRole{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.SystemUserMigrate,
			},
		},
	},
	{
		Name: "dux-project/app/system/web",
		Annotations: []*annotation.Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"app":   "web",
					"name":  "manage",
					"route": "/manage",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "GET",
					"name":   "manage.location",
					"route":  "",
				},
				Func: appSystemWeb.Location,
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "GET",
					"name":   "manage.index",
					"route":  "/",
				},
				Func: appSystemWeb.Index,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.area",
					"route": "/tools/area",
				},
				Func: appToolsAdmin.AreaRes,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "POST",
					"name":   "import",
					"route":  "",
				},
				Func: appToolsAdmin.AreaImport,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.backup",
					"route": "/tools/backup",
				},
				Func: appToolsAdmin.BackupRes,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "POST",
					"name":   "download",
					"route":  "/download/:id",
				},
				Func: appToolsAdmin.BackupDownloadList,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "POST",
					"name":   "import",
					"route":  "/import",
				},
				Func: appToolsAdmin.BackupImport,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "GET",
					"name":   "export",
					"route":  "/export",
				},
				Func: appToolsAdmin.BackupExportList,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "POST",
					"name":   "export",
					"route":  "/export",
				},
				Func: appToolsAdmin.BackupExport,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.file",
					"route": "/tools/file",
				},
				Func: appToolsAdmin.FileRes,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.fileDir",
					"route": "/tools/fileDir",
				},
				Func: appToolsAdmin.FileDirRes,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.magic",
					"route": "/tools/magic",
				},
				Func: appToolsAdmin.MagicRes,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "GET",
					"name":   "config",
					"route":  "/config",
				},
				Func: appToolsAdmin.MagicConfig,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "GET",
					"name":   "source",
					"route":  "/source",
				},
				Func: appToolsAdmin.MagicSource,
			},
			{
				Name: "Action",
				Params: map[string]any{
					"method": "GET",
					"name":   "sourceData",
					"route":  "/sourceData",
				},
				Func: appToolsAdmin.MagicSourceData,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.data",
					"route": "/tools/data",
				},
				Func: appToolsAdmin.MagicDataRes,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.magicGroup",
					"route": "/tools/magicGroup",
				},
				Func: appToolsAdmin.MagicGroupRes,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "Resource",
				Params: map[string]any{
					"app":   "admin",
					"name":  "tools.magicSource",
					"route": "/tools/magicSource",
				},
				Func: appToolsAdmin.MagicSourceRes,
			},
		},
	},
	{
		Name: "dux-project/app/tools/admin",
		Annotations: []*annotation.Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"app":   "admin",
					"name":  "upload",
					"route": "/upload",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "POST",
					"name":   "upload",
					"route":  "",
				},
				Func: appToolsAdmin.Upload,
			},
		},
	},
	{
		Name: "dux-project/app/tools/listener",
		Annotations: []*annotation.Annotation{
			{
				Name: "Listener",
				Params: map[string]any{
					"name": "tools.backup",
				},
				Func: appToolsListener.BackupListener,
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsArea{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsBackup{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsFile{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsFileDir{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsMagic{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsMagicData{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsMagicGroup{},
			},
		},
	},
	{
		Name: "dux-project/app/tools/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appToolsModels.ToolsMagicSource{},
			},
		},
	},
}

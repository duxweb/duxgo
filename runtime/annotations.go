package runtime

import (
	appSystemAdmin "dux-project/app/system/admin"
	appSystemModels "dux-project/app/system/models"
	appSystemWeb "dux-project/app/system/web"
	appToolsAdmin "dux-project/app/tools/admin"
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
}

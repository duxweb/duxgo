package runtime

import (
	appSystemAdmin "dux-project/app/system/admin"
	appSystemModels "dux-project/app/system/models"
	appSystemWeb "dux-project/app/system/web"
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
					"app":   "admin",
					"name":  "auth",
					"route": "",
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
				Func:   appSystemModels.LogLogin{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogOperate{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisit{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisitData{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisitSpider{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisitUv{},
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
}

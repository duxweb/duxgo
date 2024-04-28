package admin

import (
	model "dux-project/app/system/models"
	"fmt"
	"github.com/duxweb/go-fast/auth"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/i18n"
	"github.com/duxweb/go-fast/menu"
	"github.com/duxweb/go-fast/request"
	"github.com/duxweb/go-fast/response"
	"github.com/go-errors/errors"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"

	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// @RouteGroup(app="admin", name="auth", route="")

// Login @Route(method="POST", name="login", route="/login")
func Login(ctx echo.Context) error {

	var params struct {
		Username string `json:"username" validate:"required" validateMsg:"请输入账号"`
		Password string `json:"password" validate:"required" validateMsg:"请输入密码"`
	}
	if err := request.RequestParser(ctx, &params); err != nil {
		return err
	}

	var user model.SystemUser
	err := database.Gorm().Model(model.SystemUser{Username: params.Username}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.BusinessError("账号或密码错误")
	}
	if err != nil {
		return err
	}

	fmt.Println(user.Password, params.Password)
	isPass := helper.HashVerify(user.Password, params.Password)

	err = loginCheck(user.ID, isPass, ctx)
	if err != nil {
		return err
	}

	if !isPass {
		return response.BusinessError("账号或密码错误")
	}

	token, err := auth.NewJWT().MakeToken("admin", cast.ToString(user.ID))
	if err != nil {
		return err
	}

	rolename := "Admins"
	if len(user.Roles) > 0 {
		rolename = user.Roles[0].Name
	}

	return response.Send(ctx, response.Data{
		Data: map[string]any{
			"userInfo": map[string]any{
				"user_id":  user.ID,
				"avatar":   user.Avatar,
				"username": user.Username,
				"nickname": user.Nickname,
				"rolename": rolename,
			},
			"token":      "Bearer " + token,
			"permission": []string{},
		},
	})
}

// Check @Route(method="POST", name="check", route="/check")
func Check(ctx echo.Context) error {
	var params struct {
		Token string `json:"token" validate:"required" validateMsg:"token does not exist"`
	}
	if err := request.RequestParser(ctx, &params); err != nil {
		return err
	}
	ctx.Request().Header.Set("Authorization", params.Token)
	id := auth.NewService("admin", ctx).ID()
	if id == "" {
		return response.BusinessError("Expired or incorrect token")
	}

	var user model.SystemUser
	err := database.Gorm().Model(&model.SystemUser{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return err
	}

	if !user.Status {
		return response.BusinessError("User Disabled")
	}

	token, err := auth.NewJWT().MakeToken("admin", cast.ToString(user.ID))
	if err != nil {
		return err
	}

	rolename := "Admins"
	if len(user.Roles) > 0 {
		rolename = user.Roles[0].Name
	}

	return response.Send(ctx, response.Data{
		Data: map[string]any{
			"userInfo": map[string]any{
				"user_id":  user.ID,
				"avatar":   user.Avatar,
				"username": user.Username,
				"nickname": user.Nickname,
				"rolename": rolename,
			},
			"token":      "Bearer " + token,
			"permission": []string{},
		},
	})

}

// Menu @Route(method="GET", name="menu", route="/menu")
func Menu(ctx echo.Context) error {
	id := auth.NewService("admin", ctx).ID()
	if id == "" {
		return response.BusinessError("Expired or incorrect token")
	}

	return response.Send(ctx, response.Data{
		Data: menu.Get("admin").Get(),
	})
}

func loginCheck(id uint, isPass bool, ctx echo.Context) error {
	lasSeconds := carbon.Now().SubSeconds(60)

	var data []model.LogLogin
	err := database.Gorm().
		Model(model.LogLogin{}).
		Where("user_type", "system_user").
		Where("user_id = ?", id).
		Where("status = ?", false).
		Where("created_at >= ?", lasSeconds.StdTime()).
		Order("id desc").
		Limit(3).
		Find(&data).Error
	if err != nil {
		return err
	}
	loginCount := len(data)
	loginLast, err := lo.Last[model.LogLogin](data)

	if err == nil && loginCount >= 3 && carbon.CreateFromStdTime(loginLast.CreatedAt).AddSeconds(60).Gt(carbon.Now()) {
		return response.BusinessError(i18n.Trans.Get("system.auth.error.passwordCheck"))
	}

	userAgentString := ctx.Request().UserAgent()

	ua := useragent.Parse(userAgentString)

	database.Gorm().Create(&model.LogLogin{
		UserType: "system_user",
		UserId:   id,
		Browser:  ua.Name + " " + ua.Version,
		Ip:       ctx.RealIP(),
		Platform: ua.OS,
		Status:   isPass,
	})

	return nil
}

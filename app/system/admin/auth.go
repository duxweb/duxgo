package admin

import (
	model "dux-project/app/system/models"
	"fmt"
	"github.com/duxweb/go-fast/auth"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/i18n"
	"github.com/duxweb/go-fast/models"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/validator"
	"github.com/go-errors/errors"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"

	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// @RouteGroup(app="web", name="auth", route="/admin")

// Login @Route(method="POST", name="login", route="/login")
func Login(ctx echo.Context) error {

	var params struct {
		Username string `json:"username" validate:"required" langMessage:"system.auth.validator.username"`
		Password string `json:"password" validate:"required" langMessage:"system.auth.validator.password"`
	}
	if err := validator.RequestParser(ctx, &params); err != nil {
		return err
	}

	var user model.SystemUser
	err := database.Gorm().Model(model.SystemUser{Username: params.Username}).Preload("Roles").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.BusinessLangError("system.auth.error.login")
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
		return response.BusinessLangError("system.auth.error.login")
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
			"permission": user.GetPermission(),
		},
	})
}

// Check @Route(method="POST", name="check", route="/check")
func Check(ctx echo.Context) error {
	var params struct {
		Token string `json:"token" validate:"required" validateMsg:"token does not exist"`
	}
	if err := validator.RequestParser(ctx, &params); err != nil {
		return err
	}
	ctx.Request().Header.Set("Authorization", params.Token)
	id := auth.NewService("admin", ctx).ID()
	if id == "" {
		return response.BusinessError("Expired or incorrect token")
	}

	var user model.SystemUser
	err := database.Gorm().Model(&model.SystemUser{}).Preload("Roles").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return err
	}

	if !*user.Status {
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
			"permission": user.GetPermission(),
		},
	})

}

func loginCheck(id uint, isPass bool, ctx echo.Context) error {
	lasSeconds := carbon.Now().SubSeconds(60)

	var data []models.LogLogin
	err := database.Gorm().
		Model(models.LogLogin{}).
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
	loginLast, err := lo.Last[models.LogLogin](data)

	if err == nil && loginCount >= 3 && loginLast.CreatedAt.AddSeconds(60).Gt(carbon.Now()) {
		return response.BusinessError(i18n.Trans.Get("system.auth.error.passwordCheck"))
	}

	userAgentString := ctx.Request().UserAgent()

	ua := useragent.Parse(userAgentString)

	database.Gorm().Create(&models.LogLogin{
		UserType: "system_user",
		UserId:   id,
		Browser:  ua.Name + " " + ua.Version,
		Ip:       ctx.RealIP(),
		Platform: ua.OS,
		Status:   &isPass,
	})

	return nil
}

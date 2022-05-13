package routes

import (
	"github.com/gin-gonic/gin"
	"goApi/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			//authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)
			// 发送验证码
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)
		}

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "gin",
			})
		})
	}
}

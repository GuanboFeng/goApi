package user

import (
	"github.com/gin-gonic/gin"
	"goApi/pkg/app"
	"goApi/pkg/database"
	"goApi/pkg/paginator"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

//Get 通过ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}

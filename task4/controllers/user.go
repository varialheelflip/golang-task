package controllers

import (
	"blog_system/db"
	"blog_system/models"
	"blog_system/pkg/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// todo 参数校验 用户/邮箱已存在校验
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ServerError(c, "Failed to hash password")
		return
	}
	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		response.ServerError(c, "Failed to create user")
		return
	}

	response.Success(c, "User registered successfully")
}

func (u *UserController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	var storedUser models.User
	if err := db.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		response.BadRequest(c, "Invalid username or password")
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		response.BadRequest(c, "Invalid username or password")
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// todo 密钥配置化
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		response.ServerError(c, "Failed to generate token")
		return
	}
	response.Success(c, tokenString)
}

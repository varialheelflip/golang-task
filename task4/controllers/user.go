package controllers

import (
	"blog_system/config"
	"blog_system/db"
	"blog_system/logger"
	"blog_system/models"
	"blog_system/pkg/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, err.Error())
		return
	}
	// 校验用户名/邮箱已存在
	var judgeUsers []models.User
	db.DB.Limit(1).Where("username = ?", user.Username).Find(&judgeUsers)
	if len(judgeUsers) > 0 {
		response.Fail(c, "用户名已被注册")
		return
	}
	db.DB.Limit(1).Where("email = ?", user.Email).Find(&judgeUsers)
	if len(judgeUsers) > 0 {
		response.Fail(c, "邮箱已被注册")
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, "Failed to hash password")
		return
	}
	user.Password = string(hashedPassword)

	db.DB.Create(&user)

	response.Success(c, user.ID)
}

func (u *UserController) Login(c *gin.Context) {
	var user struct {
		Username string `json:"username" binding:"required,max=20"`
		Password string `json:"password" binding:"required,max=20"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, err.Error())
		return
	}

	var storedUser models.User
	if err := db.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		response.Fail(c, "Invalid username or password")
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		response.Fail(c, "Invalid username or password")
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GlobalConfig.JWT.SecretKey))
	if err != nil {
		response.Fail(c, "Failed to generate token")
		return
	}
	logger.LOGGER.Info("用户登录", zap.Any("userID", storedUser.ID))
	response.Success(c, tokenString)
}

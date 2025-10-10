package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const LOGIN_SECRET_KEY = "secret"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null;size:64" json:"password"`
	Email    string `gorm:"unique;not null;size:128" json:"email"`
}

func Register(c *gin.Context) {
	var user User

	// 参数解析
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// 密码加密
	if password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// 密码设置
		user.Password = string(password)
	}

	// 新增用户
	if DB.Create(&user).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "create user failed!",
		})
	}

	// 返回结构
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}

func Login(c *gin.Context) {
	ctx := context.Background()
	// 登录参数解析
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 根据用户名查询用户
	userDB, err := gorm.G[User](DB).Where("username = ?", user.Username).First(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "find user failed!",
		})
		return
	}

	// 密码比对
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	// 返回jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":       userDB.ID,
		"username":     userDB.Username,
		"reqtimestamp": time.Now().Unix(),
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(LOGIN_SECRET_KEY))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "create token failed!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
}

func GetLoginUId(c *gin.Context) uint {
	userid, exist := c.Get("userid")
	if !exist {
		log.Println("userid is not exist")
	}
	return userid.(uint)
}

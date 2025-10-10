package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null;type:text" json:"content"`
	User_id uint   `gorm:"not null;size:64" json:"userId"`
}

func GetPostParam(c *gin.Context) Post {
	post := Post{}
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(getRequestStatusTxt(requestParamError))
	}
	return post
}

func PostCreate(c *gin.Context) {
	post := GetPostParam(c)

	post.User_id = GetLoginUId(c)
	DB.Create(&post)

	c.JSON(http.StatusOK, ResSuccess())
}

func PostUpdate(c *gin.Context) {
	ctx := context.Background()

	post := GetPostParam(c)
	post.User_id = GetLoginUId(c)

	if _, err := gorm.G[Post](DB).Where("id", post.ID).Updates(ctx, post); err != nil {
		c.JSON(getRequestStatusTxt(operateSucess))
		return
	}
	c.JSON(http.StatusOK, ResSuccess())
}

func PostDelete(c *gin.Context) {
	ctx := context.Background()

	id := c.Param("id")

	if _, err := gorm.G[Post](DB).Where("id", id).Delete(ctx); err != nil {
		c.JSON(getRequestStatusTxt(operateSucess))
		return
	}
	c.JSON(http.StatusOK, ResSuccess())
}

func PostQuery(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	//post := Post{}
	post, err := gorm.G[Post](DB).Where("id = ?", id).First(ctx)

	code, msg := getRequestStatusTxt(operateSucess)

	if err != nil {
		c.JSON(http.StatusOK, ResObj{Code: code, Message: msg, Data: nil})
		return
	}

	c.JSON(http.StatusOK, ResObj{Code: code, Message: msg, Data: post})
}

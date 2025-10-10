package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Content string `gorm:"not null;type:text" json:"content"`
	UserID  uint   `json:"userID"`
	User    User
	PostID  uint `json:"postID"`
	Post    Post
}

func GetCommentsParam(c *gin.Context) Comments {
	comments := Comments{}
	if err := c.ShouldBind(&comments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return comments
}

func commentsCreate(c *gin.Context) {
	comments := GetCommentsParam(c)
	userId := GetLoginUId(c)

	comments.UserID = userId

	comments2 := Comments{
		Content: comments.Content,
		PostID:  comments.PostID,
		UserID:  comments.UserID,
	}

	DB.Create(&comments2)

	c.JSON(http.StatusOK, ResSuccess())
}

func CommentsQuery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResObj{
			Code:    requestParamError,
			Message: "无效的 id 参数",
		})
		return
	}

	var comments []Comments
	result := DB.Where("post_id = ?", id).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ResObj{
			Code:    1,
			Message: result.Error.Error(),
		})
		return
	}

	code, message := getRequestStatusTxt(operateSucess)
	c.JSON(http.StatusOK, ResObj{
		Code:    code,
		Message: message,
		Data:    comments,
	})
}

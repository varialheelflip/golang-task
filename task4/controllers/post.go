package controllers

import (
	"blog_system/db"
	"blog_system/models"
	"blog_system/pkg/response"
	"blog_system/pkg/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if !util.StrNotBlank(post.Title) || !util.StrNotBlank(post.Content) {
		response.BadRequest(c, "文章标题和内容不能为空!")
		return
	}
	post.UserID = util.GetHeaderUserId(c)
	if err := db.DB.Create(&post).Error; err != nil {
		response.ServerError(c, "Failed to create user")
		return
	}
	response.Success(c, nil)
}

// todo 文章列表查询, 支持分页

func Update(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	var oldPost models.Post
	if err := db.DB.First(&oldPost, newPost.ID).Error; err != nil {
		response.ServerError(c, "文章不存在!")
		return
	}
	userId := util.GetHeaderUserId(c)
	if oldPost.UserID != userId {
		response.BadRequest(c, "只有文章的作者才能更新自己的文章!")
		return
	}
	newPost.UserID = userId
	db.DB.Model(&newPost).Updates(models.Post{Title: newPost.Title, Content: newPost.Content})
	response.Success(c, nil)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	if err := db.DB.First(&post, id).Error; err != nil {
		response.ServerError(c, "文章不存在!")
		return
	}
	userId := util.GetHeaderUserId(c)
	if post.UserID != userId {
		response.BadRequest(c, "只有文章的作者才能删除自己的文章!")
		return
	}
	db.DB.Delete(&models.Post{}, id)
	response.Success(c, nil)
}

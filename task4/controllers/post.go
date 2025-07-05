package controllers

import (
	"blog_system/db"
	"blog_system/models"
	"blog_system/pkg/response"
	"blog_system/pkg/util"
	"github.com/gin-gonic/gin"
)

type PostController struct{}

type pageVo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
}

func (p *PostController) Create(c *gin.Context) {
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

func (p *PostController) Page(c *gin.Context) {
	var postQueryDto struct {
		PageNo   uint `form:"pageNo"`
		PageSize uint `form:"pageSize"`
		UserID   uint `form:"userId"`
	}
	if err := c.ShouldBindQuery(&postQueryDto); err != nil {
		response.BadRequest(c, "Invalid Query Param")
		return
	}

	limit := postQueryDto.PageSize
	offset := postQueryDto.PageSize * (postQueryDto.PageNo - 1)
	tx := db.DB.Model(models.Post{})
	if postQueryDto.UserID != 0 {
		tx.Where("user_id = ?", postQueryDto.UserID)
	}

	var total int64
	tx.Count(&total)
	var posts []models.Post
	tx.Offset(int(offset)).Limit(int(limit)).Find(&posts)

	result := buildPostPageVO(posts)
	response.Success(c, struct {
		Total int64    `json:"total"`
		Data  []pageVo `json:"data"`
	}{
		Total: total,
		Data:  result,
	})
}

func buildPostPageVO(postList []models.Post) (voList []pageVo) {
	for _, val := range postList {
		voList = append(voList, pageVo{
			Title:   val.Title,
			Content: val.Content,
			UserID:  val.UserID,
		})
	}
	return
}

func (p *PostController) Detail(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	if err := db.DB.First(&post, id).Error; err != nil {
		response.ServerError(c, "文章不存在!")
		return
	}
	response.Success(c, pageVo{Title: post.Title, Content: post.Content, UserID: post.UserID})
}

func (p *PostController) Update(c *gin.Context) {
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

func (p *PostController) Delete(c *gin.Context) {
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

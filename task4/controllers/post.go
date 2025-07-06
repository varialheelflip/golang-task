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
		response.Fail(c, err.Error())
		return
	}
	post.UserID = util.GetHeaderUserId(c)
	if err := db.DB.Create(&post).Error; err != nil {
		response.Fail(c, "Failed to create post")
		return
	}
	response.Success(c, post.ID)
}

func (p *PostController) Page(c *gin.Context) {
	var postQueryDto struct {
		PageNo   uint `form:"pageNo" binding:"required,min=1"`
		PageSize uint `form:"pageSize" binding:"required,min=1,max=50"`
		UserID   uint `form:"userId"`
	}
	if err := c.ShouldBindQuery(&postQueryDto); err != nil {
		response.Fail(c, "Invalid Query Param")
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
		response.Fail(c, "文章不存在!")
		return
	}
	response.Success(c, pageVo{Title: post.Title, Content: post.Content, UserID: post.UserID})
}

func (p *PostController) Update(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if newPost.ID == 0 {
		response.Fail(c, "找不到文章!")
		return
	}
	var oldPost models.Post
	if err := db.DB.First(&oldPost, newPost.ID).Error; err != nil {
		response.Fail(c, "找不到文章!")
		return
	}
	userId := util.GetHeaderUserId(c)
	if oldPost.UserID != userId {
		response.Fail(c, "只有文章的作者才能更新自己的文章!")
		return
	}
	newPost.UserID = userId
	db.DB.Model(&newPost).Updates(models.Post{Title: newPost.Title, Content: newPost.Content})
	response.Success(c, newPost.ID)
}

func (p *PostController) Delete(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	if err := db.DB.First(&post, id).Error; err != nil {
		response.Fail(c, "文章不存在!")
		return
	}
	userId := util.GetHeaderUserId(c)
	if post.UserID != userId {
		response.Fail(c, "只有文章的作者才能删除自己的文章!")
		return
	}
	db.DB.Delete(&models.Post{}, id)
	response.Success(c, id)
}

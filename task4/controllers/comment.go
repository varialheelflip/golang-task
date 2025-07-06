package controllers

import (
	"blog_system/db"
	"blog_system/models"
	"blog_system/pkg/response"
	"blog_system/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type CommentController struct{}

func (co *CommentController) Create(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if comment.Content == "" || comment.PostID == 0 {
		response.Fail(c, "评论内容为空, 或者未选择文章!")
		return
	}
	comment.UserID = util.GetHeaderUserId(c)
	db.DB.Create(&comment)
	response.Success(c, comment.ID)
}

func (co *CommentController) List(c *gin.Context) {
	id := c.Param("postId")
	if id == "" {
		response.Fail(c, "请选择文章!")
		return
	}
	var result []models.Comment
	db.DB.Where("post_id = ?", id).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("ID", "username") }).
		Find(&result)
	response.Success(c, buildCommentVo(result))
}

type commentVo struct {
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UserName  string `json:"userName"`
}

func buildCommentVo(commentList []models.Comment) []commentVo {
	var result []commentVo
	for _, val := range commentList {
		result = append(result, commentVo{
			Content:   val.Content,
			CreatedAt: val.CreatedAt.Format(time.DateTime),
			UserName:  val.User.Username,
		})
	}
	return result
}

package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Posts []Post
}

type Post struct {
	gorm.Model
	UserID   uint
	Comments []Comment
}

type Comment struct {
	gorm.Model
	PostID uint
}

func createTable(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func findUserInfo(db *gorm.DB) {
	var users []User
	db.Preload("Posts.Comments").Find(&users, 1)
	PrintToJson(users)
}

func findMaxCommentPost(db *gorm.DB) {
	var postResultView []struct {
		PostID       uint
		CommentCount uint
	}
	db.Raw("select post_id, count(post_id) as comment_count from comments group by post_id order by count(post_id) desc").Scan(&postResultView)
	if len(postResultView) <= 0 {
		fmt.Println("文章没有评论信息")
		return
	}
	// 获取评论最多的文章id(可能有多个)
	var maxPostIds []uint
	maxCount := postResultView[0].CommentCount
	for _, val := range postResultView {
		if val.CommentCount == maxCount {
			maxPostIds = append(maxPostIds, val.PostID)
		} else {
			break
		}
	}
	var maxPost []Post
	db.Preload("Comments").Find(&maxPost, maxPostIds)
	PrintToJson(maxPost)
}

func PrintToJson(o interface{}) {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	fmt.Println(err)
	fmt.Println(string(jsonData))
}

func main() {
	dsn := "root:root@tcp(192.168.200.130:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 题目1
	//createTable(db)
	// 题目2-1
	//findUserInfo(db)
	// 题目2-2
	findMaxCommentPost(db)
}

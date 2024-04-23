package controller

import (
	"github.com/gin-gonic/gin"
)

// Router 实例化router类型对象，首字母大写用于跨包调用
var Router router

// 声明router结构体w
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		// 用户相关
		POST("/user/sendsms", User.SendSMS).
		POST("/user/signup", User.Signup).
		POST("/user/login", User.Login).
		GET("/user/get", User.GetUser).
		// 帖子相关
		GET("/article/list", Article.GetArticleList).
		GET("/article/get", Article.GetArticle).
		POST("/article/create", Article.CreateArticle).
		// 评论相关
		POST("/comment/send", Comment.SendComment).
		GET("/comment/home", Comment.GetHomeComment).
		GET("/comment/list", Comment.GetComment)
}

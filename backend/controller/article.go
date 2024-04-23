package controller

import (
	"TreeHole/dao"
	"TreeHole/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Article 实例化chat类型对象，首字母大写用于跨包调用
var Article article

// 声明article结构体
type article struct{}

// CreateArticle 创建帖子
func (a article) CreateArticle(ctx *gin.Context) {
	// 拿到身份
	claims, _ := ctx.Get("claims")
	role := claims.(map[string]interface{})["role"]
	id := claims.(map[string]interface{})["id"]
	//参数绑定
	params := new(struct {
		Title string `form:"title" binding:"required"`
		Text  string `form:"text" binding:"required"`
	})
	if role.(float64) >= 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "权限不够！",
			"data": nil,
		})
		return
	}
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	createArticle, err := dao.Dao.CreateArticle(utils.GetUint(fmt.Sprint(id)), params.Text, params.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "创建成功",
		"data": createArticle,
	})
}

// GetArticleList 获取帖子列表
func (a article) GetArticleList(ctx *gin.Context) {
	// 拿到身份
	//claims, _ := ctx.Get("claims")
	//role := claims.(map[string]interface{})["role"]
	articles, err := dao.Dao.GetArticles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取帖子列表失败",
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取帖子列表成功",
		"data": articles,
	})
}

// GetArticle 获取某个帖子内容
func (a article) GetArticle(ctx *gin.Context) {
	//参数绑定
	params := new(struct {
		ID string `form:"id" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	// 拿到身份
	//claims, _ := ctx.Get("claims")
	//role := claims.(map[string]interface{})["role"]
	articleInfo, err := dao.Dao.GetArticle(params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取帖子失败",
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取帖子成功",
		"data": articleInfo,
	})
}

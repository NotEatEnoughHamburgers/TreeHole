package controller

import (
	"TreeHole/dao"
	"TreeHole/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

var User user

type user struct {
}

func (u *user) SendSMS(ctx *gin.Context) {
	params := new(struct {
		Number string `form:"number" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	// 发送验证码
	err := utils.SendMsg(params.Number)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "验证码发送成功",
		"data": nil,
	})
}

// Signup 注册
func (u *user) Signup(ctx *gin.Context) {
	// 参数绑定
	params := new(struct {
		Name   string `form:"name" binding:"required"`
		Number string `form:"number" binding:"required"`
		Pass   string `form:"pass"  binding:"required"`
		//Img    string `form:"img_url"  binding:"required"`
		Img   string `form:"img_url"`
		Email string `form:"email"`
		Code  string `form:"code" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	// 核验是否已注册
	numberUser, err := dao.Dao.CheckNumberUser(params.Number)
	if err != nil {
		if err.Error() == "用户检查出现错误:record not found" {

		} else {
			fmt.Println("核验用户是否已注册失败")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}
	}
	if numberUser.Number != "" {
		fmt.Println(numberUser.Number + "已被注册")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  numberUser.Number + "已被注册",
			"data": nil,
		})
		return
	}

	// 验证码核验错误
	if params.Code != utils.Code[params.Number] {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "验证码核验失败",
			"data": nil,
		})
		return
	}
	if params.Email == "" {
		params.Email = generateRandomEmail()
	}
	// 创建用户
	err = dao.Dao.CreateUser(params.Name, params.Number, utils.CalculateMD5Hash(params.Pass), params.Img, params.Email, 2)
	if err != nil {
		fmt.Println("注册创建用户时出现错误", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fmt.Sprint("注册创建用户时出现错误", err),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功！",
		"data": nil,
	})
	return
}

// Login 登录
func (u *user) Login(ctx *gin.Context) {
	// 参数绑定
	params := new(struct {
		Number string `form:"number" binding:"required"`
		Pass   string `form:"pass"  binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	check, err := dao.Dao.CheckUser(params.Number, utils.CalculateMD5Hash(params.Pass))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	//token, err := utils.GenerateToken(check.GetID(), int(check.Role))
	token, err := utils.GenerateToken(fmt.Sprint(check.ID), int(check.Role))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fmt.Sprint("生成token出现错误:", err.Error()),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功!",
		"data": token,
	})
	return
}

// GetUser 获取某个用户的详细信息
func (u *user) GetUser(ctx *gin.Context) {
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
	articleInfo, err := dao.Dao.GetUser(params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取用户信息失败",
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取用户信息成功",
		"data": articleInfo,
	})
}

func generateRandomEmail() string {
	// 生成随机字符串作为邮箱用户名部分
	rand.Seed(time.Now().UnixNano())
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	username := make([]byte, 8)
	for i := range username {
		username[i] = charSet[rand.Intn(len(charSet))]
	}

	// 生成随机的域名部分
	domain := "example.com" // 可以根据需要更改域名

	return fmt.Sprintf("%s@%s", string(username), domain)
}

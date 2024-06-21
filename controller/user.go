package controller

import (
	"fmt"
	"net/http"
	db "templategoapi/db"
	"templategoapi/middlewares"
	"templategoapi/model"
	"templategoapi/repo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Login
func Login(resource *db.Resource) func(c *gin.Context) {
	type infoBody struct {
		Username string `json:"username"  binding:"required"`
		Password string `json:"password"  binding:"required"`
	}
	return func(c *gin.Context) {
		var body infoBody
		if err := c.ShouldBind(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{
			"username": body.Username,
		}
		user := model.UserModelS{}
		//
		if err := repo.GetOneStatement(resource, "user", filter, nil, &user); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "Login Fail | Username or Password is incorrect | username not found",
				"erro": err.Error(),
			})

			return
		}

		filter = bson.M{
			"username":    body.Username,
			"password":    GetMD5Hash(body.Password),
			"user_delete": false,
			"active":      true,
		}
		if err := repo.GetOneStatement(resource, "user", filter, nil, &user); err != nil {
			CreateLogUseUsername(resource, c, "", body.Username, "Login Fail", "Login Fail | Username or Password is incorrect : "+err.Error())
			c.JSON(400, gin.H{
				"code": 400,
				"text": "Login Fail | Username or Password is incorrect",
				"erro": err.Error(),
			})

			return
		}

		token, _ := middlewares.GenToken(c, user)
		//c.ClientIP() เรียก IP จาก ผู้ใช้
		fmt.Println("createLog run")
		fmt.Println("IP : " + c.ClientIP())
		c.Set("username", user.Username)
		c.Set("name", user.Name)
		c.Set("user_id", user.ID)
		CreateLogUseUsername(resource, c, user.ID, user.Username, "Login Success", " Login Success")

		c.JSON(200, gin.H{
			"code":  200,
			"text":  "Login Success",
			"token": token,
		})
	}
}

// Login
func IsLogin(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {

		filter := bson.M{
			"username": c.GetString("username"),
		}
		user := model.UserModelS{}
		if err := repo.GetOneStatement(resource, "user", filter, nil, &user); err != nil {

			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		token, _ := middlewares.GenToken(c, user)
		c.JSON(200, gin.H{
			"code": 200,
			"text": "Success",
			"data": token,
		})
	}
}

// get User from token
func GetUser(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		type infoBody struct {
			User_id  string `json:"user_id"  `
			Username string `json:"username"  `
			Name     string `json:"name"  `
		}
		var payload infoBody

		payload.Name = c.GetString("username")
		payload.User_id = c.GetString("user_id")
		payload.Username = c.GetString("username")

		c.JSON(200, gin.H{
			"code":    200,
			"payload": payload,
		})
	}
}
func LogOut(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		LoginKey := c.Request.Header.Get("LoginKey")
		if LoginKey != c.GetString("LoginKey") {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "you are not login",
			})
			return
		}
		c.Set("token", nil)
		c.Set("username", nil)
		c.Set("name", nil)
		c.Set("user_id", nil)
		c.JSON(200, gin.H{
			"code": 200,
			"text": "Success",
		})
	}
}

package controller

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"templategoapi/db"
	"templategoapi/model"
	"templategoapi/repo"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type LogModel struct {
	User_ID     string    `bson:"user_id" json:"user_id"`
	Username    string    `bson:"username" json:"username"`
	Name        string    `bson:"name" json:"name"`
	IpAddress   string    `bson:"ip_address" json:"ip_address"`
	TimeStamp   time.Time `bson:"time_stamp" json:"time_stamp"`
	Activity    string    `bson:"activity" json:"activity"`
	Description string    `bson:"description" json:"description"`
}
type LogModelForUpdate struct {
	User_ID      string    `bson:"user_id" json:"user_id"`
	Username     string    `bson:"username" json:"username"`
	Name         string    `bson:"name" json:"name"`
	IpAddress    string    `bson:"ip_address" json:"ip_address"`
	TimeStamp    time.Time `bson:"time_stamp" json:"time_stamp"`
	Activity     string    `bson:"activity" json:"activity"`
	Description  string    `bson:"description" json:"description"`
	BeforeUpdate string    `bson:"before_update" json:"before_update"`
	AfterUpdate  string    `bson:"after_update" json:"after_update"`
}

func CreateLog(resource *db.Resource, c *gin.Context, ac, dc string) {
	LogData := LogModel{
		User_ID:     c.GetString("user_id"),
		Username:    c.GetString("username"),
		Name:        c.GetString("name"),
		IpAddress:   c.ClientIP(),
		TimeStamp:   time.Now(),
		Activity:    ac,
		Description: dc,
	}

	result, err := repo.CreateStatement(resource, "log", LogData)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("log Success result: ", result)

}
func CreateLogUseUsername(resource *db.Resource, c *gin.Context, uid, un, ac, dc string) {

	LogData := LogModel{
		User_ID:     uid,
		Username:    un,
		Name:        "",
		IpAddress:   c.ClientIP(),
		TimeStamp:   time.Now(),
		Activity:    ac,
		Description: dc,
	}

	result, err := repo.CreateStatement(resource, "log", LogData)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("log Success result: ", result)

}
func CreateLogForUpdate(resource *db.Resource, c *gin.Context, ac string, dc string, beforeUpdate interface{}, afterUpdate interface{}) {
	bu, err := json.Marshal(beforeUpdate)
	if err != nil {
		panic(err)
	}

	au, err := json.Marshal(afterUpdate)
	if err != nil {
		panic(err)
	}
	LogData := LogModelForUpdate{
		User_ID:      c.GetString("user_id"),
		Username:     c.GetString("username"),
		Name:         c.GetString("name"),
		IpAddress:    c.ClientIP(),
		TimeStamp:    time.Now(),
		Activity:     ac,
		Description:  dc,
		BeforeUpdate: string(bu),
		AfterUpdate:  string(au),
	}

	result, err := repo.CreateStatement(resource, "log", LogData)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("log Success result: ", result)

}
func GetAllLog(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                     value,
			"user_delete":                  false,
			"role.permission.all_log.read": true,
			"active":                       true,
		}
		filterOption1 := bson.M{}

		err := repo.GetOneStatement(resource, "user", filter1, filterOption1, &permission)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "No permission || permission Needed",
				"payload": err.Error(),
			})
			return
		}
		var entity []LogModel
		filter := bson.M{}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		skip := (page - 1) * limit
		if skip < 0 {
			skip = 0
		}
		filterOption := bson.M{"time_stamp": -1}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "log", filter, filterOption, limit, page, &entity)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code":           200,
			"total_page":     totalPages,
			"total_document": count,
			"current_paget":  page,
			"limit_data":     limit,
			"payload":        entity,
		})
	}
}

func GetLogByAccount(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Username string `bson:"username" json:"username" form:"username"`
	}
	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                 value,
			"user_delete":              false,
			"role.permission.log.read": true,
			"active":                   true,
		}
		filterOption1 := bson.M{}

		err := repo.GetOneStatement(resource, "user", filter1, filterOption1, &permission)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "No permission || permission Needed",
				"payload": err.Error(),
			})
			return
		}
		var entity []model.LogModel
		filter := bson.M{
			"username": body.Username,
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		skip := (page - 1) * limit
		if skip < 0 {
			skip = 0
		}
		filterOption := bson.M{"time_stamp": -1}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "log", filter, filterOption, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		if entity == nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": "not found",
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].Username < entity[j].Username
		})
		c.JSON(200, gin.H{
			"code":           200,
			"total_page":     totalPages,
			"total_document": count,
			"current_paget":  page,
			"limit_data":     limit,
			"payload":        entity,
		})
	}
}

// func CreateLog(resource *db.Resource, ip, ac, des string) func(c *gin.Context) {
// 	return func(c *gin.Context) {

// 		LogData := LogModel{
// 			User_ID:     c.GetString("user_id"),
// 			Username:    c.GetString("username"),
// 			IpAddress:   ip,
// 			TimeStamp:   time.Now(),
// 			Activity:    ac,
// 			Description: des,
// 		}

// 		_, err := repo.CreateStatement(resource, "log", LogData)
// 		if err != nil {
// 			fmt.Print(err)
// 			return
// 		}
// 		fmt.Println("Create LOG Success")
// 	}
// }

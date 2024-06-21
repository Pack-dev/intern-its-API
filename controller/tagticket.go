package controller

import (
	"templategoapi/db"
	"templategoapi/model"
	"templategoapi/repo"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTagTicket(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body model.TagTicketModel

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
			"username":                          value,
			"user_delete":                       false,
			"role.permission.tag_ticket.create": true,
			"active":                            true,
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

		TagTicketData := model.TagTicketModel{
			TagTicketName: body.TagTicketName,
			CreateTime:    time.Now(),
		}

		result, err := repo.CreateStatement(resource, "TagTicket", TagTicketData)
		if err != nil {
			CreateLog(resource, c, "CreateTagTicket fail", "CreateTagTicket fail | error : "+err.Error())

			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "CreateTagTicket success", "CreateTagTicket name : "+body.TagTicketName+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
func UpdateOneTagTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ID            string `bson:"_id" json:"_id" from:"_id"`
		TagTicketName string `bson:"tag_ticket_name" json:"tag_ticket_name"`
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
			"username":                          value,
			"user_delete":                       false,
			"role.permission.tag_ticket.update": true,
			"active":                            true,
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
		ID, err := primitive.ObjectIDFromHex(body.ID)
		if err != nil {
			CreateLog(resource, c, "UpdateOneTagTicket fail", "ID not Hex format | error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{
			"_id":               ID,
			"tag_ticket_delete": false}
		update := bson.M{"$set": bson.M{"tag_ticket_name": body.TagTicketName, "update_time": time.Now()}}

		var beforeUpdate model.TagTicketModel
		err = repo.GetOneStatement(resource, "TagTicket", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "TagTicket", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateOneTagTicket fail", "UpdateOneStatement fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.TagTicketModel
		err = repo.GetOneStatement(resource, "TagTicket", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateOneTagTicket success", "UpdateOneTagTicket name : "+body.TagTicketName+" success", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}

// เดี๋ยวแก้เป็น soft delete
func DeleteOneTagTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		TagTicketName string `bson:"tag_ticket_name" json:"tag_ticket_name" form:"tag_ticket_name"`
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
			"username":                          value,
			"user_delete":                       false,
			"role.permission.tag_ticket.delete": true,
			"active":                            true,
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
		filter := bson.M{
			"tag_ticket_name":   body.TagTicketName,
			"tag_ticket_delete": false}
		update := bson.M{"$set": bson.M{"tag_ticket_delete": true, "delete_time": time.Now()}}

		Result, err := repo.UpdateOneStatement(resource, "TagTicket", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneTagTicket fail", "DeleteOneTagTicket fail | error : "+err.Error())

			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Delete Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneTagTicket success", "DeleteOneTagTicket name : "+body.TagTicketName+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}
func GetAllTagTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.TagTicketModelID
		filter := bson.M{
			"tag_ticket_delete": false,
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.tag_ticket.read": true,
			"active":                          true,
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
		if err := repo.GetManyStatement(resource, "TagTicket", filter, nil, &entity); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": entity,
		})
	}
}

func DropdownGetAllTagTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.TagTicketModelID
		filter := bson.M{
			"tag_ticket_delete": false,
		}
		if err := repo.GetManyStatement(resource, "TagTicket", filter, nil, &entity); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": entity,
		})
	}
}

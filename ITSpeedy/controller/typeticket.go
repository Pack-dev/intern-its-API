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

func CreateTypeTicket(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body model.TypeTicketModel

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
			"username":                           value,
			"user_delete":                        false,
			"role.permission.type_ticket.create": true,
			"active":                             true,
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

		TypeTicketData := model.TypeTicketModel{
			TypeTicketName: body.TypeTicketName,
			CreateTime:     time.Now(),
		}

		result, err := repo.CreateStatement(resource, "typeTicket", TypeTicketData)
		if err != nil {
			CreateLog(resource, c, "CreateTypeTicket fail", "CreateTypeTicket fail | error : "+err.Error())

			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "CreateTypeTicket success", "CreateTypeTicket name : "+body.TypeTicketName+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
func UpdateOneTypeTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ID             string `bson:"_id" json:"_id" from:"_id"`
		TypeTicketName string `bson:"type_ticket_name" json:"type_ticket_name"`
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
			"username":                           value,
			"user_delete":                        false,
			"role.permission.type_ticket.update": true,
			"active":                             true,
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
			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{
			"_id":                ID,
			"type_ticket_delete": false}
		update := bson.M{"$set": bson.M{"type_ticket_name": body.TypeTicketName, "updatetime": time.Now()}}

		var beforeUpdate model.TypeTicketModel
		err = repo.GetOneStatement(resource, "typeTicket", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "typeTicket", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateOneTypeTicket fail", "UpdateOneTypeTicket fail | error : "+err.Error())

			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.TypeTicketModel
		err = repo.GetOneStatement(resource, "typeTicket", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateOneTypeTicket success", "UpdateOneTypeTicket name : "+body.TypeTicketName+" success", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}

// เดี๋ยวแก้เป็น soft delete
func DeleteOneTypeTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		TypeTicketName string `bson:"type_ticket_name" json:"type_ticket_name" form:"type_ticket_name"`
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
			"username":                           value,
			"user_delete":                        false,
			"role.permission.type_ticket.delete": true,
			"active":                             true,
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
			"type_ticket_name":   body.TypeTicketName,
			"type_ticket_delete": false}
		update := bson.M{"$set": bson.M{"type_ticket_delete": true, "delete_time": time.Now()}}

		Result, err := repo.UpdateOneStatement(resource, "typeTicket", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneTypeTicket fail", "DeleteOneTypeTicket fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Delete Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneTypeTicket success", "DeleteOneTypeTicket name : "+body.TypeTicketName+" sucess")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}
func GetAllTypeTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.TypeTicketModelID
		filter := bson.M{
			"type_ticket_delete": false,
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                         value,
			"user_delete":                      false,
			"role.permission.type_ticket.read": true,
			"active":                           true,
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
		if err := repo.GetManyStatement(resource, "typeTicket", filter, nil, &entity); err != nil {
			c.JSON(400, gin.H{
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

func DropdownGetAllTypeTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.TypeTicketModelID
		filter := bson.M{
			"type_ticket_delete": false,
		}
		if err := repo.GetManyStatement(resource, "typeTicket", filter, nil, &entity); err != nil {
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

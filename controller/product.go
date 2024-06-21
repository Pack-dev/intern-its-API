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

func CreateProduct(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body model.ProductModel

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
			"username":                       value,
			"user_delete":                    false,
			"role.permission.product.create": true,
			"active":                         true,
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
		// ให้ตั๋งทำต่อ
		ProductData := model.ProductModel{
			ProductName: body.ProductName,
			CreateTime:  time.Now(),
		}

		result, err := repo.CreateStatement(resource, "Product", ProductData)
		if err != nil {
			CreateLog(resource, c, "Create Product fail", "Create Product fail : "+err.Error())
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "Create Product success", "Create Product name : "+body.ProductName+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
func UpdateOneProduct(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ID          string `bson:"_id" json:"_id" from:"_id"`
		ProductName string `bson:"product_name" json:"product_name"`
	}
	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                       value,
			"user_delete":                    false,
			"role.permission.product.update": true,
			"active":                         true,
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
			CreateLog(resource, c, "UpdateOneProduct fail", "ID not from Hex : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{
			"_id":            ID,
			"product_delete": false}
		update := bson.M{"$set": bson.M{"product_name": body.ProductName, "update_time": time.Now()}}

		var beforeUpdate model.ProductModel
		err = repo.GetOneStatement(resource, "Product", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "Product", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateOneProduct fail", "Update Product name : "+body.ProductName+" fail : "+err.Error())

			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.ProductModel
		err = repo.GetOneStatement(resource, "Product", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateOneProduct success", "Update Product name : "+body.ProductName+" success", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}

// เดี๋ยวแก้เป็น soft delete
func DeleteOneProduct(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ProductName string `bson:"product_name" json:"product_name" form:"product_name"`
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
			"username":                       value,
			"user_delete":                    false,
			"role.permission.product.delete": true,
			"active":                         true,
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
			"product_name":   body.ProductName,
			"product_delete": false}
		update := bson.M{"$set": bson.M{"product_delete": true, "delete_time": time.Now()}}

		Result, err := repo.UpdateOneStatement(resource, "Product", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneProduct fail", "DeleteOneProduct fail : "+err.Error())

			c.JSON(400, gin.H{
				"code":    401,
				"msg":     "Delete Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneProduct success", "Delete Product name : "+body.ProductName+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}
func GetAllProduct(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.ProductModelID
		filter := bson.M{
			"product_delete": false,
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                     value,
			"user_delete":                  false,
			"role.permission.product.read": true,
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
		if err := repo.GetManyStatement(resource, "Product", filter, nil, &entity); err != nil {
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
func DropdownGetAllProduct(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.ProductModelID
		filter := bson.M{
			"product_delete": false,
		}
		if err := repo.GetManyStatement(resource, "Product", filter, nil, &entity); err != nil {
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

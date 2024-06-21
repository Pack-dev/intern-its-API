package controller

import (
	"fmt"
	"sort"
	"strconv"
	"templategoapi/db"
	"templategoapi/model"
	"templategoapi/repo"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCustomer(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var body struct {
			Prefix      string   `bson:"prefix" json:"prefix"`
			FirstName   string   `bson:"first_name" json:"first_name"`
			LastName    string   `bson:"last_name" json:"last_name"`
			PhoneNumber string   `bson:"phone_number" json:"phone_number"`
			Customer    string   `bson:"customer" json:"customer"`
			Product     []string `bson:"product" json:"product"`
		}

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		var permission model.UserModelS
		value := c.GetString("username")
		fmt.Printf("%s", value)
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.customer.create": true,
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

		var ProductModel []model.ProductModelS
		if len(body.Product) == 0 {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "Require Product",
			})
			return
		}
		for _, v := range body.Product {
			var ProductModelONE model.ProductModelS
			filter1 := bson.M{
				"product_name":   v,
				"product_delete": false,
			}
			filterOption1 := bson.M{}

			err := repo.GetOneStatement(resource, "Product", filter1, filterOption1, &ProductModelONE)
			if err != nil {
				CreateLog(resource, c, "UpdateOneCustomer fail", "Product name : "+v+" not found | "+err.Error())
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "Product name : " + v + " not found ",
					"payload": err.Error(),
				})
				return
			}
			ProductModel = append(ProductModel, ProductModelONE)
		}
		CustomerData := model.CustomerModel{
			Prefix:      body.Prefix,
			FirstName:   body.FirstName,
			LastName:    body.LastName,
			FullName:    body.FirstName + " " + body.LastName,
			PhoneNumber: body.PhoneNumber,
			Customer:    body.Customer,
			Product:     ProductModel,
			CreateTime:  time.Now(),
		}

		result, err := repo.CreateStatement(resource, "Customer", CustomerData)
		if err != nil {
			CreateLog(resource, c, "CreateCustomer fail", "CreateCustomer fail : "+err.Error())

			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "CreateCustomer success", "CreateCustomer name : "+body.FirstName+" "+body.LastName+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
func UpdateOneCustomer(resource *db.Resource) func(c *gin.Context) {
	type ProductModelS struct {
		ProductName string `bson:"product_name" json:"product_name"`
	}
	type Body struct {
		ID          string   `bson:"_id" json:"_id" from:"_id"`
		Prefix      string   `bson:"prefix" json:"prefix"`
		FirstName   string   `bson:"first_name" json:"first_name"`
		LastName    string   `bson:"last_name" json:"last_name"`
		PhoneNumber string   `bson:"phone_number" json:"phone_number"`
		Customer    string   `bson:"customer" json:"customer"`
		Product     []string `bson:"product" json:"product"`
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
		ID, err := primitive.ObjectIDFromHex(body.ID)
		if err != nil {
			CreateLog(resource, c, "UpdateOneRole fail", "ID not hex format |error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		var permission model.UserModelS
		value := c.GetString("username")
		fmt.Printf("%s", value)
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.customer.update": true,
			"active":                          true,
		}
		filterOption1 := bson.M{}

		err = repo.GetOneStatement(resource, "user", filter1, filterOption1, &permission)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "No permission || permission Needed",
				"payload": err.Error(),
			})
			return
		}
		var ProductModel []model.ProductModelS
		for _, v := range body.Product {
			var ProductModelONE model.ProductModelS
			filter1 := bson.M{
				"product_name":   v,
				"product_delete": false,
			}
			filterOption1 := bson.M{}

			err := repo.GetOneStatement(resource, "Product", filter1, filterOption1, &ProductModelONE)
			if err != nil {
				CreateLog(resource, c, "UpdateOneCustomer fail", "Product name : "+v+" not found | "+err.Error())
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "Product name : " + v + " not found ",
					"payload": err.Error(),
				})
				return
			}
			ProductModel = append(ProductModel, ProductModelONE)
		}
		filter := bson.M{
			"_id":             ID,
			"customer_delete": false}
		update := bson.M{"$set": bson.M{
			"prefix":       body.Prefix,
			"first_name":   body.FirstName,
			"last_name":    body.LastName,
			"full_name":    body.FirstName + " " + body.LastName,
			"phone_number": body.PhoneNumber,
			"update_time":  time.Now(),
			"customer":     body.Customer,
			"product":      ProductModel,
		}}

		var beforeUpdate model.CustomerModel
		err = repo.GetOneStatement(resource, "Customer", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "Customer", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateOneCustomer fail", "UpdateOneCustomer fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.CustomerModel
		err = repo.GetOneStatement(resource, "Customer", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateOneCustomer success", "UpdateOneCustomer name : "+body.FirstName+" "+body.LastName+" success", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}

func DeleteOneCustomer(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ID string `bson:"_id" json:"_id" form:"_id"`
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
		fmt.Printf("%s", value)
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.customer.delete": true,
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
		ID, err := primitive.ObjectIDFromHex(body.ID)
		if err != nil {
			CreateLog(resource, c, "DeleteOneCustomer fail", "ID not hex format |error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{"_id": ID}
		update := bson.M{"$set": bson.M{"customer_delete": true, "delete_time": time.Now()}}

		Result, err := repo.UpdateOneStatement(resource, "Customer", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneCustomer fail", "DeleteOneCustomer fail | error : "+err.Error())

			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Delete Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneCustomer success", "DeleteOneCustomer name : "+body.ID+"  success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}
func GetAllCustomer(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.CustomerModelID
		filter := bson.M{
			"customer_delete": false,
		}
		var permission model.UserModelS
		value := c.GetString("username")
		fmt.Printf("%s", value)
		filter1 := bson.M{
			"username":                      value,
			"user_delete":                   false,
			"role.permission.customer.read": true,
			"active":                        true,
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
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		skip := (page - 1) * limit
		if skip < 0 {
			skip = 0
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "Customer", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].FullName < entity[j].FullName
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

func GetOneCustomer(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ID string `json:"_id"`
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

		ID, err := primitive.ObjectIDFromHex(body.ID)
		if err != nil {
			CreateLog(resource, c, "Get one Customer Fail", " ID not from Hex | err :	"+err.Error())
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		// test login
		var CustomerModel model.CustomerModelID
		filter := bson.M{
			"_id":             ID,
			"customer_delete": false,
		}
		filterOption := bson.M{}

		err = repo.GetOneStatement(resource, "Customer", filter, filterOption, &CustomerModel)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": filter,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"payload": CustomerModel,
		})
	}
}

func DropdownGetAllCustomer(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.CustomerModelID
		filter := bson.M{
			"customer_delete": false,
		}
		if err := repo.GetManyStatement(resource, "Customer", filter, nil, &entity); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].FullName < entity[j].FullName
		})
		c.JSON(200, gin.H{
			"code":    200,
			"payload": entity,
		})
	}
}

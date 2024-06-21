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

func CreateRole(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body model.RoleModel

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
			"username":                    value,
			"user_delete":                 false,
			"role.permission.role.create": true,
			"active":                      true,
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

		RoleData := model.RoleModel{
			Rolename:   body.Rolename,
			Permission: body.Permission,
			CreateTime: time.Now(),
		}

		result, err := repo.CreateStatement(resource, "role", RoleData)
		if err != nil {
			CreateLog(resource, c, "CreateRole fail", "CreateRole fail : "+err.Error())

			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "CreateRole success", "CreateRole name : "+body.Rolename+" success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
func UpdateOneRole(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		ID       string `bson:"_id" json:"_id" from:"_id"`
		Rolename string `bson:"role_name" json:"role_name"`
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
		filter1 := bson.M{
			"username":                    value,
			"user_delete":                 false,
			"role.permission.role.update": true,
			"active":                      true,
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

		filter := bson.M{
			"_id":         ID,
			"role_delete": false}
		update := bson.M{"$set": bson.M{"role_name": body.Rolename, "update_time": time.Now()}}

		var beforeUpdate model.RoleModel
		err = repo.GetOneStatement(resource, "role", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "role", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateOneRole fail", "UpdateOneRole fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.RoleModel
		err = repo.GetOneStatement(resource, "role", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateOneRole success", "UpdateOneRole name : "+body.Rolename+" success", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}
func UpdateOneRoleByRole_name(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Rolename   string               `bson:"role_name" json:"role_name"`
		Permission model.RolePermission `bson:"permission" json:"permission"`
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
			"username":                    value,
			"user_delete":                 false,
			"role.permission.role.update": true,
			"active":                      true,
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
			"role_name":   body.Rolename,
			"role_delete": false}
		update := bson.M{"$set": bson.M{"role_name": body.Rolename, "permission": body.Permission, "update_time": time.Now()}}

		var beforeUpdate model.RoleModel
		err = repo.GetOneStatement(resource, "role", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "role", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateOneRole fail", "UpdateOneRole fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var updatedRole model.RoleModel
		err = repo.GetOneStatement(resource, "role", filter, nil, &updatedRole)
		if err != nil {
			CreateLog(resource, c, "Get updated role fail", "Get updated role fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Get updated role fail",
				"payload": err.Error(),
			})
			return
		}

		userFilter := bson.M{
			"role.role_name": body.Rolename,
			"user_delete":    false,
			"active":         true,
		}
		userUpdate := bson.M{"$set": bson.M{"role": updatedRole}}

		err = repo.UpdateManyStatement(resource, "user", userFilter, userUpdate)
		if err != nil {
			CreateLog(resource, c, "Update related users fail", "Update related users fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update related users fail",
				"payload": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 200,
			"text": "Update role and related users success",
		})
		var afterUpdate model.RoleModel
		err = repo.GetOneStatement(resource, "role", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateOneRole success", "UpdateOneRole name : "+body.Rolename+" success", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}

func DeleteOneRole(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Rolename string `bson:"role_name" json:"role_name" form:"role_name"`
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
			"username":                    value,
			"user_delete":                 false,
			"role.permission.role.delete": true,
			"active":                      true,
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
			"role_name":   body.Rolename,
			"role_delete": false}
		update := bson.M{"$set": bson.M{"role_delete": true, "delete_time": time.Now()}}

		Result, err := repo.UpdateOneStatement(resource, "role", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneRole fail", "DeleteOneRole fail | error : "+err.Error())

			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Delete Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneRole success", "DeleteOneRole name : "+body.Rolename+"  success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})

	}
}
func GetAllRole(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.RoleModel
		filter := bson.M{
			"role_delete": false,
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                  value,
			"user_delete":               false,
			"role.permission.role.read": true,
			"active":                    true,
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

		if err := repo.GetManyStatement(resource, "role", filter, nil, &entity); err != nil {
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
func GetRoleByRole_name(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Rolename string `bson:"role_name" json:"role_name" form:"role_name"`
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
		var entity model.RoleModel

		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                  value,
			"user_delete":               false,
			"role.permission.role.read": true,
			"active":                    true,
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
			"role_name":   body.Rolename,
			"role_delete": false,
		}
		if err := repo.GetOneStatement(resource, "role", filter, nil, &entity); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"msg":     "Get Role Fail",
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

func DropdownGetAllRole(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.RoleModel
		filter := bson.M{
			"role_delete": false,
		}
		if err := repo.GetManyStatement(resource, "role", filter, nil, &entity); err != nil {
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

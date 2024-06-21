package controller

import (
	"templategoapi/db"
	"templategoapi/repo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Test(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"SERVICE": "G2EPAY",
			"RsaKey":  c.GetString("RsaKey"),
		})
	}
}
func UndeleteAllUser(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {

		filter := bson.M{"user_delete": true}
		update := bson.M{"$set": bson.M{
			"user_delete": false,
		}}

		err := repo.UpdateManyStatement(resource, "user", filter, update)
		if err != nil {

			c.JSON(400, gin.H{
				"code":    400,
				"text":    "undelete user Fail",
				"payload": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"msg":     "undelete user success",
			"payload": nil,
		})
	}
}

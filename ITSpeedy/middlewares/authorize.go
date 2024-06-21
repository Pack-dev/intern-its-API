package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Headerdata Header data

// AuthRequired to check the authentication key in HTTP Header
func Authorization(level int16) gin.HandlerFunc {
	return func(c *gin.Context) {

		// roleName, _ := c.Get("roleName")
		roleLevel, _ := c.Get("level")
		roleLevelInt := roleLevel.(int16)
		fmt.Println(roleLevelInt)
		if roleLevelInt < level {
			c.AbortWithStatusJSON(401, gin.H{
				"code": 401,
				"msg":  "สิทธิ์ของคุณไม่สามารถเข้าใช้งานได้",
			})
			return
		}
		c.Next()
	}
}

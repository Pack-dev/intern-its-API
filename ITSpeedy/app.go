package templategoapi

import (
	"net/http"
	"templategoapi/db"
	"templategoapi/route"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	color.Green("Server starting...")

	r := gin.Default()
	r.Use(CORS)

	resource, err := db.CreateResource()
	if err != nil {
		color.Red("Connection database failure, Please check connection")
		color.Cyan(err.Error())
		logrus.Error(err)
	}
	defer resource.Close()

	// Route(r)
	route.NewRoute(r, resource)

	// SET ROUTE
	//	routes.NewRoutes(router, bot)

	r.Run(":27017")

}

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	} else {
		c.AbortWithStatus(http.StatusOK)
		return
	}
}

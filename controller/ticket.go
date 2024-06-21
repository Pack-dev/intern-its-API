package controller

import (
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

func CreateTicket(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var ticket struct {
			TitleTicket    string   `bson:"title_ticket" json:"title_ticket"`
			TagTicket      []string `bson:"tag_ticket" json:"tag_ticket"`
			Username       string   `bson:"username" json:"username"`
			ReportCustomer string   `bson:"report_customer" json:"report_customer"`
			Product        []string `bson:"product" json:"product"`
			TypeTicket     string   `bson:"type_ticket" json:"type_ticket"`
			Description    string   `bson:"description" json:"description"`
			PictureReport  []string `bson:"picture_report" json:"picture_report"`
		}

		if err := c.ShouldBind(&ticket); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
				"msg":     "ShouldBind Error",
			})
			return
		}

		//ดึง Type
		var TypeTicketModelS model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_name":   ticket.TypeTicket,
			"type_ticket_delete": false,
		}
		filterOption := bson.M{}

		err := repo.GetOneStatement(resource, "typeTicket", filter, filterOption, &TypeTicketModelS)
		if err != nil {
			CreateLog(resource, c, "Type Ticket Not Found", "Type Ticket Not Found")
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
				"msg":     "Type Ticket Error | type name :" + ticket.TypeTicket,
			})
			return
		}

		//ดึง Tag
		var TagticketModel []model.TagTicketModelS
		if len(ticket.TagTicket) == 0 {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "Require TagTicket",
			})
			return
		} else {
			for _, v := range ticket.TagTicket {
				var TagticketModelONE model.TagTicketModelS
				filter1 := bson.M{
					"tag_ticket_name":   v,
					"tag_ticket_delete": false,
				}
				filterOption1 := bson.M{}

				err = repo.GetOneStatement(resource, "TagTicket", filter1, filterOption1, &TagticketModelONE)
				if err != nil {
					CreateLog(resource, c, "CreateTicket fail", "Tag Ticket name : "+v+" not found | "+err.Error())
					c.JSON(400, gin.H{
						"code":    400,
						"msg":     "Product name : " + v + " not found ",
						"payload": err.Error(),
					})
					return
				}
				TagticketModel = append(TagticketModel, TagticketModelONE)
			}
		}

		var ReportCustomer primitive.ObjectID
		var CustomerModel model.CustomerTicketModel

		if ticket.ReportCustomer != "" {
			ReportCustomer, err = primitive.ObjectIDFromHex(ticket.ReportCustomer)
			if err != nil {
				c.JSON(500, gin.H{
					"code":    500,
					"text":    "ObjectIDFromHex Fail",
					"payload": err.Error(),
				})
				return
			}

			filter = bson.M{
				"_id":             ReportCustomer,
				"customer_delete": false,
			}
			filterOption = bson.M{}

			err = repo.GetOneStatement(resource, "Customer", filter, filterOption, &CustomerModel)
			if err != nil {
				c.JSON(400, gin.H{
					"code":    400,
					"payload": err.Error(),
					"msg":     "Customer Error",
				})
				return
			}
		}
		//ดึง Customer
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                      value,
			"user_delete":                   false,
			"role.permission.ticket.create": true,
			"active":                        true,
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
		//ดึง User

		var UserModel model.UserTicketModel
		filter = bson.M{
			"username":    ticket.Username,
			"user_delete": false,
		}
		filterOption = bson.M{}

		err = repo.GetOneStatement(resource, "user", filter, filterOption, &UserModel)
		if err != nil {

			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
				"msg":     "Username Error",
			})
			return
		}
		var ProductModel []model.ProductModelS
		if len(ticket.Product) == 0 {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "Require Product",
			})
			return
		} else {
			for _, v := range ticket.Product {
				var ProductModelONE model.ProductModelS
				filter3 := bson.M{
					"product_name":   v,
					"product_delete": false,
				}
				filterOption3 := bson.M{}

				err = repo.GetOneStatement(resource, "Product", filter3, filterOption3, &ProductModelONE)
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
		}

		var passto model.PassToModel
		passto.Role = "Customer Service"

		// แก้type and tag
		TicketData := model.TicketModel{
			TitleTicket:    ticket.TitleTicket,
			TagTicket:      TagticketModel,
			Username:       UserModel,
			Product:        ProductModel,
			ReportCustomer: CustomerModel,
			TypeTicket:     TypeTicketModelS,
			Description:    ticket.Description,
			PictureReport:  ticket.PictureReport, // เพิ่ม Product ของ User
			CreateTime:     time.Now(),
			DeleteTime:     time.Time{},
			UpdateTime:     time.Time{},
			SuccessTime:    time.Time{},
			TicketStatus:   "Pending",
			PassTo:         passto,
			TicketCancel:   false,
			TicketDelete:   false,
		}

		result, err := repo.CreateStatement(resource, "ticket", TicketData)
		if err != nil {
			CreateLog(resource, c, "Create Ticket Failed", "Create Ticket Failed | err : "+err.Error())
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "Create Ticket Success", "Create Ticket name : "+ticket.TitleTicket+" Success ")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}

func GetAllTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.all_ticket.read": true,
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel": false,
			"ticket_delete": false,
			"ticket_status": bson.M{"$in": bson.A{"Pending", "In fix", "Success"}},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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

// ส่งเคสให้โปรแกรมเมอร์
func PassToProgramer(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id       string `bson:"_id" json:"_id" form:"_id"`
		Username string `bson:"username" json:"username"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}
		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "PassToProgramer fail", "ID not from Hex  | error : "+err.Error())
			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		//Row role
		var PassToModel model.PassToModel
		filter1 := bson.M{
			"username":       body.Username,
			"user_delete":    false,
			"role.role_name": "Programer",
			"active":         true,
		}
		filterOption1 := bson.M{}

		err = repo.GetOneStatement(resource, "user", filter1, filterOption1, &PassToModel.User)
		if err != nil {
			CreateLog(resource, c, "Not found user in collection", "Not found user in collection | err : "+err.Error())
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "get User error not found user",
				"payload": err.Error(),
			})
			return
		}

		filter := bson.M{
			"_id":               Id,
			"ticket_cancel":     false,
			"ticket_delete":     false,
			"ticket_status":     "Pending",
			"pass_to.role_name": "Customer Service",
		}
		PassToModel.Role = "Programer"
		update := bson.M{"$set": bson.M{
			"update_time":   time.Now(),
			"pass_to":       PassToModel,
			"ticket_status": "In fix",
		}}
		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
		if err != nil {
			CreateLog(resource, c, "Send work to Programer fail", "Send work to Programer fail | err : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "Send work to Programer success ", "Send work to : "+body.Username+" Programer success ")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
			"user":    PassToModel,
		})
	}
}

func GetAllTicketProgramer(resource *db.Resource) func(c *gin.Context) {
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":         false,
			"ticket_delete":         false,
			"ticket_status":         bson.M{"$in": bson.A{"In fix"}},
			"pass_to.user.username": body.Username,
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
		filterOption := bson.M{"create_time": -1}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, filterOption, limit, page, &entity)
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
			return entity[i].TitleTicket < entity[j].TitleTicket
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

func GetCancelTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.all_ticket.read": true,
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel": true,
			"ticket_delete": false,
			"ticket_status": "Canceled",
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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

func GetInfixTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.all_ticket.read": true,
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel": false,
			"ticket_delete": false,
			"ticket_status": "In fix",
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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

func GetDeleteTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.all_ticket.read": true,
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_delete": true,
			"ticket_cancel": false,
			"ticket_status": "Delete",
		}
		if err := repo.GetManyStatement(resource, "ticket", filter, nil, &entity); err != nil {
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
func GetPendingTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.all_ticket.read": true,
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_delete": false,
			"ticket_cancel": false,
			"ticket_status": "Pending",
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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

func GetSuccessTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                        value,
			"user_delete":                     false,
			"role.permission.all_ticket.read": true,
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel": false,
			"ticket_delete": false,
			"ticket_status": "Success",
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func SuccessOneTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id              string   `bson:"_id" json:"_id"`
		Solution        string   `bson:"solution" json:"solution"`
		PictureSolution []string `bson:"picture_solution" json:"picture_solution"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}
		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "SuccessOneTicket fail", "ID not from Hex  | error : "+err.Error())
			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{"_id": Id, "ticket_status": "In fix"}
		update := bson.M{"$set": bson.M{
			"ticket_status":    "Success",
			"success_time":     time.Now(),
			"solution":         body.Solution,
			"picture_solution": body.PictureSolution,
		}}
		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
		if err != nil {
			CreateLog(resource, c, "SuccessOneTicket success ", "Send solve problem failed ")
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "SuccessOneTicket success ", "Send solve problem successfuly ")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}

func CsUpdateTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id            string   `bson:"_id" json:"_id" form:"_id"`
		TitleTicket   string   `bson:"title_ticket" json:"title_ticket"`
		TagTicket     []string `bson:"tag_ticket" json:"tag_ticket"`
		Product       []string `bson:"product" json:"product"`
		TypeTicket    string   `bson:"type_ticket" json:"type_ticket"`
		Description   string   `bson:"description" json:"description"`
		PictureReport []string `bson:"picture_report" json:"picture_report"`
		PassTo        string   `bson:"pass_to" json:"pass_to"`
		Username      string   `bson:"username" json:"username"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}

		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "CsUpdateTicket fail", "ID not from Hex  | error : "+err.Error())
			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return

		}
		var TypeTicketModelS model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_name":   body.TypeTicket,
			"type_ticket_delete": false,
		}
		filterOption := bson.M{}

		err = repo.GetOneStatement(resource, "typeTicket", filter, filterOption, &TypeTicketModelS)
		if err != nil {
			CreateLog(resource, c, "Type Ticket Not Found", "Type Ticket Not Found")
			c.JSON(400, gin.H{
				"code":    400,
				"payload": filter,
			})
			return
		}
		//ดึง Tag
		var TagticketModel []model.TagTicketModelS
		for _, v := range body.TagTicket {
			var TagticketModelONE model.TagTicketModelS
			filter1 := bson.M{
				"tag_ticket_name":   v,
				"tag_ticket_delete": false,
			}
			filterOption1 := bson.M{}

			err = repo.GetOneStatement(resource, "TagTicket", filter1, filterOption1, &TagticketModelONE)
			if err != nil {
				CreateLog(resource, c, "UserUpdateTicket fail", "Tag ticket name : "+v+" not found | "+err.Error())
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "Tag ticket name : " + v + " not found ",
					"payload": err.Error(),
				})
				return
			}
			TagticketModel = append(TagticketModel, TagticketModelONE)
		}
		var ProductModel []model.ProductModelS
		for _, v := range body.Product {
			var ProductModelONE model.ProductModelS
			filter3 := bson.M{
				"product_name":   v,
				"product_delete": false,
			}
			filterOption3 := bson.M{}

			err = repo.GetOneStatement(resource, "Product", filter3, filterOption3, &ProductModelONE)
			if err != nil {
				CreateLog(resource, c, "UserUpdateTicket fail", "Product name : "+v+" not found | "+err.Error())
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "Product name : " + v + " not found ",
					"payload": err.Error(),
				})
				return
			}
			ProductModel = append(ProductModel, ProductModelONE)
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                      value,
			"user_delete":                   false,
			"role.permission.ticket.update": true,
			"active":                        true,
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
		var PassToModel model.PassToModel
		filter1 = bson.M{
			"username":       body.Username,
			"user_delete":    false,
			"role.role_name": "Programer",
			"active":         true,
		}
		filterOption1 = bson.M{}

		err = repo.GetOneStatement(resource, "user", filter1, filterOption1, &PassToModel.User)
		if err != nil {
			CreateLog(resource, c, "CsUpdateTicket fail", "ID user not from Hex "+body.Username+" | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "get User error not found user " + body.Username + " ",
				"payload": err.Error(),
			})
			return
		}
		PassToModel.Role = "Programer"
		filter = bson.M{
			"_id":           Id,
			"ticket_cancel": false,
			"ticket_delete": false}
		update := bson.M{"$set": bson.M{
			"title_ticket":                 body.TitleTicket,
			"tag_ticket":                   TagticketModel,
			"product":                      ProductModel,
			"type_ticket.type_ticket_name": body.TypeTicket,
			"description":                  body.Description,
			"ticket_status":                "In fix",
			"picture_report":               body.PictureReport,
			"update_time":                  time.Now(),
			"pass_to":                      PassToModel,
		}}
		var beforeUpdate model.TicketModel
		_ = repo.GetOneStatement(resource, "ticket", filter, nil, &beforeUpdate)
		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
		if err != nil {
			CreateLog(resource, c, "CSUpdateTicketCS fail", "Customer Service Update Ticket fail | err : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}
		var afterUpdate model.TicketModel
		_ = repo.GetOneStatement(resource, "ticket", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UserUpdateTicketCS Success", "UserUpdateTicketCS name : "+body.TitleTicket+" Success ", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}
func GetOneTicket(resource *db.Resource) func(c *gin.Context) {
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
			CreateLog(resource, c, "Get one Ticket Fail", " ID not from Hex | err :	"+err.Error())
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		// test login
		var TicketModel model.TicketModelID
		filter := bson.M{
			"_id":           ID,
			"ticket_delete": false,
		}
		filterOption := bson.M{}

		err = repo.GetOneStatement(resource, "ticket", filter, filterOption, &TicketModel)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": filter,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"payload": TicketModel,
		})
	}
}

func GetAllTicketByAccount(resource *db.Resource) func(c *gin.Context) {
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":     false,
			"ticket_delete":     false,
			"ticket_status":     bson.M{"$in": bson.A{"Pending", "In fix", "Success"}},
			"username.username": primitive.Regex{Pattern: body.Username, Options: ""},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func GetCancelTicketByAccount(resource *db.Resource) func(c *gin.Context) {
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":     true,
			"ticket_delete":     false,
			"ticket_status":     bson.M{"$in": bson.A{"Canceled"}},
			"username.username": primitive.Regex{Pattern: body.Username, Options: ""},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func GetPendingTicketByAccount(resource *db.Resource) func(c *gin.Context) {
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":     false,
			"ticket_delete":     false,
			"ticket_status":     bson.M{"$in": bson.A{"Pending"}},
			"username.username": primitive.Regex{Pattern: body.Username, Options: ""},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func GetInfixTicketByAccount(resource *db.Resource) func(c *gin.Context) {
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":     false,
			"ticket_delete":     false,
			"ticket_status":     bson.M{"$in": bson.A{"In fix"}},
			"username.username": primitive.Regex{Pattern: body.Username, Options: ""},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func GetSuccessTicketByAccount(resource *db.Resource) func(c *gin.Context) {
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
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":     false,
			"ticket_delete":     false,
			"ticket_status":     bson.M{"$in": bson.A{"Success"}},
			"username.username": primitive.Regex{Pattern: body.Username, Options: ""},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func GetPendingTicketCS(resource *db.Resource) func(c *gin.Context) {
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
		var UserModel model.PassToModel
		filter1 := bson.M{
			"username":       body.Username,
			"user_delete":    false,
			"role.role_name": "Customer Service",
			"active":         true,
		}
		filterOption1 := bson.M{}

		err := repo.GetOneStatement(resource, "user", filter1, filterOption1, &UserModel.User)
		if err != nil {
			CreateLog(resource, c, "Not found user in collection", "Not found user in collection | err : "+err.Error())
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "get User error not found user",
				"payload": err.Error(),
			})
			return
		}
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel":     false,
			"ticket_delete":     false,
			"ticket_status":     bson.M{"$in": bson.A{"Pending"}},
			"pass_to.role_name": "Customer Service",
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 100
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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
func UserUpdateTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id            string   `bson:"_id" json:"_id" form:"_id"`
		TitleTicket   string   `bson:"title_ticket" json:"title_ticket"`
		TagTicket     []string `bson:"tag_ticket" json:"tag_ticket"`
		Product       []string `bson:"product" json:"product"`
		TypeTicket    string   `bson:"type_ticket" json:"type_ticket"`
		Description   string   `bson:"description" json:"description"`
		PictureReport []string `bson:"picture_report" json:"picture_report"`
		TicketStatus  string   `bson:"ticket_status" json:"ticket_status"`
		PassTo        string   `bson:"pass_to" json:"pass_to"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}

		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "UserUpdateTicket fail", "ID not from Hex  | error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return

		}
		var TypeTicketModelS model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_name":   body.TypeTicket,
			"type_ticket_delete": false,
		}
		filterOption := bson.M{}

		err = repo.GetOneStatement(resource, "typeTicket", filter, filterOption, &TypeTicketModelS)
		if err != nil {
			CreateLog(resource, c, "Type Ticket Not Found", "Type Ticket Not Found")
			c.JSON(400, gin.H{
				"code":    400,
				"payload": filter,
			})
			return
		}
		//ดึง Tag
		var TagticketModel []model.TagTicketModelS
		for _, v := range body.TagTicket {
			var TagticketModelONE model.TagTicketModelS
			filter1 := bson.M{
				"tag_ticket_name":   v,
				"tag_ticket_delete": false,
			}
			filterOption1 := bson.M{}

			err = repo.GetOneStatement(resource, "TagTicket", filter1, filterOption1, &TagticketModelONE)
			if err != nil {
				CreateLog(resource, c, "UserUpdateTicket fail", "Tag ticket name : "+v+" not found | "+err.Error())
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "Tag ticket name : " + v + " not found ",
					"payload": err.Error(),
				})
				return
			}
			TagticketModel = append(TagticketModel, TagticketModelONE)
		}
		var ProductModel []model.ProductModelS
		for _, v := range body.Product {
			var ProductModelONE model.ProductModelS
			filter3 := bson.M{
				"product_name":   v,
				"product_delete": false,
			}
			filterOption3 := bson.M{}

			err = repo.GetOneStatement(resource, "Product", filter3, filterOption3, &ProductModelONE)
			if err != nil {
				CreateLog(resource, c, "UserUpdateTicket fail", "Product name : "+v+" not found | "+err.Error())
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "Product name : " + v + " not found ",
					"payload": err.Error(),
				})
				return
			}
			ProductModel = append(ProductModel, ProductModelONE)
		}
		var PassToModel model.PassToModel
		PassToModel.Role = "Customer Service"
		filter = bson.M{
			"_id":           Id,
			"ticket_cancel": false,
			"ticket_delete": false}
		update := bson.M{"$set": bson.M{
			"title_ticket":                 body.TitleTicket,
			"tag_ticket":                   TagticketModel,
			"type_ticket.type_ticket_name": body.TypeTicket,
			"description":                  body.Description,
			"ticket_status":                "Pending",
			"picture_report":               body.PictureReport,
			"product":                      ProductModel,
			"update_time":                  time.Now(),
			"pass_to":                      PassToModel,
		}}

		var beforeUpdate model.TicketModel
		err = repo.GetOneStatement(resource, "ticket", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
		if err != nil {
			CreateLog(resource, c, "UserUpdateTicket fail", "UserUpdateTicket fail ")

			c.JSON(400, gin.H{
				"code":    400,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.TicketModel
		err = repo.GetOneStatement(resource, "ticket", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UserUpdateTicket Success", "UserUpdateTicket name : "+body.TitleTicket+" Success ", beforeUpdate, afterUpdate)

		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}
func DeleteOneTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id string `bson:"_id" json:"_id"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                      value,
			"user_delete":                   false,
			"role.permission.ticket.delete": true,
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
		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "DeleteOneTicket fail", "ID not from Hex  | error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{"_id": Id}
		update := bson.M{"$set": bson.M{
			"ticket_status": "Delete",
			"ticket_delete": true,
			"delete_time":   time.Now(),
		}}
		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneTicket fail", "DeleteOneTicket fail ")
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneTicket Success", "DeleteOneTicket Success ")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}
func CancelOneTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id string `bson:"_id" json:"_id"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}
		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "CancelOneTicket fail", "ID not from Hex  | error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{"_id": Id}
		update := bson.M{"$set": bson.M{
			"ticket_status": "Canceled",
			"ticket_cancel": true,
			"cancel_time":   time.Now(),
		}}
		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
		if err != nil {
			CreateLog(resource, c, "CancelOneTicket fail", "CancelOneTicket fail | error : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "CancelOneTicket Success", "CancelOneTicket Success")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}
func SearchTicket(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		// Id            string `bson:"_id" json:"_id"`
		TitleTicket string `bson:"title_ticket" json:"title_ticket"`
		TagTicket   string `bson:"tag_ticket" json:"tag_ticket"`
		TypeTicket  string `bson:"type_ticket" json:"type_ticket"`
		// OrderBy     string //เก็บไว้ตั้งค่าว่าจะเรียงตามอะไร
	}
	return func(c *gin.Context) {

		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"text": "ShouldBind Fail",
				"erro": err.Error(),
			})
			return
		}

		var entity []model.TicketModel
		filter := bson.M{
			"ticket_cancel":                false,
			"ticket_delete":                false,
			"ticket_status":                bson.M{"$in": bson.A{"Pending", "In fix", "Success"}},
			"title_ticket":                 primitive.Regex{Pattern: body.TitleTicket, Options: ""},
			"tag_ticket.tag_ticket_name":   primitive.Regex{Pattern: body.TagTicket, Options: ""},
			"type_ticket.type_ticket_name": primitive.Regex{Pattern: body.TypeTicket, Options: ""},
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleTicket < entity[j].TitleTicket
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

package controller

// func UserUpdateTicket(resource *db.Resource) func(c *gin.Context) {
// 	type Body struct {
// 		Id            string   `bson:"_id" json:"_id" form:"_id"`
// 		TitleTicket   string   `bson:"title_ticket" json:"title_ticket"`
// 		TagTicket     []string `bson:"tag_ticket" json:"tag_ticket"`
// 		Product       []string `bson:"product" json:"product"`
// 		TypeTicket    string   `bson:"type_ticket" json:"type_ticket"`
// 		Description   string   `bson:"description" json:"description"`
// 		PictureReport []string `bson:"picture_report" json:"picture_report"`
// 		TicketStatus  string   `bson:"ticket_status" json:"ticket_status"`
// 		PassTo        string   `bson:"pass_to" json:"pass_to"`
// 	}

// 	return func(c *gin.Context) {
// 		var body Body
// 		if err := c.ShouldBind(&body); err != nil {
// 			c.JSON(400, gin.H{
// 				"code": 400,
// 				"text": "ShouldBind Fail",
// 				"erro": err.Error(),
// 			})
// 			return
// 		}

// 		Id, err := primitive.ObjectIDFromHex(body.Id)
// 		if err != nil {
// 			CreateLog(resource, c, "UserUpdateTicket fail", "ID not from Hex  | error : "+err.Error())

// 			c.JSON(500, gin.H{
// 				"code":    500,
// 				"text":    "ObjectIDFromHex Fail",
// 				"payload": err.Error(),
// 			})
// 			return

// 		}
// 		var TypeTicketModelS model.TypeTicketModelS
// 		filter := bson.M{
// 			"type_ticket_name":   body.TypeTicket,
// 			"type_ticket_delete": false,
// 		}
// 		filterOption := bson.M{}

// 		err = repo.GetOneStatement(resource, "typeTicket", filter, filterOption, &TypeTicketModelS)
// 		if err != nil {
// 			CreateLog(resource, c, "Type Ticket Not Found", "Type Ticket Not Found")
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"payload": filter,
// 			})
// 			return
// 		}
// 		//ดึง Tag
// 		var TagticketModel []model.TagTicketModelS
// 		for _, v := range body.TagTicket {
// 			var TagticketModelONE model.TagTicketModelS
// 			filter1 := bson.M{
// 				"tag_ticket_name":   v,
// 				"tag_ticket_delete": false,
// 			}
// 			filterOption1 := bson.M{}

// 			err = repo.GetOneStatement(resource, "TagTicket", filter1, filterOption1, &TagticketModelONE)
// 			if err != nil {
// 				CreateLog(resource, c, "UserUpdateTicket fail", "Tag ticket name : "+v+" not found | "+err.Error())
// 				c.JSON(400, gin.H{
// 					"code":    400,
// 					"msg":     "Tag ticket name : " + v + " not found ",
// 					"payload": err.Error(),
// 				})
// 				return
// 			}
// 			TagticketModel = append(TagticketModel, TagticketModelONE)
// 		}
// 		var ProductModel []model.ProductModelS
// 		for _, v := range body.Product {
// 			var ProductModelONE model.ProductModelS
// 			filter3 := bson.M{
// 				"product_name":   v,
// 				"product_delete": false,
// 			}
// 			filterOption3 := bson.M{}

// 			err = repo.GetOneStatement(resource, "Product", filter3, filterOption3, &ProductModelONE)
// 			if err != nil {
// 				CreateLog(resource, c, "UserUpdateTicket fail", "Product name : "+v+" not found | "+err.Error())
// 				c.JSON(400, gin.H{
// 					"code":    400,
// 					"msg":     "Product name : " + v + " not found ",
// 					"payload": err.Error(),
// 				})
// 				return
// 			}
// 			ProductModel = append(ProductModel, ProductModelONE)
// 		}
// 		var PassToModel model.PassToModel
// 		PassToModel.Role = "Customer Service"
// 		filter = bson.M{
// 			"_id":           Id,
// 			"ticket_cancel": false,
// 			"ticket_delete": false}
// 		update := bson.M{"$set": bson.M{
// 			"title_ticket":                 body.TitleTicket,
// 			"tag_ticket":                   TagticketModel,
// 			"type_ticket.type_ticket_name": body.TypeTicket,
// 			"description":                  body.Description,
// 			"ticket_status":                "Pending",
// 			"picture_report":               body.PictureReport,
// 			"product":                      ProductModel,
// 			"update_time":                  time.Now(),
// 			"pass_to":                      PassToModel,
// 		}}

// 		var beforeUpdate model.TicketModel
// 		err = repo.GetOneStatement(resource, "ticket", filter, nil, &beforeUpdate)

// 		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
// 		if err != nil {
// 			CreateLog(resource, c, "UserUpdateTicket fail", "UserUpdateTicket fail ")

// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "Update Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}

// 		var afterUpdate model.TicketModel
// 		err = repo.GetOneStatement(resource, "ticket", filter, nil, &afterUpdate)

// 		CreateLogForUpdate(resource, c, "UserUpdateTicket Success", "UserUpdateTicket name : "+body.TitleTicket+" Success ", beforeUpdate, afterUpdate)

// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": Result,
// 		})
// 	}
// }
// func DeleteOneTicket(resource *db.Resource) func(c *gin.Context) {
// 	type Body struct {
// 		Id string `bson:"_id" json:"_id"`
// 	}

// 	return func(c *gin.Context) {
// 		var body Body
// 		if err := c.ShouldBind(&body); err != nil {
// 			c.JSON(400, gin.H{
// 				"code": 400,
// 				"text": "ShouldBind Fail",
// 				"erro": err.Error(),
// 			})
// 			return
// 		}
// 		var permission model.UserModelS
// 		value := c.GetString("username")
// 		filter1 := bson.M{
// 			"username":                      value,
// 			"user_delete":                   false,
// 			"role.permission.ticket.delete": true,
// 			"active":                        true,
// 		}
// 		filterOption1 := bson.M{}

// 		err := repo.GetOneStatement(resource, "user", filter1, filterOption1, &permission)
// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "No permission || permission Needed",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		Id, err := primitive.ObjectIDFromHex(body.Id)
// 		if err != nil {
// 			CreateLog(resource, c, "DeleteOneTicket fail", "ID not from Hex  | error : "+err.Error())

// 			c.JSON(500, gin.H{
// 				"code":    500,
// 				"text":    "ObjectIDFromHex Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		filter := bson.M{"_id": Id}
// 		update := bson.M{"$set": bson.M{
// 			"ticket_status": "Delete",
// 			"ticket_delete": true,
// 			"delete_time":   time.Now(),
// 		}}
// 		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
// 		if err != nil {
// 			CreateLog(resource, c, "DeleteOneTicket fail", "DeleteOneTicket fail ")
// 			c.JSON(400, gin.H{
// 				"code":    401,
// 				"text":    "Update Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		CreateLog(resource, c, "DeleteOneTicket Success", "DeleteOneTicket Success ")
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": Result,
// 		})
// 	}
// }
// func CancelOneTicket(resource *db.Resource) func(c *gin.Context) {
// 	type Body struct {
// 		Id string `bson:"_id" json:"_id"`
// 	}

// 	return func(c *gin.Context) {
// 		var body Body
// 		if err := c.ShouldBind(&body); err != nil {
// 			c.JSON(400, gin.H{
// 				"code": 400,
// 				"text": "ShouldBind Fail",
// 				"erro": err.Error(),
// 			})
// 			return
// 		}
// 		Id, err := primitive.ObjectIDFromHex(body.Id)
// 		if err != nil {
// 			CreateLog(resource, c, "CancelOneTicket fail", "ID not from Hex  | error : "+err.Error())

// 			c.JSON(500, gin.H{
// 				"code":    500,
// 				"text":    "ObjectIDFromHex Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		filter := bson.M{"_id": Id}
// 		update := bson.M{"$set": bson.M{
// 			"ticket_status": "Canceled",
// 			"ticket_cancel": true,
// 			"cancel_time":   time.Now(),
// 		}}
// 		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
// 		if err != nil {
// 			CreateLog(resource, c, "CancelOneTicket fail", "CancelOneTicket fail | error : "+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    401,
// 				"text":    "Update Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		CreateLog(resource, c, "CancelOneTicket Success", "CancelOneTicket Success")
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": Result,
// 		})
// 	}
// }

// // func ChangeTicketStatus(resource *db.Resource) func(c *gin.Context) {
// // 	type Body struct {
// // 		Id           string `bson:"_id" json:"_id"`
// // 		TicketStatus string `bson:"ticket_status" json:"ticket_status"`
// // 	}

// // 	return func(c *gin.Context) {
// // 		var body Body
// // 		if err := c.ShouldBind(&body); err != nil {
// // 			c.JSON(400, gin.H{
// // 				"code": 400,
// // 				"text": "ShouldBind Fail",
// // 				"erro": err.Error(),
// // 			})
// // 			return
// // 		}
// // 		Id, err := primitive.ObjectIDFromHex(body.Id)
// // 		if err != nil {
// // 			c.JSON(500, gin.H{
// // 				"code":    500,
// // 				"text":    "ObjectIDFromHex Fail",
// // 				"payload": err.Error(),
// // 			})
// // 			return
// // 		}
// // 		var TS string
// // 		if body.TicketStatus == "1" {
// // 			TS = "Pending"
// // 		} else if body.TicketStatus == "2" {
// // 			TS = "In Fix"
// // 		} else if body.TicketStatus == "3" {
// // 			TS = "Success"
// // 		} else if body.TicketStatus == "4" {
// // 			TS = "Cancel"
// // 		} else {
// // 			c.JSON(400, gin.H{
// // 				"code":    400,
// // 				"text":    "Ticket Status Fail (1 : Pending,2 : In Fix,3 : Success,4 : Cancel)",
// // 				"payload": body.TicketStatus,
// // 			})
// // 			return
// // 		}
// // 		filter := bson.M{
// // 			"_id":           Id,
// // 			"ticket_cancel": false,
// // 			"ticket_delete": false}
// // 		update := bson.M{"$set": bson.M{
// // 			"update_time":   time.Now(),
// // 			"ticket_status": TS,
// // 		}}
// // 		Result, err := repo.UpdateOneStatement(resource, "ticket", filter, update)
// // 		if err != nil {
// // 			c.JSON(400, gin.H{
// // 				"code":    401,
// // 				"text":    "Update Fail",
// // 				"payload": err.Error(),
// // 			})
// // 			return
// // 		}

// //			c.JSON(200, gin.H{
// //				"code":    200,
// //				"payload": Result,
// //			})
// //		}
// //	}
// func SearchTicket(resource *db.Resource) func(c *gin.Context) {
// 	type Body struct {
// 		// Id            string `bson:"_id" json:"_id"`
// 		TitleTicket string `bson:"title_ticket" json:"title_ticket"`
// 		TagTicket   string `bson:"tag_ticket" json:"tag_ticket"`
// 		TypeTicket  string `bson:"type_ticket" json:"type_ticket"`
// 		// OrderBy     string //เก็บไว้ตั้งค่าว่าจะเรียงตามอะไร
// 	}
// 	return func(c *gin.Context) {

// 		var body Body
// 		if err := c.ShouldBind(&body); err != nil {
// 			c.JSON(400, gin.H{
// 				"code": 400,
// 				"text": "ShouldBind Fail",
// 				"erro": err.Error(),
// 			})
// 			return
// 		}

// 		var entity []model.TicketModel
// 		filter := bson.M{
// 			"ticket_cancel":                false,
// 			"ticket_delete":                false,
// 			"ticket_status":                bson.M{"$in": bson.A{"Pending", "In fix", "Success"}},
// 			"title_ticket":                 primitive.Regex{Pattern: body.TitleTicket, Options: ""},
// 			"tag_ticket.tag_ticket_name":   primitive.Regex{Pattern: body.TagTicket, Options: ""},
// 			"type_ticket.type_ticket_name": primitive.Regex{Pattern: body.TypeTicket, Options: ""},
// 		}
// 		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
// 		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
// 		if limit == 0 {
// 			limit = 20
// 		}
// 		if page == 0 {
// 			page = 1
// 		}
// 		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "ticket", filter, nil, limit, page, &entity)
// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		sort.SliceStable(entity, func(i, j int) bool {
// 			return entity[i].TitleTicket < entity[j].TitleTicket
// 		})
// 		c.JSON(200, gin.H{
// 			"code":           200,
// 			"total_page":     totalPages,
// 			"total_document": count,
// 			"current_paget":  page,
// 			"limit_data":     limit,
// 			"payload":        entity,
// 		})
// 	}
// }

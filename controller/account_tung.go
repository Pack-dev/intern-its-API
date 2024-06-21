package controller

//	func ShowTicket(resource *db.Resource) func(c *gin.Context) {
//		return func(c *gin.Context) {
//			var entity []model.TicketModel
//			filter := bson.M{
//				"TicketStatus": true,
//			}
//			if err := repo.GetManyStatement(resource, "ticket", filter, nil, &entity); err != nil {
//				c.JSON(200, gin.H{
//					"code":    400,
//					"payload": err.Error(),
//				})
//				return
//			}
//			c.JSON(200, gin.H{
//				"code":    200,
//				"payload": entity,
//			})
//		}
//	}
// func DeleteOneUser(resource *db.Resource) func(c *gin.Context) {
// 	type Body struct {
// 		Username string `bson:"username" json:"username" form:"username"`
// 	}
// 	return func(c *gin.Context) {
// 		var body Body
// 		if err := c.ShouldBind(&body); err != nil {
// 			c.JSON(400, gin.H{
// 				"code": 400,
// 				"erro": err.Error(),
// 			})
// 			return
// 		}
// 		var permission model.UserModelS
// 		value := c.GetString("username")
// 		filter1 := bson.M{
// 			"username":                    value,
// 			"user_delete":                 false,
// 			"role.permission.user.update": true,
// 			"active":                      true,
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
// 		fmt.Print("username : " + body.Username)
// 		filter := bson.M{
// 			"username":    body.Username,
// 			"user_delete": false}
// 		update := bson.M{"$set": bson.M{
// 			"user_delete": true,
// 			"delete_time": time.Now()}}

// 		Result, err := repo.UpdateOneStatement(resource, "user", filter, update)
// 		if err != nil {
// 			CreateLog(resource, c, "DeleteOneUser Fail", " Delete User name : "+body.Username+" Fail | err :	"+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    401,
// 				"text":    "DeleteOneUser Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		CreateLog(resource, c, "DeleteOneUser Success", " Delete user name : "+body.Username+" Success")
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": Result,
// 		})

// 	}
// } //เพิ่ม Update ของ Getalluser
// func UpdateOneUser(resource *db.Resource) func(c *gin.Context) {

// 	type RoleModel struct {
// 		Rolename   string                `bson:"role_name" json:"role_name"`
// 		Permission model.PermissionModel `bson:"permission" json:"permission"`
// 	}
// 	type ProductModelS struct {
// 		ProductName string `bson:"product_name" json:"product_name"`
// 	}
// 	type Body struct {
// 		ID          string   `bson:"_id" json:"_id"`
// 		Prefix      string   `bson:"prefix" json:"prefix"`
// 		Name        string   `bson:"name" json:"name"`
// 		LastName    string   `bson:"last_name" json:"last_name"`
// 		FullName    string   `bson:"full_name" json:"full_name"`
// 		PhoneNumber string   `bson:"phone_number" json:"phone_number"`
// 		ImgUrl      string   `bson:"img_url" json:"img_url"`
// 		Role        string   `bson:"role" json:"role"`
// 		Customer    string   `bson:"customer" json:"customer"`
// 		Product     []string `bson:"product" json:"product"`
// 		Active      bool     `bson:"active" json:"active"`
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
// 		fmt.Println("ID : " + body.ID)
// 		ID, err := primitive.ObjectIDFromHex(body.ID)
// 		if err != nil {
// 			CreateLog(resource, c, "UpdateOneCustomer Fail", " ID not from Hex | err :	"+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "ObjectIDFromHex Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		var permission model.UserModelS
// 		value := c.GetString("username")
// 		filter1 := bson.M{
// 			"username":                    value,
// 			"user_delete":                 false,
// 			"role.permission.user.update": true,
// 			"active":                      true,
// 		}
// 		filterOption1 := bson.M{}

// 		err = repo.GetOneStatement(resource, "user", filter1, filterOption1, &permission)
// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "No permission || permission Needed",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		//Row role
// 		var RoleModel model.RoleModeli
// 		filter := bson.M{
// 			"role_name":   body.Role,
// 			"role_delete": false,
// 		}
// 		filterOption := bson.M{}

// 		err = repo.GetOneStatement(resource, "role", filter, filterOption, &RoleModel)
// 		if err != nil {
// 			CreateLog(resource, c, "UpdateOneCustomer fail", "Role name : "+body.Role+" not found | "+err.Error())

// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"msg":     "Role not found",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}

// 		var ProductModel []model.ProductModelS
// 		var ProductModelONE model.ProductModelS
// 		if len(body.Product) > 0 && body.Product[0] != "" {
// 			for _, v := range body.Product {
// 				filter1 := bson.M{
// 					"product_name":   v,
// 					"product_delete": false,
// 				}
// 				filterOption1 := bson.M{}

// 				err = repo.GetOneStatement(resource, "Product", filter1, filterOption1, &ProductModelONE)
// 				if err != nil {
// 					CreateLog(resource, c, "Update Account fail", "Product name : "+v+" not found | "+err.Error())
// 					c.JSON(400, gin.H{
// 						"code":    400,
// 						"msg":     "Product name : " + v + " not found ",
// 						"payload": err.Error(),
// 					})
// 					return
// 				}
// 				ProductModel = append(ProductModel, ProductModelONE)
// 			}
// 		}

// 		filter = bson.M{"_id": ID, "user_delete": false}
// 		update := bson.M{"$set": bson.M{
// 			"prefix":       body.Prefix,
// 			"name":         body.Name,
// 			"last_name":    body.LastName,
// 			"full_name":    body.Name + " " + body.LastName,
// 			"phone_number": body.PhoneNumber,
// 			"img_url":      body.ImgUrl,
// 			"role":         RoleModel,
// 			"update_time":  time.Now(),
// 			"customer":     body.Customer,
// 			"product":      ProductModel,
// 			"active":       body.Active,
// 		}}

// 		var beforeUpdate model.UserModelS
// 		err = repo.GetOneStatement(resource, "user", filter, nil, &beforeUpdate)

// 		Result, err := repo.UpdateOneStatement(resource, "user", filter, update)
// 		if err != nil {
// 			CreateLog(resource, c, "UpdateOneUser Fail", " Update User name : "+body.FullName+" Fail | err :	"+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "Update Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}

// 		var afterUpdate model.UserModelS
// 		err = repo.GetOneStatement(resource, "user", filter, nil, &afterUpdate)

// 		CreateLogForUpdate(resource, c, "UpdateOneUser Success", " Update User name : "+body.FullName+" Success", beforeUpdate, afterUpdate)
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": Result,
// 		})
// 	}
// }

// // ของ Profile
// func UpdateProfile(resource *db.Resource) func(c *gin.Context) {

// 	type RoleModel struct {
// 		Rolename   string                `bson:"role_name" json:"role_name"`
// 		Permission model.PermissionModel `bson:"permission" json:"permission"`
// 	}
// 	type ProductModelS struct {
// 		ProductName string `bson:"product_name" json:"product_name"`
// 	}
// 	type Body struct {
// 		ID          string          `bson:"_id" json:"_id"`
// 		Prefix      string          `bson:"prefix" json:"prefix"`
// 		Name        string          `bson:"name" json:"name"`
// 		LastName    string          `bson:"last_name" json:"last_name"`
// 		FullName    string          `bson:"full_name" json:"full_name"`
// 		PhoneNumber string          `bson:"phone_number" json:"phone_number"`
// 		ImgUrl      string          `bson:"img_url" json:"img_url"`
// 		Role        string          `bson:"role" json:"role"`
// 		Customer    string          `bson:"customer" json:"customer"`
// 		Product     []ProductModelS `bson:"product" json:"product"`
// 	}
// 	//อิอิ
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
// 		fmt.Println("ID : " + body.ID)
// 		ID, err := primitive.ObjectIDFromHex(body.ID)
// 		if err != nil {
// 			CreateLog(resource, c, "UpdateOneCustomer Fail", " ID not from Hex | err :	"+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "ObjectIDFromHex Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}

// 		//Row role
// 		var RoleModel model.RoleModeli
// 		filter := bson.M{
// 			"role_name":   body.Role,
// 			"role_delete": false,
// 		}
// 		filterOption := bson.M{}

// 		err = repo.GetOneStatement(resource, "role", filter, filterOption, &RoleModel)
// 		if err != nil {
// 			CreateLog(resource, c, "UpdateOneCustomer fail", "Role name : "+body.Role+" not found | "+err.Error())

// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"msg":     "Role not found",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}

// 		var ProductModel []model.ProductModelS
// 		for _, v := range body.Product {
// 			var ProductModelONE model.ProductModelS
// 			filter1 := bson.M{
// 				"product_name":   v.ProductName,
// 				"product_delete": false,
// 			}
// 			filterOption1 := bson.M{}

// 			err = repo.GetOneStatement(resource, "Product", filter1, filterOption1, &ProductModelONE)
// 			if err != nil {
// 				CreateLog(resource, c, "UpdateOneCustomer fail", "Product name : "+v.ProductName+" not found | "+err.Error())
// 				c.JSON(400, gin.H{
// 					"code":    400,
// 					"msg":     "Product name : " + v.ProductName + " not found ",
// 					"payload": err.Error(),
// 				})
// 				return
// 			}
// 			ProductModel = append(ProductModel, ProductModelONE)
// 		}

// 		filter = bson.M{"_id": ID, "user_delete": false}
// 		update := bson.M{"$set": bson.M{
// 			"prefix":       body.Prefix,
// 			"name":         body.Name,
// 			"last_name":    body.LastName,
// 			"full_name":    body.Name + " " + body.LastName,
// 			"phone_number": body.PhoneNumber,
// 			"img_url":      body.ImgUrl,
// 			"role":         RoleModel,
// 			"update_time":  time.Now(),
// 			"customer":     body.Customer,
// 			"product":      ProductModel,
// 		}}

// 		var beforeUpdate model.UserModelS
// 		err = repo.GetOneStatement(resource, "user", filter, nil, &beforeUpdate)

// 		Result, err := repo.UpdateOneStatement(resource, "user", filter, update)
// 		if err != nil {
// 			CreateLog(resource, c, "UpdateOneUser Fail", " Update User name : "+body.FullName+" Fail | err :	"+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "Update Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}

// 		var afterUpdate model.UserModelS
// 		err = repo.GetOneStatement(resource, "user", filter, nil, &afterUpdate)

// 		CreateLogForUpdate(resource, c, "UpdateOneUser Success", " Update User name : "+body.FullName+" Success", beforeUpdate, afterUpdate)
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": Result,
// 		})
// 	}
// }

// func GetOneUser(resource *db.Resource) func(c *gin.Context) {
// 	type Body struct {
// 		ID string `json:"_id"`
// 	}

// 	return func(c *gin.Context) {
// 		var body Body
// 		if err := c.ShouldBind(&body); err != nil {
// 			c.JSON(400, gin.H{
// 				"code": 400,
// 				"erro": err.Error(),
// 			})
// 			return
// 		}

// 		ID, err := primitive.ObjectIDFromHex(body.ID)
// 		if err != nil {
// 			CreateLog(resource, c, "Get one User Fail", " ID not from Hex | err :	"+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "ObjectIDFromHex Fail",
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		// test login
// 		var userModel model.UserModelS
// 		filter := bson.M{
// 			"_id":         ID,
// 			"user_delete": false,
// 		}
// 		filterOption := bson.M{}

// 		err = repo.GetOneStatement(resource, "user", filter, filterOption, &userModel)
// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"payload": filter,
// 			})
// 			return
// 		}
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": userModel,
// 		})
// 	}
// }

// func GetAllUser(resource *db.Resource) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var entity []model.UserModelS
// 		filter := bson.M{
// 			"user_delete": false,
// 		}
// 		var permission model.UserModelS
// 		value := c.GetString("username")
// 		filter1 := bson.M{
// 			"username":                  value,
// 			"user_delete":               false,
// 			"role.permission.user.read": true,
// 			"active":                    true,
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
// 		active, _ := strconv.ParseBool(c.GetHeader("active"))
// 		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
// 		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
// 		if limit == 0 {
// 			limit = 5
// 		}
// 		if page == 0 {
// 			page = 1
// 		}
// 		// skip := (page - 1) * limit
// 		// if skip < 0 {
// 		// 	skip = 0
// 		// }
// 		filter = bson.M{"active": active}
// 		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "user", filter, nil, limit, page, &entity)
// 		fmt.Println(page)
// 		fmt.Println(limit)
// 		fmt.Println(entity)
// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		sort.SliceStable(entity, func(i, j int) bool {
// 			return entity[i].Username < entity[j].Username
// 		})
// 		c.JSON(200, gin.H{
// 			"code":           200,
// 			"active":         active,
// 			"total_page":     totalPages,
// 			"total_document": count,
// 			"current_paget":  page,
// 			"limit_data":     limit,
// 			"payload":        entity,
// 		})
// 	}
// }

// func SearchAccount(resource *db.Resource) func(c *gin.Context) {
// 	type RoleModel struct {
// 		RoleName string `bson:"role_name" json:"role_name"`
// 	}
// 	type Body struct {
// 		Username    string    `bson:"username" json:"username"`
// 		Prefix      string    `bson:"prefix" json:"prefix"` // คำนำหน้าชื่อ
// 		Name        string    `bson:"name" json:"name"`
// 		LastName    string    `bson:"last_name" json:"last_name"`
// 		FullName    string    `bson:"full_name" json:"full_name"`
// 		Role        RoleModel `bson:"role" json:"role"` // สิทธิ์การใช้งาน
// 		PhoneNumber string    `bson:"phone_number" json:"phone_number"`
// 		Customer    string    `bson:"customer" json:"customer"`
// 		Product     string    `bson:"product" json:"product"`
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

// 		var entity []model.UserModel
// 		filter := bson.M{
// 			"active":         true,
// 			"user_delete":    false,
// 			"username":       primitive.Regex{Pattern: body.Username, Options: ""},
// 			"prefix":         primitive.Regex{Pattern: body.Prefix, Options: ""},
// 			"name":           primitive.Regex{Pattern: body.Name, Options: ""},
// 			"last_name":      primitive.Regex{Pattern: body.LastName, Options: ""},
// 			"full_name":      primitive.Regex{Pattern: body.FullName, Options: ""},
// 			"role.role_name": primitive.Regex{Pattern: body.Role.RoleName, Options: ""},
// 			"phone_number":   primitive.Regex{Pattern: body.PhoneNumber, Options: ""},
// 		}
// 		if err := repo.GetManyStatement(resource, "user", filter, nil, &entity); err != nil {
// 			CreateLog(resource, c, "Search User Fail", " Search User Fail | error : "+err.Error())
// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"payload": err.Error(),
// 			})
// 			return
// 		}
// 		if entity == nil {
// 			CreateLog(resource, c, "SearchUser not found", " Search User not found")

// 			c.JSON(400, gin.H{
// 				"code":    400,
// 				"text":    "not found",
// 				"payload": body,
// 			})
// 			return
// 		}
// 		sort.SliceStable(entity, func(i, j int) bool {
// 			return entity[i].Username < entity[j].Username
// 		})
// 		CreateLog(resource, c, "SearchUser Success", " SearchUser Success")
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"payload": entity,
// 		})
// 		//
// 	}
// }

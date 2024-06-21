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

func KnowledgeFromTicket(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.TicketModelID
		filter := bson.M{
			"ticket_cancel": false,
			"ticket_delete": false,
			"ticket_status": "Success",
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                       value,
			"user_delete":                    false,
			"role.permission.knowledge.read": true,
			"active":                         true,
		}
		filterOption1 := bson.M{}

		err := repo.GetOneStatement(resource, "user", filter1, filterOption1, &permission)
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

func CreateKnowledge(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var knowledge struct {
			TitleKnowledge    string   `bson:"title_knowledge" json:"title_knowledge"`
			SolutionKnowledge string   `bson:"solution_knowledge" json:"solution_knowledge"`
			PictureKnowledge  []string `bson:"picture_knowledge" json:"picture_knowledge"`
			TagTicket         []string `bson:"tag_ticket" json:"tag_ticket"`
			TypeTicket        string   `bson:"type_ticket" json:"type_ticket"`
			Username          string   `bson:"username" json:"username"`
		}
		if err := c.ShouldBind(&knowledge); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                         value,
			"user_delete":                      false,
			"role.permission.knowledge.create": true,
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
		//ดึง Type
		var TypeTicketModelS model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_name":   knowledge.TypeTicket,
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
		for _, v := range knowledge.TagTicket {
			var TagticketModelONE model.TagTicketModelS
			filter := bson.M{
				"tag_ticket_name":   v,
				"tag_ticket_delete": false,
			}
			filterOption := bson.M{}

			err := repo.GetOneStatement(resource, "TagTicket", filter, filterOption, &TagticketModelONE)
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
		//ดึง User (เดี๋ยวมาทำต่อ)
		var UserModel model.UserModelPassTo
		filter2 := bson.M{
			"username":    knowledge.Username,
			"user_delete": false,
			"active":      true,
		}
		filterOption2 := bson.M{}

		err2 := repo.GetOneStatement(resource, "user", filter2, filterOption2, &UserModel)
		if err2 != nil {

			c.JSON(400, gin.H{
				"code":    400,
				"payload": filter,
			})
			return
		}
		// แก้type and tag
		TicketData := model.KnowledgeModel{
			TitleKnowledge:    knowledge.TitleKnowledge,
			SolutionKnowledge: knowledge.SolutionKnowledge,
			PictureKnowledge:  knowledge.PictureKnowledge,
			TagTicket:         TagticketModel,
			Username:          UserModel,
			TypeTicket:        TypeTicketModelS,
			CreateTime:        time.Now(),
			DeleteTime:        time.Time{},
			UpdateTime:        time.Time{},
			KnowledgeDelete:   false,
		}

		result, err := repo.CreateStatement(resource, "Knowledge", TicketData)
		if err != nil {
			CreateLog(resource, c, "Create Forum Failed", "Create Forum Failed | err : "+err.Error())
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "Create Forum Success", "Create Forum name : "+knowledge.TitleKnowledge+" Success ")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
func GetAllKnowledge(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		var entity []model.KnowledgeModelID

		var permission model.UserModelS
		value := c.GetString("username")
		filter1 := bson.M{
			"username":                       value,
			"user_delete":                    false,
			"role.permission.knowledge.read": true,
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
			"knowledge_delete": false,
		}
		limit, _ := strconv.ParseInt(c.GetHeader("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.GetHeader("page"), 10, 64)
		if limit == 0 {
			limit = 20
		}
		if page == 0 {
			page = 1
		}
		totalPages, count, err := repo.GetManyStatementLimitPagination(resource, "Knowledge", filter, nil, limit, page, &entity)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}
		sort.SliceStable(entity, func(i, j int) bool {
			return entity[i].TitleKnowledge < entity[j].TitleKnowledge
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
func UpdateKnowledge(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Id                string   `bson:"_id" json:"_id" form:"_id"`
		TitleKnowledge    string   `bson:"title_knowledge" json:"title_knowledge"`
		SolutionKnowledge string   `bson:"solution_knowledge" json:"solution_knowledge"`
		PictureKnowledge  []string `bson:"picture_knowledge" json:"picture_knowledge"`
		TagTicket         []string `bson:"tag_ticket" json:"tag_ticket"`
		TypeTicket        string   `bson:"type_ticket" json:"type_ticket"`
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
			"username":                         value,
			"user_delete":                      false,
			"role.permission.knowledge.update": true,
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
		var TagticketModel []model.TagTicketModelS
		for _, v := range body.TagTicket {
			var TagticketModelONE model.TagTicketModelS
			filter := bson.M{
				"tag_ticket_name":   v,
				"tag_ticket_delete": false,
			}
			filterOption := bson.M{}

			err := repo.GetOneStatement(resource, "TagTicket", filter, filterOption, &TagticketModelONE)
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
		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "UpdateKnowledge fail", "ID not from Hex  | error : "+err.Error())
			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return

		}
		filter := bson.M{
			"_id":              Id,
			"knowledge_delete": false}
		update := bson.M{"$set": bson.M{
			"title_knowledge":              body.TitleKnowledge,
			"solution_knowledge":           body.SolutionKnowledge,
			"picture_knowledge":            body.PictureKnowledge,
			"tag_ticket":                   TagticketModel,
			"type_ticket.type_ticket_name": body.TypeTicket,
			"update_time":                  time.Now(),
		}}

		var beforeUpdate model.KnowledgeModel
		err = repo.GetOneStatement(resource, "Knowledge", filter, nil, &beforeUpdate)

		Result, err := repo.UpdateOneStatement(resource, "Knowledge", filter, update)
		if err != nil {
			CreateLog(resource, c, "UpdateKnowledge fail", "Update Forum fail | err : "+err.Error())
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}

		var afterUpdate model.KnowledgeModel
		err = repo.GetOneStatement(resource, "Knowledge", filter, nil, &afterUpdate)

		CreateLogForUpdate(resource, c, "UpdateKnowledge Success", "Update Forum name : "+body.TitleKnowledge+" Success ", beforeUpdate, afterUpdate)
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}
func DeleteOneKnowledge(resource *db.Resource) func(c *gin.Context) {
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
			"username":                         value,
			"user_delete":                      false,
			"role.permission.knowledge.delete": true,
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
		Id, err := primitive.ObjectIDFromHex(body.Id)
		if err != nil {
			CreateLog(resource, c, "DeleteOneKnowledge fail", "ID not from Hex  | error : "+err.Error())

			c.JSON(500, gin.H{
				"code":    500,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		filter := bson.M{"_id": Id}
		update := bson.M{"$set": bson.M{
			"knowledge_delete": true,
			"delete_time":      time.Now(),
		}}
		Result, err := repo.UpdateOneStatement(resource, "Knowledge", filter, update)
		if err != nil {
			CreateLog(resource, c, "DeleteOneKnowledge fail", "DeleteOneKnowledge fail ")
			c.JSON(400, gin.H{
				"code":    401,
				"text":    "Update Fail",
				"payload": err.Error(),
			})
			return
		}
		CreateLog(resource, c, "DeleteOneKnowledge Success", "Delete Forum ID : "+body.Id+" Success ")
		c.JSON(200, gin.H{
			"code":    200,
			"payload": Result,
		})
	}
}
func SearchKnowledge(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		TitleTicket    string `bson:"title_search" json:"title_search"`
		TagTicket      string `bson:"tag_ticket" json:"tag_ticket"`
		TypeTicket     string `bson:"type_ticket" json:"type_ticket"`
		TitleKnowledge string `bson:"title_knowledge" json:"title_knowledge"`
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
		var entity1 []model.KnowledgeModel
		filter := bson.M{
			"knowledge_delete":             false,
			"title_knowledge":              primitive.Regex{Pattern: body.TitleTicket, Options: ""},
			"tag_ticket.tag_ticket_name":   primitive.Regex{Pattern: body.TagTicket, Options: ""},
			"type_ticket.type_ticket_name": primitive.Regex{Pattern: body.TypeTicket, Options: ""},
		}
		if err := repo.GetManyStatement(resource, "Knowledge", filter, nil, &entity1); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		sort.SliceStable(entity1, func(i, j int) bool {
			return entity1[i].TitleKnowledge < entity1[j].TitleKnowledge
		})

		var entity []model.TicketModel
		filter = bson.M{
			"ticket_delete":                false,
			"ticket_status":                bson.M{"$in": bson.A{"Success"}},
			"title_ticket":                 primitive.Regex{Pattern: body.TitleTicket, Options: ""},
			"tag_ticket.tag_ticket_name":   primitive.Regex{Pattern: body.TagTicket, Options: ""},
			"type_ticket.type_ticket_name": primitive.Regex{Pattern: body.TypeTicket, Options: ""},
		}
		if err := repo.GetManyStatement(resource, "ticket", filter, nil, &entity); err != nil {
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
			"code":     200,
			"payload1": entity,
			"payload2": entity1,
		})
	}
}
func GetOneKnowledge(resource *db.Resource) func(c *gin.Context) {
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
			CreateLog(resource, c, "Get one Knowledge Fail", " ID not from Hex | err :	"+err.Error())
			c.JSON(400, gin.H{
				"code":    400,
				"text":    "ObjectIDFromHex Fail",
				"payload": err.Error(),
			})
			return
		}
		// test login
		var KnowledgeModel model.KnowledgeModelID
		filter := bson.M{
			"_id":              ID,
			"knowledge_delete": false,
		}
		filterOption := bson.M{}

		err = repo.GetOneStatement(resource, "Knowledge", filter, filterOption, &KnowledgeModel)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"payload": filter,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"payload": KnowledgeModel,
		})
	}
}

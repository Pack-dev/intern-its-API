package controller

import (
	"templategoapi/db"
	"templategoapi/model"
	"templategoapi/repo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AggregateY(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{"ticket_delete": false},
			// "$unwind": bson.M{"path": "$payload.description"},
		}

		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"DateCreate": bson.M{
						"$dateToString": bson.M{
							"format": "%Y",
							"date":   "$create_time",
						},
					},
				},
				"data":  bson.M{"$push": bson.M{"_id": "$_id", "title_ticket": "$title_ticket", "username": "$username", "product": "$product", "report_customer": "$report_customer", "description": "$description", "ticket_status": "$ticket_status", "pass_to": "$pass_to", "solution": "$solution"}},
				"count": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateYM(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{"ticket_delete": false, "ticket_cancel": false},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"DateCreate": bson.M{
						"$dateToString": bson.M{
							"format": "%m-%Y",
							"date":   "$create_time",
						},
					},
				},
				"data":  bson.M{"$push": bson.M{"_id": "$_id", "title_ticket": "$title_ticket", "username": "$username", "product": "$product", "report_customer": "$report_customer", "description": "$description", "ticket_status": "$ticket_status", "pass_to": "$pass_to", "solution": "$solution"}},
				"count": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateYMD(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{"ticket_delete": false},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"DateCreate": bson.M{
						"$dateToString": bson.M{
							"format": "%d-%m-%Y",
							"date":   "$create_time",
						},
					},
				},
				// "data":  bson.M{"$push": bson.M{"_id": "$_id", "title_ticket": "$title_ticket", "username": "$username", "product": "$product", "report_customer": "$report_customer", "description": "$description", "ticket_status": "$ticket_status", "pass_to": "$pass_to", "solution": "$solution"}},
				"count": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateWeek(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{"ticket_delete": false},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"week": bson.M{
						"$week": "$create_time",
					},
				},
				"data":  bson.M{"$push": bson.M{"_id": "$_id", "title_ticket": "$title_ticket", "username": "$username", "product": "$product", "report_customer": "$report_customer", "description": "$description", "ticket_status": "$ticket_status", "pass_to": "$pass_to", "solution": "$solution"}},
				"count": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateCountPendingYM(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "Pending",
			},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%m-%Y",
						"date":   "$create_time",
					},
				},
				"CountDocument": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateCountInfixYM(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "In fix",
			},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%m-%Y",
						"date":   "$create_time",
					},
				},
				"CountDocument": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateCountSuccessYM(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "Success",
			},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%m-%Y",
						"date":   "$create_time",
					},
				},
				"CountDocument": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateCountSuccessYMD(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "Success",
			},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%d-%m-%Y",
						"date":   "$create_time",
					},
				},
				"CountDocument": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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
func AggregateTicket_Count_All(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var result []struct {
			All_Ticket           int `json:"all_ticket"`
			All_Ticket_Success   int `json:"all_ticket_success"`
			All_Ticket_Infix     int `json:"all_ticket_infix"`
			All_Ticket_Pending   int `json:"all_ticket_pending"`
			All_Ticket_Unsuccess int `json:"all_ticket_unsuccess"`
		}
		// var entity bson.M

		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
			},
		}

		group := bson.M{
			"$count": "all_ticket",
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &result); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		filter = bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "Success"},
		}

		group = bson.M{
			"$count": "all_ticket_success",
		}
		pipeline = []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &result); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		filter = bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "In fix"},
		}

		group = bson.M{
			"$count": "all_ticket_infix",
		}
		pipeline = []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &result); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		filter = bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "Pending"},
		}

		group = bson.M{
			"$count": "all_ticket_pending",
		}
		pipeline = []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &result); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"payload": err.Error(),
			})
			return
		}

		result[0].All_Ticket_Unsuccess = result[0].All_Ticket_Infix + result[0].All_Ticket_Pending
		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})

	}
}
func Aggregate_TypeTicket_Count_YMD(resource *db.Resource) func(c *gin.Context) {
	// Mo mean Model

	type TypeTicketMo struct {
		TypeTicketName string   `json:"type_ticket_name"`
		TypeTicketData []bson.M `json:"type_ticket_data"`
	}

	return func(c *gin.Context) {

		//get type ticket
		var entity_type_ticket []model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_delete": false,
		}
		if err := repo.GetManyStatement(resource, "typeTicket", filter, nil, &entity_type_ticket); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"msg":     "get type ticket fail",
				"payload": err.Error(),
			})
			return
		}
		var TypeTicketEntity []TypeTicketMo
		for _, v := range entity_type_ticket {
			//get count type ticket
			var entity []bson.M
			filter := bson.M{
				"$match": bson.M{
					"ticket_delete":                false,
					"type_ticket.type_ticket_name": v.TypeTicketName,
				},
			}
			group := bson.M{
				"$group": bson.M{
					"_id": bson.M{
						"$dateToString": bson.M{
							"format": "%d-%m-%Y",
							"date":   "$create_time",
						},
					},
					"CountDocument": bson.M{"$sum": 1},
				},
			}
			pipeline := []bson.M{filter, group}
			if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "get count type ticket fail",
					"payload": err.Error(),
				})
				return
			}

			if entity == nil {
				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					TypeTicketData: nil,
				})
			} else {

				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					TypeTicketData: entity,
				})
			}

		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": TypeTicketEntity,
		})

	}
}
func Aggregate_TypeTicket_Count_YM(resource *db.Resource) func(c *gin.Context) {
	// Mo mean Model

	type TypeTicketMo struct {
		TypeTicketName string   `json:"type_ticket_name"`
		TypeTicketData []bson.M `json:"type_ticket_data"`
	}

	return func(c *gin.Context) {

		//get type ticket
		var entity_type_ticket []model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_delete": false,
		}
		if err := repo.GetManyStatement(resource, "typeTicket", filter, nil, &entity_type_ticket); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"msg":     "get type ticket fail",
				"payload": err.Error(),
			})
			return
		}
		var TypeTicketEntity []TypeTicketMo
		for _, v := range entity_type_ticket {
			//get count type ticket
			var entity []bson.M
			filter := bson.M{
				"$match": bson.M{
					"ticket_delete":                false,
					"type_ticket.type_ticket_name": v.TypeTicketName,
				},
			}
			group := bson.M{
				"$group": bson.M{
					"_id": bson.M{
						"$dateToString": bson.M{
							"format": "%m-%Y",
							"date":   "$create_time",
						},
					},
					"CountDocument": bson.M{"$sum": 1},
				},
			}
			pipeline := []bson.M{filter, group}
			if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "get count type ticket fail",
					"payload": err.Error(),
				})
				return
			}

			if entity == nil {
				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					TypeTicketData: nil,
				})
			} else {

				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					TypeTicketData: entity,
				})
			}

		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": TypeTicketEntity,
		})

	}
}
func Aggregate_TypeTicket_Count_Y(resource *db.Resource) func(c *gin.Context) {
	// Mo mean Model

	type TypeTicketMo struct {
		TypeTicketName string   `json:"type_ticket_name"`
		TypeTicketData []bson.M `json:"type_ticket_data"`
	}

	return func(c *gin.Context) {

		//get type ticket
		var entity_type_ticket []model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_delete": false,
		}
		if err := repo.GetManyStatement(resource, "typeTicket", filter, nil, &entity_type_ticket); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"msg":     "get type ticket fail",
				"payload": err.Error(),
			})
			return
		}
		var TypeTicketEntity []TypeTicketMo
		for _, v := range entity_type_ticket {
			//get count type ticket
			var entity []bson.M
			filter := bson.M{
				"$match": bson.M{
					"ticket_delete":                false,
					"type_ticket.type_ticket_name": v.TypeTicketName,
				},
			}
			group := bson.M{
				"$group": bson.M{
					"_id": bson.M{
						"$dateToString": bson.M{
							"format": "%Y",
							"date":   "$create_time",
						},
					},
					"CountDocument": bson.M{"$sum": 1},
				},
			}
			pipeline := []bson.M{filter, group}
			if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "get count type ticket fail",
					"payload": err.Error(),
				})
				return
			}

			if entity == nil {
				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					TypeTicketData: nil,
				})
			} else {

				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					TypeTicketData: entity,
				})
			}

		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": TypeTicketEntity,
		})

	}
}
func Aggregate_TypeTicket_Count_All(resource *db.Resource) func(c *gin.Context) {
	// Mo mean Model

	type TypeTicketMo struct {
		TypeTicketName string `json:"type_ticket_name"`
		CountDocument  int    `json:"count_document"`
	}

	return func(c *gin.Context) {

		//get type ticket
		var entity_type_ticket []model.TypeTicketModelS
		filter := bson.M{
			"type_ticket_delete": false,
		}
		if err := repo.GetManyStatement(resource, "typeTicket", filter, nil, &entity_type_ticket); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"msg":     "get type ticket fail",
				"payload": err.Error(),
			})
			return
		}
		var TypeTicketEntity []TypeTicketMo
		for _, v := range entity_type_ticket {
			//get count type ticket

			var entity []struct {
				CountDocument int `json:"count_document"`
			}
			filter := bson.M{
				"$match": bson.M{
					"ticket_delete":                false,
					"type_ticket.type_ticket_name": v.TypeTicketName,
				},
			}
			group := bson.M{
				"$count": "count_document",
			}
			pipeline := []bson.M{filter, group}
			if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
				c.JSON(400, gin.H{
					"code":    400,
					"msg":     "get count type ticket fail",
					"payload": err.Error(),
				})
				return
			}

			if entity == nil {
				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					CountDocument:  0,
				})
			} else {

				TypeTicketEntity = append(TypeTicketEntity, TypeTicketMo{
					TypeTicketName: v.TypeTicketName,
					CountDocument:  entity[0].CountDocument,
				})
			}
		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": TypeTicketEntity,
		})

	}
}

func AggregateCountPendingYMD(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "Pending",
			},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%d-%m-%Y",
						"date":   "$create_time",
					},
				},
				"CountDocument": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

func AggregateCountInfixYMD(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		var entity []bson.M
		filter := bson.M{
			"$match": bson.M{
				"ticket_delete": false,
				"ticket_status": "In fix",
			},
		}
		group := bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%d-%m-%Y",
						"date":   "$create_time",
					},
				},
				"CountDocument": bson.M{"$sum": 1},
			},
		}
		pipeline := []bson.M{filter, group}
		if err := repo.AggregateStatement(resource, "ticket", pipeline, &entity); err != nil {
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

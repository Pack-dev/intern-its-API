package model

import "time"

type TypeTicketModelS struct {
	TypeTicketName string `bson:"type_ticket_name" json:"type_ticket_name"`
}
type TypeTicketModel struct {
	TypeTicketName   string    `bson:"type_ticket_name" json:"type_ticket_name"`
	CreateTime       time.Time `bson:"create_time" json:"create_time"`
	UpdateTime       time.Time `bson:"update_time" json:"update_time"`
	DeleteTime       time.Time `bson:"delete_time" json:"delete_time"`
	TypeTicketDelete bool      `bson:"type_ticket_delete" json:"type_ticket_delete"`
}

type TypeTicketModelID struct {
	Id               string    `bson:"_id" json:"_id"`
	TypeTicketName   string    `bson:"type_ticket_name" json:"type_ticket_name"`
	CreateTime       time.Time `bson:"create_time" json:"create_time"`
	UpdateTime       time.Time `bson:"update_time" json:"update_time"`
	DeleteTime       time.Time `bson:"delete_time" json:"delete_time"`
	TypeTicketDelete bool      `bson:"type_ticket_delete" json:"type_ticket_delete"`
}

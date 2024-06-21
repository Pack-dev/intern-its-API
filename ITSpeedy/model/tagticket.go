package model

import "time"

type TagTicketModelS struct {
	TagTicketName string `bson:"tag_ticket_name" json:"tag_ticket_name"`
}

type TagTicketModel struct {
	TagTicketName    string    `bson:"tag_ticket_name" json:"tag_ticket_name"`
	CreateTime       time.Time `bson:"create_time" json:"create_time"`
	UpdateTime       time.Time `bson:"update_time" json:"update_time"`
	DeleteTime       time.Time `bson:"delete_time" json:"delete_time"`
	TypeTicketDelete bool      `bson:"tag_ticket_delete" json:"tag_ticket_delete"`
}

type TagTicketModelID struct {
	Id               string    `bson:"_id" json:"_id"`
	TagTicketName    string    `bson:"tag_ticket_name" json:"tag_ticket_name"`
	CreateTime       time.Time `bson:"create_time" json:"create_time"`
	UpdateTime       time.Time `bson:"update_time" json:"update_time"`
	DeleteTime       time.Time `bson:"delete_time" json:"delete_time"`
	TypeTicketDelete bool      `bson:"tag_ticket_delete" json:"tag_ticket_delete"`
}

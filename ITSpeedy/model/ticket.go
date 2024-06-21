package model

import (
	"time"
)

type TicketModelID struct {
	ID              string              `bson:"_id" json:"_id"`
	TitleTicket     string              `bson:"title_ticket" json:"title_ticket"`
	TagTicket       []TagTicketModelS   `bson:"tag_ticket" json:"tag_ticket"`
	Username        UserTicketModel     `bson:"username" json:"username"`
	Product         []ProductModelS     `bson:"product" json:"product"`
	ReportCustomer  CustomerTicketModel `bson:"report_customer" json:"report_customer"`
	TypeTicket      TypeTicketModelS    `bson:"type_ticket" json:"type_ticket"`
	Description     string              `bson:"description" json:"description"`
	PictureReport   []string            `bson:"picture_report" json:"picture_report"`
	CreateTime      time.Time           `bson:"create_time" json:"create_time"`
	DeleteTime      time.Time           `bson:"delete_time" json:"delete_time"`
	UpdateTime      time.Time           `bson:"update_time" json:"update_time"`
	SuccessTime     time.Time           `bson:"success_time" json:"success_time"`
	CancelTime      time.Time           `bson:"cancel_time" json:"cancel_time"`
	TicketStatus    string              `bson:"ticket_status" json:"ticket_status"`
	PassTo          PassToModel         `bson:"pass_to" json:"pass_to"`
	TicketCancel    bool                `bson:"ticket_cancel" json:"ticket_cancel"`
	TicketDelete    bool                `bson:"ticket_delete" json:"ticket_delete"`
	Solution        string              `bson:"solution" json:"solution"`
	PictureSolution []string            `bson:"picture_solution" json:"picture_solution"`
}
type TicketModel struct {
	TitleTicket     string              `bson:"title_ticket" json:"title_ticket"`
	TagTicket       []TagTicketModelS   `bson:"tag_ticket" json:"tag_ticket"`
	Username        UserTicketModel     `bson:"username" json:"username"`
	Product         []ProductModelS     `bson:"product" json:"product"`
	ReportCustomer  CustomerTicketModel `bson:"report_customer" json:"report_customer"`
	TypeTicket      TypeTicketModelS    `bson:"type_ticket" json:"type_ticket"`
	Description     string              `bson:"description" json:"description"`
	PictureReport   []string            `bson:"picture_report" json:"picture_report"`
	CreateTime      time.Time           `bson:"create_time" json:"create_time"`
	DeleteTime      time.Time           `bson:"delete_time" json:"delete_time"`
	UpdateTime      time.Time           `bson:"update_time" json:"update_time"`
	SuccessTime     time.Time           `bson:"success_time" json:"success_time"`
	CancelTime      time.Time           `bson:"cancel_time" json:"cancel_time"`
	TicketStatus    string              `bson:"ticket_status" json:"ticket_status"`
	PassTo          PassToModel         `bson:"pass_to" json:"pass_to"`
	TicketCancel    bool                `bson:"ticket_cancel" json:"ticket_cancel"`
	TicketDelete    bool                `bson:"ticket_delete" json:"ticket_delete"`
	Solution        string              `bson:"solution" json:"solution"`
	PictureSolution []string            `bson:"picture_solution" json:"picture_solution"`
}

type TicketModelP struct {
	TitleTicket     string           `bson:"title_ticket" json:"title_ticket"`
	TagTicket       TagTicketModelS  `bson:"tag_ticket" json:"tag_ticket"`
	Username        string           `bson:"username" json:"username"`
	TypeTicket      TypeTicketModelS `bson:"type_ticket" json:"type_ticket"`
	Description     string           `bson:"description" json:"description"`
	PictureReport   string           `bson:"picture_report" json:"picture_report"`
	CreateTime      time.Time        `bson:"create_time" json:"create_time"`
	DeleteTime      time.Time        `bson:"delete_time" json:"delete_time"`
	UpdateTime      time.Time        `bson:"update_time" json:"update_time"`
	SuccessTime     time.Time        `bson:"success_time" json:"success_time"`
	CancelTime      time.Time        `bson:"cancel_time" json:"cancel_time"`
	TicketStatus    string           `bson:"ticket_status" json:"ticket_status"`
	PassTo          PassToModel      `bson:"pass_to" json:"pass_to"`
	TicketCancel    bool             `bson:"ticket_cancel" json:"ticket_cancel"`
	TicketDelete    bool             `bson:"ticket_delete" json:"ticket_delete"`
	Solution        string           `bson:"solution" json:"solution"`
	PictureSolution string           `bson:"picture_solution" json:"picture_solution"`
}
type UserTicketModel struct {
	Username    string          `bson:"username" json:"username"`
	FullName    string          `bson:"full_name" json:"full_name"`
	Product     []ProductModelS `bson:"product" json:"product"`
	PhoneNumber string          `bson:"phone_number" json:"phone_number"`
	Customer    string          `bson:"customer" json:"customer"`
}

type CustomerTicketModel struct {
	ID          string          `bson:"_id" json:"_id"`
	FullName    string          `bson:"full_name" json:"full_name"`
	PhoneNumber string          `bson:"phone_number" json:"phone_number"`
	Customer    string          `bson:"customer" json:"customer"`
	Product     []ProductModelS `bson:"product" json:"product"`
}

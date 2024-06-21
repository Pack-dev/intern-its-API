package model

import "time"

type CustomerModelS struct {
	FullName string `bson:"full_name" json:"full_name"`
}
type CustomerModel struct {
	Prefix         string          `bson:"prefix" json:"prefix"`
	FirstName      string          `bson:"first_name" json:"first_name"`
	LastName       string          `bson:"last_name" json:"last_name"`
	FullName       string          `bson:"full_name" json:"full_name"`
	PhoneNumber    string          `bson:"phone_number" json:"phone_number"`
	Customer       string          `bson:"customer" json:"customer"`
	Product        []ProductModelS `bson:"product" json:"product"`
	CreateTime     time.Time       `bson:"create_time" json:"create_time"`
	UpdateTime     time.Time       `bson:"update_time" json:"update_time"`
	DeleteTime     time.Time       `bson:"delete_time" json:"delete_time"`
	CustomerDelete bool            `bson:"customer_delete" json:"customer_delete"`
}

type CustomerModelID struct {
	Id             string          `bson:"_id" json:"_id"`
	Prefix         string          `bson:"prefix" json:"prefix"`
	FirstName      string          `bson:"first_name" json:"first_name"`
	LastName       string          `bson:"last_name" json:"last_name"`
	FullName       string          `bson:"full_name" json:"full_name"`
	PhoneNumber    string          `bson:"phone_number" json:"phone_number"`
	Customer       string          `bson:"customer" json:"customer"`
	Product        []ProductModelS `bson:"product" json:"product"`
	CreateTime     time.Time       `bson:"create_time" json:"create_time"`
	UpdateTime     time.Time       `bson:"update_time" json:"update_time"`
	DeleteTime     time.Time       `bson:"delete_time" json:"delete_time"`
	CustomerDelete bool            `bson:"customer_delete" json:"customer_delete"`
}

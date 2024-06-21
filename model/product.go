package model

import "time"

type ProductModelS struct {
	ProductName string `bson:"product_name" json:"product_name"`
}

type ProductModel struct {
	ProductName   string    `bson:"product_name" json:"product_name"`
	CreateTime    time.Time `bson:"create_time" json:"create_time"`
	UpdateTime    time.Time `bson:"update_time" json:"update_time"`
	DeleteTime    time.Time `bson:"delete_time" json:"delete_time"`
	ProductDelete bool      `bson:"product_delete" json:"product_delete"`
}
type ProductModelID struct {
	Id            string    `bson:"_id" json:"_id"`
	ProductName   string    `bson:"product_name" json:"product_name"`
	CreateTime    time.Time `bson:"create_time" json:"create_time"`
	UpdateTime    time.Time `bson:"update_time" json:"update_time"`
	DeleteTime    time.Time `bson:"delete_time" json:"delete_time"`
	ProductDelete bool      `bson:"product_delete" json:"product_delete"`
}

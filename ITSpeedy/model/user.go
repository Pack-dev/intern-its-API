package model

import (
	"time"
)

type AccountModel struct {
	Datetime time.Time `bson:"date_time" json:"date_time"`
	Username string    `bson:"username" json:"username"`
	Password string    `bson:"password" json:"password"`
	Name     string    `bosn:"name" json:"name"`
}
type UserModelS struct {
	ID       string `bson:"_id" json:"_id"` // รหัสผู้ใช้งาน
	Username string `bson:"username" json:"username"`
	//Password    string    `bson:"password" json:"password"`
	Prefix          string          `bson:"prefix" json:"prefix"` // คำนำหน้าชื่อ
	Name            string          `bson:"name" json:"name"`
	LastName        string          `bson:"last_name" json:"last_name"`
	FullName        string          `bson:"full_name" json:"full_name"`
	Role            RoleModeli      `bson:"role" json:"role"` // สิทธิ์การใช้งาน
	PhoneNumber     string          `bson:"phone_number" json:"phone_number"`
	ImgUrl          string          `bson:"img_url" json:"img_url"`
	Active          bool            `bson:"active" json:"active"`
	Customer        string          `bson:"customer" json:"customer"`
	Product         []ProductModelS `bson:"product" json:"product"`
	CreateTime      time.Time       `bson:"create_time" json:"create_time"` // วันที่สร้าง
	LastLogin       time.Time       `bson:"last_login" json:"last_login"`   // วันที่เข้าใช้งานล่าสุด
	UpdateTime      time.Time       `bson:"update_time" json:"update_time"` // วันที่แก้ไขล่าสุด
	DeleteTime      time.Time       `bson:"delete_time" json:"delete_time"` // วันที่ลบ
	UserDelete      bool            `bson:"user_delete" json:"user_delete"`
	DefaultPassword bool            `bson:"default_password" json:"default_password"`
	// jwt.StandardClaims
}

type UserModel struct {
	Username        string          `bson:"username" json:"username"`
	Password        string          `bson:"password" json:"password"`
	Prefix          string          `bson:"prefix" json:"prefix"` // คำนำหน้าชื่อ
	Name            string          `bson:"name" json:"name"`
	LastName        string          `bson:"last_name" json:"last_name"`
	FullName        string          `bson:"full_name" json:"full_name"`
	Role            RoleModeli      `bson:"role" json:"role"` // สิทธิ์การใช้งาน
	PhoneNumber     string          `bson:"phone_number" json:"phone_number"`
	ImgUrl          string          `bson:"img_url" json:"img_url"`
	CreateTime      time.Time       `bson:"create_time" json:"create_time"` // วันที่สร้าง
	LastLogin       time.Time       `bson:"last_login" json:"last_login"`   // วันที่เข้าใช้งานล่าสุด
	UpdateTime      time.Time       `bson:"update_time" json:"update_time"` // วันที่แก้ไขล่าสุด
	DeleteTime      time.Time       `bson:"delete_time" json:"delete_time"` // วันที่ลบ
	Customer        string          `bson:"customer" json:"customer"`
	Product         []ProductModelS `bson:"product" json:"product"`
	Active          bool            `bson:"active" json:"active"`
	UserDelete      bool            `bson:"user_delete" json:"user_delete"`
	DefaultPassword bool            `bson:"default_password" json:"default_password"`
}
type RoleModeli struct {
	Rolename   string         `bson:"role_name" json:"role_name"`
	Permission RolePermission `bson:"permission" json:"permission"`
}
type RoleModel struct {
	Rolename   string         `bson:"role_name" json:"role_name"`
	Permission RolePermission `bson:"permission" json:"permission"`
	CreateTime time.Time      `bson:"create_time" json:"create_time"` // วันที่สร้าง
	UpdateTime time.Time      `bson:"update_time" json:"update_time"` // วันที่แก้ไขล่าสุด
	DeleteTime time.Time      `bson:"delete_time" json:"delete_time"` // วันที่ลบ
	RoleDelete bool           `bson:"role_delete" json:"role_delete"`
}
type PermissionModel struct {
	CreateTicket bool `bson:"create_ticket" json:"create_ticket"`
	UpdateTicket bool `bson:"update_ticket" json:"update_ticket"`
	DeleteTicket bool `bson:"delete_ticket" json:"delete_ticket"`
}
type PassToModel struct {
	Role string          `bson:"role_name" json:"role_name"`
	User UserModelPassTo `bson:"user" json:"user"`
}
type UserModelPassTo struct {
	Username string `bson:"username" json:"username"`
	FullName string `bson:"full_name" json:"full_name"`
}

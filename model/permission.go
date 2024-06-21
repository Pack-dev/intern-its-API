package model

type RolePermission struct {
	Customer   Permission `bson:"customer" json:"customer"`
	User       Permission `bson:"user" json:"user"`
	Product    Permission `bson:"product" json:"product"`
	TagTicket  Permission `bson:"tag_ticket" json:"tag_ticket"`
	TypeTicket Permission `bson:"type_ticket" json:"type_ticket"`
	Knowledge  Permission `bson:"knowledge" json:"knowledge"`
	Role       Permission `bson:"role" json:"role"`
	Ticket     Permission `bson:"ticket" json:"ticket"`
	AllTicket  Permission `bson:"all_ticket" json:"all_ticket"`
	Log        Permission `bson:"log" json:"log"`
	AllLog     Permission `bson:"all_log" json:"all_log"`
	Dashboard  Permission `bson:"dashboard" json:"dashboard"`
}

type Permission struct {
	Create bool `bson:"create" json:"create"`
	Read   bool `bson:"read" json:"read"`
	Update bool `bson:"update" json:"update"`
	Delete bool `bson:"delete" json:"delete"`
}

package model

import "time"

type LogModel struct {
	User_ID     string    `bson:"user_id" json:"user_id"`
	Username    string    `bson:"username" json:"username"`
	Name        string    `bson:"name" json:"name"`
	IpAddress   string    `bson:"ip_address" json:"ip_address"`
	TimeStamp   time.Time `bson:"time_stamp" json:"time_stamp"`
	Activity    string    `bson:"activity" json:"activity"`
	Description string    `bson:"description" json:"description"`
}

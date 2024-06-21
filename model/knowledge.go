package model

import "time"

type KnowledgeModel struct {
	TitleKnowledge    string            `bson:"title_knowledge" json:"title_knowledge"`
	SolutionKnowledge string            `bson:"solution_knowledge" json:"solution_knowledge"`
	PictureKnowledge  []string          `bson:"picture_knowledge" json:"picture_knowledge"`
	TagTicket         []TagTicketModelS `bson:"tag_ticket" json:"tag_ticket"`
	Username          UserModelPassTo   `bson:"username" json:"username"`
	TypeTicket        TypeTicketModelS  `bson:"type_ticket" json:"type_ticket"`
	CreateTime        time.Time         `bson:"create_time" json:"create_time"`
	DeleteTime        time.Time         `bson:"delete_time" json:"delete_time"`
	UpdateTime        time.Time         `bson:"update_time" json:"update_time"`
	KnowledgeDelete   bool              `bson:"knowledge_delete" json:"knowledge_delete"`
}

type KnowledgeModelID struct {
	Id                string            `bson:"_id" json:"_id"`
	TitleKnowledge    string            `bson:"title_knowledge" json:"title_knowledge"`
	SolutionKnowledge string            `bson:"solution_knowledge" json:"solution_knowledge"`
	PictureKnowledge  []string          `bson:"picture_knowledge" json:"picture_knowledge"`
	TagTicket         []TagTicketModelS `bson:"tag_ticket" json:"tag_ticket"`
	Username          UserModelPassTo   `bson:"username" json:"username"`
	TypeTicket        TypeTicketModelS  `bson:"type_ticket" json:"type_ticket"`
	CreateTime        time.Time         `bson:"create_time" json:"create_time"`
	DeleteTime        time.Time         `bson:"delete_time" json:"delete_time"`
	UpdateTime        time.Time         `bson:"update_time" json:"update_time"`
	KnowledgeDelete   bool              `bson:"knowledge_delete" json:"knowledge_delete"`
}

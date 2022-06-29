package infrastructure

import "gorm.io/gorm"

type PfUser struct {
	gorm.Model
	FirstName  string       `gorm:"not null;"`
	LastName   string       `gorm:"not null;"`
	Username   string       `gorm:"unique;not null"`
	Email      string       `gorm:"not null"`
	PfProjects []*PfProject `gorm:"many2many:pf_user_projects;"`
	Boards     []*PfBoard   `gorm:"many2many:pf_user_boards;"`
	Comments   []PfComment  `gorm:"foreignKey:PfUserID"`
}

type PfProject struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
	Image       string
	Kind        string
	PfUsers     []*PfUser `gorm:"many2many:pf_user_projects;"`
	PfBoards    []PfBoard
}

type PfBoard struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
	PfProjectID uint
	PfProject   PfProject
	Users       []*PfUser  `gorm:"many2many:pf_user_boards;"`
	Tickets     []PfTicket `gorm:"foreignKey:PfBoardID"`
}

type PfTicket struct {
	gorm.Model
	Title        string
	Description  string
	Votes        uint
	Status       string `gorm:"index"`
	PfBoardID    uint
	PfBoard      PfBoard
	PfCategoryID uint
	PfCategory   PfCategory
	PfComments   []PfComment `gorm:"foreignKey:PfTicketID"`
}

type PfCategory struct {
	gorm.Model
	Name string
}

type PfComment struct {
	gorm.Model
	Comment    string
	PfUserID   uint
	PfUser     PfUser
	PfTicketID uint
	PfTicket   PfTicket
}

type PfParentChildComment struct {
	gorm.Model
	PfParentCommentID uint
	PfChildCommentID  uint
	PfParentComment   PfComment
	PfChildComment    PfComment
}

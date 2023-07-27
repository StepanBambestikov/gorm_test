package main

type Order struct {
	ID            *string     `gorm:"type:uuid;primaryKey;unique"`
	Comment       *string     `gorm:"column:comment;not null"`
	RequisiteList []Requisite `gorm:"many2many:order_requisite_relation;"`
}

type Requisite struct {
	ID      *string `gorm:"type:uuid;primaryKey;not null;unique"`
	Comment *string `gorm:"column:comment;not null"`
}

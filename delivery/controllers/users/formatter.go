package controllers

import "project-e-commerces/entities"

type RegisterReqFormat struct {
	Name     string
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `gorm:"default:member"`
}

type LoginRequestFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseFormat struct {
	Message string        `json:"message"`
	Data    entities.User `json:"Data"`
	Token   string        `json:"Token"`
}

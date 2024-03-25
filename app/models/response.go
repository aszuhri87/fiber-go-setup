package models

import (
	"github.com/google/uuid"
)

type ListResponseOk struct {
	Data    []User `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type List2ResponseOk struct {
	Data    string `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type DataResponseOk struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type StatusOk struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type DataReturn struct {
	Data User `json:"data"`
}

type ResponseData struct {
	ID       uuid.UUID `json:id`
	Name     string    `json:"name"`
	Username string    `json:"username"`
}

type ResponseToken struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type ResponseStatus struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

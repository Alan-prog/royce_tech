package models

import "database/sql"

const (
	SqlNoRows = "no rows in result set"
)

type AliveResponse struct {
	Text string `json:"text"`
}

type CreateUserRequest struct {
	Name        *string `json:"name"`
	DOB         *string `json:"dob"`
	Address     *string `json:"address"`
	Description *string `json:"description"`
}

type SingleUserDataDbResponse struct {
	ID          int
	Name        string
	DOB         sql.NullTime
	Address     string
	Description string
	CreateAt    sql.NullTime
	UpdatedAt   sql.NullTime
}

type SingleUserData struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	DOB         *string `json:"dob"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	CreateAt    string  `json:"createAt"`
	UpdatedAt   *string `json:"updatedAt"`
}

type AllUsersData []*SingleUserData

type UpdateUserData struct {
	ID          int
	Name        *string
	DOB         *string
	Address     *string
	Description *string
}

package models

type AliveResponse struct {
	Text string
}

type CreateUserRequest struct{
	Name *string
	DOB *string
	Address *string
	Description *string
}

type SingleUserData struct{
	ID int32
	Name *string
	DOB *string
	Address *string
	Description *string
	CreateAt string
	UpdatedAt *string
}

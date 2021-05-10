package royce_tech_app

import (
	"context"
	"github.com/jackc/pgx"
	"my_projects/royce_tech/pkg/models"
)

type Royce interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest)(output int32, err error)
	GetSingleUser(ctx context.Context, id int32)(output models.SingleUserData, err error)
}

type royce struct {
	db *pgx.Conn
}

func (r *royce) Alive() (output models.AliveResponse, err error) {
	output.Text = "service is okay"
	return
}

func (r *royce) CreateUser(ctx context.Context, input *models.CreateUserRequest)(output int32, err error){
	const(
		dbRequest = `insert into human_resources (name, dob, address, description, created_at) values 
    ($1,$2,$3,$4,current_date) returning id`
	)

	err = r.db.QueryRow(dbRequest, input.Name, input.DOB, input.Address, input.Description).Scan(&output)
	if err != nil

	return
}

func (r *royce) GetSingleUser(ctx context.Context, id int32)(output models.SingleUserData, err error){

	return
}

func NewRoyce(db *pgx.Conn) Royce {
	return &royce{
		db:db,
	}
}

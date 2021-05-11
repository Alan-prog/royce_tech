package royce_tech_app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
	"net/http"
)

type Royce interface {
	Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output int, err tools.ErrorMessage)
	GetSingleUser(ctx context.Context, id int) (output models.SingleUserData, err tools.ErrorMessage)
	DeleteUser(ctx context.Context, input int) (err tools.ErrorMessage)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (err tools.ErrorMessage)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err tools.ErrorMessage)
}

type royce struct {
	db *pgx.Conn
}

func (r *royce) Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage) {
	output.Text = "service is okay"
	return
}

func (r *royce) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output int, err tools.ErrorMessage) {
	const (
		dbRequest = `insert into human_resources (name, dob, address, description, created_at) values 
    ($1,$2,$3,$4,current_date) returning id`
	)

	er := r.db.QueryRowEx(ctx, dbRequest, nil, input.Name, input.DOB, input.Address, input.Description).Scan(&output)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while adding new user to DB", http.StatusInternalServerError)
	}

	return
}

func (r *royce) GetSingleUser(ctx context.Context, id int) (output models.SingleUserData, err tools.ErrorMessage) {
	const (
		dbRequest = `select id, name, dob, address, description, created_at, updated_at
			from human_resources where id = $1 and visibility = true;`
	)

	var (
		preResponse models.SingleUserDataDbResponse
	)

	er := r.db.QueryRowEx(ctx, dbRequest, nil, id).Scan(&preResponse.ID, &preResponse.Name, &preResponse.DOB, &preResponse.Address, &preResponse.Description,
		&preResponse.CreateAt, &preResponse.UpdatedAt)
	if er != nil {
		if er.Error() == models.SqlNoRows {
			err = tools.NewErrorMessage(er, fmt.Sprintf("There is no user with id = %d", id), http.StatusBadRequest)
		} else {
			err = tools.NewErrorMessage(er, "Error while making request to db to get single user",
				http.StatusInternalServerError)
		}
	}
	return singleUserDataRepack(&preResponse), err
}

func (r *royce) DeleteUser(ctx context.Context, id int) (err tools.ErrorMessage) {
	const (
		dbRequest = `update human_resources set visibility = false where id = $1;`
	)

	_, er := r.db.ExecEx(ctx, dbRequest, nil, id)
	if er != nil {
		err = tools.NewErrorMessage(er, fmt.Sprintf("Error while deleting from db user with id: %d", id),
			http.StatusBadRequest)
	}

	return
}

func (r *royce) UpdateUser(ctx context.Context, input *models.UpdateUserData) (err tools.ErrorMessage) {
	const (
		dbRequest = `update human_resources set name = $1, dob = $2, description = $3, address = $4, 
			updated_at = current_date where id = $5;`
	)

	if input.Name == nil && input.Address == nil && input.Description == nil && input.DOB == nil {
		return
	}

	data, err := r.GetSingleUser(ctx, input.ID)
	if err != nil {
		return
	}

	if input.DOB != nil {
		data.DOB = input.DOB
	}
	if input.Description != nil {
		data.Description = *input.Description
	}
	if input.Address != nil {
		data.Address = *input.Address
	}
	if input.Name != nil {
		data.Name = *input.Name
	}

	_, er := r.db.ExecEx(ctx, dbRequest, nil, data.Name, data.DOB, data.Description, data.Address, data.ID)
	if er != nil {
		err = tools.NewErrorMessage(er, fmt.Sprintf("Error while updating id db user with id: %d", input.ID),
			http.StatusInternalServerError)
	}
	return
}

func (r *royce) GetAllUser(ctx context.Context) (response models.AllUsersData, err tools.ErrorMessage) {
	const (
		dbRequest = `select id,name,dob, address, description, created_at,updated_at 
			from human_resources where visibility = true order by id;`
	)

	rows, er := r.db.QueryEx(ctx, dbRequest, nil)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while getting all users", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var respElement models.SingleUserDataDbResponse

		if er = rows.Scan(&respElement.ID, &respElement.Name, &respElement.DOB,
			&respElement.Address, &respElement.Description, &respElement.CreateAt,
			&respElement.UpdatedAt); er != nil {
			err = tools.NewErrorMessage(er, "Error while scanning query rows in get all users",
				http.StatusInternalServerError)
			return
		}
		elem := singleUserDataRepack(&respElement)
		response = append(response, &elem)
	}

	return
}

func NewRoyce(db *pgx.Conn) Royce {
	return &royce{
		db: db,
	}
}

func singleUserDataRepack(input *models.SingleUserDataDbResponse) (output models.SingleUserData) {
	if input == nil {
		return
	}

	output.ID = input.ID
	output.Name = input.Name
	output.Address = input.Address
	output.Description = input.Description
	if input.DOB.Valid {
		output.DOB = new(string)
		*output.DOB = input.DOB.Time.String()
		if len(*output.DOB) > 10 {
			*output.DOB = (*output.DOB)[:10]
		}
	}
	if input.CreateAt.Valid {
		output.CreateAt = input.CreateAt.Time.String()
		if len(output.CreateAt) > 10 {
			output.CreateAt = output.CreateAt[:10]
		}
	}
	if input.UpdatedAt.Valid {
		output.UpdatedAt = new(string)
		*output.UpdatedAt = input.UpdatedAt.Time.String()
		if len(*output.UpdatedAt) > 10 {
			*output.UpdatedAt = (*output.UpdatedAt)[:10]
		}
	}
	return
}

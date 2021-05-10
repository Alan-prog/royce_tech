package royce_tech_app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"my_projects/royce_tech/pkg/models"
)

type Royce interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output int, err error)
	GetSingleUser(ctx context.Context, id int) (output models.SingleUserData, err error)
	DeleteUser(ctx context.Context, input int) (err error)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (err error)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err error)
}

type royce struct {
	db *pgx.Conn
}

func (r *royce) Alive() (output models.AliveResponse, err error) {
	output.Text = "service is okay"
	return
}

func (r *royce) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output int, err error) {
	const (
		dbRequest = `insert into human_resources (name, dob, address, description, created_at) values 
    ($1,$2,$3,$4,current_date) returning id`
	)

	err = r.db.QueryRowEx(ctx, dbRequest, nil, input.Name, input.DOB, input.Address, input.Description).Scan(&output)
	if err != nil {
		err = errors.Wrap(err, "error while making request to db")
	}

	return
}

func (r *royce) GetSingleUser(ctx context.Context, id int) (output models.SingleUserData, err error) {
	const (
		dbRequest = `select id, name, dob, address, description, created_at, updated_at
			from human_resources where id = $1 and visibility = true;`
	)

	var (
		preResponse models.SingleUserDataDbResponse
	)

	err = r.db.QueryRowEx(ctx, dbRequest, nil, id).Scan(&preResponse.ID, &preResponse.Name, &preResponse.DOB, &preResponse.Address, &preResponse.Description,
		&preResponse.CreateAt, &preResponse.UpdatedAt)
	if err != nil {
		if err.Error() == models.SqlNoRows {
			err = errors.New(fmt.Sprintf("there is no user with id = %d", id))
		} else {
			err = errors.Wrap(err, "error while making request to db")
		}
	}
	return singleUserDataRepack(&preResponse), err
}

func (r *royce) DeleteUser(ctx context.Context, id int) (err error) {
	const (
		dbRequest = `update human_resources set visibility = false where id = $1;`
	)

	_, err = r.db.ExecEx(ctx, dbRequest, nil, id)
	if err != nil {
		err = errors.Wrap(err, "error while deleting user from db")
	}

	return
}

func (r *royce) UpdateUser(ctx context.Context, input *models.UpdateUserData) (err error) {
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

	_, err = r.db.ExecEx(ctx, dbRequest, nil, data.Name, data.DOB, data.Description, data.Address, data.ID)
	if err != nil {
		err = errors.Wrap(err, "error while updating user info in db")
	}

	return
}

func (r *royce) GetAllUser(ctx context.Context) (response models.AllUsersData, err error) {
	const (
		dbRequest = `select id,name,dob, address, description, created_at,updated_at 
			from human_resources where visibility = true order by id;`
	)

	rows, err := r.db.QueryEx(ctx, dbRequest, nil)
	if err != nil {
		err = errors.Wrap(err, "error while getting all users")
		return
	}

	for rows.Next() {
		var respElement models.SingleUserDataDbResponse

		if err = rows.Scan(&respElement.ID, &respElement.Name, &respElement.DOB,
			&respElement.Address, &respElement.Description, &respElement.CreateAt,
			&respElement.UpdatedAt); err != nil {
			err = errors.Wrap(err, "error while scanning query rows")
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

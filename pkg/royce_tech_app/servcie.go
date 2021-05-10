package royce_tech_app

import (
	"github.com/jackc/pgx"
	"my_projects/royce_tech/pkg/models"
)

type Royce interface {
	Alive() (output models.AliveResponse, err error)
}

type royce struct {
	db *pgx.Conn
}

func (r *royce) Alive() (output models.AliveResponse, err error) {
	output.Text = "service is okay"
	return
}

func NewRoyce(db *pgx.Conn) Royce {
	return &royce{
		db:db,
	}
}

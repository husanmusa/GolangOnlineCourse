package repository

import (
	"database/sql"

	"lesson28/models"

	"github.com/google/uuid"
)

type UserRepo struct {
	DB *sql.DB
	//db *mongo.Collection
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db} //db: db.Collecion("user")
}

func (u *UserRepo) CreateUser(user models.User) error {
	id := uuid.NewString()

	_, err := u.DB.Exec(`INSERT INTO users (id, name, role, email)
	 VALUES ($1, $2, $3, $4)`,
		id, user.Name, user.Role, user.Email)
	if err != nil {
		return err
	}

	return nil
}

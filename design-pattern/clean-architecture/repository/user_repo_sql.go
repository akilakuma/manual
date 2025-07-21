package repository

import (
	"clean-architecture/domain"
	"database/sql"
)

type UserRepoSQL struct {
	DB *sql.DB
}

func NewUserRepoSQL(db *sql.DB) *UserRepoSQL {
	return &UserRepoSQL{
		DB: db,
	}
}

func (r *UserRepoSQL) Save(user *domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, email) VALUES (?, ?)", user.ID, user.Email)
	return err
}

func (r *UserRepoSQL) FindByEmail(email string) (*domain.User, error) {
	result, err := r.DB.Query("SELECT *  FROM WHERE email = ?", email)
	if result != nil {
		for result.Next() {
			var user domain.User
			scanErr := result.Scan(&user.ID, &user.Email)
			return &domain.User{
				ID: user.Email,
			}, scanErr
		}
	}
	return nil, err
}

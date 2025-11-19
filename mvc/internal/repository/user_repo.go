package repository

import (
	"database/sql"
	"mvc/internal/model"
)

// tao interface
type UserRepository interface {
	CreateUser(user *model.Users) (*model.Users, error)
	GetUserByID(id int) (*model.Users, error)
	GetAllUser() ([]model.Users, error)
	GetUserByUsername(username string) (bool, error)
	GetUserByEmail(email string) (bool, error)
	UpdateUserByID(user *model.Users) (*model.Users, error)
	DeleteUserByID(id int) error
}

// tao struct trien khai UserRepository
type UserRepo struct {
	DB *sql.DB
}

// NewUserRepo la khoi tao 1 repository moi
func NewUserRepo(db *sql.DB) UserRepository {
	return &UserRepo{
		DB: db,
	}
}

// Create
func (r *UserRepo) CreateUser(user *model.Users) (*model.Users, error) {
	result, err := r.DB.Exec("insert into users (username,email,password,age) values (?,?,?,?)", user.Username, user.Email, user.PasswordHard, user.Age)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetUserByID(int(id))
}

// Get 1 by ID
func (r *UserRepo) GetUserByID(id int) (*model.Users, error) {
	row := r.DB.QueryRow("select id,username,email,age,created_at,updated_at from users where id = ?", id)
	var person model.Users
	err := row.Scan(&person.ID, &person.Username, &person.Email, &person.Age, &person.CreatedAt, &person.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// Get All
func (r *UserRepo) GetAllUser() ([]model.Users, error) {
	rows, err := r.DB.Query("select id,username,email,age,created_at,updated_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []model.Users
	for rows.Next() {
		var person model.Users
		if err := rows.Scan(&person.ID, &person.Username, &person.Email, &person.Age, &person.CreatedAt, &person.UpdatedAt); err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}

func (r *UserRepo) GetUserByUsername(username string) (bool, error) {
	var exists int
	err := r.DB.QueryRow("select 1 from users where username = ? limit 1", username).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *UserRepo) GetUserByEmail(email string) (bool, error) {
	var exists int
	err := r.DB.QueryRow("Select 1 from users where email=? limit 1", email).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
func (r *UserRepo) UpdateUserByID(user *model.Users) (*model.Users, error) {
	_, err := r.DB.Exec("update users set username=?, email=?, password=?, age=? where id=?", user.Username, user.Email, user.PasswordHard, user.Age, user.ID)
	if err != nil {
		return nil, err
	}
	return r.GetUserByID(user.ID)
}

func (r *UserRepo) DeleteUserByID(id int) error {
	_, err := r.DB.Exec("delete from users where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

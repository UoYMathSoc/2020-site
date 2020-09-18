package models

import (
	"context"
	"fmt"
	"github.com/UoYMathSoc/2020-site/database"
	"golang.org/x/crypto/bcrypt"
)

// UserModel is used when accessing the site's users
type UserModel struct {
	Model
}

// NewUserModel returns a new UserModel with access to the database
func NewUserModel(q *database.Queries) *UserModel {
	return &UserModel{Model{q}}
}

func (m *UserModel) Get(id int32) (*database.User, []database.Position, error) {
	user, err := m.querier.GetUser(context.Background(), id)
	if err != nil {
		return &user, nil, err
	}

	positions, err := m.querier.GetUserPositions(context.Background(), id)
	if err != nil {
		return &user, positions, err
	}

	return &user, positions, nil
}

func (m *UserModel) Register(username string, password string) error {
	id, err := m.querier.CreateUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("could not generate hash for provided password: %w", err)
	}
	err = m.querier.SetUsersPass(context.Background(), database.SetUsersPassParams{
		ID:       id,
		Password: string(hashedPass),
	})
	if err != nil {
		return fmt.Errorf("could not set provided password: %w", err)
	}
	return nil
}

func (m *UserModel) Validate(username string, password string) (int32, error) {
	id, err := m.querier.FindUserUsername(context.Background(), username)
	if err != nil {
		return id, fmt.Errorf("could not find user: %w", err)
	}
	userPass, err := m.querier.GetUsersPass(context.Background(), id)
	if err != nil {
		return id, fmt.Errorf("could not find the password for specified user: %w", err)
	}
	return id, bcrypt.CompareHashAndPassword([]byte(userPass.Password), []byte(password))
}

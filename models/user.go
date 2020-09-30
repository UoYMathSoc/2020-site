package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/UoYMathSoc/2020-site/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserModel is used when accessing the site's users
type UserModel struct {
	Model
}

type User struct {
	ID       int
	Username string
	Name     string
	Bio      string
}

type Position struct {
	ID       int
	Name     string
	Alias    string
	FromDate time.Time
	TillDate time.Time
}

func (s *Session) GetUser(id int) (*User, error) {
	user, err := s.querier.GetUser(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       int(user.ID),
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio.String,
	}, nil
}

func (s *Session) GetUserPositions(id int) ([]Position, error) {
	positions, err := s.querier.GetUserPositions(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	var result []Position
	for _, position := range positions {
		p, err := s.querier.GetPosition(context.Background(), position.CommitteeID)
		if err != nil {
			break
		}
		position := Position{
			ID:       int(p.ID),
			Name:     p.Name.String,
			Alias:    p.Alias,
			FromDate: position.FromDate,
			TillDate: position.TillDate.Time,
		}
		result = append(result, position)
	}
	return result, nil
}

func (m *UserModel) Get(id int32) (*database.User, *Position, error) {
	user, err := m.querier.GetUser(context.Background(), id)
	if err != nil {
		return &user, nil, err
	}

	positions, err := m.querier.GetUserPositions(context.Background(), id)
	if err != nil {
		return &user, nil, err
	}

	userPositions, err := m.querier.GetUserPosition

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

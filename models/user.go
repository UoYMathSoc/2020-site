package models

import (
	"context"
	"fmt"
	"time"

	"github.com/UoYMathSoc/2020-site/database"
	"golang.org/x/crypto/bcrypt"
)

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

type UserStore struct {
	querier database.Querier
}

func NewUserStore(q database.Querier) UserStore {
	return UserStore{q}
}

func (us *UserStore) Get(id int) (*User, error) {
	user, err := us.querier.GetUser(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	sanitiseUser(&user)
	return &User{
		ID:       int(user.ID),
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio.String,
	}, nil
}

func (us *UserStore) GetPositions(id int) ([]Position, error) {
	positions, err := us.querier.GetUserPositions(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	var result []Position
	for _, position := range positions {
		p, err := us.querier.GetPosition(context.Background(), position.CommitteeID)
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

func (us *UserStore) Validate(username, password string) (int, error) {
	id, err := us.querier.FindUserUsername(context.Background(), username)
	if err != nil {
		return -1, fmt.Errorf("could not find user: %w", err)
	}
	creds, err := us.querier.GetUsersPass(context.Background(), id)
	err = bcrypt.CompareHashAndPassword([]byte(creds.Password), []byte(password))
	if err != nil {
		return -1, fmt.Errorf("could not validate user: %w", err)
	}
	return int(id), nil
}

func (us *UserStore) Create(user *User) error {
	return nil //TODO: create this function
}

func sanitiseUser(user *database.User) {
	if !user.Bio.Valid {
		user.Bio.String = ""
		user.Bio.Valid = true
	}
}

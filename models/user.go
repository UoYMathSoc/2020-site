package models

import (
	"context"
	"github.com/UoYMathSoc/2020-site/database"
	"time"
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

func (s *Session) GetUser(id int) (*User, error) {
	user, err := s.querier.GetUser(context.Background(), int32(id))
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

func sanitiseUser(user *database.User) {
	if !user.Bio.Valid {
		user.Bio.String = ""
	}
}

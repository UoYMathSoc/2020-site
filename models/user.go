package models

import (
	"context"
	"fmt"
	"sort"
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
	ordering int
	Name     string
	Alias    string
	FromDate time.Time
	TillDate time.Time
}

type Positions []Position

func (p Positions) Len() int              { return len(p) }
func (p Positions) Swap(i, j int)         { p[i], p[j] = p[j], p[i] }
func (p Positions) Before(i, j int) bool  { return p[i].ordering < p[j].ordering }
func (p Positions) Earlier(i, j int) bool { return p[i].FromDate.After(p[j].FromDate) }

func (p Positions) ByOrder() Positions {
	sort.Slice(p, p.Before)
	return p
}

func (p Positions) ByDate() Positions {
	sort.Slice(p, p.Earlier)
	return p
}

type Officer struct {
	Position Position
	User     User
}

type Committee []Officer

type UserStore struct {
	querier database.Querier
}

func NewUserStore(q database.Querier) UserStore {
	return UserStore{q}
}

func (us *UserStore) Get(id int) (User, Positions, error) {
	user, err := us.getUser(id)
	if err != nil {
		return User{}, nil, err
	}
	positions, err := us.getPositions(id)
	if err != nil {
		return User{}, nil, err
	}
	return user, positions, nil
}

func (us *UserStore) GetByUsername(username string) (User, Positions, error) {
	id, err := us.querier.FindUserUsername(context.Background(), username)
	if err != nil {
		return User{}, nil, err
	}
	return us.Get(int(id))
}

func (us *UserStore) getUser(id int) (User, error) {
	user, err := us.querier.GetUser(context.Background(), int32(id))
	if err != nil {
		return User{}, err
	}

	sanitiseUser(&user)
	return User{
		ID:       int(user.ID),
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio.String,
	}, nil
}

func (us *UserStore) getPositions(id int) (Positions, error) {
	positions, err := us.querier.GetUserPositions(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	var result Positions
	for _, position := range positions {
		p, err := us.querier.GetPosition(context.Background(), position.CommitteeID)
		if err != nil {
			return nil, err
		}
		position := Position{
			ID:       int(p.ID),
			ordering: int(p.Ordering),
			Name:     p.Name.String,
			Alias:    p.Alias,
			FromDate: position.FromDate,
			TillDate: position.TillDate.Time,
		}
		result = append(result, position)
	}

	return result.ByOrder(), nil
}

func (us *UserStore) GetCommittee() (executive Committee, committee Committee, err error) {
	positions, err := us.querier.ListPositions(context.Background())
	if err != nil {
		return nil, nil, err
	}

	for _, position := range positions {
		users, err := us.querier.GetPositionUsers(context.Background(), position.ID)
		if err != nil {
			return nil, nil, err
		}
		for _, user := range users {
			if user.Outgoing || (user.TillDate.Time.Before(time.Now()) && user.TillDate.Valid) {
				continue
			}
			id := int(user.UserID)
			u, _, err := us.Get(id)
			if err != nil {
				return nil, nil, err
			}
			c := Officer{
				Position: Position{
					ID:       int(position.ID),
					Name:     position.Name.String,
					Alias:    position.Alias,
					FromDate: user.FromDate,
					TillDate: user.TillDate.Time,
				},
				User: u,
			}
			if position.Executive.Bool {
				executive = append(executive, c)
			} else {
				committee = append(committee, c)
			}
		}
	}
	return executive, committee, nil
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

package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type SessionStore struct {
	sessions.CookieStore
	querier     database.Querier
	oauthConfGl oauth2.Config
}

func NewSessionStore(c *structs.Config, querier database.Querier) *SessionStore {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  c.PageContext.FullURL + "/callback-gl",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	sessions := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	return &SessionStore{
		querier:     querier,
		oauthConfGl: config,
		CookieStore: *sessions,
	}
}

func (ss *SessionStore) GenToken(code string) (*oauth2.Token, error) {
	return ss.oauthConfGl.Exchange(context.Background(), code)
}

func (ss *SessionStore) AuthCodeURL(state string) string {
	return ss.oauthConfGl.AuthCodeURL(state)
}

func (ss *SessionStore) GetUsername(token *oauth2.Token) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var data map[string]interface{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	username := strings.TrimSuffix(data["email"].(string), "@"+data["hd"].(string))
	return username, nil
}

func NewSessionFromConfig(db structs.Database) *SessionStore {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &SessionStore{querier: database.New(conn)}
}

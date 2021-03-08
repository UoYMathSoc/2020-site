package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/utils"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	config := &structs.Config{}
	_, err := toml.DecodeFile("../../config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	q := utils.NewDBFromConfig(config.Database)
	us := models.NewUserStore(q)

	var username, password string
	for {
		username, password, err = credentials()
		if err == nil {
			break
		}
		log.Println(err)
		log.Println("Please try again")
	}

	id, err := us.Create(username, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User", username, "successfully created with ID:", id)
}

func credentials() (username, password string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err = reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}

	password = string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

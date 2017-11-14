package main

import (
	"fmt"

	"github.com/zainulbr/go-skype/skype"
)

func main() {
	client := skype.NewClient("client_id", "client_secret", "scope_url")
	_, err := client.Authorization.Authorize()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	fmt.Printf("Bot token: %s \n", client.Token)
}

package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Access auth service from the default app
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	uid := "DMBaKvVXTmOtwO0NFsAY0YfMRM63"
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %#v\n", u)

	//	token, err := client.CustomToken(ctx, uid)
	claims := map[string]interface{}{
		"premiumAccount": true,
	}
	token, err := client.CustomTokenWithClaims(ctx, uid, claims)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Got custom token: %#v\n", token)
}

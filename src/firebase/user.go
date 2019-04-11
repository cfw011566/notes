package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
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

	/*
		uid := "cf.wang@metropia.com"
		u, err := client.GetUserByEmail(ctx, uid)
		if err != nil {
			log.Fatalf("error getting user %s: %v\n", uid, err)
		}
		log.Printf("Successfully fetched user data: %#v\n", u)
		log.Printf("UserInfo: %#v\n", u.UserInfo)

			params := (&auth.UserToCreate{}).
				Email("cf.wang@metropia.com").
				EmailVerified(false).
				PhoneNumber("+886932869917").
				Password("Ab123456").
				DisplayName("CF Wang").
				Disabled(false)
			u, err := client.CreateUser(ctx, params)
			if err != nil {
				log.Fatalf("error creating user: %v\n", err)
			}
			log.Printf("Successfully created user: %v\n", u)
	*/

	uid := "DMBaKvVXTmOtwO0NFsAY0YfMRM63"
	params := (&auth.UserToUpdate{}).
		Password("aB123456")
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
}

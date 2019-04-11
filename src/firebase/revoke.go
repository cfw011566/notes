package main

import (
	"context"
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

	uid := "vjnlNudfVdcKK3MguJDj0qT0hZG2"
	/*
		if err := client.RevokeRefreshTokens(ctx, uid); err != nil {
			log.Fatalf("error revoking tokens for user: %v, %v\n", uid, err)
		}
	*/
	// accessing the user's TokenValidAfter
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("user = %v\n", u)
	log.Printf("user meta = %v\n", u.UserMetadata)
	timestamp := u.TokensValidAfterMillis / 1000
	log.Printf("the refresh tokens were revoked at: %d (UTC seconds) ", timestamp)

	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdkMmY5ZjNmYjgzZDYzMzc0OTdiNmY3Y2QyY2ZmNGRmYTVjMmU4YjgiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vcXVvdGVzLTVkYzg4IiwiYXVkIjoicXVvdGVzLTVkYzg4IiwiYXV0aF90aW1lIjoxNTU0OTUxNDExLCJ1c2VyX2lkIjoidmpubE51ZGZWZGNLSzNNZ3VKRGowcVQwaFpHMiIsInN1YiI6InZqbmxOdWRmVmRjS0szTWd1SkRqMHFUMGhaRzIiLCJpYXQiOjE1NTQ5NTE0MTEsImV4cCI6MTU1NDk1NTAxMSwiZW1haWwiOiJjZi53YW5nKzFAbWV0cm9waWEuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImNmLndhbmcrMUBtZXRyb3BpYS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.MgMlqXveztogutY-SjjlwrBT65DJ3cD245g5gkfWhHtgEhkL6oIXo1AC_fSz_uQ7yxZJzL6lS2kiE0FNZE_uCQyGDv7-ksDMvhtBF9C3DGLGqfavD3CJZsuDKT2nUti0CS6y2gqtZ2ODuizuWGb0HOU1qR-fTGmzk9NCEnLIsjTpAqxO-Iv9ZjomtV8URE1SY_TzYhGLwxdvWGV2kuYxmAvCB7q2CXN7Sx0_teRupHaXVogosjc39WoZ1NzDEocXkCmOFH1GMrxOgAACVSSanCZma9IZVswx5k-YVhkpe--HM7lGR__6AV28Kazr0xg5jV_sO6UKtJJLgszda6S1ag"
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("VerifyIDTokenAndCheckRevoked error = %v\n", err)
		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Inform the user to reauthenticate or signOut() the user.
		} else {
			// Token is invalid
		}
	}
	log.Printf("Verified ID token: %v\n", token)

	token, err = client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		log.Printf("VerifyIDTokenAndCheckRevoked error = %v\n", err)
		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Inform the user to reauthenticate or signOut() the user.
		} else {
			// Token is invalid
		}
	}
	log.Printf("Verified ID token and Check revoked: %v\n", token)
}

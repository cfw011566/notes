package main

import (
	"context"
	"log"
	"time"

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

	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdkMmY5ZjNmYjgzZDYzMzc0OTdiNmY3Y2QyY2ZmNGRmYTVjMmU4YjgiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiQ0YgV2FuZyIsInByZW1pdW0iOmZhbHNlLCJzY29yZSI6MCwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL3F1b3Rlcy01ZGM4OCIsImF1ZCI6InF1b3Rlcy01ZGM4OCIsImF1dGhfdGltZSI6MTU1NTA0NDUzMCwidXNlcl9pZCI6IkRNQmFLdlZYVG1PdHdPME5Gc0FZMFlmTVJNNjMiLCJzdWIiOiJETUJhS3ZWWFRtT3R3TzBORnNBWTBZZk1STTYzIiwiaWF0IjoxNTU1MDQ0NTMwLCJleHAiOjE1NTUwNDgxMzAsImVtYWlsIjoiY2Yud2FuZ0BtZXRyb3BpYS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInBob25lX251bWJlciI6Iis4ODY5MzI4Njk5MTciLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIis4ODY5MzI4Njk5MTciXSwiZW1haWwiOlsiY2Yud2FuZ0BtZXRyb3BpYS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.uga_FrUDRc84bOtgJU9J13NydZ0EwEO4EKZfSEEA7ktmLSu2EZ-8VOHpYx8U2Jif8jt3S-CDZkWC1Y7S4Id6b9SPxdBxrrOQoRGTMIvSiuS8FF8jRVZN2cjXp_klLvCqEFbQ7dsXYrZ17UCCi1n9RauHOruKOfPLYGE74nJj6p-eF-5dI5ZBCRukPingjUymYVfZS7H8ozhhYcDYU_guXP5Tug2TlChVGNskvg-QMU9ory-Ex0LHhbf4Va7a0xB67pIKoFnCYVQC4jclcVseMLqXohjQKdf-S6hAoz_gcLR_e-AMnk6TBhT9Rcd1GYp5FLBzzHu2U3dwqnaA4W3a1w"

	expire, _ := time.ParseDuration("12h34m56s")
	log.Printf("expire: %v\n", expire)
	cookie, err := client.SessionCookie(ctx, idToken, expire)
	if err != nil {
		log.Printf("SessionCookie error = %v\n", err)
	}
	log.Printf("cookie: %v\n", cookie)

	token, err := client.VerifySessionCookie(ctx, cookie)
	if err != nil {
		log.Printf("VerifySessionCookie error = %v\n", err)
	}
	log.Printf("session token: %v\n", token)

	oldCookie := "eyJhbGciOiJSUzI1NiIsImtpZCI6InNrSUJOZyJ9.eyJpc3MiOiJodHRwczovL3Nlc3Npb24uZmlyZWJhc2UuZ29vZ2xlLmNvbS9xdW90ZXMtNWRjODgiLCJuYW1lIjoiQ0YgV2FuZyIsInByZW1pdW0iOmZhbHNlLCJzY29yZSI6MCwiYXVkIjoicXVvdGVzLTVkYzg4IiwiYXV0aF90aW1lIjoxNTU1MDQ0NTMwLCJ1c2VyX2lkIjoiRE1CYUt2VlhUbU90d08wTkZzQVkwWWZNUk02MyIsInN1YiI6IkRNQmFLdlZYVG1PdHdPME5Gc0FZMFlmTVJNNjMiLCJpYXQiOjE1NTUwNDUyOTYsImV4cCI6MTU1NTA1MjY4MCwiZW1haWwiOiJjZi53YW5nQG1ldHJvcGlhLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicGhvbmVfbnVtYmVyIjoiKzg4NjkzMjg2OTkxNyIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsicGhvbmUiOlsiKzg4NjkzMjg2OTkxNyJdLCJlbWFpbCI6WyJjZi53YW5nQG1ldHJvcGlhLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.dvH_CSwwIP2l57K8oDnQAzlSWrnlMX5AZ3k51d2XkkwnZeXn28bhWXbJhUIVoXoL1a5QG5XmqkpEYcoPM_tzva2OP6PYa1KKYumNaL66uX_BtjAcFyTxKYMfooXdIwWt7bB-KOzTZKJfp8J0UFNYUZ7t-R5buCij9eYVIPYAphIdFNGbMpFf3783tTkTgkmjukQIuPPslaq-_BeZVKkfRumdwVVjpq59UQWlYf6dRWw5UjizmQb-jehPJ0OBIRGuKDKaYH-4y9VsEjUKddBSJFuEGtlc-AW9W8IGyHwYvdDmA_SvK6AImDuENgaho-uPtsIcIk6WOEGKsUpCem8Tgg"
	token, err = client.VerifySessionCookie(ctx, oldCookie)
	if err != nil {
		log.Printf("VerifySessionCookie error = %v\n", err)
	}
	log.Printf("session token: %v\n", token)
}

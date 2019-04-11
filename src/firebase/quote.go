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

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// quote := getQuote()
	quote := map[string]interface{}{
		"author": "golang on Mac",
		"quote":  "Go for Firebase Admin SDK ...",
	}
	log.Print(quote)
	result, err := client.Collection("sampleData").Doc("inspiration").Set(ctx, quote)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)

	dsnap, err := client.Collection("sampleData").Doc("inspiration").Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	m := dsnap.Data()
	// Do somthing with data!
	fmt.Println(m)

	defer client.Close()
}

package main

import (
	"fmt"
	"github.com/cameronstanley/go-reddit"
)

func main() {
	// Create a new authenticator with your app's client ID, secret, and redirect URI
	// A random string representing state and a list of requested OAuth scopes are required
	authenticator := reddit.NewAuthenticator("g1uQYjsE3h1_GA", "59pYu2YpxOAS3pxO1_-Eqa8QNxo", "http://127.0.0.1",
		"Debian:badboyd app:0.1.0 (by /u/<reddit protrandat>)", "code", reddit.ScopeRead, reddit.ScopeIdentity)

	// Instruct your user to visit the URL retrieved from GetAuthenticationURL in their web browser
	url := authenticator.GetAuthenticationURL()
	fmt.Printf("Please proceed to %s\n", url)

	// After the user grants permission for your client, they will be redirected to the supplied redirect_uri with a code and state as URL parameters
	// Gather these values from the user on the console
	// Note: this can be automated by having a web server listen on the redirect_uri and parsing the state and code params
	fmt.Print("Enter state: ")
	var state string
	fmt.Scanln(&state)
	//
	fmt.Print("Enter code: ")
	var code string
	fmt.Scanln(&code)

	// Exchange the code for an access token
	token, err := authenticator.GetToken(state, code)
	if err != nil {
		panic(err)
	}
	// Create a new client using the access token and a user agent string to identify your application
	client := authenticator.GetAuthClient(token)
	acc, err := client.GetMe()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *acc)

	sr, err := client.GetGoldSubreddits()
	if err != nil {
		panic(err)
	}

	for _, r := range sr {
		fmt.Printf("subreddit %+v", *r)
	}
}

package main

import "fmt"
import "os"
import "log"


import (
	"github.com/joho/godotenv"
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

type Credentials struct {
	//api keys for twitter bot
	ConsumerKey		  string 
	ConsumerSecret	  string
	AccessToken		  string
	AccessTokenSecret string

}


func main() {
	fmt.Println("Go Twitter Bot")
	//load env creds 
	err := godotenv.Load("keys.env")
	if err != nil {
		fmt.Printf("ERROR GETTING FILE ")
	}
	creds := Credentials{
        AccessToken:       os.Getenv("ACCESS_TOKEN"),
        AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
        ConsumerKey:       os.Getenv("CONSUMER_KEY"),
        ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}
	fmt.Printf("%+v\n", creds)
	client, err := getClients(&creds)
	if err != nil {
		log.Println("issue with twitter client")
		log.Println(err); 
	}
	fmt.Printf("%+v\n", client)
}

//passing api keys 
func getClients(creds *Credentials) (*twitter.Client, error){
	//pass the api key 
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// pass in the acsess token 
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)	
	//verifying creds
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus: twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	//we can retrieve the user and verify if the credentials
    // we have used successfully allow us to log in!
    user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        return nil, err
    }
	// returns the information of the account 
	log.Printf("User's ACCOUNT:\n%+v\n", user)
	tweet, resp, err := client.Statuses.Update("commit change & feed Milo m8", nil)
	if err != nil {
    log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
    return client, nil
}
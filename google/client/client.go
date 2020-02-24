package client

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const credentialsFile = "credentials.json"
const scopes = "https://www.googleapis.com/auth/spreadsheets.readonly"

// Retrieve a token, saves the token, then returns the generated client.
func Get(tokenFile string) *http.Client {
	c, err := config()
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(c)
		saveToken(tokenFile, tok)
	}

	return c.Client(context.Background(), tok)
}

func GetFromToken(tok *oauth2.Token) (*http.Client, error) {
	c, err := config()
	if err != nil {
		return nil, err
	}

	return c.Client(context.Background(), tok), nil
}

func config() (*oauth2.Config, error) {
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(b, scopes, drive.DriveMetadataReadonlyScope)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	_ = json.NewEncoder(f).Encode(token)
}

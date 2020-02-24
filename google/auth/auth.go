package auth

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"log"
)

const credentialsFile = "credentials.json"
const scopes = "https://www.googleapis.com/auth/spreadsheets.readonly"

func GetAuthURL() string {
	return config().AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func Authenticate(authCode string) (*oauth2.Token, error) {
	tok, err := config().Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func config() *oauth2.Config {
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, scopes, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config
}

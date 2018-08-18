package platforms

import (
	"github.com/matrix-org/gomatrix"
)

type Riot struct {
	client *gomatrix.Client
}

func NewRiot(conf map[string]string) (Platform, error) {
	url, ok := conf["url"]
	if !ok {
		url = "https://matrix.org"
	}

	userID, ok := conf["user"]
	if !ok {
		return nil, fieldNotFound("user")
	}

	accessToken := conf["token"]
	if !ok {
		return nil, fieldNotFound("token")
	}

	client, err := gomatrix.NewClient(url, userID, accessToken)
	if err != nil {
		return nil, err
	}

	r := Riot{
		client: client,
	}

	return r, nil
}

func (r Riot) Send(room, msg string) error {
	_, err := r.client.SendText(room, msg)
	return err
}

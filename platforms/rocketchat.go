package platforms

import (
	"fmt"
	"net/url"

	"github.com/detached/gorocket/api"
	"github.com/detached/gorocket/rest"
)

type RocketChat struct {
	client *rest.Client
}

func NewRocketChat(conf map[string]string) (Platform, error) {
	r := RocketChat{}

	endpoint, ok := conf["url"]
	if !ok {
		return nil, fieldNotFound("url")
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	email, ok := conf["email"]
	if !ok {
		return nil, fieldNotFound("password")
	}

	password, ok := conf["password"]
	if !ok {
		return nil, fieldNotFound("password")
	}

	r.client = rest.NewClient(u.Hostname(), u.Port(), false, false)
	if err := r.client.Login(api.UserCredentials{
		Email:    email,
		Password: password,
	}); err != nil {
		return nil, err
	}

	return r, nil
}

func (r RocketChat) GetChannel(room string) (*api.Channel, error) {
	publicChannels, err := r.client.GetPublicChannels()
	if err != nil {
		return nil, fmt.Errorf("Failed to query publicChannels: %s", err)
	}

	joinedChannels, err := r.client.GetJoinedChannels()
	if err != nil {
		return nil, fmt.Errorf("Failed to query joinedChannels: %s", err)
	}

	channels := append(publicChannels, joinedChannels...)
	for _, channel := range channels {
		if channel.Name == room {
			return &channel, nil
		}
	}

	return nil, fmt.Errorf("Room not found: %s", room)
}

func (r RocketChat) Send(room, msg string) error {
	channel, err := r.GetChannel(room)
	if err != nil {
		return err
	}

	return r.client.Send(channel, msg)
}

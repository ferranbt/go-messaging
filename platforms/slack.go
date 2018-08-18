package platforms

import "github.com/nlopes/slack"

type Slack struct {
	client *slack.Client
}

func NewSlack(conf map[string]string) (Platform, error) {

	token, ok := conf["token"]
	if !ok {
		return nil, fieldNotFound("token")
	}

	api := slack.New(token)
	if _, err := api.AuthTest(); err != nil {
		return nil, err
	}

	r := Slack{
		client: api,
	}

	return r, nil
}

func (r Slack) Send(room, msg string) error {
	_, _, err := r.client.PostMessage(room, msg, slack.PostMessageParameters{})
	return err
}

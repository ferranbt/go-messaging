package platforms

import (
	"fmt"
	"regexp"
)

type Config map[string]string

var (
	re = regexp.MustCompile("([[:graph:]]*)=([[:graph:]]*)")
)

func parse(cfg string) (Config, error) {
	config := Config{}
	for _, match := range re.FindAllStringSubmatch(cfg, -1) {
		key, value := match[1], match[2]
		if _, ok := config[key]; ok {
			return config, fmt.Errorf("Key '%s' repeated", key)
		}
		config[key] = value
	}

	return config, nil
}

type Platform interface {
	Send(room, msg string) error
}

type factory func(attr map[string]string) (Platform, error)

var Platforms = map[string]factory{
	"riot":       NewRiot,
	"rocketchat": NewRocketChat,
}

func NewPlatform(conf string) (Platform, error) {
	args, err := parse(conf)
	if err != nil {
		return nil, err
	}

	name, ok := args["platform"]
	if !ok {
		return nil, fmt.Errorf("Platform name not found on conf")
	}

	platform, ok := Platforms[name]
	if !ok {
		return nil, fmt.Errorf("Platform %s not found", name)
	}

	return platform(args)
}

func fieldNotFound(field string) error {
	return fmt.Errorf("field '%s' not found", field)
}

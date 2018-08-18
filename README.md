
# Go-Messaging

`go-messaging` is a Go library to send messages to different messaging platforms via the same interface. The configuration for each platform is provided as a list of key/value tuples. This library has been inspired by [go-discover](https://github.com/hashicorp/go-discover).

## Usage

```
conf := "platform=riot url=<url> user=<user> token=<token>"
platform, err := platforms.NewPlatform(conf)
if err != nil {
	panic(err)
}

platform.Send("<room/channel/user>", "Hello World")
```

## Platforms

The following messaging platforms are available in go-messaging

### Riot

```
platform=riot url=<url> user=<user> token=<token>
```

### RocketChat

```
platform=rocketchat url=<url> email=<email> password=<password>
```

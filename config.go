package idcloudhost

type Config struct {

	// API of idcloudhost.com uses tokens to allow access to
	// the API. You can register a new API token at our user
	// interface.
	// https://app.idcloudhost.com/
	ApiKey string `yaml:"api_key"`
}

const BASEAPI = "https://api.idcloudhost.com"

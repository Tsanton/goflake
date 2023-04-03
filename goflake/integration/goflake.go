package integration

import (
	"log"

	g "github.com/tsanton/goflake-client/goflake"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Goflake() *g.GoflakeClient {
	cfg := g.SnowflakeDsnConfig{
		Account:              u.GetEnv("SNOWFLAKE_ACCOUNT", ""),
		User:                 u.GetEnv("SNOWFLAKE_USER", ""),
		Password:             u.GetEnv("SNOWFLAKE_PASSWORD", ""),
		BrowserAuth:          u.GetEnv("SNOWFLAKE_BROWSER_AUTH", false),
		PrivateKeyPath:       u.GetEnv("SNOWFLAKE_PRIVATEKEY_PATH", ""),
		PrivateKey:           u.GetEnv("SNOWFLAKE_PRIVATEKEY", ""),
		PrivateKeyPassphrase: u.GetEnv("SNOWFLAKE_PRIVATEKEY_PASSPHRASE", ""),
		OauthAccessToken:     u.GetEnv("SNOWFLAKE_OAUTH_ACCESS_TOKEN", ""),
		Region:               u.GetEnv("SNOWFLAKE_REGION", ""),
		Role:                 u.GetEnv("SNOWFLAKE_ROLE", ""),
		OauthRefreshToken:    u.GetEnv("SNOWFLAKE_OAUTH_REFRESH_TOKEN", ""),
		OauthClientID:        u.GetEnv("SNOWFLAKE_OAUTH_CLIENT_ID", ""),
		OauthClientSecret:    u.GetEnv("SNOWFLAKE_OAUTH_CLIENT_SECRET", ""),
		OauthEndpoint:        u.GetEnv("SNOWFLAKE_OAUTH_ENDPOINT", ""),
		OauthRedirectURL:     u.GetEnv("SNOWFLAKE_OAUTH_REDIRECT_URL", ""),
		Host:                 u.GetEnv("SNOWFLAKE_HOST", ""),
		Protocol:             u.GetEnv("SNOWFLAKE_PROTOCOL", "https"),
		Port:                 u.GetEnv("SNOWFLAKE_PORT", 443),
		Warehouse:            u.GetEnv("SNOWFLAKE_WAREHOUSE", ""),
	}
	cli, err := g.NewClient(cfg)
	if err != nil {
		log.Panicf("unable to create client: %s", err.Error())
	}

	return cli
}

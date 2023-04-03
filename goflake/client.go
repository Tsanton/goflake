package goflake

import (
	"fmt"
	"log"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider"
	"github.com/jmoiron/sqlx"
	_ "github.com/snowflakedb/gosnowflake"
)

type GoflakeClient struct {
	db *sqlx.DB
}

// Snowflake Data Source Name Config
type SnowflakeDsnConfig struct {
	Account              string
	User                 string
	Password             string
	BrowserAuth          bool
	PrivateKeyPath       string
	PrivateKey           string
	PrivateKeyPassphrase string
	OauthAccessToken     string
	Region               string
	Role                 string
	OauthRefreshToken    string
	OauthClientID        string
	OauthClientSecret    string
	OauthEndpoint        string
	OauthRedirectURL     string
	Host                 string
	Protocol             string
	Port                 int
	Warehouse            string
}

func NewClient(cfg SnowflakeDsnConfig) (*GoflakeClient, error) {
	db, err := Open(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not build dsn for snowflake connection err = %w", err)
	}

	return &GoflakeClient{db: db}, nil
}

func Open(cfg SnowflakeDsnConfig) (*sqlx.DB, error) {

	if cfg.OauthRefreshToken != "" {
		accessToken, err := provider.GetOauthAccessToken(cfg.OauthEndpoint, cfg.OauthClientID, cfg.OauthClientSecret, provider.GetOauthData(cfg.OauthRefreshToken, cfg.OauthRedirectURL))
		if err != nil {
			return nil, fmt.Errorf("could not retrieve access token from refresh token")
		}
		cfg.OauthAccessToken = accessToken
	}

	dsn, err := provider.DSN(
		cfg.Account,
		cfg.User,
		cfg.Password,
		cfg.BrowserAuth,
		cfg.PrivateKeyPath,
		cfg.PrivateKey,
		cfg.PrivateKeyPassphrase,
		cfg.OauthAccessToken,
		cfg.Region,
		cfg.Role,
		cfg.Host,
		cfg.Protocol,
		cfg.Port,
		cfg.Warehouse,
	)
	if err != nil {
		return nil, fmt.Errorf("could not build dsn for snowflake connection err = %w", err)
	}

	db, err := sqlx.Connect("snowflake", dsn)
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	db = db.Unsafe()

	return db, nil
}

func (c *GoflakeClient) Close() {
	c.db.Close()
}

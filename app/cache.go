package app

import (
	"crypto/tls"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/valkey-io/valkey-go"
)

var rdb valkey.Client

func InitRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		logrus.Fatal("REDIS_URL environment variable is not set")
	}

	parsedURL, err := url.Parse(redisURL)
	if err != nil {
		logrus.Fatalf("Failed to parse Redis URL: %v", err)
	}

	host := parsedURL.Host
	password, _ := parsedURL.User.Password()

	rdb, err = valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{host},
		Password:    password,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})
	if err != nil {
		logrus.Fatalf("Failed to connect to Redis: %v", err)
	}

	logrus.Info("Connected to Redis")
}

package main

import (
	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"os"
	"fmt"
	"strings"
)

func dumpPath(c *api.Client, prefix string) {
	data, err := c.Logical().Read("secret/vx/" + prefix)
	if err != nil {
		logrus.Fatal(err)
	}
	for key, value := range data.Data {
		fmt.Printf("export %s_%s=\"%s\"\n", strings.ToUpper(prefix), strings.ToUpper(key), value.(string))
	}
}
func main() {
	role := os.Getenv("APPROLE_ID")
	if role == "" {
		logrus.Fatal("APPROLE_ID env var is required.")
	}
	secret_id := os.Getenv("APPROLE_SECRET")
	if secret_id == "" {
		logrus.Fatal("APPROLE_SECRET env var is required.")
	}

	config := api.DefaultConfig()
	config.HttpClient = &http.Client{
		Timeout:   1 * time.Second,
		Transport: http.DefaultTransport,
	}
	config.Address = "https://identity.cloud.vx-labs.net/"
	c, err := api.NewClient(config)
	if err != nil {
		logrus.Fatal(err)
	}
	secret, err := c.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   role,
		"secret_id": secret_id,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	c.SetToken(secret.Auth.ClientToken)
	data, err := c.Logical().List("secret/vx/")
	if err != nil {
		logrus.Fatal(err)
	}
	for _, key := range data.Data["keys"].([]interface{}) {
		dumpPath(c, key.(string))
	}
}

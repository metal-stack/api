package main

import (
	"context"
	"log/slog"
	"os"

	"connectrpc.com/connect"
	"github.com/metal-stack/api/go/client"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
)

func main() {

	token := os.Getenv("API_TOKEN")
	project := os.Getenv("PROJECT_ID")
	baseurl := os.Getenv("METAL_APISERVER_URL")

	c := client.New(client.DialConfig{
		BaseURL: baseurl,
		Token:   token,
	})

	ctx := context.Background()
	resp, err := c.Apiv2().IP().List(ctx, connect.NewRequest(&apiv2.IPServiceListRequest{
		Project: project,
	}))

	if err != nil {
		panic(err)
	}

	for _, ip := range resp.Msg.Ips {
		slog.Info("ip", "address", ip.Ip, "id", ip.Uuid, "name", ip.Name)
	}
}

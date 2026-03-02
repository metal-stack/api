package client

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	v2 "github.com/metal-stack/api/go/metalstack/api/v2"
	infra "github.com/metal-stack/api/go/metalstack/infra/v2"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	minInterval     = 5 * time.Second
	maxInterval     = time.Hour
	defaultInterval = 5 * time.Minute
)

// PingConfig is used to configure ping
type PingConfig struct {
	// ComponentType should be set to the type of the microservice
	ComponentType v2.ComponentType
	// Identifier helps to identify multiple instances of a microservice
	// Usually the hostname or the podname could be used for that
	// If omitted, hostname is used
	Identifier *string
	// StartedAt contains the starttime when this go process was started
	StartedAt time.Time
	// Interval at which the ping should happen
	// If not specified, or shorter than 1min, or longer than 1h, it defaults to 5min
	Interval time.Duration
	// Version contains all version details about this microservice
	// You should use https://github.com/metal-stack/v to get all information.
	Version v2.Version
}

// Ping should be called from every microservice which talks to the metal-apiserver
// It will stay and ping at config.Interval in the background until context is canceled
func (c *client) Ping(ctx context.Context, config *PingConfig) {
	if config == nil {
		return
	}

	if config.Interval < minInterval || config.Interval > maxInterval {
		config.Interval = defaultInterval
	}

	var identifier string
	if config.Identifier == nil {
		hostname, err := os.Hostname()
		if err != nil {
			suffix := uuid.NewString()
			identifier = "unknown-" + suffix
		} else {
			identifier = hostname
		}
	} else {
		identifier = *config.Identifier
	}

	req := &infra.ComponentServicePingRequest{
		Type:       config.ComponentType,
		Identifier: identifier,
		StartedAt:  timestamppb.New(config.StartedAt),
		Interval:   durationpb.New(config.Interval),
		Version:    &config.Version,
	}

	c.config.Log.Debug("ping", "config", config)

	ticker := time.NewTicker(config.Interval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
				defer cancel()
				_, err := c.Infrav2().Component().Ping(pingCtx, req)
				if err != nil {
					c.config.Log.Error("ping", "error", err)
				}
			}
		}
	}()
}

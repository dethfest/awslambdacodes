package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger zerolog.Logger
	HUDSONROCK *hudsonrock.Client
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "hudsonrock"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	// TODO: Add your client initialization here
	c, err := hudsonrock.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger: logger,
		HUDSONROCK: c,
	}, nil
}

package couchdb

import (
	"context"
	"fmt"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/driver"
	"github.com/flimzy/kivik/driver/couchdb/chttp"
)

func (c *client) ConfigContext(_ context.Context) (driver.Config, error) {
	return &config{client: c}, nil
}

type config struct {
	*client
}

var _ driver.Config = &config{}

func (c *config) GetAllContext(ctx context.Context) (map[string]map[string]string, error) {
	conf := map[string]map[string]string{}
	_, err := c.DoJSON(ctx, kivik.MethodGet, "/_config", nil, &conf)
	return conf, err
}

func (c *config) SetContext(ctx context.Context, secName, key, value string) error {
	return c.DoError(ctx, kivik.MethodPut, fmt.Sprintf("/_config/%s/%s", secName, key), &chttp.Options{JSON: value})
}

func (c *config) DeleteContext(ctx context.Context, secName, key string) error {
	return c.DoError(ctx, kivik.MethodDelete, fmt.Sprintf("/_config/%s/%s", secName, key), nil)
}

func (c *config) GetSectionContext(ctx context.Context, secName string) (map[string]string, error) {
	sec := map[string]string{}
	_, err := c.DoJSON(ctx, kivik.MethodGet, fmt.Sprintf("/_config/%s", secName), nil, &sec)
	return sec, err
}

func (c *config) GetContext(ctx context.Context, secName, key string) (string, error) {
	var value string
	_, err := c.DoJSON(ctx, kivik.MethodGet, fmt.Sprintf("_config/%s/%s", secName, key), nil, &value)
	return value, err
}

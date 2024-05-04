package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bakito/go-awtrix/client/types"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	UpdateCustomApp(name string, builder *types.AppBuilder) error
	Notify(builder *types.AppBuilder) error
	UpdateSettings(settings *types.SettingsBuilder) error

	Loop() (*types.Loop, error)
}

func New(url string) Client {
	return &client{
		client: resty.New().SetBaseURL(url),
	}
}

type client struct {
	client *resty.Client
}

func (c *client) Notify(builder *types.AppBuilder) error {
	b, _ := json.Marshal(builder.Build())
	print(string(b))

	return c.handleOKResult(c.client.R().SetBody(builder.Build()).Post("/api/notify"))
}

func (c *client) UpdateCustomApp(name string, builder *types.AppBuilder) error {
	return c.handleOKResult(c.client.R().SetBody(builder.Build()).SetQueryParam("name", name).Post("/api/custom"))
}

func (c *client) UpdateSettings(builder *types.SettingsBuilder) error {
	return c.handleOKResult(c.client.R().SetBody(builder.Build()).Post("/api/settings"))
}

func (c *client) Loop() (*types.Loop, error) {
	loop := &types.Loop{}
	return loop, c.handleResult(c.client.R().SetResult(loop).Get("/api/loop"))
}

func (c *client) handleResult(resp *resty.Response, errIn error) error {
	if errIn != nil {
		return errIn
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d (%s)", resp.StatusCode(), resp.Status())
	}
	return nil
}

func (c *client) handleOKResult(resp *resty.Response, errIn error) error {
	err := c.handleResult(resp, errIn)
	if err != nil {
		return err
	}
	body := string(resp.Body())
	if body != "OK" {
		return fmt.Errorf("unexpected response: '%s'", body)
	}
	return nil
}

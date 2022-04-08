package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) ListStories(query *model.ListStoriesParams) (*model.ListStoresResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListStoresResponse
	resp, err := c.Get("/stories", data, &res)

	return res, resp, err
}

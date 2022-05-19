package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) ListBugs(query *model.ListBugsParams) (*model.ListBugsResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListBugsResponse
	resp, err := c.Get("/bugs", data, &res)

	return res, resp, err
}

func (c *Client) UpdateBug(query *model.UpdateBugsParams) (*model.UpdateBugResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}
	var res *model.UpdateBugResponse
	resp, err := c.Post("/bugs", data, &res)
	return res, resp, err
}

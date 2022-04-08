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

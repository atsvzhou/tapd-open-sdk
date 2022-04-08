package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) GetReleases(query *model.ListReleasesParams) (*model.ListReleaseResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListReleaseResponse
	resp, err := c.Get("/releases", data, &res)

	return res, resp, err
}

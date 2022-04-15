package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) ListIterations(query *model.ListIterationsParams) (*model.ListIterationsResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListIterationsResponse
	resp, err := c.Get("/iterations", data, &res)

	return res, resp, err
}

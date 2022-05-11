package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) QueryUsers(query *model.QueryUsersParams) (*model.QueryUsersResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.QueryUsersResponse
	resp, err := c.Get("/workspaces/users", data, &res)

	return res, resp, err
}

package tapd

import (
	"net/http"
	"tapd-open-sdk/model"
)

func (c *Client) ListProjects(query *model.ListProjectsParams) (*model.ListProjectsResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListProjectsResponse
	resp, err := c.Get("/workspaces/projects", data, &res)
	return res, resp, err
}

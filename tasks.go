package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) ListTasks(query *model.ListTasksParams) (*model.ListTasksResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListTasksResponse
	resp, err := c.Get("/tasks", data, &res)

	return res, resp, err
}

func (c *Client) UpdateTask(query *model.UpdateTaskParams) (*model.UpdateTaskResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.UpdateTaskResponse
	resp, err := c.Post("/tasks", data, &res)

	return res, resp, err
}

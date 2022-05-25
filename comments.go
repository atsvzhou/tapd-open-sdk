package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) ListComments(query *model.ListCommentsParams) (*model.ListCommentsResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListCommentsResponse
	resp, err := c.Get("/comments", data, &res)

	return res, resp, err
}

func (c *Client) UpdateComment(query *model.UpdateCommentParams) (*model.UpdateCommentsResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.UpdateCommentsResponse
	resp, err := c.Post("/comments", data, &res)

	return res, resp, err
}

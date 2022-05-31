package tapd

import (
	"github.com/atsvzhou/tapd-open-sdk/model"
	"net/http"
)

func (c *Client) ListStoryCategories(query *model.ListStoryCategoriesParams) (*model.ListStoryCategoriesResponse, *http.Response, error) {
	data, err := c.HandleParams(query)
	if err != nil {
		return nil, nil, err
	}

	var res *model.ListStoryCategoriesResponse
	resp, err := c.Get("/story_categories", data, &res)

	return res, resp, err
}

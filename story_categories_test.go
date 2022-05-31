package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListStoryCategories(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListStoryCategoriesParams{
		WorkspaceId: "",
	}
	listStoryCategories, _, err := client.ListStoryCategories(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(listStoryCategories.Data))
	content, err := json.Marshal(&listStoryCategories)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

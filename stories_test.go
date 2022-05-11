package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListStories(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListStoriesParams{
		WorkspaceId: "",
	}
	listStories, _, err := client.ListStories(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(listStories.Data))
	content, err := json.Marshal(&listStories)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

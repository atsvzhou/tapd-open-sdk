package tapd

import (
	"encoding/json"
	"fmt"
	"tapd-open-sdk/model"
	"testing"
)

func TestListStories(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListStoriesParams{WorkspaceId: 0}

	listStories, _, err := client.ListStories(params)
	if err != nil {
		t.Error(err)
	}

	content, err := json.Marshal(&listStories)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

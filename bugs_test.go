package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListBugs(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListBugsParams{
		WorkspaceId: "",
		//CurrentOwner: "",
	}

	listStories, _, err := client.ListBugs(params)
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

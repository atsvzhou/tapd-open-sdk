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

	listBugs, _, err := client.ListBugs(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(listBugs.Data))
	content, err := json.Marshal(&listBugs)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

func TestUpdateBugs(t *testing.T) {
	client := NewClient("", "")
	params := &model.UpdateBugsParams{
		WorkspaceId: "",
		Title:       "testwer11",
		Id:          "",
		//Status:      "reopened",

	}

	data, _, err := client.UpdateBug(params)
	if err != nil {
		t.Error(err)
	}
	content, err := json.Marshal(&data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

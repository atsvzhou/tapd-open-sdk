package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListTasks(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListTasksParams{
		WorkspaceId: "",
	}

	listTasks, _, err := client.ListTasks(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(listTasks.Data))
	content, err := json.Marshal(&listTasks)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

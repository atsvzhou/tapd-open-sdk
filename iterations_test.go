package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListIterations(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListIterationsParams{
		WorkspaceId: "",
	}

	listIterations, _, err := client.ListIterations(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(listIterations.Data))
	content, err := json.Marshal(&listIterations)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

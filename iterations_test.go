package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListIterations(t *testing.T) {
	//client := NewClient("", "")
	client := NewClient("", "")
	params := &model.ListIterationsParams{
		WorkspaceId: "",
		Status:      "open",
	}

	listIterations, _, err := client.ListIterations(params)
	if err != nil {
		t.Error(err)
	}
	for _, v := range listIterations.Data {
		fmt.Println(v.Iteration.Status)
	}
	fmt.Println(len(listIterations.Data))
	content, err := json.Marshal(&listIterations)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

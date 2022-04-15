package tapd

import (
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListProjects(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListProjectsParams{CompanyId: ""}

	listProjects, _, err := client.ListProjects(params)
	if err != nil {
		t.Error(err)
	}
	for _, v := range listProjects.Data {
		if v.Workspace.Status != "closed" {
			fmt.Println(v.Workspace)
		}
	}
	//content, err := json.Marshal(&listProjects)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(string(content))

}

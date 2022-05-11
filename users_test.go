package tapd

import (
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestQueryUsers(t *testing.T) {
	client := NewClient("", "")
	params := &model.QueryUsersParams{
		WorkspaceId: "",
	}

	queryUsers, _, err := client.QueryUsers(params)
	if err != nil {
		t.Error(err)
	}
	for _, v := range queryUsers.Data {
		for _, w := range v.UserWorkspace.RoleID {
			if w == "" {
				fmt.Println(v.UserWorkspace.User)
			}

		}
	}
	//content, err := json.Marshal(&queryUsers.Data)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(string(content))

}

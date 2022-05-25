package tapd

import (
	"encoding/json"
	"fmt"
	"github.com/atsvzhou/tapd-open-sdk/model"
	"testing"
)

func TestListComments(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListCommentsParams{
		WorkspaceId: "",
		EntryType:   "",
		EntryId:     "",
	}

	listComments, _, err := client.ListComments(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(listComments.Data))
	content, err := json.Marshal(&listComments)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

func TestUpdateComment(t *testing.T) {
	client := NewClient("", "")
	params := &model.UpdateCommentParams{
		WorkspaceId: "",
		Description: "这是另外一条测试评论",
		Author:      "",
		EntryType:   "bug",
		EntryId:     "",
	}

	updateComment, _, err := client.UpdateComment(params)
	if err != nil {
		t.Error(err)
	}
	content, err := json.Marshal(&updateComment)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(content))

}

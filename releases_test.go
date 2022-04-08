package tapd

import (
	"encoding/json"
	"fmt"
	"tapd-open-sdk/model"
	"testing"
)

func TestGetReleases(t *testing.T) {
	client := NewClient("", "")
	params := &model.ListReleasesParams{
		WorkspaceId: 0,
	}

	getReleases, _, err := client.GetReleases(params)
	if err != nil {
		t.Error(err)
	}
	for _, v := range getReleases.Data {
		fmt.Println(v.Release)
	}
	content, err := json.Marshal(&getReleases)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(content))

}

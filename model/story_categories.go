package model

type ListStoryCategoriesParams struct {
	WorkspaceId string `json:"workspace_id"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Order       string `json:"order,omitempty"`
	ParentId    string `json:"parent_id,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Page        int    `json:"page,omitempty"`
}
type ListStoryCategoriesResponse struct {
	BaseResponse
	Data []CategoryData `json:"data"`
}

type Category struct {
	ID          string `json:"id"`
	WorkspaceID string `json:"workspace_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}

type CategoryData struct {
	Category Category `json:"Category"`
}

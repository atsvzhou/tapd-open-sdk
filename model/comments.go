package model

type ListCommentsParams struct {
	WorkspaceId string `json:"workspace_id"`
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Author      string `json:"author,omitempty"`
	EntryType   string `json:"entry_type,omitempty"`
	EntryId     string `json:"entry_id,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Page        int    `json:"page,omitempty"`
	Order       string `json:"order,omitempty"`
}
type ListCommentsResponse struct {
	BaseResponse
	Data []CommentsData `json:"data"`
}

type CommentsData struct {
	Comment Comment `json:"Comment"`
}

type Comment struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	EntryType   string `json:"entry_type"`
	EntryID     string `json:"entry_id"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	WorkspaceID string `json:"workspace_id"`
}

type UpdateCommentParams struct {
	WorkspaceId string `json:"workspace_id"`
	Description string `json:"description"`
	Author      string `json:"author"`
	EntryType   string `json:"entry_type"`
	EntryId     string `json:"entry_id"`
}

type UpdateCommentsResponse struct {
	BaseResponse
	Data CommentsData `json:"data"`
}

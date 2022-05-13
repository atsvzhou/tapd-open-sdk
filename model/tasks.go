package model

type ListTasksParams struct {
	WorkspaceId string `json:"workspace_id"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Status      string `json:"status,omitempty"`
	IterationId string `json:"iteration_id,omitempty"`
	StoryId     string `json:"story_id,omitempty"`
	Description string `json:"description,omitempty"`
	Order       string `json:"order,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Page        int    `json:"page,omitempty"`
	Owner       string `json:"owner,omitempty"`
}

type ListTasksResponse struct {
	Status int            `json:"status"`
	Data   []ListTaskData `json:"data"`
	Info   string         `json:"info"`
}

type ListTaskData struct {
	Task ListTask `json:"task"`
}

type ListTask struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	WorkspaceID     string      `json:"workspace_id"`
	Creator         string      `json:"creator"`
	Created         string      `json:"created"`
	Modified        string      `json:"modified"`
	Status          string      `json:"status"`
	Owner           string      `json:"owner"`
	Cc              interface{} `json:"cc"`
	Begin           interface{} `json:"begin"`
	Due             interface{} `json:"due"`
	StoryID         string      `json:"story_id"`
	IterationID     string      `json:"iteration_id"`
	Priority        string      `json:"priority"`
	Progress        string      `json:"progress"`
	Completed       interface{} `json:"completed"`
	EffortCompleted string      `json:"effort_completed"`
	EffortTotal     interface{} `json:"effort_total"`
	Exceed          string      `json:"exceed"`
	Remain          string      `json:"remain"`
	Effort          string      `json:"effort"`
}

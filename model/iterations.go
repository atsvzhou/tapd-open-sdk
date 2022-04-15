package model

type ListIterationsParams struct {
	WorkspaceId string `json:"workspace_id"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Page        int    `json:"page,omitempty"`
}

type ListIterationsResponse struct {
	Status int                  `json:"status"`
	Data   []ListIterationsData `json:"data"`
	Info   string               `json:"info"`
}

type ListIterations struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	WorkspaceID  string      `json:"workspace_id"`
	Startdate    string      `json:"startdate"`
	Enddate      string      `json:"enddate"`
	Status       string      `json:"status"`
	ReleaseID    string      `json:"release_id"`
	Description  string      `json:"description"`
	Creator      string      `json:"creator"`
	Created      string      `json:"created"`
	Modified     string      `json:"modified"`
	Completed    interface{} `json:"completed"`
	Releaseowner interface{} `json:"releaseowner"`
	Launchdate   interface{} `json:"launchdate"`
	Notice       interface{} `json:"notice"`
	Releasename  interface{} `json:"releasename"`
}

type ListIterationsData struct {
	Iteration ListIterations `json:"iteration"`
}

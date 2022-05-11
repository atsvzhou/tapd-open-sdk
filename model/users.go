package model

type QueryUsersParams struct {
	WorkspaceId string `json:"workspace_id"`
	RoleId      string `json:"role_id"`
}

type QueryUsersResponse struct {
	Status int    `json:"status"`
	Data   []Data `json:"data"`
	Info   string `json:"info"`
}

type UserWorkspace struct {
	User   string   `json:"user"`
	RoleID []string `json:"role_id"`
}

type Data struct {
	UserWorkspace UserWorkspace `json:"UserWorkspace"`
}

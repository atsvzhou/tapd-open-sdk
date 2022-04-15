package model

type ListProjectsParams struct {
	CompanyId string `json:"company_id"`
}

type ListProjectsResponse struct {
	BaseResponse
	Data []ListProjectsData `json:"data"`
}
type ListProjectsData struct {
	Workspace ListProjects `json:"workspace"`
}
type ListProjects struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	PrettyName  string      `json:"pretty_name"`
	Description interface{} `json:"description"`
	Status      string      `json:"status"`
	ParentId    string      `json:"parent_id"`
	Secrecy     string      `json:"secrecy"`
	Created     string      `json:"created"`
	CreatorId   string      `json:"creator_id"`
	Creator     string      `json:"creator"`
	BeginDate   string      `json:"begin_date"`
	EndDate     string      `json:"end_date"`
	MemberCount int         `json:"member_count"`
}

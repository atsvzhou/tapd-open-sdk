package model

type ListReleasesParams struct {
	Id          int    `json:"id,omitempty"`
	WorkspaceId string `json:"workspace_id,required"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	//Startdate   time.Time `json:"startdate,omitempty"`
	//Enddate     time.Time `json:"enddate,omitempty"`
	Creator string `json:"creator,omitempty"`
	//Created     time.Time `json:"created,omitempty"`
	//Modified    time.Time `json:"modified,omitempty"`
	Status string `json:"status,omitempty"` //done,open
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Order  string `json:"order,omitempty"`
	Fields string `json:"fields,omitempty"`
}
type ListReleaseResponse struct {
	BaseResponse
	Data []ListReleasesData `json:"data"`
}

type ListReleases struct {
	Id          string `json:"id"`
	WorkspaceId string `json:"workspace_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Startdate   string `json:"startdate"`
	Enddate     string `json:"enddate"`
	Creator     string `json:"creator"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	Status      string `json:"status"`
}

type ListReleasesData struct {
	Release ListReleases `json:"release"`
}

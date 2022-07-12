package model

type ListBugsParams struct {
	WorkspaceId   string `json:"workspace_id"`
	ReleaseId     string `json:"release_id,omitempty"` //发布计划
	Priority      string `json:"priority,omitempty"`
	Severity      string `json:"severity,omitempty"`
	Status        string `json:"status,omitempty"`
	IterationId   string `json:"iteration_id,omitempty"`
	VersionReport string `json:"version_report,omitempty"`
	Order         string `json:"order,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Page          int    `json:"page,omitempty"`
	Description   string `json:"description,omitempty"`
	CurrentOwner  string `json:"current_owner,omitempty"`
}

type UpdateBugsParams struct {
	WorkspaceId  string `json:"workspace_id"`
	Id           string `json:"id"`
	Title        string `json:"title,omitempty"`
	Priority     string `json:"priority,omitempty"`
	Severity     string `json:"severity,omitempty"`
	Status       string `json:"status,omitempty"`
	Module       string `json:"module,omitempty"`
	CurrentOwner string `json:"current_owner,omitempty"`
	Te           string `json:"te,omitempty"`
	De           string `json:"de,omitempty"`
	Auditer      string `json:"auditer,omitempty"`
	Confirmer    string `json:"confirmer,omitempty"`
	Fixer        string `json:"fixer,omitempty"`
	Closer       string `json:"closer,omitempty"`
	Lastmodify   string `json:"lastmodify,omitempty"`
}

type ListBugsResponse struct {
	BaseResponse
	Data []ListBugsData `json:"data"`
}

type ListBugsData struct {
	Bug ListBugs `json:"bug"`
}

type ListBugs struct {
	Id               string      `json:"id"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	Priority         string      `json:"priority"`
	Severity         string      `json:"severity"`
	Module           string      `json:"module"`
	Status           string      `json:"status"`
	Reporter         string      `json:"reporter"`
	Deadline         interface{} `json:"deadline"`
	Created          string      `json:"created"`
	Bugtype          string      `json:"bugtype"`
	Resolved         interface{} `json:"resolved"`
	Closed           string      `json:"closed"`
	Modified         string      `json:"modified"`
	Lastmodify       string      `json:"lastmodify"`
	Auditer          string      `json:"auditer"`
	De               string      `json:"de"`
	Fixer            string      `json:"fixer"`
	VersionTest      string      `json:"version_test"`
	VersionReport    string      `json:"version_report"`
	VersionClose     string      `json:"version_close"`
	VersionFix       string      `json:"version_fix"`
	BaselineFind     string      `json:"baseline_find"`
	BaselineJoin     string      `json:"baseline_join"`
	BaselineClose    string      `json:"baseline_close"`
	BaselineTest     string      `json:"baseline_test"`
	Sourcephase      string      `json:"sourcephase"`
	Te               string      `json:"te"`
	CurrentOwner     string      `json:"current_owner"`
	IterationId      string      `json:"iteration_id"`
	Resolution       string      `json:"resolution"`
	Source           string      `json:"source"`
	Originphase      string      `json:"originphase"`
	Confirmer        string      `json:"confirmer"`
	Milestone        string      `json:"milestone"`
	Participator     string      `json:"participator"`
	Closer           string      `json:"closer"`
	Platform         string      `json:"platform"`
	Os               string      `json:"os"`
	Testtype         string      `json:"testtype"`
	Testphase        string      `json:"testphase"`
	Frequency        string      `json:"frequency"`
	Cc               string      `json:"cc"`
	RegressionNumber string      `json:"regression_number"`
	Flows            string      `json:"flows"`
	Feature          interface{} `json:"feature"`
	Testmode         string      `json:"testmode"`
	Estimate         interface{} `json:"estimate"`
	IssueId          interface{} `json:"issue_id"`
	CreatedFrom      interface{} `json:"created_from"`
	InProgressTime   interface{} `json:"in_progress_time"`
	VerifyTime       interface{} `json:"verify_time"`
	RejectTime       interface{} `json:"reject_time"`
	ReopenTime       interface{} `json:"reopen_time"`
	AuditTime        interface{} `json:"audit_time"`
	SuspendTime      interface{} `json:"suspend_time"`
	Due              interface{} `json:"due"`
	Begin            interface{} `json:"begin"`
	WorkspaceId      string      `json:"workspace_id"`
}

type UpdateBugResponse struct {
	Status int           `json:"status"`
	Data   UpdateBugData `json:"data"`
	Info   string        `json:"info"`
}
type UpdateBugData struct {
	Bug UpdateBug `json:"Bug"`
}
type UpdateBug struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	Priority         string      `json:"priority"`
	Severity         string      `json:"severity"`
	Module           string      `json:"module"`
	Status           string      `json:"status"`
	Reporter         string      `json:"reporter"`
	Deadline         string      `json:"deadline"`
	Created          string      `json:"created"`
	Bugtype          string      `json:"bugtype"`
	Resolved         string      `json:"resolved"`
	Closed           string      `json:"closed"`
	Modified         string      `json:"modified"`
	Lastmodify       string      `json:"lastmodify"`
	Auditer          string      `json:"auditer"`
	De               string      `json:"de"`
	Fixer            string      `json:"fixer"`
	VersionTest      string      `json:"version_test"`
	VersionReport    string      `json:"version_report"`
	VersionClose     string      `json:"version_close"`
	VersionFix       string      `json:"version_fix"`
	BaselineFind     string      `json:"baseline_find"`
	BaselineJoin     string      `json:"baseline_join"`
	BaselineClose    string      `json:"baseline_close"`
	BaselineTest     string      `json:"baseline_test"`
	Sourcephase      string      `json:"sourcephase"`
	Te               string      `json:"te"`
	CurrentOwner     string      `json:"current_owner"`
	IterationID      string      `json:"iteration_id"`
	Resolution       string      `json:"resolution"`
	Source           string      `json:"source"`
	Originphase      string      `json:"originphase"`
	Confirmer        string      `json:"confirmer"`
	Milestone        string      `json:"milestone"`
	Participator     string      `json:"participator"`
	Closer           string      `json:"closer"`
	Platform         string      `json:"platform"`
	Os               string      `json:"os"`
	Testtype         string      `json:"testtype"`
	Testphase        string      `json:"testphase"`
	Frequency        string      `json:"frequency"`
	Cc               string      `json:"cc"`
	RegressionNumber string      `json:"regression_number"`
	Flows            string      `json:"flows"`
	Feature          string      `json:"feature"`
	Testmode         string      `json:"testmode"`
	Estimate         interface{} `json:"estimate"`
	IssueID          string      `json:"issue_id"`
	CreatedFrom      interface{} `json:"created_from"`
	InProgressTime   interface{} `json:"in_progress_time"`
	VerifyTime       interface{} `json:"verify_time"`
	RejectTime       interface{} `json:"reject_time"`
	ReopenTime       interface{} `json:"reopen_time"`
	AuditTime        interface{} `json:"audit_time"`
	SuspendTime      interface{} `json:"suspend_time"`
	Due              interface{} `json:"due"`
	Begin            interface{} `json:"begin"`
	ReleaseID        string      `json:"release_id"`
	WorkspaceID      string      `json:"workspace_id"`
}

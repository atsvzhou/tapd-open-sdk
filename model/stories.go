package model

type ListStoriesParams struct {
	WorkspaceId int `json:"workspace_id"` //
	//Id             int    `json:"id"`               //ID	支持多ID查询
	//Name           string `json:"name"`             //标题	支持模糊匹配
	//Priority       string `json:"priority"`         //优先级	支持枚举查询
	//BusinessValue  int    `json:"business_value"`   //业务价值
	//Status         string `json:"status"`           //状态	支持枚举查询
	//Version        string `json:"version"`          //版本
	//Module         string `json:"module"`           //模块
	//TestFocus      string `json:"test_focus"`       //测试重点
	//WorkitemTypeId string `json:"workitem_type_id"` //需求类别
	//Size           int    `json:"size"`             //规模
	//Owner          string `json:"owner"`            //当前处理人	支持模糊匹配
	//Cc             string `json:"cc"`               //当前处理人	支持模糊匹配
	//Creator        string `json:"creator"`          //创建人	支持多人员查询
	//Developer      string `json:"developer"`        //开发人员
	//
}
type ListStoresResponse struct {
	BaseResponse
	Data []ListStoresData `json:"data"`
}

type ListStores struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	WorkspaceID     string      `json:"workspace_id"`
	Creator         string      `json:"creator"`
	Created         string      `json:"created"`
	Modified        string      `json:"modified"`
	Status          string      `json:"status"`
	Owner           string      `json:"owner"`
	Cc              string      `json:"cc"`
	Begin           interface{} `json:"begin"`
	Due             interface{} `json:"due"`
	Size            interface{} `json:"size"`
	Priority        string      `json:"priority"`
	Developer       string      `json:"developer"`
	IterationID     string      `json:"iteration_id"`
	TestFocus       string      `json:"test_focus"`
	Type            string      `json:"type"`
	Source          string      `json:"source"`
	Module          string      `json:"module"`
	Version         string      `json:"version"`
	Completed       string      `json:"completed"`
	CategoryID      string      `json:"category_id"`
	Path            string      `json:"path"`
	ParentID        string      `json:"parent_id"`
	ChildrenID      string      `json:"children_id"`
	AncestorID      string      `json:"ancestor_id"`
	BusinessValue   interface{} `json:"business_value"`
	Effort          string      `json:"effort"`
	EffortCompleted string      `json:"effort_completed"`
	Exceed          string      `json:"exceed"`
	Remain          string      `json:"remain"`
	ReleaseID       string      `json:"release_id"`
}

type ListStoresData struct {
	Story ListStores `json:"story"`
}

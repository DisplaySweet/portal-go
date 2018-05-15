package portal

type PermissionGroup struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectid"`
	CompanyID string `json:"companyid"`
	GroupType string `json:"type"`
	GroupName string `json:"groupname"`
}

const permissiongroupEndpoint = "permissiongroups"

//TODO:  this is a stub

//GET

//GET id

//POST

//DELETE

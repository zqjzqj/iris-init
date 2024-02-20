package tplStruct

import "iris-init/model"

var Str2ModelMap = map[string]any{
	"Admin":           model.Admin{},
	"Department":      model.Department{},
	"DepartmentAdmin": model.DepartmentAdmin{},
	"Roles":           model.Roles{},
	"Project":         model.Project{},
	"ProjectAudit":    model.ProjectAudit{},
	"ApprovalProcess": model.ApprovalProcess{},
}

package tplStruct

import "iris-init/model"

var Str2ModelMap = map[string]any{
	"Admin":           model.Admin{},
	"ApprovalProcess": model.ApprovalProcess{},
	"Organizer":       model.Organizer{},
	"Project":         model.Project{},
	"ProjectAudit":    model.ProjectAudit{},
	"ProjectOrdinary": model.ProjectOrdinary{},
	"Roles":           model.Roles{},
}

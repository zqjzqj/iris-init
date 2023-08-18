package tplStruct

import "iris-init/model"

var Str2ModelMap = map[string]any{
	"Course":             model.Course{},
	"CourseSection":      model.CourseSection{},
	"Institution":        model.Institution{},
	"InstitutionAccount": model.InstitutionAccount{},
	"Question":           model.Question{},
	"QuestionAnswers":    model.QuestionAnswers{},
	"Students":           model.Students{},
	"Teacher":            model.Teacher{},
}

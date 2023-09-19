package tplStruct

import "iris-init/model"

var Str2ModelMap = map[string]any{
	"Admin":                model.Admin{},
	"Class":                model.Class{},
	"Course":               model.Course{},
	"CourseReferences":     model.CourseReferences{},
	"CourseSection":        model.CourseSection{},
	"CourseSectionStudy":   model.CourseSectionStudy{},
	"ExamPaper":            model.ExamPaper{},
	"ExamPaperStudy":       model.ExamPaperStudy{},
	"Institution":          model.Institution{},
	"InstitutionAccount":   model.InstitutionAccount{},
	"LearnRecords":         model.LearnRecords{},
	"LearnStatistics":      model.LearnStatistics{},
	"Question":             model.Question{},
	"QuestionAnswers":      model.QuestionAnswers{},
	"QuestionCollection":   model.QuestionCollection{},
	"QuestionErr":          model.QuestionErr{},
	"QuestionFeedback":     model.QuestionFeedback{},
	"QuestionRecords":      model.QuestionRecords{},
	"QuestionRecordsItems": model.QuestionRecordsItems{},
	"StudentClass":         model.StudentClass{},
	"StudentCourseScore":   model.StudentCourseScore{},
	"StudentExtInfo":       model.StudentExtInfo{},
	"StudentToken":         model.StudentToken{},
	"Students":             model.Students{},
	"Teacher":              model.Teacher{},
}

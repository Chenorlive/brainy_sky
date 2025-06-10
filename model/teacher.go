package model

import "github.com/gofrs/uuid"

type TeacherSubjectClass struct {
	Base
	TeacherID      uuid.UUID     `json:"teacher_id" `
	SubjectID      uuid.UUID     `json:"subject_id" `
	Teacher        *User         `json:"teacher" gorm:"foreignKey:TeacherID"`
	ClassID        uuid.UUID     `json:"class_id" `
	Class          *Class        `json:"class" gorm:"foreignKey:ClassID"`
	Subject        *Subject      `json:"subject" gorm:"foreignKey:SubjectID"`
	AcadimicYearID *string       `json:"acadimic_year_id" `
	AcadimicYear   *AcadimicYear `json:"acadimic_year" gorm:"foreignKey:AcadimicYearID"`
}

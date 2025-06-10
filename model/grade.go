package model

import "github.com/gofrs/uuid"

type Grade struct {
	Base
	StudentClassID      uuid.UUID           `json:"student_class_id" `
	StudentClass        *StudentClass       `json:"student_class" gorm:"foreignKey:StudentClassID"`
	TeacherSubjectID    uuid.UUID           `json:"teacher_subject_id" `
	TeacherSubjectClass TeacherSubjectClass `json:"teacher_subject" gorm:"foreignKey:TeacherSubjectID"`
	PeriodID            *uuid.UUID          `json:"period_id" `
	Period              *Period             `json:"peroid"  gorm:"foreignKey:PeriodID"`
	Grade               int                 `json:"grade" gorm:"colume:not null"`
	GradeLetter         *string             `json:"grade_letter"`
}

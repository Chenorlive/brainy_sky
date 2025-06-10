package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type School struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
	Address     *string `json:"address"`
}

type AcadimicYear struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active" `
}

type AcadimicSemesterType struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
	IsVirtual   bool    `json:"is_virtual" `
}

type AcadimicSemester struct {
	Base
	AcadimicSemesterTypeID uuid.UUID             `json:"acadimic_semester_type_id" `
	AcadimicYearID         uuid.UUID             `json:"acadimic_year_id" `
	AcadimicSemesterType   *AcadimicSemesterType `json:"acadimic_semester_type" gorm:"foreignKey:AcadimicSemesterTypeID"`
	AcadimicYear           *AcadimicYear         `json:"acadimic_year" gorm:"foreignKey:AcadimicYearID"`
	IsActive               *bool                 `json:"is_active" `
	StartDate              time.Time             `json:"start_date" `
	EndDate                time.Time             `json:"end_date" `
}

type PeriodType struct {
	Base
	Name        *string `json:"name" `
	Description *string `json:"description"`
	IsVirtual   *bool   `json:"is_virtual" `
}

type Period struct {
	Base
	PeriodTypeID           uuid.UUID             `json:"period_type_id" `
	AcadimicSemesterTypeID uuid.UUID             `json:"acadimic_semester_type_ID" `
	PeriodType             *PeriodType           `json:"period_type" gorm:"foreignKey:PeriodTypeID"`
	AcadimicSemesterType   *AcadimicSemesterType `json:"acadimic_semester_type" gorm:"foreignKey:AcadimicSemesterTypeID"`
}

type ClassType struct {
	Base
	Name        *string `json:"name" `
	Description *string `json:"description"`
}

type Class struct {
	Base
	ClassTypeID uuid.UUID  `json:"class_type_id" `
	Name        *string    `json:"name" `
	Description *string    `json:"description"`
	ClassType   *ClassType `json:"class_type" gorm:"foreignKey:ClassTypeID"`
}

type Subject struct {
	Base
	Name        *string `json:"name" `
	Description *string `json:"description"`
}

type SchedulePeriod struct {
	Base
	Name        string    `json:"name" `
	Description *string   `json:"description"`
	StartTime   time.Time `json:"start_time" `
	EndTime     time.Time `json:"end_time" `
}

type Schedule struct {
	Base
	SchedulePeriodID uuid.UUID       `json:"schedule_period_id" `
	ClassID          uuid.UUID       `json:"class_id" `
	SubjectID        uuid.UUID       `json:"subject_id" `
	Day              string          `json:"day" gorm:"type:day"`
	SchedulePeriod   *SchedulePeriod `json:"schedule_period" gorm:"foreignKey:SchedulePeriodID"`
	Class            *Class          `json:"class" gorm:"foreignKey:ClassID"`
	Subject          *Subject        `json:"subject" gorm:"foreignKey:SubjectID"`
}

type Day string

const (
	Monday    Day = "Monday"
	Tuesday   Day = "Tuesday"
	Wednesday Day = "Wednesday"
	Thursday  Day = "Thursday"
	Friday    Day = "Friday"
	Saturday  Day = "Saturday"
	Sunday    Day = "Sunday"
)

func (p *Day) Scan(value interface{}) error {
	if val, ok := value.(string); ok {
		*p = Day(val)
		return nil
	}
	return nil
}

func (p Day) Value() (interface{}, error) {
	return string(p), nil
}

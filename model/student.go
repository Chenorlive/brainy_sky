package model

import "github.com/gofrs/uuid"

type StudentClass struct {
	Base
	ClassID            uuid.UUID         `json:"class_id" `
	StudentID          uuid.UUID         `json:"student_id" `
	Class              *Class            `json:"class" gorm:"foreignKey:ClassID"`
	Student            *User             `json:"student" gorm:"foreignKey:StudentID"`
	AcadimicSemesterID uuid.UUID         `json:"acadimic_semester_id" `
	AcadimicSemester   *AcadimicSemester `json:"acadimic_semester" gorm:"foreignKey:AcadimicSemesterID"`
}

type StudentClassReport struct {
	Base
	StudentClassID uuid.UUID     `json:"student_class_id" `
	StudentClass   *StudentClass `json:"student_class" gorm:"foreignKey:StudentClassID"`
	Status         *string       `json:"status" gorm:"type:status;default:'Ongoing'"` // Status can be Passed, Failed, or Summer
	Comment        *string       `json:"comment"`
	Average        *float64      `json:"average"`
}

type StudentBillReport struct {
	Base
	StudentClassID uuid.UUID     `json:"student_class_id" `
	StudentClass   *StudentClass `json:"student_class" gorm:"foreignKey:StudentClassID"`
	Amount         *float64      `json:"amount" `
	Balance        *float64      `json:"balance" `
	Comment        *string       `json:"comment"`
	IsCompleted    bool          `json:"is_completed" `
}

// Parent model
type Parent struct {
	Base
	UserID uuid.UUID `json:"user_id" `
	User   *User     `json:"user" gorm:"foreignKey:UserID"`
}

type ParentStudent struct {
	Base
	ParentID  uuid.UUID `json:"parent_id" `
	StudentID uuid.UUID `json:"student_id" `
	Parent    *Parent   `json:"parent" gorm:"foreignKey:ParentID"`
	Student   *User     `json:"student" gorm:"foreignKey:StudentID"`
}

type Status string

const (
	Passed  Status = "Passed"
	Failed  Status = "Failed"
	Summer  Status = "Summer"
	Ongoing Status = "Ongoing"
)

func (p *Status) Scan(value interface{}) error {
	if val, ok := value.(string); ok {
		*p = Status(val)
		return nil
	}
	return nil
}

func (p Status) Value() (interface{}, error) {
	return string(p), nil
}

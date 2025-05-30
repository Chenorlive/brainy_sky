package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Base struct {
	BaseModel
	CreatedByID *string `json:"created_by" gorm:"created_by"`
	UpdatedByID *string `json:"updated_by" gorm:"updated_by"`
	CreatedBy   *User   `json:"created_by_user" gorm:"foreignKey:CreatedByID"`
	UpdatedBy   *User   `json:"updated_by_user" gorm:"foreignKey:UpdatedByID"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	base.ID = uuid
	//fmt.Printf("UUID: %s\n", base.ID)
	scope.Set("ID", base.ID)
	return nil
}

type Gender struct {
	Base
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Nationality struct {
	Base
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type MariageStatus struct {
	Base
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

var Models = []interface{}{
	&Gender{}, &Nationality{}, &MariageStatus{},
	&User{},
	&Permission{}, &Role{}, &UserRole{}, &RolePermission{},
	&School{}, &AcadimicYear{}, &AcadimicSemesterType{}, &AcadimicSemester{},
	&PeriodType{}, &Period{}, &ClassType{}, &Class{}, &Subject{},
	&SchedulePeriod{}, &Schedule{},
	&BillItem{}, &Installment{}, &ClassBillInstillment{}, &ClassBillItem{},
	&ClassBillSummary{}, &PaymentMethod{}, &Payment{},
	&StudentClass{}, &StudentClassReport{}, &StudentBillReport{},
	&Parent{}, &ParentStudent{},
	&Grade{},
	&TeacherSubjectClass{},
}

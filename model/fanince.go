package model

import "github.com/gofrs/uuid"

type BillItem struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
}

type ClassBillItem struct {
	Base
	ClassID        uuid.UUID     `json:"class_id" `
	BillItemID     uuid.UUID     `json:"bill_item_id" `
	AcadimicYearID uuid.UUID     `json:"acadimic_year_id" `
	AcadimicYear   *AcadimicYear `json:"acadimic_year" gorm:"foreignKey:AcadimicYearID"`
	Class          *Class        `json:"class" gorm:"foreignKey:ClassID"`
	BillItem       *BillItem     `json:"bill_item" gorm:"foreignKey:BillItemID"`
	Amount         *float64      `json:"amount" `
	Coment         *string       `json:"comment"`
}

type ClassBillSummary struct {
	Base
	ClassID        uuid.UUID     `json:"class_id" `
	AcadimicYearID uuid.UUID     `json:"acadimic_year_id" `
	AcadimicYear   *AcadimicYear `json:"acadimic_year" gorm:"foreignKey:AcadimicYearID"`
	Class          *Class        `json:"class" gorm:"foreignKey:ClassID"`
	TotalAmount    *float64      `json:"total_amount" `
}

type Installment struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
}

type ClassBillInstillment struct {
	Base
	InstallmentID      uuid.UUID         `json:"installmentID" `
	Installment        *Installment      `json:"installment" gorm:"foreignKey:InstallmentID"`
	ClassBillSummaryID uuid.UUID         `json:"class_bill_summary_id" `
	ClassBillSummary   *ClassBillSummary `json:"class_bill_summary" gorm:"foreignKey:ClassBillSummaryID"`
	Amount             float64           `json:"amount" `
	Comment            *string           `json:"comment"`
}

type PaymentMethod struct {
	Base
	Name        *string `json:"name" `
	Description *string `json:"description"`
}

type Payment struct {
	Base
	StudentClassID  uuid.UUID      `json:"student_class_id" `
	StudentClass    *StudentClass  `json:"student_class" gorm:"foreignKey:StudentClassID"`
	Amount          float64        `json:"amount" `
	PaymentMethodID *uuid.UUID     `json:"payment_method_id" `
	PaymentMethod   *PaymentMethod `json:"payment_method" gorm:"foreignKey:PaymentMethodID"`
	Comment         *string        `json:"comment"`
}

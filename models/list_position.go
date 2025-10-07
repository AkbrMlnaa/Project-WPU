package models

import (
	"github.com/AkbrMlnaa/Project-WPU/models/types"
	"github.com/google/uuid"
)

type ListPosition struct {
	InternalID int64 `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID uuid.UUID`json:"public_id" db:"public_id" gorm:"public_id"`
	BoardID int64 `json:"board_id" db:"board_id" gorm:"column:board_id"`
	ListOrder types.UUIDArray `json:"list_order"`
}
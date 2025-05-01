package sql

import (
	"time"

	"gorm.io/gorm"
)

type BaceModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

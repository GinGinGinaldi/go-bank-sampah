package models

type Waste struct {
	ID       uint    `gorm:"primaryKey"`
	UserID   uint    `gorm:"not null"`
	Type     string  `gorm:"not null"`
	Volume   float64 `gorm:"not null"` // dalam liter
	Location string  `gorm:"not null"`
}

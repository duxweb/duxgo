package models

// Config @AutoMigrate()
type Config struct {
	Name  string `gorm:"size:250" json:"name"`
	Value string `json:"value"`
}

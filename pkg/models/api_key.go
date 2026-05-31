package models

type APIKey struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64  `gorm:"column:created_at;default:0"`
	UpdatedAt int64  `gorm:"column:updated_at;default:0"`
	DeletedAt int64  `gorm:"column:deleted_at;default:0"`
	ModelID   int64  `gorm:"column:model_id;not null"`
	Key       string `gorm:"column:key;type:varchar(128);not null;index:key,unique"`
}

func (APIKey) TableName() string {
	return "api_keys"
}

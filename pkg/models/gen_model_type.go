package models

type GenModelType struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64  `gorm:"column:created_at;default:0"`
	UpdatedAt int64  `gorm:"column:updated_at;default:0"`
	Name      string `gorm:"column:name;type:varchar(128);not null"`
	Code      string `gorm:"column:code;type:varchar(64);not null;index:code,unique"`
}

func (GenModelType) TableName() string {
	return "gen_model_types"
}

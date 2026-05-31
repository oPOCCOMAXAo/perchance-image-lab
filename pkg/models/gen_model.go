package models

type GenModel struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64  `gorm:"column:created_at;default:0"`
	UpdatedAt int64  `gorm:"column:updated_at;default:0"`
	TypeID    int64  `gorm:"column:type_id;not null"`
	Name      string `gorm:"column:name;type:varchar(128);not null"`
}

func (GenModel) TableName() string {
	return "gen_models"
}

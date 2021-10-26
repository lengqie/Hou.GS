package model

type Url struct {
	ID  uint `gorm:"primaryKey;AUTO_INCREMENT"`
    OriginalUrl string
	Key string
}

func (Url) TableName() string {
	return "url"
}
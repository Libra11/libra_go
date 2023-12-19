package blogs

type Tag struct {
	Id   int64 `gorm:"primary_key"`
	Name string
}

func (*Tag) TableName() string {
	return "tags"
}

package blogs

type Category struct {
	Id   int64 `gorm:"primary_key"`
	Name string
}

func (*Category) TableName() string {
	return "categories"
}

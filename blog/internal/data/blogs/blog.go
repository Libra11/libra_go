package blogs

type Blog struct {
	Id        int64 `gorm:"primary_key"`
	Title     string
	Content   string
	AudioFile string
	Author    string
	CreateAt  int64
	UpdateAt  int64
	ImgUrl    string
	Desc      string
	Category  []Category `gorm:"many2many:blog_categories;"`
	Tags      []Tag      `gorm:"many2many:blog_tags;"`
}

func (*Blog) TableName() string {
	return "blogs"
}

package blogs

type Word struct {
	Id         int64 `gorm:"primary_key"`
	Word       string
	Definition string
	Example    string
	Phrase     string
	Phonetic   string
	CreateAt   int64
	UpdateAt   int64
}

func (*Word) TableName() string {
	return "words"
}

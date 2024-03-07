package dao

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/database/gorms"
	"libra.com/blog/internal/dto"
)

type WordDao struct {
	conn *gorms.GormConn
}

func (b WordDao) GetWordTranslate(ctx context.Context, word string) (*blogs.Word, error) {
	//TODO implement me
	panic("implement me")
}

func (b WordDao) AddWord(ctx context.Context, word *blogs.Word) (int64, error) {
	var exitWord blogs.Word
	err := b.conn.Session(ctx).First(&exitWord, "word = ?", word.Word).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err := b.conn.Session(ctx).Save(word).Error
			if err != nil {
				return 0, err
			}
			return word.Id, nil
		} else {
			return 0, err
		}
	} else {
		// update
		word.Id = exitWord.Id
		err := b.conn.Session(ctx).Save(word).Error
		if err != nil {
			return 0, err
		}
		return word.Id, nil
	}

}

func (b WordDao) GetWords(ctx context.Context, page, pageSize int64) (*dto.WordList, error) {
	var res = &dto.WordList{
		Words: make([]*dto.WordInfo, 0),
		Total: 0,
	}
	var words []*blogs.Word
	err := b.conn.Session(ctx).Model(&blogs.Word{}).Count(&res.Total).Error
	if err != nil {
		return nil, err
	}
	err = b.conn.Session(ctx).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&words).Error
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&res.Words, &words)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b WordDao) DeleteWord(ctx context.Context, id int64) error {
	return b.conn.Session(ctx).Delete(&blogs.Word{}, "id = ?", id).Error
}

func (b WordDao) GetWordById(ctx context.Context, id int64) (*blogs.Word, error) {
	var word = &blogs.Word{}
	err := b.conn.Session(ctx).First(word, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return word, nil
}

func NewWordDao() *WordDao {
	return &WordDao{conn: gorms.New()}
}

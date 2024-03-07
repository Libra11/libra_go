package word

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"libra.com/blog/internal/dao"
	"libra.com/blog/internal/data/blogs"
	"libra.com/blog/internal/repo"
	"libra.com/blog/pkg/model"
	"libra.com/common/crawler"
	"libra.com/common/errs"
	"libra.com/grpc/service/word"
)

type ServiceWord struct {
	word.UnimplementedWordServiceServer
	cache repo.Cache
	blog  repo.WordRepo
}

func (l *ServiceWord) AddWord(ctx context.Context, message *word.AddWordRequest) (*word.AddWordResponse, error) {
	w := &blogs.Word{}
	_ = copier.Copy(w, message.Word)
	id, err := l.blog.AddWord(ctx, w)
	if err != nil {
		zap.L().Error("db AddWord error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormSetError)
	}
	return &word.AddWordResponse{
		Id: id,
	}, nil
}

func (l *ServiceWord) GetWords(ctx context.Context, message *word.GetWordsRequest) (*word.GetWordsResponse, error) {
	res, err := l.blog.GetWords(ctx, message.Page, message.PageSize)
	if err != nil {
		zap.L().Error("db GetWords error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormGetError)
	}
	wordsMsg := &word.GetWordsResponse{}
	if res != nil {
		_ = copier.Copy(&wordsMsg, &res)
	}
	return wordsMsg, nil
}

func (l *ServiceWord) DeleteWord(ctx context.Context, message *word.DeleteWordRequest) (*word.DeleteWordResponse, error) {
	err := l.blog.DeleteWord(ctx, message.Id)
	if err != nil {
		zap.L().Error("db DeleteWord error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormDeleteError)
	}
	return &word.DeleteWordResponse{}, nil
}

func (l *ServiceWord) GetWordById(ctx context.Context, message *word.GetWordByIdRequest) (*word.GetWordByIdResponse, error) {
	res, err := l.blog.GetWordById(ctx, message.Id)
	if err != nil {
		zap.L().Error("db GetWordById error", zap.Error(err))
		return nil, errs.GrpcErr(model.GormGetError)
	}
	a := &word.Word{}
	_ = copier.Copy(&a, &res)
	wordMsg := &word.GetWordByIdResponse{
		Word: a,
	}
	return wordMsg, nil
}

func (l *ServiceWord) GetWordTranslate(ctx context.Context, message *word.GetWordTranslateRequest) (*word.GetWordTranslateResponse, error) {
	translate := crawler.GetTranslate(message.Word)
	var w = &word.Word{
		Word:       translate.Title,
		Definition: translate.Definition,
		Phrase:     translate.Phrase,
		Example:    translate.Example,
		Phonetic:   translate.Phonetic,
	}
	return &word.GetWordTranslateResponse{
		Word: w,
	}, nil
}

func New() *ServiceWord {
	return &ServiceWord{
		cache: dao.Rc,
		blog:  dao.NewWordDao(),
	}
}

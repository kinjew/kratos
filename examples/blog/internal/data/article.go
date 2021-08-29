package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/examples/blog/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	/*
		ps, err := ar.data.db.Article.Query().All(ctx)
		if err != nil {
			return nil, err
		}
		rv := make([]*biz.Article, 0)
		for _, p := range ps {
			rv = append(rv, &biz.Article{
				Id:        p.ID,
				Title:     p.Title,
				Content:   p.Content,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
			})
		}
		return rv, nil
	*/
	var users []*biz.Article
	result := ar.data.db.Find(&users)
	fmt.Println(users, result)
	//rv := make([]*biz.Article, 0)
	return users, nil

}

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	/*
		p, err := ar.data.db.Article.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return &biz.Article{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}, nil
	*/

	var article *biz.Article
	ar.data.db.Where("id = ?", id).First(&article)

	return article, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	/*
		_, err := ar.data.db.Article.
			Create().
			SetTitle(article.Title).
			SetContent(article.Content).
			Save(ctx)
		return err
	*/

	user := biz.Article{Title: article.Title, Content: article.Content}
	tx := ar.data.db.Create(&user)
	return tx.Error

}

func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	/*
		p, err := ar.data.db.Article.Get(ctx, id)
		if err != nil {
			return err
		}
		_, err = p.Update().
			SetTitle(article.Title).
			SetContent(article.Content).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		return err

	*/

	// 根据 `map` 更新属性
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	var article1 *biz.Article
	tx := ar.data.db.Model(article1).Where("id = ?", id).Updates(map[string]interface{}{"title": article.Title, "content": article.Content})
	return tx.Error

}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	//return ar.data.db.Article.DeleteOneID(id).Exec(ctx)
	var article1 biz.Article
	tx := ar.data.db.Where("id = ?", id).Delete(&article1)

	return tx.Error
}

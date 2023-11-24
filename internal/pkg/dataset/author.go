package dataset

import (
	"poetry/internal/app"
	"poetry/internal/app/model"

	"github.com/tidwall/gjson"
)

var (
	authors = datasets[*model.Author]{
		{
			files: []string{
				"五代诗词/nantang/authors.json",
			},
			dynasty: "五代十国",
			parser: func(row gjson.Result) *model.Author {
				return &model.Author{
					Name: row.Get("name").String(),
					Desc: row.Get("desc").String(),
				}
			},
		},
		{
			files: []string{
				"全唐诗/authors.tang.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) *model.Author {
				return &model.Author{
					Name: row.Get("name").String(),
					Desc: row.Get("desc").String(),
				}
			},
		},
		{
			files: []string{
				"全唐诗/authors.song.json",
			},
			dynasty: "宋",
			parser: func(row gjson.Result) *model.Author {
				return &model.Author{
					Name: row.Get("name").String(),
					Desc: row.Get("desc").String(),
				}
			},
		},
		{
			files: []string{
				"宋词/author.song.json",
			},
			dynasty: "宋",
			parser: func(row gjson.Result) *model.Author {
				return &model.Author{
					Name:      row.Get("name").String(),
					Desc:      row.Get("description").String(),
					ShortDesc: row.Get("short_description").String(),
				}
			},
		},
	}
)

func insertAuthors(app *app.App, set *dataset[*model.Author], file string, result gjson.Result) error {
	app.Log.Infof("正在读取author文件: %s", file)

	dynasty := &model.Dynasty{
		Name: set.dynasty,
	}
	if err := insertDynasty(app, dynasty); err != nil {
		return err
	}

	authors := make(model.Authors, 0)
	for _, row := range result.Array() {
		author := set.parser(row)
		author.DynastyId = dynasty.Id

		if author.Dynasty != nil {
			if err := insertDynasty(app, author.Dynasty); err != nil {
				return err
			}
			author.DynastyId = author.Dynasty.Id
		}
		authors = append(authors, author)
	}
	return app.DB.CreateInBatches(authors, 1000).Error
}

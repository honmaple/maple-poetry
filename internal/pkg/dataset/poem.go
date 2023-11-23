package dataset

import (
	"poetry/internal/app"
	"poetry/internal/app/model"

	"github.com/tidwall/gjson"
)

var (
	poems = datasets[*model.Poem]{
		{
			name: "花间集",
			path: "五代诗词/huajianji",
			files: []string{
				"*.json",
			},
			dynasty: "五代十国",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Section: row.Get("rhythmic").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Note:    resultsToString(row.Get("notes").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "南唐",
			path: "五代诗词/nantang",
			files: []string{
				"poems.json",
			},
			dynasty: "五代十国",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Section: row.Get("rhythmic").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Note:    resultsToString(row.Get("notes").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "元曲",
			path: "元曲",
			files: []string{
				"yuanqu.json",
			},
			dynasty: "元",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "唐诗",
			path: "全唐诗",
			files: []string{
				"poet.tang.*.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "宋诗",
			path: "全唐诗",
			files: []string{
				"poet.song.*.json",
			},
			dynasty: "宋",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "宋词",
			path: "宋词",
			files: []string{
				"ci.song.*.json",
			},
			dynasty: "宋",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("rhythmic").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "幽梦影",
			path: "幽梦影",
			files: []string{
				"youmengying.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Note:    resultsToString(row.Get("comment").Array()),
					Title:   "幽梦影",
					Content: row.Get("content").String(),
					Author: &model.Author{
						Name: "张潮",
					},
				}
			},
		},
		{
			name: "御定全唐诗",
			path: "御定全唐诗/json",
			files: []string{
				"*.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Note:    resultsToString(row.Get("notes").Array()),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Chapter: row.Get("volume").String(),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "曹操诗集",
			path: "曹操诗集",
			files: []string{
				"caocao.json",
			},
			dynasty: "东汉",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: "曹操",
					},
				}
			},
		},
		{
			name: "楚辞",
			path: "楚辞",
			files: []string{
				"chuci.json",
			},
			dynasty: "战国",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Section: row.Get("section").String(),
					Content: resultsToString(row.Get("content").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "水墨唐诗",
			path: "水墨唐诗",
			files: []string{
				"shuimotangshi.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Note:    row.Get("prologue").String(),
					Title:   row.Get("title").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "纳兰性德",
			path: "纳兰性德",
			files: []string{
				"纳兰性德诗集.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Content: resultsToString(row.Get("para").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "论语",
			path: "论语",
			files: []string{
				"lunyu.json",
			},
			dynasty: "春秋",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("chapter").String(),
					Content: resultsToString(row.Get("paragraphs").Array()),
					Author: &model.Author{
						Name: "孔子",
					},
				}
			},
		},
		{
			name: "诗经",
			path: "诗经",
			files: []string{
				"shijing.json",
			},
			dynasty: "西周",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Chapter: row.Get("chapter").String(),
					Section: row.Get("section").String(),
					Content: resultsToString(row.Get("content").Array()),
					Author: &model.Author{
						Name: "佚名",
					},
				}
			},
		},
	}
)

func insertPoems(app *app.App, set *dataset[*model.Poem], file string, result gjson.Result) error {
	app.Log.Infof("正在读取poem文件: %s", file)

	dynasty := &model.Dynasty{
		Name: set.dynasty,
	}
	if err := insertDynasty(app, dynasty); err != nil {
		return err
	}

	collection := &model.Collection{
		Name:      set.name,
		DynastyId: dynasty.Id,
	}
	if err := insertCollection(app, collection); err != nil {
		return err
	}

	poems := make(model.Poems, 0)
	for _, row := range result.Array() {
		poem := set.parser(row)
		poem.DynastyId = dynasty.Id
		poem.CollectionId = collection.Id

		if poem.Dynasty != nil {
			if err := insertDynasty(app, poem.Dynasty); err != nil {
				return err
			}
			poem.DynastyId = poem.Dynasty.Id
		}
		if poem.Author != nil {
			poem.Author.DynastyId = poem.DynastyId

			if err := insertAuthor(app, poem.Author); err != nil {
				return err
			}
			poem.AuthorId = poem.Author.Id
		}
		poems = append(poems, poem)
	}
	return app.DB.CreateInBatches(poems, 1000).Error
}

package dataset

import (
	"fmt"
	"poetry/internal/app"
	"poetry/internal/app/model"

	"github.com/tidwall/gjson"
)

var (
	poems = datasets[*model.Poem]{
		{
			name: "诗经",
			files: []string{
				"诗经/shijing.json",
			},
			dynasty: "西周",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Chapter: fmt.Sprintf("%s · %s", row.Get("chapter").String(), row.Get("section").String()),
					Content: resultsToString(row.Get("content").Array()),
					Author: &model.Author{
						Name: "佚名",
					},
				}
			},
		},
		{
			name: "论语",
			files: []string{
				"论语/lunyu.json",
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
			name: "楚辞",
			files: []string{
				"楚辞/chuci.json",
			},
			dynasty: "战国",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:   row.Get("title").String(),
					Chapter: row.Get("section").String(),
					Content: resultsToString(row.Get("content").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "曹操诗集",
			files: []string{
				"曹操诗集/caocao.json",
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
			name: "花间集",
			files: []string{
				"五代诗词/huajianji/*.json",
			},
			dynasty: "五代十国",
			parser: func(row gjson.Result) *model.Poem {
				title := row.Get("title").String()
				chapter := row.Get("rhythmic").String()
				return &model.Poem{
					Title:      trimTitle(title, chapter),
					Chapter:    chapter,
					Content:    resultsToString(row.Get("paragraphs").Array()),
					Annotation: resultsToString(row.Get("notes").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "南唐二主词",
			desc: `《南唐二主词》，系南唐中主李璟、后主李煜撰。约成书于南宋，后世续有辑补，又有后人编写了各种版本\n
南唐二主李璟、李煜是中国词史上极少数受到社会各阶层民众普遍喜爱的词人。其词突破五代花间词堆金砌玉的壁垒，前者多用比兴，妙能沉郁，后者全用赋体，超放自然，丝毫没有情感的做作。尤其是李煜，后期遭遇亡国之痛，词作纯从血泪中迸出，绝少雕琢，有很高的审美价值。`,
			files: []string{
				"五代诗词/nantang/poetrys.json",
			},
			dynasty: "五代十国",
			parser: func(row gjson.Result) *model.Poem {
				title := row.Get("title").String()
				chapter := row.Get("rhythmic").String()
				return &model.Poem{
					Title:      trimTitle(title, chapter),
					Chapter:    chapter,
					Content:    resultsToString(row.Get("paragraphs").Array()),
					Annotation: resultsToString(row.Get("notes").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "唐诗",
			files: []string{
				"全唐诗/poet.tang.*.json",
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
			name: "水墨唐诗",
			files: []string{
				"水墨唐诗/shuimotangshi.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:      row.Get("title").String(),
					Content:    resultsToString(row.Get("paragraphs").Array()),
					Annotation: row.Get("prologue").String(),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "御定全唐诗",
			files: []string{
				"御定全唐诗/json/*.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:      row.Get("title").String(),
					Chapter:    row.Get("volume").String(),
					Content:    resultsToString(row.Get("paragraphs").Array()),
					Annotation: resultsToString(row.Get("notes").Array()),
					Author: &model.Author{
						Name: row.Get("author").String(),
					},
				}
			},
		},
		{
			name: "宋诗",
			files: []string{
				"全唐诗/poet.song.*.json",
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
			files: []string{
				"宋词/ci.song.*.json",
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
			name: "元曲",
			files: []string{
				"元曲/yuanqu.json",
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
			name: "幽梦影",
			files: []string{
				"幽梦影/youmengying.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) *model.Poem {
				return &model.Poem{
					Title:      "幽梦影",
					Content:    row.Get("content").String(),
					Annotation: resultsToString(row.Get("comment").Array()),
					Author: &model.Author{
						Name: "张潮",
					},
				}
			},
		},
		{
			name: "纳兰性德",
			files: []string{
				"纳兰性德/纳兰性德诗集.json",
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
		for _, tag := range poem.Tags {
			if err := insertTag(app, tag); err != nil {
				return err
			}
		}
		poems = append(poems, poem)
	}
	return app.DB.CreateInBatches(poems, 1000).Error
}

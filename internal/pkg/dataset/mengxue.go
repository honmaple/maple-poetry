package dataset

import (
	"fmt"
	"strings"

	"poetry/internal/app"
	"poetry/internal/app/model"

	"github.com/tidwall/gjson"
)

var (
	mengxues = datasets[model.Poems]{
		{
			name: "三字經",
			path: "蒙学",
			files: []string{
				"sanzijing-traditional.json",
			},
			dynasty: "南宋",
			parser: func(row gjson.Result) model.Poems {
				return model.Poems{
					{
						Title:   row.Get("title").String(),
						Content: resultsToString(row.Get("paragraphs").Array()),
						Author: &model.Author{
							Name: row.Get("author").String(),
						},
					},
				}
			},
		},
		{
			name: "百家姓",
			path: "蒙学",
			files: []string{
				"baijiaxing.json",
			},
			dynasty: "北宋",
			parser: func(row gjson.Result) model.Poems {
				origins := row.Get("origin").Array()

				notes := make([]string, len(origins))
				for i, origin := range origins {
					notes[i] = fmt.Sprintf("%s：%s", origin.Get("surname").String(), origin.Get("place").String())
				}
				return model.Poems{
					{
						Title:   row.Get("title").String(),
						Content: resultsToString(row.Get("paragraphs").Array()),
						Note:    strings.Join(notes, "\n"),
						Author: &model.Author{
							Name: row.Get("author").String(),
						},
					},
				}
			},
		},
		{
			name: "千字文",
			path: "蒙学",
			files: []string{
				"qianziwen.json",
			},
			dynasty: "南北朝",
			parser: func(row gjson.Result) model.Poems {
				return model.Poems{
					{
						Title:   row.Get("title").String(),
						Content: resultsToString(row.Get("paragraphs").Array()),
						Note:    resultsToString(row.Get("spells").Array()),
						Author: &model.Author{
							Name: row.Get("author").String(),
						},
					},
				}
			},
		},
		{
			name: "弟子規",
			path: "蒙学",
			files: []string{
				"dizigui.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				title := row.Get("title").String()
				author := row.Get("author").String()

				poems := make(model.Poems, 0)
				for _, c := range row.Get("content").Array() {
					poem := &model.Poem{
						Title:   title,
						Chapter: c.Get("chapter").String(),
						Content: resultsToString(c.Get("paragraphs").Array()),
						Author: &model.Author{
							Name: author,
						},
					}
					poems = append(poems, poem)
				}
				return poems
			},
		},
		{
			name: "幼學瓊林",
			path: "蒙学",
			files: []string{
				"youxueqionglin.json",
			},
			dynasty: "明",
			parser: func(row gjson.Result) model.Poems {
				author := row.Get("author").String()

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					title := first.Get("title").String()
					for _, second := range first.Get("content").Array() {
						poem := &model.Poem{
							Title:   title,
							Chapter: second.Get("chapter").String(),
							Content: resultsToString(second.Get("paragraphs").Array()),
							Author: &model.Author{
								Name: author,
							},
						}
						poems = append(poems, poem)
					}
				}
				return poems
			},
		},
		{
			name: "朱子家訓",
			path: "蒙学",
			files: []string{
				"zhuzijiaxun.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				return model.Poems{
					{
						Title:   row.Get("title").String(),
						Content: resultsToString(row.Get("paragraphs").Array()),
						Author: &model.Author{
							Name: row.Get("author").String(),
						},
					},
				}
			},
		},
		{
			name: "古文觀止",
			path: "蒙学",
			files: []string{
				"guwenguanzhi.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					title := first.Get("title").String()
					for _, second := range first.Get("content").Array() {
						poem := &model.Poem{
							Title:   title,
							Chapter: second.Get("chapter").String(),
							Content: resultsToString(second.Get("paragraphs").Array()),
							Author: &model.Author{
								Name: second.Get("author").String(),
							},
						}
						poems = append(poems, poem)
					}
				}
				return poems
			},
		},
		{
			name: "聲律啟蒙",
			path: "蒙学",
			files: []string{
				"shenglvqimeng.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				author := row.Get("author").String()

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					title := first.Get("title").String()
					for _, second := range first.Get("content").Array() {
						poem := &model.Poem{
							Title:   title,
							Chapter: second.Get("chapter").String(),
							Content: resultsToString(second.Get("paragraphs").Array()),
							Author: &model.Author{
								Name: author,
							},
						}
						poems = append(poems, poem)
					}
				}
				return poems
			},
		},
		{
			name: "文字蒙求",
			path: "蒙学",
			files: []string{
				"wenzimengqiu.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					poem := &model.Poem{
						Title:   first.Get("title").String(),
						Content: resultsToString(first.Get("paragraphs").Array()),
						Author: &model.Author{
							Name: "王筠",
							Desc: row.Get("author").String(),
						},
					}
					poems = append(poems, poem)
				}
				return poems
			},
		},
		{
			name: "增廣賢文",
			path: "蒙学",
			files: []string{
				"zengguangxianwen.json",
			},
			dynasty: "明",
			parser: func(row gjson.Result) model.Poems {
				title := row.Get("title").String()
				author := row.Get("author").String()

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					poem := &model.Poem{
						Title:   title,
						Chapter: first.Get("chapter").String(),
						Content: resultsToString(first.Get("paragraphs").Array()),
						Author: &model.Author{
							Name: author,
						},
					}
					poems = append(poems, poem)
				}
				return poems
			},
		},
	}
)

func insertMengxues(app *app.App, set *dataset[model.Poems], file string, result gjson.Result) error {
	app.Log.Infof("正在读取mengxue文件: %s", file)

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
	if desc := result.Get("abstract").String(); desc != "" {
		collection.Desc = desc
	}
	if err := insertCollection(app, collection); err != nil {
		return err
	}

	poems := set.parser(result)
	for _, poem := range poems {
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
	}
	return app.DB.CreateInBatches(poems, 1000).Error
}

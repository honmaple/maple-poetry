package dataset

import (
	"fmt"
	"regexp"
	"strings"

	"poetry/internal/app"
	"poetry/internal/app/model"

	"github.com/tidwall/gjson"
)

var (
	mengxues = datasets[model.Poems]{
		{
			name: "三字經",
			files: []string{
				"蒙学/sanzijing-traditional.json",
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
			files: []string{
				"蒙学/baijiaxing.json",
			},
			dynasty: "北宋",
			parser: func(row gjson.Result) model.Poems {
				origins := row.Get("origin").Array()

				anns := make([]string, len(origins))
				for i, origin := range origins {
					anns[i] = fmt.Sprintf("%s：%s", origin.Get("surname").String(), origin.Get("place").String())
				}
				return model.Poems{
					{
						Title:      row.Get("title").String(),
						Content:    resultsToString(row.Get("paragraphs").Array()),
						Annotation: strings.Join(anns, "\n"),
						Author: &model.Author{
							Name: row.Get("author").String(),
						},
					},
				}
			},
		},
		{
			name: "千字文",
			files: []string{
				"蒙学/qianziwen.json",
			},
			dynasty: "南北朝",
			parser: func(row gjson.Result) model.Poems {
				return model.Poems{
					{
						Title:      row.Get("title").String(),
						Content:    resultsToString(row.Get("paragraphs").Array()),
						Annotation: resultsToString(row.Get("spells").Array()),
						Author: &model.Author{
							Name: row.Get("author").String(),
						},
					},
				}
			},
		},
		{
			name: "弟子規",
			files: []string{
				"蒙学/dizigui.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				title := row.Get("title").String()
				author := row.Get("author").String()

				poems := make(model.Poems, 0)
				for _, c := range row.Get("content").Array() {
					poem := &model.Poem{
						Title:   c.Get("chapter").String(),
						Chapter: title,
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
			files: []string{
				"蒙学/youxueqionglin.json",
			},
			dynasty: "明",
			parser: func(row gjson.Result) model.Poems {
				author := row.Get("author").String()

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					title := first.Get("title").String()
					for _, second := range first.Get("content").Array() {
						poem := &model.Poem{
							Title:   second.Get("chapter").String(),
							Chapter: title,
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
			files: []string{
				"蒙学/zhuzijiaxun.json",
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
			name: "千家詩",
			files: []string{
				"蒙学/qianjiashi.json",
			},
			dynasty: "南宋",
			parser: func(row gjson.Result) model.Poems {
				// "author": "（唐）王維",
				re := regexp.MustCompile(`^（(.+?)）(.+)`)

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					tags := model.Tags{
						{
							Name: first.Get("type").String(),
							Kind: model.TAG_KIND_STYLE,
						},
					}
					for _, second := range first.Get("content").Array() {
						var (
							title = second.Get("chapter").String()

							dynasty *model.Dynasty
							author  = &model.Author{
								Name: second.Get("author").String(),
							}
						)
						if match := re.FindStringSubmatch(author.Name); match != nil {
							dynasty = &model.Dynasty{
								Name: match[1],
							}
							author.Name = match[2]
						}

						length := len(poems)
						for _, para := range second.Get("paragraphs").Array() {
							if para.Type == gjson.String {
								break
							}
							poem := &model.Poem{
								Title:   title + " " + para.Get("subchapter").String(),
								Content: resultsToString(para.Get("paragraphs").Array()),
								Author:  author,
								Dynasty: dynasty,
								Tags:    tags,
							}
							poems = append(poems, poem)
						}
						if len(poems) == length {
							poem := &model.Poem{
								Title:   title,
								Content: resultsToString(second.Get("paragraphs").Array()),
								Author:  author,
								Dynasty: dynasty,
								Tags:    tags,
							}
							poems = append(poems, poem)
						}
					}
				}
				return poems
			},
		},
		{
			name: "古文觀止",
			files: []string{
				"蒙学/guwenguanzhi.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				// "author": "先秦：左丘明 "
				re := regexp.MustCompile(`^(.+?)：(.*?)\s*?$`)

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					title := first.Get("title").String()
					for _, second := range first.Get("content").Array() {
						var (
							dynasty *model.Dynasty
							author  = &model.Author{
								Name: second.Get("author").String(),
							}
						)
						if match := re.FindStringSubmatch(author.Name); match != nil {
							dynasty = &model.Dynasty{
								Name: match[1],
							}
							author.Name = match[2]
						}

						poem := &model.Poem{
							Title:   second.Get("chapter").String(),
							Chapter: title,
							Content: resultsToString(second.Get("paragraphs").Array()),
							Author:  author,
							Dynasty: dynasty,
						}
						poems = append(poems, poem)
					}
				}
				return poems
			},
		},
		{
			name: "唐詩三百首",
			files: []string{
				"蒙学/tangshisanbaishou.json",
			},
			dynasty: "唐",
			parser: func(row gjson.Result) model.Poems {
				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					tags := model.Tags{
						{
							Name: first.Get("type").String(),
							Kind: model.TAG_KIND_STYLE,
						},
					}
					for _, second := range first.Get("content").Array() {
						poem := &model.Poem{
							Title:   second.Get("chapter").String(),
							Content: resultsToString(second.Get("paragraphs").Array()),
							Author: &model.Author{
								Name: second.Get("author").String(),
							},
							Tags: tags,
						}
						poems = append(poems, poem)
					}
				}
				return poems
			},
		},
		{
			name: "聲律啟蒙",
			files: []string{
				"蒙学/shenglvqimeng.json",
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
			files: []string{
				"蒙学/wenzimengqiu.json",
			},
			dynasty: "清",
			parser: func(row gjson.Result) model.Poems {
				title := row.Get("title").String()
				author := &model.Author{
					Name: "王筠",
					Desc: row.Get("author").String(),
				}

				poems := make(model.Poems, 0)
				for _, first := range row.Get("content").Array() {
					poem := &model.Poem{
						Title:   first.Get("title").String(),
						Chapter: title,
						Content: resultsToString(first.Get("paragraphs").Array()),
						Author:  author,
					}
					poems = append(poems, poem)
				}
				return poems
			},
		},
		{
			name: "增廣賢文",
			files: []string{
				"蒙学/zengguangxianwen.json",
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
		for _, tag := range poem.Tags {
			if err := insertTag(app, tag); err != nil {
				return err
			}
		}
	}
	return app.DB.CreateInBatches(poems, 1000).Error
}

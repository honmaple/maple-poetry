package dataset

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"poetry/internal/app"
	"poetry/internal/app/model"

	"github.com/tidwall/gjson"
)

type dataset[V any] struct {
	name    string
	path    string
	files   []string
	parser  func(gjson.Result) V
	dynasty string
}

func (set *dataset[V]) init(app *app.App, fn func(*app.App, *dataset[V], string, gjson.Result) error) error {
	root := app.Config.GetString("dataset.path")
	for _, file := range set.files {
		matches, err := filepath.Glob(filepath.Join(root, set.path, file))
		if err != nil {
			return err
		}
		for _, match := range matches {
			content, err := ioutil.ReadFile(match)
			if err != nil {
				return err
			}
			if err := fn(app, set, match, gjson.ParseBytes(content)); err != nil {
				return err
			}
		}
	}
	return nil
}

type datasets[V any] []dataset[V]

func (sets datasets[V]) init(app *app.App, fn func(*app.App, *dataset[V], string, gjson.Result) error) error {
	for _, set := range sets {
		if err := set.init(app, fn); err != nil {
			return err
		}
	}
	return nil
}

func resultsToString(results []gjson.Result) string {
	strs := make([]string, len(results))
	for i, result := range results {
		strs[i] = result.String()
	}
	return strings.Join(strs, "\n")
}

func insertCollection(app *app.App, collection *model.Collection) error {
	if collection == nil {
		return nil
	}
	q := app.DB.Where("name = ?", collection.Name)
	if collection.DynastyId > 0 {
		q = q.Where("dynasty_id = ?", collection.DynastyId)
	} else if collection.Dynasty != nil {
		q = q.Where("dynasty_id = ?", app.DB.Model(model.Dynasty{}).Select("id").Where("name = ?", collection.Dynasty.Name))
	}
	result := q.FirstOrCreate(collection)
	return result.Error
}

func insertDynasty(app *app.App, dynasty *model.Dynasty) error {
	if dynasty == nil {
		return nil
	}
	result := app.DB.Where("name = ?", dynasty.Name).FirstOrCreate(dynasty)
	return result.Error
}

func insertAuthor(app *app.App, author *model.Author) error {
	if author == nil {
		return nil
	}
	q := app.DB.Where("name = ?", author.Name)
	if author.DynastyId > 0 {
		q = q.Where("dynasty_id = ?", author.DynastyId)
	} else if author.Dynasty != nil {
		q = q.Where("dynasty_id = ?", app.DB.Model(model.Dynasty{}).Select("id").Where("name = ?", author.Dynasty.Name))
	}
	result := q.FirstOrCreate(author)
	return result.Error
}

func Init(app *app.App) error {
	if err := authors.init(app, insertAuthors); err != nil {
		return err
	}

	if err := poems.init(app, insertPoems); err != nil {
		return err
	}

	if err := mengxues.init(app, insertMengxues); err != nil {
		return err
	}
	return nil
}

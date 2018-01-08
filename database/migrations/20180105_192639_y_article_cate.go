package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YArticleCate_20180105_192639 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YArticleCate_20180105_192639{}
	m.Created = "20180105_192639"

	migration.Register("YArticleCate_20180105_192639", m)
}

// Run the migrations
func (m *YArticleCate_20180105_192639) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_article_cate ("+
		"art_id integer REFERENCES y_articles(id)," +
		"cate_id integer REFERENCES y_categories(id)" +
		")")
}

// Reverse the migrations
func (m *YArticleCate_20180105_192639) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

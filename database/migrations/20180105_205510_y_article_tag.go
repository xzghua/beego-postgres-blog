package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YArticleTag_20180105_205510 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YArticleTag_20180105_205510{}
	m.Created = "20180105_205510"

	migration.Register("YArticleTag_20180105_205510", m)
}

// Run the migrations
func (m *YArticleTag_20180105_205510) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_article_tag ("+
		"art_id integer REFERENCES y_articles(id)," +
		"tag_id integer REFERENCES y_tags(id)" +
		")")
}

// Reverse the migrations
func (m *YArticleTag_20180105_205510) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

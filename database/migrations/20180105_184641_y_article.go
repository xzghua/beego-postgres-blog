package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YArticle_20180105_184641 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YArticle_20180105_184641{}
	m.Created = "20180105_184641"

	migration.Register("YArticle_20180105_184641", m)
}

// Run the migrations
func (m *YArticle_20180105_184641) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_articles("+
		"id SERIAL primary key," +
		"title varchar(255) NOT NULL," +
		"content text default null," +
		"body_original text default null," +
		"user_id integer NOT NULL," +
		"password smallint default 0," +
		"note varchar(255) default NULL," +
		"read_status smallint DEFAULT 1," +
		"top boolean default false," +
		"abstract varchar(500) default null," +
		"view_num integer default 0," +
		"created_at timestamp default current_timestamp," +
		"updated_at timestamp default current_timestamp," +
		"deleted_at timestamp default null" +
		");create  index article_index on y_articles (user_id)")

}

// Reverse the migrations
func (m *YArticle_20180105_184641) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE y_article")
}

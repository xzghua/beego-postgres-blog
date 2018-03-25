package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YTag_20180105_194406 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YTag_20180105_194406{}
	m.Created = "20180105_194406"

	migration.Register("YTag_20180105_194406", m)
}

// Run the migrations
func (m *YTag_20180105_194406) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_tags("+
		"id SERIAL primary key," +
		"name  varchar(255) NOT NULL," +
		"tag_num integer default 0" +
		");create  index name on y_tags (name)")
}

// Reverse the migrations
func (m *YTag_20180105_194406) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

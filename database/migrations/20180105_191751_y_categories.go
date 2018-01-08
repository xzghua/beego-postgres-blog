package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YCategories_20180105_191751 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YCategories_20180105_191751{}
	m.Created = "20180105_191751"

	migration.Register("YCategories_20180105_191751", m)
}

// Run the migrations
func (m *YCategories_20180105_191751) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_categories("+
		"id SERIAL primary key," +
		"name varchar(255) NOT NULL," +
		"display_name varchar(255) NOT NULL," +
		"parent_id  int default 0," +
		"description varchar(2048) default null," +
		"created_at timestamp default current_timestamp," +
		"updated_at timestamp default current_timestamp" +
		");" +
		"create index y_categories_parent_id_index on y_categories (parent_id);" +
		"create index y_categories_name_index on y_categories (name);")
}

// Reverse the migrations
func (m *YCategories_20180105_191751) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

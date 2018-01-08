package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YLinks_20180105_205730 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YLinks_20180105_205730{}
	m.Created = "20180105_205730"

	migration.Register("YLinks_20180105_205730", m)
}

// Run the migrations
func (m *YLinks_20180105_205730) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_links("+
		"id SERIAL primary key," +
		"name varchar(255) NOT NULL," +
		"link varchar(255) NOT NULL," +
		"sort smallint default 1," +
		"created_at timestamp default current_timestamp," +
		"updated_at timestamp default current_timestamp" +
		")")
}

// Reverse the migrations
func (m *YLinks_20180105_205730) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

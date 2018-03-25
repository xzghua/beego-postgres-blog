package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20180105_163508 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20180105_163508{}
	m.Created = "20180105_163508"

	migration.Register("User_20180105_163508", m)
}

// Run the migrations
func (m *User_20180105_163508) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_users(" +
		"id serial primary key," +
		"name varchar(255) NOT NULL," +
		"email varchar(255) unique NOT NULL," +
		"password varchar(255) NOT NULL," +
		"created_at timestamp default current_timestamp," +
		"updated_at timestamp default current_timestamp" +
		");create index index on y_users (email)")
}

// Reverse the migrations
func (m *User_20180105_163508) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user")
}

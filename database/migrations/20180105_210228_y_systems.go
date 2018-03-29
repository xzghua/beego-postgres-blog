package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type YSystems_20180105_210228 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &YSystems_20180105_210228{}
	m.Created = "20180105_210228"

	migration.Register("YSystems_20180105_210228", m)
}

// Run the migrations
func (m *YSystems_20180105_210228) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE y_systems("+
		"id SERIAL primary key," +
		"title varchar(255) default '叶落山城秋'," +
		"s_title varchar(255) default '叶落山城秋'," +
		"description varchar(255) default '叶落山城秋博客'," +
		"seo_key varchar(255) default '叶落山城秋'," +
		"seo_des varchar(255) default '叶落山城秋'," +
		"record_number varchar(30) default '京ICP备 88888888号'," +
		"created_at timestamp default current_timestamp," +
		"updated_at timestamp default current_timestamp" +
		")")
}

// Reverse the migrations
func (m *YSystems_20180105_210228) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE y_systems")
}

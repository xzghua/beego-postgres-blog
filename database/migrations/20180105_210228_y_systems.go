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
		"comment_type smallint default 1, " +
		"github_client_secret varchar(50) default '29x9s92s0d0s9w9sx9d9w9szqa9wq9'," +
		"github_client_id varchar(30) default '200x0ws2k3982ks'," +
		"github_name varchar(100) default 'github_name'," +
		"github_repo varchar(100) default 'github_repo'," +
		"cy_app_id varchar(30) default '2kks9w9s'," +
		"cy_app_key varchar(50) default 'kski29sk298al843hasd83ijsu38j2js82',"	+
		"cdn_type smallint default 1," +
		"cdn_url varchar(255) default 'https://www.iphpt.com'," +
		"created_at timestamp default current_timestamp," +
		"updated_at timestamp default current_timestamp" +
		")")
}

// Reverse the migrations
func (m *YSystems_20180105_210228) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE y_systems")
}

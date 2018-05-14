package common



type PostCreate struct {
	Title 			string 	`form:"title" valid:"Required;MaxSize(33)"`
	Abstract 		string	`form:"abstract" valid:"Required;MaxSize(100)"`
	Category 		int64 	`form:"category" valid:"Required"`
	Tag 			string 	`form:"tag" valid:"Required"`
	Content 		string 	`form:"editormd-html-code" valid:"Required"`
	BodyOriginal 	string 	`form:"content" valid:"Required"`
}

type CateRequest struct {
	ParentId		string 	`form:"parentId" valid:"Required"`
	Name 			string 	`form:"name" valid:"Required;MaxSize(30)"`
	DisplayName 	string 	`form:"displayName" valid:"Required;MaxSize(30)"`
	Description 	string 	`form:"description" valid:"MaxSize(150)"`
}

type LinkCreate struct {
	Name 	string 	`form:"name" valid:"Required;MaxSize(23)"`
	Link 	string 	`form:"link" valid:"Required;MaxSize(100)"`
	Sort 	int64   `form:"ordering" valid:"Required"`
}


type SystemUpdate struct {
	Title				string 		`form:"title" valid:"Required;MaxSize(23)"`
	STitle				string 		`form:"s_title" `
	Description 		string 		`form:"description" `
	SeoKey				string 		`form:"seo_key" `
	SeoDes				string 		`form:"seo_des" `
	RecordNumber		string 		`form:"record_number" `
	CommentType			int 		`form:"comment_type"`
	GithubClientSecret 	string		`form:"github_client_secret"`
	GithubClientId		string		`form:"github_client_id"`
	GithubName			string		`form:"github_name"`
	GithubRepo			string		`form:"github_repo"`
	CyAppId				string 		`form:"cy_app_id"`
	CyAppKey			string 		`form:"cy_app_key"`
	CdnType 			int 		`form:"cdn_type"`
	CdnUrl				string		`form:"cdn_url"`
}

type TagRequest struct {
	Name 	string `form:"name" valid:"Required;MaxSize(30)"`
}

type RegisterRequest struct {
	Name		string 	`form:"name" valid:"Required;MaxSize(50)"`
	Email		string 	`form:"email" valid:"Required;MaxSize(50);Email"`
	Password	string 	`form:"password" valid:"Required;MaxSize(30)"`
}

type LoginRequest struct {
	Name		string 	`form:"name" valid:"Required;MaxSize(50)"`
	Password	string 	`form:"password" valid:"Required;MaxSize(30)"`
}
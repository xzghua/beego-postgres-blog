package validation

import (
	"strconv"
)

var key = map[string][2]string {
	"Title"			: { "001", "标题" },
	"Category"		: { "002", "节点" },
	"Tag" 			: { "003", "标签" },
	"Content"		: { "004", "内容" },
	"Abstract" 		: { "005", "摘要" },
	"Name"			: {	"006", "name名" 	 },
	"DisplayName" 	: { "007", "别名" },
	"Description"	: { "008", "描述" },
	"Sort"			: { "009", "排序值" },
	"Link"			: { "010", "链接" },
	"STitle"		: { "011", "SEO标题" },
	"SeoKey"		: { "012", "SEO标题" },
	"SeoDes"		: { "013", "SEO标题" },
	"RecordNumber"	: { "014", "SEO标题" },
	"Email"			: { "015", "用户名" },
	"Password"		: { "016", "密码" },
}

var condition = map[string][2]string {
	"Required"	: { "001", "不能为空" },
	"MaxSize"	: { "002", "超出最大范围" },
	"Email"		: { "003", "邮箱格式不正确" },
}


var env = map[string][2]string {
	"Post"		: { "4001", "帖子" },
	"Cate"		: { "4002", "节点" },
	"Tag"		: { "4003", "标签" },
	"Link"		: {	"4004", "友链" },
	"System"	: { "4005", "设置" },
	"Register"	: { "4006", "注册" },
	"Login"		: { "4007", "登录" },
}


func CustomErrCodeAndMessage (errEnv string,errKey string,errCondition string) (int,string) {
	codeEnv,okOne := env[errEnv]
	codeKey,okTwo := key[errKey]
	codeCondition,okThree := condition[errCondition]
	if !okOne || !okTwo || !okThree {
		return 4000000000,"系统内部错误"
	}
	codeString := codeEnv[0] + codeKey[0] + codeCondition[0]
	message := codeKey[1] + codeCondition[1]
	code,_:= strconv.Atoi(codeString)
	return code,message
}


package services

import "bee-go-myBlog/models"

func GetMyAllPost(page int64) (ml []interface{}, err error){
	post,err := models.AllArticle(page)

	var userIds []int64
	if post != nil {
		for key, val := range post {
			userId := val.(map[string]interface{})["Id"].(int64)
			userIds = append(userIds,userId)
			user,_ := models.GetUsersById(userId)
			post[key].(map[string]interface{})["user_name"] = user.Name
			post[key].(map[string]interface{})["user_id"] = user.Id
		}
	}
	return post,err
}

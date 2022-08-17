package api

import (
	"blok/common"
	"blok/models"
	"blok/service"
	"blok/utils"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter,r *http.Request) {
	//获取用户id，判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//POST  save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)

	}

}
func (*Api) SearchPost(w http.ResponseWriter,r *http.Request)  {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w,searchResp)
}

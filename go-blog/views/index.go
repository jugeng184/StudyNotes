package views

import (
	"blok/common"
	"blok/service"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter,r *http.Request)  {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err :=r.ParseForm();err!=nil{
		index.WriteError(w,errors.New("系统错误，请联系管理员!!"))
        return
	}
	pageStr :=r.Form.Get("page")
	page:=1
	if pageStr!=""{
		page,_=strconv.Atoi(pageStr)
	}
	pageSize:=10
	hr,err := service.GetAllIndexInfo(page,pageSize)
	if err != nil {
		log.Println("Index获取数据出错：",err)
		index.WriteError(w,errors.New("系统错误，请联系管理员!!"))
	}
	index.WriteData(w,hr)
}


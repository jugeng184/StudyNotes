package views

import (
	"blok/common"
	"blok/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request)  {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w,wr)
}
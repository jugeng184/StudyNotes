package views

import (
	"blok/common"
	"blok/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter,r *http.Request)  {
	login := common.Template.Login

	login.WriteData(w,config.Cfg.Viewer)
}
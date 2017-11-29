package service

import (
	"net/http"
	"strconv"

	"github.com/Mensu/cloudgo-data/entities"

	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			panic(err)
		}
		if len(req.Form["username"]) == 0 || len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		if len(req.Form["departname"]) > 0 && len(req.Form["departname"][0]) > 0 {
			u.DepartName = req.Form["departname"][0]
		}
		var service entities.IUserInfoService
		if useOrm := len(req.Form["orm"]) > 0; useOrm {
			service = &entities.UserInfoOrmService
		} else {
			service = &entities.UserInfoService
		}
		service.Save(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			panic(err)
		}
		var service entities.IUserInfoService
		if useOrm := len(req.Form["orm"]) > 0; useOrm {
			service = &entities.UserInfoOrmService
		} else {
			service = &entities.UserInfoService
		}
		if len(req.Form["userid"]) > 0 && len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			u := service.FindByID(int(i))
			formatter.JSON(w, http.StatusBadRequest, u)
			return
		}
		ulist := service.FindAll()
		formatter.JSON(w, http.StatusOK, ulist)
	}
}

package controller

import (
	"Sit/app/model"
	"encoding/json"
	"net/http"

	"html/template"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	u := filepath.Join("public", "html", "usersDPage.html")
	common := filepath.Join("public", "html", "common.html")

	tmpl, err := template.ParseFiles(u, common)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// tmpl, err := template.ParseFiles(u), common)
	// if err != nil {
	// 	http.Error(w, err.Error(), 400)
	// 	return
	// }

	err = tmpl.ExecuteTemplate(w, "users", users)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	if name == "" || surname == "" {
		http.Error(w, "Имя и фамилия не могут быть пустыми", 400)
		return
	}

	user := model.NewUser(name, surname)
	err := user.Add()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = json.NewEncoder(w).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")

	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = user.Delete()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.NewEncoder(w).Encode("Пользователь был успешно удален")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user.Name = name
	user.Surname = surname

	err = user.Update()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.NewEncoder(w).Encode("Пользователь был успешно изменен")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

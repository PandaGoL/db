package main

import (
	"Sit/app/controller"
	"Sit/app/server"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	//создаем и запускаем в работу роутер для обслуживания запросов
	r := httprouter.New()
	routes(r)
	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	log.Fatal(http.ListenAndServe(":4444", r))
}

func routes(r *httprouter.Router) {
	//путь к папке с внешними файлами: html, js, css, изображения итд.
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	//что стоит выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
	r.POST("/user/add", controller.AddUser)
	r.DELETE("/user/delete/:userId", controller.DeleteUser)
	r.POST("/user/update/userId", controller.UpdateUser)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//"database/sql"
	"github.com/facebookgo/inject"
	"github.com/r-a-x/digo/controller"
	"os"
	"net/http"
)

type App struct {

	Movie *controller.Movie `inject:""`
	//DB sql.DB

}

//func (a *App) init(){
//	a.Router = mux.NewRouter()
//	//a.DB = sql.R
//}

func main(){

	var g inject.Graph

	var app App

	Router := mux.NewRouter()

	_ = g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: Router},
	)

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app.Movie.Register()
	fmt.Printf("The value of the movies is",app.Movie.Router)

	fmt.Printf("Welocome to the Main method!! Lets start working !!!")

	http.ListenAndServe("8080",Router)

}


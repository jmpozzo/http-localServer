package main

import (
	"fmt"
	"net/http"
	//"html/template"
)
var path string 
func main (){
	http.HandleFunc("/", httpServerRetoma)
	fmt.Println("Hola mundillo")
	http.ListenAndServe(":8080",nil)
}

func httpServerRetoma (w http.ResponseWriter, r *http.Request){
	req := r;
	fmt.Println(req.URL)
	fmt.Println(req.RemoteAddr)
	switch req.Method{
		case "GET":
			path = req.URL.Path
			if path == "/" {
				fmt.Println(path)
				path="."
			} else {
				path = "."+path
			}
			http.ServeFile(w,r,path)
		break
		case "POST":
			fmt.Println("Todavia no")
		break
	}
}
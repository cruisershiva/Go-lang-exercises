
package main

import (
	"fmt"   //
	"log"
	"net/http" //it contains most of the functions to create web server
)

func formHandler(w http.ResponseWriter, r *http.Request){

	if err:=r.ParseForm(); err != nil{
	fmt.Fprintf(w,"parseform() err: %v", err)
	return
	}
	fmt.Fprintf(w,"POST request successful\n")
	name:= r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n", name)
	fmt.Fprintf(w,"Address = %s\n", address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/hello"{
	http.Error(w, "404 not found",http.StatusNotFound)
}
if r.Method!="GET"{
	http.Error(w, "method not supported",http.StatusNotFound)
	return 
}
fmt.Fprintf(w, "Hello!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static")) //we are basically telling golang to check static directory 
	http.Handle("/",fileServer) //haandle root route
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err:=http.ListenAndServe(":8000",nil);err!=nil{
		log.Fatal(err) //handle error
	}

}


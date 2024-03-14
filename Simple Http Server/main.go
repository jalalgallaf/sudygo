package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/",func(w http.ResponseWriter,r* http.Request){
		n,err:=fmt.Fprint(w,"Hello world!!")
		if err !=nil{
			fmt.Println("Error in code ",err)
		}

		fmt.Println(fmt.Sprintf("Number of bite written is %d",n))
	})

	fmt.Println("Hello")


	http.ListenAndServe(":8080",nil);


}

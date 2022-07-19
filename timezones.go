package main

import(
	"fmt"
	"math/rand"
)

func Now(w http.ResponseWriter,_ *http.Request){
	//print to client "w"
	now:=time.Now().UTC()
	//map key string, value string
	p:=make(map[string]string)
	p["now"]= now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")//set data type
	json.NewEncoder(w).Encode(p)//convert map into json data (key: value)
}
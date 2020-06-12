package main
import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.RemoteAddr, ":")

	//fmt.Printf("%s",  a[0])
	_,err:=fmt.Fprintf(w,"%s",  a[0])
	if err!=nil{
		fmt.Println("Error:",err)
	}
}

func main() {
	http.HandleFunc("/myip", handler)
	err:=http.ListenAndServe(":8080", nil)
	if err!=nil{
		fmt.Println("Error:",err)
	}
}

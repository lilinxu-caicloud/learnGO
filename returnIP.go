package main
import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.RemoteAddr, ":")
	//fmt.Printf("%s",  a[0])
	fmt.Fprintf(w,"%s",  a[0])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

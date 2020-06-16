package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var (
	userNumber = 0
	//userStruct[] User
	userMap = make(map[string]User)
)

//struct of users
type User struct {
	identifier string //unique identifier for every user
	name       string
	age        string
	NO         int
}

//show ip&por, request time, request method, request path
func showMassage(w http.ResponseWriter, r *http.Request) {

	//ip := strings.Split(r.RemoteAddr, ":")//get ip
	ip := r.RemoteAddr      //get ip&port
	time := time.Now()      //get time
	method := r.Method      // get request method
	pathURI := r.RequestURI //get URI
	pathURL := r.URL.Path   //get URL
	fmt.Println("Request ip:port:", ip)
	fmt.Println("Request time", time)
	//fmt.Println("Request time",time)
	fmt.Println("Request method:", method)
	fmt.Println("Request URI:", pathURI)
	fmt.Println("Request URL path:", pathURL)

	//show on the web
	////print ip&port
	//_,errIp:=fmt.Fprintf(w,"Request ip:port:%s\n",ip)
	//if errIp!=nil{
	//	fmt.Println("Ip error:",errIp)
	//}
	////print time
	//_,errTime:=fmt.Fprintf(w,"Request time:%s\n",time)
	//if errTime!=nil{
	//	fmt.Println("Time error:",errTime)
	//}
	////print method
	//_,errMethod:=fmt.Fprintf(w,"Request method:%s\n",method)
	//if errMethod!=nil{
	//	fmt.Println("Method error:",errMethod)
	//}
	////print URI
	//_,errURI:=fmt.Fprintf(w,"Request URI:%s\n",pathURI)
	//if errURI!=nil{
	//	fmt.Println("Method error:",errURI)
	//}
	////print URL path
	//_,errURL:=fmt.Fprintf(w,"Request URL path:%s\n",pathURL)
	//if errURL!=nil{
	//	fmt.Println("Method error:",errURL)
	//}

}

//add user
func addUser(w http.ResponseWriter, r *http.Request) {
	addIdentifier := r.FormValue("addIdentifier")
	addName := r.FormValue("addName")
	addAge := r.FormValue("addAge")
	userNumber++
	var newUser User
	if len(addIdentifier) == 0 || len(addName) == 0 || len(addAge) == 0 {
		io.WriteString(w, "The user cannot be created because there lacks necessary information.")

	} else {
		newUser.identifier = addIdentifier
		newUser.name = addName
		newUser.age = addAge
		newUser.NO = userNumber
		userMap[addIdentifier] = newUser //use identifier to creat map
	}
	showMassage(w, r)
}

//show a specific user, the user is identified by the attribute 'identifier'
func checkMyUser(w http.ResponseWriter, r *http.Request) {
	checkIdentifier := r.FormValue("checkIdentifier")
	checkUser, exist := userMap[checkIdentifier]
	if exist == false {
		io.WriteString(w, "The user does not exist.")
	} else {
		io.WriteString(w, "The user's identifier is ")
		io.WriteString(w, checkUser.identifier)
		io.WriteString(w, "\n")
		io.WriteString(w, "The user's name is ")
		io.WriteString(w, checkUser.name)
		io.WriteString(w, "\n")
		io.WriteString(w, "The user's age is ")
		io.WriteString(w, checkUser.age)
		io.WriteString(w, "\n")
		io.WriteString(w, "The user's ID is ")
		io.WriteString(w, strconv.Itoa(checkUser.NO))
		io.WriteString(w, "\n")
	}
	showMassage(w, r)
}

//delete a specific user, the user is identified by the attribute 'identifier'
func deleteUser(w http.ResponseWriter, r *http.Request) {
	dIdentifier := r.FormValue("dIdentifier")
	_, exist := userMap[dIdentifier]
	if exist == false {
		io.WriteString(w, "The user does not exist.")
	} else {
		if len(dIdentifier) == 0 {
			io.WriteString(w, "No user is deleted.")

		} else {
			delete(userMap, dIdentifier)
			io.WriteString(w, "Delete the user whose identifier is ")
			io.WriteString(w, dIdentifier)
		}

	}
	showMassage(w, r)
}

//change information of a specific user, the user is identified by the attribute 'identifier'
func changeInfo(w http.ResponseWriter, r *http.Request) {
	changeIdentifier := r.FormValue("cIdentifier")
	changeName := r.FormValue("cName")
	changeAge := r.FormValue("cAge")
	var tempUser User                     //use a temp struct object
	_, exist := userMap[changeIdentifier] //only copy
	if exist == false {
		io.WriteString(w, "The user does not exist.")
	} else {
		tempUser.NO = userMap[changeIdentifier].NO
		tempUser.identifier = userMap[changeIdentifier].identifier
		if len(changeName) == 0 && len(changeAge) == 0 {
			io.WriteString(w, "There is no change for the user.")
		} else {
			if len(changeName) != 0 {
				tempUser.name = changeName
			}
			if len(changeAge) != 0 {
				tempUser.age = changeAge
			}
		}
	}
	userMap[changeIdentifier] = tempUser
	showMassage(w, r)
}

//show all the users
func checkAllUser(w http.ResponseWriter, r *http.Request) {
	for key := range userMap {
		thisUser, _ := userMap[key]
		io.WriteString(w, "NO.")
		io.WriteString(w, strconv.Itoa(thisUser.NO))
		io.WriteString(w, "\n")
		io.WriteString(w, "The user's identifier is ")
		io.WriteString(w, thisUser.identifier)
		io.WriteString(w, "\n")
		io.WriteString(w, "The user's name is ")
		io.WriteString(w, thisUser.name)
		io.WriteString(w, "\n")
		io.WriteString(w, "The user's age is ")
		io.WriteString(w, thisUser.age)
		io.WriteString(w, "\n\n")
	}
	showMassage(w, r)

}

func main() {
	//userServe := http.NewServeMux()
	//http.HandleFunc("/myip", showMassage)
	http.HandleFunc("/addUser", addUser)       //创建
	http.HandleFunc("/myUser", checkMyUser)    //查看
	http.HandleFunc("/deleteUser", deleteUser) //删除
	http.HandleFunc("/changeInfo", changeInfo) //修改
	http.HandleFunc("/allUser", checkAllUser)  //查看所有用户
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

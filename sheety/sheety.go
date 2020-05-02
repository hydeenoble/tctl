package sheety

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "time"
	"github.com/joho/godotenv"
	"encoding/json"
	// "text/tabwriter"
	"tctl/model"	
)

func init (){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func CreateTask(task string){
	
	// requestParam := fmt.Sprintf(`{
	// 	"task": {
	// 		"task": "%v",
	// 		"time": "%v",
	// 		"status": "backlog"
	// 	}
	// }`, task, time.Now().Format("2006/01/02 15:04:05"))

	requestParam := &model.SheetyTask{
		Task: &model.Task{
			Task: task,
		},
	}
	requestParam.Default()

	requestParamString, err := json.Marshal(requestParam)

	if err != nil {
        log.Fatal(err)
	}
	
	response, err := http.Post(os.Getenv("API_URL"), "application/json", 
	bytes.NewBuffer([]byte(string(requestParamString))))
	

	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
	}
	
	defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}

func GetTasks(status string){
	response, err := http.Get(os.Getenv("API_URL")+"?status="+status)
	
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
	}
	
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
	resp := &model.SheetyTasks{
		Tasks: &[]model.Task{},
	}

	err = json.Unmarshal([]byte(string(responseData)), resp)
	if err != nil {
        log.Fatal(err)
	}
	
	resp.Output()
}
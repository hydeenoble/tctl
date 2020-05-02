package sheety

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

func CreateTask(tasks string){
	
	requestParam := fmt.Sprintf(`{
		"task": {
			"task": "%v",
			"time": "%v",
			"status": "backlog"
		}
	}`, tasks, time.Now().Format("2006/01/02 15:04:05"))

	response, err := http.Post(os.Getenv("API_URL"), "application/json", 
	bytes.NewBuffer([]byte(requestParam)))
	

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
		Tasks: &[]model.Tasks{},
	}

	err = json.Unmarshal([]byte(string(responseData)), resp)
	if err != nil {
        log.Fatal(err)
	}
	
	resp.Output()
}
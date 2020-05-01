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
)

type Tasks struct{
	task string
	time time.Time
	status string
}

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
			"status": "pending"
		}
	}`, tasks, time.Now())

	response, err := http.Post(os.Getenv("API_URL"),
	"application/json", 
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

func GetTasks(){
	response, err := http.Get(os.Getenv("API_URL"))
	
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
	}
	
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}
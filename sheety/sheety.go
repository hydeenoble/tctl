package sheety

import (
	"bytes"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Tasks struct{
	task string
	age string
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
			"name": "%v",
			"age": "1122",
			"sex": "male"
		}
	}`, tasks)

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
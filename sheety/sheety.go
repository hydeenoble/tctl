package sheety

import (
	"bytes"
	// "encoding/json"
	// "github.com/tidwall/gjson"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateTask(tasks string){
	
	requestParam := fmt.Sprintf(`{
		"task": {
			"name": "%v",
			"age": "1122",
			"sex": "male"
		}
	}`, tasks)

	response, err := http.Post("https://v2-api.sheety.co/8428d9964266c130d65343f01a3d5916/tctl/tasks",
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
	response, err := http.Get("https://v2-api.sheety.co/8428d9964266c130d65343f01a3d5916/tctl/tasks")
	
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
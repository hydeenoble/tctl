package sheety

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateTask(tasks string){
	
	requestParam := fmt.Sprintf(`{
		"tasks": {
			"name": "%d",
			"age": "1122",
			"sex": "male",
		}
	}`, tasks)

	parsedRequestParam, _ := gjson.Parse(requestParam).Value().(map[string]interface{}) 

	var parsedRequestParam = []byte()
	// requestBody, err := json.Marshal(parsedRequestParam)
// 	requestBody, err := json.Marshal(map[string]string{
	// 	"tasks": {
	// 		"name": tasks,
	// 		"age": "1122",
	// 		"sex": "male",
	// 	},
	// })

	// if err != nil {
    //     fmt.Print(err.Error())
    //     os.Exit(1)
    // }

	req, err := http.NewRequest("POST",
	"https://v2-api.sheety.co/8428d9964266c130d65343f01a3d5916/tctl/tasks", 
	bytes.NewBuffer(parsedRequestParam))
	
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

// 	if err != nil {
//         fmt.Print(err.Error())
//         os.Exit(1)
// 	}
	
// 	defer response.Body.Close()

//     responseData, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(string(responseData))
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
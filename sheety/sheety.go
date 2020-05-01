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
	"text/tabwriter"
	"tctl/model"
	"tctl/helper"
	
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
			"status": "pending"
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
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 4, ' ', 0)
	fmt.Fprintln(w, "TASK\tSTATUS\tAGE")
	for i := 0; i < len(*resp.Tasks); i++ {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v", 
		(*resp.Tasks)[i].Task, (*resp.Tasks)[i].Status, 
		helper.TimeToAgeConverter((*resp.Tasks)[i].Time)))
	}
	fmt.Fprint(w)
	w.Flush()
	if len(*resp.Tasks) < 1 {
		fmt.Println("No tasks found")
	}
	
}
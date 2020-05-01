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
	"strings"
	"text/tabwriter"
	
)

type Tasks struct{
	Task string `json:"task"`
	Time TransformTime `json:"time"`
	Status string `json:"status"`
}

type SheetyTasks struct {
	Tasks *[]Tasks `json:"tasks"`
}

type TransformTime struct {
	time.Time
}

func (tt *TransformTime) UnmarshalJSON(input []byte) error {
    strInput := string(input)
    strInput = strings.Trim(strInput, `"`)
    newTime, err := time.Parse("2006/01/02 15:04:05", strInput)
    if err != nil {
        return err
    }

    tt.Time = newTime
    return nil
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
	
	resp := &SheetyTasks{
		Tasks: &[]Tasks{},
	}

	err = json.Unmarshal([]byte(string(responseData)), resp)
	if err != nil {
        log.Fatal(err)
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 4, ' ', 0)
	fmt.Fprintln(w, "TASK\tSTATUS\tAGE")
	for i := 0; i < len(*resp.Tasks); i++ {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v", (*resp.Tasks)[i].Task, (*resp.Tasks)[i].Status, (*resp.Tasks)[i].Time))
	}
	fmt.Fprintln(w)
	w.Flush()
}
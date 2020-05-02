package model

import (
	"time"
    "strings"
    "text/tabwriter"
    "os"
    "fmt"
	"math"
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

func timeToAgeConverter(timestamp TransformTime) string {
	now, _ := time.Parse("2006/01/02 15:04:05", time.Now().Format("2006/01/02 15:04:05"))
	duration := now.Sub(timestamp.Time).String()
	parsedDuration, _ := time.ParseDuration(duration)
	seconds := parsedDuration.Seconds()
	return timeFormater(seconds)
}

func timeFormater(seconds float64) string {
	if((seconds/60) < 1){
		return fmt.Sprintf("%gs", math.Round(seconds))
	}else if((seconds/60) >= 1 && (seconds/60) < 60) {
		return fmt.Sprintf("%gm", math.Round(seconds/60))
	}else if((seconds/3600) >= 1 && (seconds/3600) < 24) {
		return fmt.Sprintf("%gh", math.Round(seconds/3600))
	}else{
		return fmt.Sprintf("%gd", math.Round(seconds/86400))
	}
}

func (st SheetyTasks) Output() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 4, ' ', 0)
	fmt.Fprintln(w, "TASK\tSTATUS\tAGE")
	for i := 0; i < len(*st.Tasks); i++ {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v", 
		(*st.Tasks)[i].Task, (*st.Tasks)[i].Status, 
		timeToAgeConverter((*st.Tasks)[i].Time)))
	}
	fmt.Fprint(w)
	w.Flush()
	if len(*st.Tasks) < 1 {
		fmt.Println("No tasks found")
	}
}
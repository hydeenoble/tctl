package model

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Task struct {
	ID interface{} `json:"id,omitempty"`
	Task   string `json:"task,omitempty"`
	Time   string `json:"time,omitempty"`
	Status string `json:"status,omitempty"`
}

type SheetyTasks struct {
	Tasks *[]Task `json:"tasks"`
}

type SheetyTask struct {
	Task *Task `json:"task"`
}

func (st SheetyTask) Default() {
	st.Task.Status = "backlog"
	st.Task.Time = time.Now().Format("2006/01/02 15:04:05")
}

func timeToAgeConverter(timestamp string) string {
	now, _ := time.Parse("2006/01/02 15:04:05", time.Now().Format("2006/01/02 15:04:05"))
	parsedTimestamp, _ := time.Parse("2006/01/02 15:04:05", timestamp)
	duration := now.Sub(parsedTimestamp).String()
	parsedDuration, _ := time.ParseDuration(duration)
	seconds := parsedDuration.Seconds()
	return timeFormater(seconds)
}

func timeFormater(seconds float64) string {
	if (seconds / 60) < 1 {
		return fmt.Sprintf("%gs", math.Round(seconds))
	} else if (seconds/60) >= 1 && (seconds/60) < 60 {
		return fmt.Sprintf("%gm", math.Round(seconds/60))
	} else if (seconds/3600) >= 1 && (seconds/3600) < 24 {
		return fmt.Sprintf("%gh", math.Round(seconds/3600))
	} else {
		return fmt.Sprintf("%gd", math.Round(seconds/86400))
	}
}

func (st SheetyTask) Output() {
	fmt.Println("Task '" + st.Task.Task + "' created.")
}

func (st SheetyTasks) Output() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tTASK\tSTATUS\tAGE")
	for i := 0; i < len(*st.Tasks); i++ {
		// fmt.Println(strconv((*st.Tasks)[i].ID))
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v\t%v",
			(*st.Tasks)[i].ID,
			(*st.Tasks)[i].Task, 
			strings.Title((*st.Tasks)[i].Status),
			timeToAgeConverter((*st.Tasks)[i].Time)))
	}
	fmt.Fprint(w)
	w.Flush()
	if len(*st.Tasks) < 1 {
		fmt.Println("No tasks found")
	}
}

func GenerateID(task string) string {
	if len(task) < 20 {
		return strings.ReplaceAll(task[:len(task)], " ", "-") + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	return strings.ReplaceAll(task[:20], " ", "-") + "-" + strconv.FormatInt(time.Now().Unix(), 10)

}

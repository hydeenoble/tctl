package model

import (
	"time"
	"strings"
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
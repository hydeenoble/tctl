package helper

import (
	"fmt"
	"tctl/model"
	"time"
)

func TimeToAgeConverter(timestamp model.TransformTime) time.Duration {
	now, _ := time.Parse("2006/01/02 15:04:05", time.Now().Format("2006/01/02 15:04:05"))
	duration := now.Sub(timestamp.Time).String()
	age, _ := time.ParseDuration(duration)
	fmt.Println(age.Seconds())
	return now.Sub(timestamp.Time)
}
package helper

import (
	"fmt"
	"tctl/model"
	"time"
	"math"
)

func TimeToAgeConverter(timestamp model.TransformTime) string {
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
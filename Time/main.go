package main

import (
	"fmt"
	"time"
)

func main() {
	// create current time or   present time we use Now to method 
	// presenttime := time.Now()
	// fmt.Println(presenttime);
	// // to make it in good format we use Formart method  and inorder to find that it must be written in this format  as golang suggest   01-02-2006 Monday 15:04:05
	// fmt.Println(presenttime.Format("01-02-2006 Monday"))

	// // to create date   
	// createdate := time.Date(2002,time.October,07,23,02,23,0,time.UTC)
	// fmt.Println(createdate)
	// fmt.Println("up above")
    // P:= fmt.Println
	// createdate2 :=  time.Date(2021,time.March,02,12,03,03,0,time.UTC)
	// P(createdate2)
	// fmt.Println(createdate.Format("01-02-2006 15:04:05 Monday"))
    // saido :=time.Date(2024,time.February,05,02,34,34,12,time.UTC)
	// fmt.Println(saido.Format("02-02-2006 15:04:05 Monday"))

	// currentDate := time.Now()
	// interval :=1 
	// looptime := 8

	// for  i:= 0 ; i< looptime; i++{
	// 	NextTime := currentDate.AddDate(0,0,interval)
	// 	fmt.Printf("from %s To %s\n",currentDate.Format("2/1/2006"),NextTime.Format("2/1/2006"))
	// 	currentDate = NextTime
	// } 
//  currentDate := time.Now()
//  loopTime   := 8
//  intervalTime := 1
//  var TimeSlot []map[string]interface {}

//  for i := 0; i<loopTime; i++{
// 	// time slice 
// 	nextTime := currentDate.AddDate(0,0,intervalTime)
// 	timeslote := map[string]interface{}{
// 		"From":currentDate.Format("2/1/2006"),
// 		"To":nextTime.Format("2/6/2006"),
// 	}
// 	TimeSlot = append(TimeSlot,timeslote)
// 	currentDate = nextTime
	
//  }
//    for _ , slot := range TimeSlot{
// 	fmt.Printf("From %s To %s \n", slot["From"], slot["To"])
//    }
// 

// for time interval  eg from  6:30  to 8:30

Timelayout := "15:04" // 24-hours format 

startTime, _ := time.Parse(Timelayout,"6:30")
endTime,_ := time.Parse(Timelayout, "8:30")
intervalMinutes := 10 * time.Minute 

Now := time.Now()

startTime =time.Date (Now.Year(), Now.Month(),Now.Day(),startTime.Hour(), startTime.Minute(),0,0,Now.Location())
endTime = time.Date(Now.Year(),Now.Month(),Now.Day(), endTime.Hour(),endTime.Minute(),0,0,Now.Location())
 var Timeslote []map[string]interface {}
for t := startTime; t.Before(endTime); t = t.Add(intervalMinutes){
	timeslot := map[string]interface{}{
		"From":t.Format("03:04 PM"),
	}
	Timeslote = append(Timeslote, timeslot)
}
for _, slot := range Timeslote{
	fmt.Printf("Time %s \n", slot["From"])
}

}
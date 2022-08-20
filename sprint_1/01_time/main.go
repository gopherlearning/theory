package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("\n// Task 1/8")
	timeStr := now.Format("1.2.06 3:4:5 -07 MST")
	fmt.Println(timeStr)
	// -> 3.21.22 10:29:50 +03 MSK

	fmt.Println("\n// Task 2/8")
	fmt.Println(now.Format("Mon 02 Jan 2006 15:04:05 MST"))
	// -> Mon 21 Mar 2022 10:38:55 MSK

	fmt.Println("\n// Task 3/8")
	fmt.Println(now.Format(time.UnixDate))
	// -> Mon Mar 21 11:10:41 MSK 2022
	fmt.Println(now.Format(time.RFC822))
	// -> 21 Mar 22 11:10 MSK

	fmt.Println("\n// Task 4/8")
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	layout := "2006-01-02T15:04:05-07:00"
	currentTime, err := time.Parse(layout, currentTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentTime)
	// -> 2021-09-19 15:59:41 +0300 MSK

	fmt.Println("\n// Task 5/8")
	now = time.Now()
	fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
	fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
	fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))
	// -> Is 2022-03-21 11:19:08.166315157 +0300 MSK m=+0.000122638 before 2021-09-19 15:59:41 +0300 MSK ? Answer: false
	// -> Is 2022-03-21 11:19:08.166315157 +0300 MSK m=+0.000122638 after 2021-09-19 15:59:41 +0300 MSK ? Answer: true
	// -> Is 2022-03-21 11:19:08.166315157 +0300 MSK m=+0.000122638 equal 2021-09-19 15:59:41 +0300 MSK ? Answer: false

	fmt.Println("\n// Task 6/8")
	now = time.Now()
	truncTime := now.Truncate(time.Hour * 23)
	fmt.Println(truncTime)
	// -> 2022-03-21 00:00:00 +0300 MSK

	fmt.Println("\n// Task 7/8")
	birthday := time.Date(1993, time.November, 26, 1, 0, 0, 0, time.Local)
	birthday100 := birthday.AddDate(100, 0, 0)
	now = time.Now().Truncate(time.Hour * 23)
	days := int(birthday100.Sub(now).Hours() / 24)
	fmt.Println(days)
	// -> 26183

	fmt.Println("\n// Task 8/8")
	now = time.Now()
	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(int((<-ticker.C).Sub(now).Seconds()))
	}
	// -> 2
	// -> 4
	// -> 6
	// -> 8
	// -> 10
	// -> 12
	// -> 14
	// -> 16
	// -> 18
	// -> 20

}

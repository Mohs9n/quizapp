package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	timeLimit := flag.Int("time", 30, "the time limit for each question")
	csvFile   := flag.String("csv", "problems.csv", "Pass the csv file with the problems")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}

	cr := csv.NewReader(file)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%v", record)
		question(record, *timeLimit)
	}
}

func question(qs []string, limit int) {
	ch := make(chan int)
	fmt.Printf("%v: ", qs[0])
	go timer(ch, limit)
	var ans int
	fmt.Scan(&ans)
	if cor, _ := strconv.Atoi(qs[1]); ans != cor {
		fmt.Printf("Incorrect answer!\nthe correct answer is: %v\n", cor)
		os.Exit(0)
	} else {
		ch <- 1
	}
}

// there is a timer with a channel in time
// time.Timer
// used by time.NewTimer(duration)
// can be used with select in calling for loop with the timer channel instead of new channel
func timer(ch chan int, limit int) {
	limit = limit * 1000
	for {
		select {
		case <-ch:
			return
		default:
			if limit > 0 {
				time.Sleep(time.Millisecond)
				limit--
			} else {
				os.Exit(0)
			}
		}
	}
}
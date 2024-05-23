package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type workHistory struct {
	duration int
	date     string
	time     string
	fraction int
}

// writeHistory like csv file
func (w workHistory) String() string {
	// duration,date,time,fraction
	return fmt.Sprintf(
		"%d,%s,%s,%d", w.duration, w.date, w.time, w.fraction,
	)
}

const historyHeader = "Duration,Date,Fraction"
const historyFile = "/.thirdTime_history.csv"

// writeHistory write history to file, if file not exist create it
// if file exist append to it, use home directory to store file
func writeHistory(w workHistory) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(home+historyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(w.String() + "\n"); err != nil {
		return err
	}
	return nil
}

// showHistory read history from file and print it
func showHistory() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	f, err := os.Open(home + historyFile)
	if err != nil {
		return err
	}
	defer f.Close()
	// print all lines from file
	r := csv.NewReader(bufio.NewReader(f))
	fmt.Println(historyHeader)
	avg := 0
	rows := 0
	var wg sync.WaitGroup

	whChannel := make(chan workHistory, 100)
	wg.Add(1)

	// go dailyAverage(wh_channel)
	go func() {
		defer wg.Done()
		dailyAverage(whChannel)
	}()

	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		duration, date, time, fraction := record[0], record[1], record[2], record[3]

		fmt.Printf("%s,%s,%s,%s\n", duration, date, time, fraction)
		intDuration, _ := strconv.Atoi(duration)
		intFraction, _ := strconv.Atoi(fraction)

		wh := workHistory{
			duration: intDuration,
			date:     date,
			time:     time,
			fraction: intFraction,
		}

		whChannel <- wh

		avg += intDuration
		rows++

	}
	close(whChannel)

	fmt.Printf("Average time: %d minutes\n", avg/rows)
	wg.Wait()

	return nil
}

// create a goroutine to compute daily average using workHistory channel
func dailyAverage(history <-chan workHistory) map[string]int {
	sum := make(map[string]int)
	dayCount := make(map[string]int)
	for h := range history {
		dayCount[h.date]++
		sum[h.date] += h.duration

	}
	avg := make(map[string]int)
	fmt.Println("Daily Average")
	for k, v := range sum {
		avg[k] = v / dayCount[k]
		fmt.Printf("%s: %d minutes\n", k, avg[k])
	}

	return avg

}

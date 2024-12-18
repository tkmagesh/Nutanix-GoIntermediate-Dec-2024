package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := readData([]string{"data1.dat", "data2.dat"})
	evenCh, oddCh := Splitter(dataCh)
	evenSumCh := Sum(evenCh)
	oddSumCh := Sum(oddCh)
	done := Merger(evenSumCh, oddSumCh)
	<-done
}

func readData(fileNames []string) <-chan int {
	dataCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for _, fileName := range fileNames {
			wg.Add(1)
			go Source(fileName, dataCh, wg)
		}
		wg.Wait()
		close(dataCh)
	}()
	return dataCh
}

func Source(fileName string, dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if no, err := strconv.Atoi(line); err == nil {
			dataCh <- no
		}
	}
}

func Splitter(dataCh <-chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		defer close(evenCh)
		defer close(oddCh)
		for no := range dataCh {
			if no%2 == 0 {
				evenCh <- no
			} else {
				oddCh <- no
			}
		}
	}()
	return evenCh, oddCh
}

func Sum(ch <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		defer close(sumCh)
		var total int
		for val := range ch {
			total += val
		}
		sumCh <- total
	}()
	return sumCh
}

func Merger(evenSumCh, oddSumCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create("result.txt")
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		for range 2 {
			select {
			case evenSum := <-evenSumCh:
				fmt.Fprintf(file, "Even Total : %d\n", evenSum)
			case oddSum := <-oddSumCh:
				fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
			}
		}
		// close(doneCh)
		doneCh <- struct{}{}
	}()
	return doneCh
}

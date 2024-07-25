package main

import (
	"bufio"
	"os"
	"strconv"
	"sync"
	"time"
)

func doHeavyWork(n int) int {
	time.Sleep(time.Second)
	return n
}

func producer(ch chan int) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			ch <- doHeavyWork(num)
		}(i)
	}
	wg.Wait()
	close(ch)
}

func consumer(ch chan int, w *bufio.Writer) {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for v := range ch {
		wg.Add(1)
		go func(w *bufio.Writer, val int) {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			w.WriteString(strconv.Itoa(val) + "\n")
		}(w, v)
	}
	wg.Wait()
}

func main() {
	ch := make(chan int)

	go producer(ch)

	file, _ := os.Create("src/concurrency/channels/output.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	consumer(ch, writer)
}

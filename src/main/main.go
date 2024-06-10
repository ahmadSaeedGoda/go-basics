package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime/pprof"
	"sync"
	"time"
	"unsafe"
)

func main() {
	// display_defer_order()
	// use_chan()
	// use_chan_with_select()
	// profile()
	// profile_workload()
	// struct_to_json()
	// json_to_struct()
	// race_conditions()
	// use_reflect()
	// use_unsafe()
	// use_worker_pool()
}

func display_defer_order() {
	defer fmt.Println("End.")
	defer fmt.Println("world")
	fmt.Println("Starting the program...")
	fmt.Println("hello")
}

func use_chan() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pong"
		close(messages) // Close the channel after sending messages
	}()

	for msg := range messages { // Use range to automatically break when the channel is closed
		fmt.Println(msg)
	}
}

func use_chan_with_select() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pong"
		close(messages) // Close the channel after sending messages
	}()

	for {
		select {
		case msg, open := <-messages:
			if !open {
				fmt.Println("Channel closed!")
				return
			}
			fmt.Println(msg)
		default:
			fmt.Println("Waiting for message...")
			time.Sleep(500 * time.Millisecond) // Sleep to avoid busy-waiting
		}
	}
}

func profile() {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// Code to be profiled
	use_chan_with_select()
}

func profile_workload() {
	// Create a file to store the profile data
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
		return
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Execute the function to profile
	use_chan_with_select()

	fmt.Println("CPU profiling completed")
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func struct_to_json() {
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	fmt.Println(string(jsonData))
}

func json_to_struct() {
	jsonData := []byte(`{"name":"Alice","age":30}`)
	var p Person
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", p)
}

// Concurrency
var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func race_conditions() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func use_reflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x).Elem()

	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())

	if v.Kind() == reflect.Float64 {
		v.SetFloat(7.1)
	}

	fmt.Println("value altered:", v)
	fmt.Println("x = ", v)
}

func use_unsafe() {
	var i int = 42
	var f float64 = *(*float64)(unsafe.Pointer(&i))
	fmt.Println(f)
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work by sleeping for a second
		results <- job * 2
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func use_worker_pool() {
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg) // Pass w as argument to ensure each goroutine gets a unique ID
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}
}

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "  finished job", j)
		results <- j * 2
	}
}

func main() {

	// buffer = 10 for channels
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Run 3 workers

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Start 10 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// No more jobs
	close(jobs)
	fmt.Println("All jobs submitted!")

	// Read 10 results
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("job %d, res=%d\n", a, res)
	}
}

package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

// Write a Go program that does the following:
//
//	Start a worker pool with 5 goroutines.
//	Send integers 1 through 10 as jobs into a jobs channel.
//	Each worker should read jobs from the jobs channel, compute the square of the integer, and send the result into a results channel.
//	Use a context.Context with a 2 second timeout to control cancellation.
//	Collect results from the results channel, sort them in ascending order, and print the sorted slice.
//	Ensure the program exits cleanly (no goroutine leaks) and closes channels where appropriate.
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: context cancelled, exiting\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			res := job * job

			// try to send result but also respect the context
			select {
			case <-ctx.Done():
				fmt.Printf("worker %d: context cancelled before sending result for job %d\n", id, job)
				return
			case results <- res:
				fmt.Printf("worker %d: job=%d result=%d\n", id, job, res)

			}

			// cancellable wait
			select {
			case <-time.After(4 * time.Second):
			case <-ctx.Done():
				fmt.Printf("worker %d: context cancelled during sleep\n", id)
				return
			}
		}
	}
}

func main() {
	const numJobs int = 10
	const numWorker int = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create a context with timeout
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// start worker
	var wg sync.WaitGroup
	for w := 1; w <= numWorker; w++ {
		wg.Add(1)
		go worker(ctxTimeout, w, jobs, results, &wg)
	}

	// send jobs to the worker pool
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	sl := []int{}
	for res := range results {
		sl = append(sl, res)
	}

	sort.Ints(sl)

	// useful final message: if we timed out, we likely have fewer than numJobs results
	if len(sl) < numJobs {
		fmt.Printf("collected %d/%d results (context likely timed out)\n", len(sl), numJobs)
	} else {
		fmt.Printf("collected all %d results\n", len(sl))
	}
	fmt.Println("results:", sl)

}

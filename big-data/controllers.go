package main

import (
	"bufio"
	"context"
	"os"
	"sync"
)

func concurrent_processing(file string, numWorkers, batchSize int) (res result) {
	res = result{donationMonthFreq: map[string]int{}}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	// reader creates and returns a channel that receives batches of rows
	// from the file
	reader := func(ctx context.Context, rowsBatch *[]string) <-chan []string {
		out := make(chan []string)

		scanner := bufio.NewScanner(f)

		go func() {
			defer close(out) // close channel when completed sending all rows

			for {
				scanned := scanner.Scan()

				select {
				case <-ctx.Done():
					return
				// default appends a new row to the rowsBatch until the batch size is reached, which it then sends the data to the out channel
				default:
					row := scanner.Text()
					if len(*rowsBatch) == batchSize || !scanned {
						out <- *rowsBatch
						*rowsBatch = []string{}
					}
					*rowsBatch = append(*rowsBatch, row)
				}

				if !scanned {
					return
				}
			}
		}()

		return out
	}

	// worker picks up a batch from the reader, processes each batch and sends out the processed data
	worker := func(ctx context.Context, rowBatch <-chan []string) <-chan processed {
		out := make(chan processed)

		go func() {
			defer close(out)

			p := processed{}

			for rowBatch := range rowBatch {
				for _, row := range rowBatch {
					firstName, fullName, month := processRow(row)
					p.fullNames = append(p.fullNames, fullName)
					p.firstNames = append(p.firstNames, firstName)
					p.months = append(p.months, month)
					p.numRows++
				}
			}
			// Send to output channel once processing is done
			out <- p
		}()

		return out
	}

	// combiner merges incoming processed data from the workers
	combiner := func(ctx context.Context, inputs ...<-chan processed) <-chan processed {
		out := make(chan processed)

		var wg sync.WaitGroup
		mulitplexer := func(p <-chan processed) {
			defer wg.Done()

			for in := range p {
				select {
				case <-ctx.Done():
				case out <- in:
				}
			}
		}

		// Add the number of wait groups per parallel execution
		wg.Add(len(inputs))
		for _, in := range inputs {
			go mulitplexer(in)
		}

		// close channel after all inputs channels are closed
		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start a reader
	rowsBatch := []string{}
	rowsCh := reader(ctx, &rowsBatch)

	// Start a processer
	workersCh := make([]<-chan processed, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workersCh[i] = worker(ctx, rowsCh)
	}

	firstNameCount := map[string]int{}
	fullNameCount := map[string]bool{}

	// Read from processer
	for processed := range combiner(ctx, workersCh...) {

		// add number of rows processed by worker
		res.numRows += processed.numRows

		// add months processed by worker
		for _, month := range processed.months {
			res.donationMonthFreq[month]++
		}

		// use full names to count people
		for _, fullName := range processed.fullNames {
			fullNameCount[fullName] = true
		}
		res.peopleCount = len(fullNameCount)

		// update most common first name based on processed results
		for _, firstName := range processed.firstNames {
			firstNameCount[firstName]++

			if firstNameCount[firstName] > res.commonNameCount {
				res.commonName = firstName
				res.commonNameCount = firstNameCount[firstName]
			}
		}
	}

	return res
}

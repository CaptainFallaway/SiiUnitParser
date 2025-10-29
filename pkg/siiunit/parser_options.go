package siiunit

// This shit is really useless IMO but it's a nice pattern I wanted to explore more

import "fmt"

const DEFAULT_WORKER_COUNT = 4

func getOptions(opts ...ParserOption) *parserOptions {
	options := &parserOptions{
		workerCount: DEFAULT_WORKER_COUNT,
	}

	var err error
	for _, opt := range opts {
		err = opt(options)
		if err != nil {
			panic(err)
		}
	}

	return options
}

type parserOptions struct {
	workerCount int
}

type ParserOption func(*parserOptions) error

// OptGoroutinesCount is if you want to control the ammount to goroutines the parser uses.
// Minimum 2 goroutines.
func OptWorkerCount(count int) ParserOption {
	return func(po *parserOptions) error {
		if count < 1 {
			return fmt.Errorf("OptWorkerCount: need at least 1 worker")
		}

		po.workerCount = count

		return nil
	}
}

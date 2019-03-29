package main

// WorkerQueue will be channel of work request
var WorkerQueue chan chan ClientRequest

// StartDispatcher Will Start the worker dispatcher
func StartDispatcher(nWorkers int) {
	WorkerQueue = make(chan chan ClientRequest, nWorkers)

	// Now lets create our workers
	for index := 0; index < nWorkers; index++ {
		worker := NewWorker(index+1, WorkerQueue)
		worker.ReadandPut()
	}

	go func() {
		for {
			select {
			case work := <-RequestQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()

			}
		}
	}()

}

package main

import "fmt"

type Worker struct {
	ID          int
	Work        chan ClientRequest
	WorkerQueue chan chan ClientRequest
	QuitChan    chan bool
}

// NewWorker will create a worker object
func NewWorker(id int, workerQueue chan chan ClientRequest) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan ClientRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
	return worker
}

// ReadandPut will read input message and process it
func (w *Worker) ReadandPut() {
	go func() {
		for {
			// Add ourselves into the worker queue
			w.WorkerQueue <- w.Work
			select {
			case work := <-w.Work:
				fmt.Printf("worker %d received work requat", w.ID)
				// http post, save it to the db and reply back with manipulated message
				go FetchandWrite(ResponseQueue)
				ResponseQueue <- ClientResponse{work.Message, work.Conn}
			case <-w.QuitChan:
				// We have been asked to stop
				fmt.Println("Worker Stopped", w.ID)
			}
		}
	}()
}

// FetchandWrite will fetch and write the message to the duplex socket
func FetchandWrite(response chan ClientResponse) {
	writeRes := <-response
	fmt.Println("Output message is: ", writeRes.Dataout)
	writeRes.Conn.Write([]byte(writeRes.Dataout))
}

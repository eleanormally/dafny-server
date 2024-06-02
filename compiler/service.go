package compiler

import (
	"fmt"
	"time"
)

type CompilerService interface {
	AddCodeInstanceToQueue(inst CodeInstance) (int, error)
	GetQueueSize() int
}

type compilerServiceInternal struct {
	codeChan      chan CodeInstance
	queueSize     int
	queueSizeChan chan QueueSizeRequest
}

type QueueSizeRequest struct {
	Modification int
	Response     chan int
}

type CodeInstance struct {
	Files     []DafnyFile
	Requester string
}

type DafnyFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (c *compilerServiceInternal) AddCodeInstanceToQueue(inst CodeInstance) (int, error) {
	go func() {
		c.codeChan <- inst
	}()
	return c.increaseQueueSize(), nil
}

func (c *compilerServiceInternal) GetQueueSize() int {
	return c.queueSize
}

func (c *compilerServiceInternal) increaseQueueSize() int {
	response := make(chan int)
	c.queueSizeChan <- QueueSizeRequest{
		Modification: 1,
		Response:     response,
	}
	return <-response
}

func (c *compilerServiceInternal) decreaseQueueSize() int {
	response := make(chan int)
	c.queueSizeChan <- QueueSizeRequest{
		Modification: -1,
		Response:     response,
	}
	return <-response
}

func (c *compilerServiceInternal) StartQueueIncrease() {
	for request := range c.queueSizeChan {
		c.queueSize += request.Modification
		request.Response <- c.queueSize
	}
}

func (c *compilerServiceInternal) StartCompilationQueue() {
	for inst := range c.codeChan {
		c.decreaseQueueSize()
		err := prepareCompilationEnvironment(inst)
		if err != nil {
			fmt.Printf("Error preparing compilation environment: %s\n", err.Error())
		}
		time.Sleep(time.Second)
		//TODO: compilation
		fmt.Printf("Compiled request by %s, queue now %d\n", inst.Requester, c.GetQueueSize())
	}
}

func StartCompilerService() (CompilerService, error) {
	c := compilerServiceInternal{
		codeChan:      make(chan CodeInstance),
		queueSize:     0,
		queueSizeChan: make(chan QueueSizeRequest),
	}

	// syncronized queue size
	go c.StartQueueIncrease()
	go c.StartCompilationQueue()
	return &c, nil
}

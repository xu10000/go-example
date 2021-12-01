package main

import (
	"fmt"
)

type mailController struct {
	sender   string
	mailChan chan string
	stopChan chan struct{}
}

func (m *mailController) sendMail(context string) {
	m.mailChan <- context
}

func (m *mailController) watchMail() {
	for c := range m.mailChan {
		fmt.Println("------ get mail from ", m.sender, " mail context ", c)
	}
	m.stopChan <- struct{}{}
}

func main() {

	mailController := mailController{
		"xzt",
		make(chan string),
		make(chan struct{}),
	}
	go mailController.watchMail()
	mailController.sendMail("hi, good morning")
	close(mailController.mailChan)
	select {
	case <-mailController.stopChan:
		fmt.Println("------ process exit success")
	}
}

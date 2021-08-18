package main

import (
	"fmt"
	"github.com/longjoy/logAgent/kafka"
	"github.com/longjoy/logAgent/tail_log"
	"time"
)

func main()  {
	err := kafka.Init()
	if err != nil {
		fmt.Println("kafka 链接失败")
		return
	}

	err = tail_log.Init("./my.log")
	if err != nil {
		fmt.Println("tail_log 链接失败")
		return
	}
	run()
}

func run()  {
	for {
		select {
			case line := <-tail_log.ReadChan():
				kafka.SendToMessage("my_log", line.Text)
		default:
			time.Sleep(time.Second)
		}


	}
}
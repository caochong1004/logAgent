package tail_log

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailClient *tail.Tail
	LogChin chan string
)
func Init(addr string) (err error) {
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	}
	tailClient, err = tail.TailFile(addr, config)
	if err != nil {
		fmt.Println("tail file failed", err)
		return err
	}
	return nil
}



func ReadChan() <-chan *tail.Line {
	return tailClient.Lines
}

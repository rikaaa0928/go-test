package go_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestConcurrencyWrite(t *testing.T) {
	c := make(chan string, 2)
	go func() {
		for i := 0; i < 10; i++ {
			str := <-c
			t.Log(str)
		}
	}()
	for i := 0; i < 10; i++ {
		writeConcurrency(c, strconv.FormatInt(int64(i), 10))
	}
	for i := 0; i < 10; i++ {
		writeConcurrency(c, strconv.FormatInt(int64(i), 10))
	}
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		str := <-c
		t.Log(str)
	}
	for i := 0; i < 10; i++ {
		writeConcurrency(c, strconv.FormatInt(int64(i), 10))
	}
	time.Sleep(time.Second * 4)
	for i := 0; i < 2; i++ {
		str := <-c
		t.Log(str)
	}
}

func writeConcurrency(c chan<- string, str string) {
	t := time.NewTimer(time.Millisecond * 10)
	select {
	case c <- str:
	case <-t.C:
		go func() {
			nt := time.NewTimer(time.Second * 3)
			select {
			case c <- str:
			case <-nt.C:
				fmt.Println("drop " + str)
			}
		}()
	}
}

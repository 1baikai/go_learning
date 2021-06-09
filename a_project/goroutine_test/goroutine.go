package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "1")
	_ = tr.Event(context.Background(), "2")
	_ = tr.Event(context.Background(), "3")
	_ = tr.Event(context.Background(), "4")
	_ = tr.Event(context.Background(), "5")
	_ = tr.Event(context.Background(), "6")
	_ = tr.Event(context.Background(), "7")
	_ = tr.Event(context.Background(), "8")
	_ = tr.Event(context.Background(), "9")
	_ = tr.Event(context.Background(), "10")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	tr.Shutdown(ctx)

}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()

	}
}
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case a := <-t.stop:
		fmt.Println("q", a)
	case b := <-ctx.Done():
		fmt.Println("b", b)
	}
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// gorutine_wg()
	// gorutine_chan()
	// sliceSum()
	// ping()
	/*-----*/
	// hsrv.HttpSrv()
	// ContextWithTimeOut()
	// ContextWithValue()
	ContextWithCancel()
}

func ContextWithTimeOut() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Done task")
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout")
	}
}

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Canceled")
			return
		}
	}
}

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func ContextWithValue() {
	type key int
	const EmailKey key = 0
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, EmailKey, "petya@mail.ru")
	if userEmail, ok := ctxWithValue.Value(EmailKey).(string); ok {
		fmt.Println(userEmail)
	} else {
		fmt.Println("No value")
	}
	/*
		userEmail, ok := ctxWithValue.Value(EmailKey).(string)
		if ok {
			fmt.Println(userEmail)
		} else {
			fmt.Println("No value")
		}
	*/

}

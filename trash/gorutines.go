package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func gorutine_chan() {
	code := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCodeCh(code)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for res := range code {
		fmt.Printf("Код ответа: %d\n", res)
	}
}

func getHttpCodeCh(codeCh chan int) {
	resp, err := http.Get("https://ya.ru")
	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}
	codeCh <- resp.StatusCode
}

func gorutine_wg() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go getHttpCode(&wg)
		go func() {
			getHttpCode()
			wg.Done()
		}()
	}
	// time.Sleep(time.Millisecond * 300)
	wg.Wait()
	fmt.Println("Программа работала -", time.Since(t))
}

func getHttpCode() {
	resp, err := http.Get("https://ya.ru")
	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}
	fmt.Printf("Код ответа: %d\n", resp.StatusCode)
}

// func getHttpCode(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	resp, err := http.Get("https://ya.ru")
// 	if err != nil {
// 		fmt.Printf("Ошибка: %s", err.Error())
// 	}
// 	fmt.Printf("Код ответа: %d\n", resp.StatusCode)
// }

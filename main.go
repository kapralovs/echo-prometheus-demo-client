package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	c := new(http.Client)

	noteURL := "http://localhost:8080/note/get/"
	userURL := "http://localhost:8080/user/get/"

	for i := 1; i < 4; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go func(i int, wg *sync.WaitGroup) {
			resp, err := c.Get(fmt.Sprintf("%s%d", noteURL, i))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
			wg.Done()
		}(i, wg)
		go func(i int, wg *sync.WaitGroup) {
			resp, err := c.Get(fmt.Sprintf("%s%d", userURL, i))
			if err != nil {
				log.Fatal(err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			fmt.Println(string(body))

			// c.Get(fmt.Sprintf("%s%d\n", userURL, i))
			fmt.Printf("%s%d\n", noteURL, i)
			fmt.Printf("%s%d\n", userURL, i)
			wg.Done()
		}(i, wg)
		wg.Wait()
	}

	start := time.Now()
	for {
		go func() {
			id := rand.Intn(4)
			if id == 0 {
				return
			}
			resp, err := c.Get(fmt.Sprintf("%s%d", noteURL, id))
			if err != nil {
				log.Fatal(err)
			}
			// defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
			resp.Body.Close()

			resp, err = c.Get(fmt.Sprintf("%s%d", userURL, id))
			if err != nil {
				log.Fatal(err)
			}
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
			resp.Body.Close()

			// c.Get(fmt.Sprintf("%s%d\n", userURL, i))
			fmt.Printf("%s%d\n", noteURL, id)
			fmt.Printf("%s%d\n", userURL, id)
			if time.Since(start) > (time.Second * 10) {
				return
			}
		}()

		if time.Since(start) > (time.Second * 10) {
			return
		}
	}
}

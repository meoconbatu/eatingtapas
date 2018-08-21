package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func serve(dishes chan dish) {
	dishes <- dish{name: "chorizo", morsels: random(5, 10)}
	dishes <- dish{name: "chopitos", morsels: random(5, 10)}
	dishes <- dish{name: "pimientos de padrón", morsels: random(5, 10)}
	dishes <- dish{name: "croquetas", morsels: random(5, 10)}
	dishes <- dish{name: "patatas bravas", morsels: random(5, 10)}
}

type dish struct {
	name    string
	morsels int
}

var users = []string{"Alice", "Bob", "Charlie", "Dave"}

func main() {
	fmt.Println("Bon appétit!")
	dishes := make(chan dish, 5)
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(5)
	go serve(dishes)
	for _, user := range users {
		go func(user string) {
			for {
				select {
				case dish := <-dishes:
					if dish.morsels > 0 {
						fmt.Printf("%s is enjoying some %s\n", user, dish.name)
						time.Sleep(time.Duration(random(30, 180)) * time.Second)
						dish.morsels--
						dishes <- dish
					} else {
						wg.Done()
					}
				case <-done:
					return
				}
			}
		}(user)
	}
	wg.Wait()
	close(done)
	fmt.Println("That was delicious!")
}
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

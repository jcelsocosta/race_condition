package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type buffer struct {
	items []string
	mu    sync.Mutex
}

func (buff *buffer) add(item string) {
	buff.mu.Lock()
	defer buff.mu.Unlock()
	if len(buff.items) < 5 {
		buff.items = append(buff.items, item)
		// fmt.Println("Foi adicionado o item " + item)
	} else {
		fmt.Println("O Buffer não pode armazenar nenhum item mais está com a capacidade máxima")
		os.Exit(0)
	}
}

func (buff *buffer) get() string {
	buff.mu.Lock()
	defer buff.mu.Unlock()
	if len(buff.items) == 0 {
		return ""
	}
	target := buff.items[0]

	buff.items = buff.items[1:]
	return target
}

var wg sync.WaitGroup

func main() {
	buff := buffer{}
	wg.Add(2)

	go producer(&buff)
	go consumer(&buff)
	wg.Wait()
}

func producer(buff *buffer) {
	defer wg.Done()
	for index := 1; ; index++ {
		str := strconv.Itoa(index) + "@email.com"
		buff.add(str)
		time.Sleep(5 * time.Millisecond) // Adiciona um pequeno atraso para simular produção
	}
}

func consumer(buff *buffer) {
	defer wg.Done()
	for {
		data := buff.get()

		if data != "" {
			fmt.Println("Enviado um email com a nova senha de acesso para: " + data)
		}
	}
}

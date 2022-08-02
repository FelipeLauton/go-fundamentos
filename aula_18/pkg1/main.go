package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Esse é nosso programa do nosso exercicio de compilação cruzada, foi compilado num windows, e agora está rodando num sistema:", runtime.GOARCH, runtime.GOOS)

}

package main

//複数パッケージをインストールするには、()で囲う
import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("こんにちは", time.Now())
	fmt.Println(user.Current())
}

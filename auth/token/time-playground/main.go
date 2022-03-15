// 这是一个playground，验证了1个time.Duration相当于1ns
package main

import (
	"fmt"
	"time"
)

func main() {
	var m time.Duration = 1e11
	fmt.Println(m.Seconds())

}

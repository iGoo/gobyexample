// [_环境变量_](http://zh.wikipedia.org/wiki/%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F)
// 是一个在[为 Unix 程序传递配置信息](http://www.12factor.net/config)的普遍方式。
// 让我们来看看如何设置，获取并列举环境变量。

package main

import "os"
import "strings"
import "fmt"

func main() {

	// 使用 `os.Setenv` 来设置一个键值队。使用 `os.Getenv`
	// 获取一个键对应的值。如果键不存在，将会返回一个空字符
	// 串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 使用 `os.Environ` 来列出所有环境变量键值队。这个函数
	// 会返回一个 `KEY=value` 形式的字符串切片。你可以使用
	// `strings.Split` 来得到键和值。这里我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

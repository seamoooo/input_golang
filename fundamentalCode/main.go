package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// レシーバーの値を更新するためポインタレシーバーの使用
// getterやsetterを書かフィールドに直接アクセスすることが推奨される
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// レシーバーの値がコピーされて渡される(値渡し)
func (c Counter) String() string {
	return fmt.Sprintf("合計:%d 更新:%v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())

	// レシーバでポインタで渡していないとエラーになるはずだが
	// (&c).Increment()に変換される
	c.Increment()
	fmt.Println(c.String())

	// 定義した変数から引数を渡すこともできる
	// メソッド値
	f1 := c.String
	fmt.Println(f1())
}

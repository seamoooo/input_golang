package main

import "fmt"

type Foo struct {
	Filed1 string
	Filed2 string
}

// 関数でポインタを受け取り内部で変更するとデータの流れが見えずらくなる
func MakeFoo(f *Foo) error {
	f.Filed1 = "val"
	f.Filed2 = "val2"

	return nil
}

// ポインタをもらわずに構造体ごと返して、呼び出し元に保存する方が良い
func MakeFoo2() (Foo, error) {
	f := Foo{
		Filed1: "val",
		Filed2: "val2",
	}

	return f, nil
}

func main() {
	foo := Foo{
		Filed1: "v",
		Filed2: "v",
	}

	MakeFoo(&foo) //むやみにポインタを使用して変更しない
	l, _ := MakeFoo2()

	fmt.Println(foo)
	fmt.Println(l)
}

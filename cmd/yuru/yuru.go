package main

import (
	"fmt"
	"os"

	"github.com/secondarykey/yuru"
)

func init() {
}

func main() {

	cmds := os.Args
	var name string
	if len(cmds) > 2 {
		name = cmds[1]
	}

	//ファイルの指定
	if name == "" {
		name = "yuru.xml"
	}

	conf,err := yuru.LoadConfig(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(conf.BoardData)

	max := yuru.Max(conf)
	result,err := yuru.Search(conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	//再検索を行うかを判定
	if !result.Max(max) {
		fmt.Println("最大コンボが見つからなかったので、再検索するとかも可能にする")
	}
}


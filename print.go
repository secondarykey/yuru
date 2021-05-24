package yuru

import (
	"fmt"

	"golang.org/x/xerrors"
)

func Print(name string) error {

	conf, err := LoadConfig(name)
	if err != nil {
		return xerrors.Errorf("config.Load() error: %w", err)
	}

	fmt.Println(conf.BoardData)

	max := Max(conf)
	result, err := Search(conf)
	if err != nil {
		return xerrors.Errorf("yuru.Search() error: %w", err)
	}

	fmt.Println(result)

	//再検索を行うかを判定
	if !result.Max(max) {
		fmt.Println("最大コンボが見つからなかったので、再検索するとかも可能にする")
	}
	return nil
}

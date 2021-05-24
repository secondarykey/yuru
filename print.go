package yuru

import (
	"fmt"

	"github.com/secondarykey/yuru/config"
	"github.com/secondarykey/yuru/logic"

	"golang.org/x/xerrors"
)

func Print(name string) error {

	err := config.Set(name)
	if err != nil {
		return xerrors.Errorf("config.Set() error: %w", err)
	}

	conf := config.Get()

	fmt.Println(conf.BoardData)

	max := logic.Max()
	result, err := logic.Search()
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

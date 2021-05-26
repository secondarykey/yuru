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
	T := conf.Turn
	B := conf.Beam
	R := conf.StartR
	C := conf.StartC

	board := config.GetDefaultBoard()
	fmt.Println(board)

	max := logic.Max(board)
	result, err := logic.Search(board, T, B, R, C)
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

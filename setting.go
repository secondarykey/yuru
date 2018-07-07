package yuru

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	DIRECTION = "NESW"
)

var (
	DR [4]int = [4]int{-1, 0, 1, 0}
	DC [4]int = [4]int{0, 1, 0, -1}
	N  int    = len(DR)
)

type Config struct {
	Max   bool      `xml:"max,attr"`
	StartR int      `xml:"startR,attr"`
	StartC int      `xml:"startC,attr"`
	Turn  int       `xml:"turn"`
	Beam  int       `xml:"beam"`
	Board BoardInfo `xml:"board"`

	BoardData Board
}

type BoardInfo struct {
	R int    `xml:"r,attr"`
	C int    `xml:"c,attr"`
	B string `xml:",chardata"`
}

func LoadConfig(file string) (*Config,error) {

	var conf Config

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil,err
	}

	err = xml.Unmarshal(data, &conf)
	if err != nil {
		return nil,err
	}

	//始点指示がおかしい
	if conf.StartR < 0 || conf.StartR > conf.Board.R ||
	   conf.StartC < 0 || conf.StartC > conf.Board.C {
	   return nil,fmt.Errorf("start R,C error")
	}

        //盤面データの生成
	board := make([][]int, conf.Board.R)
	for idx := 0; idx < conf.Board.R; idx++ {
		board[idx] = make([]int, conf.Board.C)
	}

	idx := 0
	r := csv.NewReader(strings.NewReader(conf.Board.B))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			continue
		}

		if len(record) == 0 {
			continue
		}

		if len(record) > conf.Board.C {
			return nil,fmt.Errorf("Error CSV Format C[%d],%v", len(record), record)
		}

		if idx >= conf.Board.R {
			return nil,fmt.Errorf("Error CSV Format R[%d]", idx)
		}

		for c := 0; c < conf.Board.C; c++ {
			board[idx][c], err = strconv.Atoi(strings.Trim(record[c], " "))
			if err != nil {
				return nil,fmt.Errorf("Error CSV Data Format [%s]", record[c])
			}
		}
		idx++
	}
	conf.BoardData = board

	return &conf,nil
}

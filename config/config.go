package config

import (
	_ "embed"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/secondarykey/yuru/dto"

	"golang.org/x/xerrors"
)

//go:embed _default.xml
var defaultFile []byte

var gConf *Config

func init() {
	var conf Config
	gConf = &conf
}

type Config struct {
	Max    bool          `xml:"max,attr"`
	StartR int           `xml:"startR,attr"`
	StartC int           `xml:"startC,attr"`
	Turn   int           `xml:"turn"`
	Beam   int           `xml:"beam"`
	Board  dto.BoardInfo `xml:"board"`

	BoardData dto.Board
}

func Get() *Config {
	return gConf
}

func Set(f string) error {

	var conf Config

	data, err := getDefault(f)
	if err != nil {
		return xerrors.Errorf("loadFile() error: %w", err)
	}

	err = xml.Unmarshal(data, &conf)
	if err != nil {
		return xerrors.Errorf("xml.Unmarshal() error: %w", err)
	}

	//始点指示がおかしい
	if conf.StartR < 0 || conf.StartR > conf.Board.R ||
		conf.StartC < 0 || conf.StartC > conf.Board.C {
		return fmt.Errorf("start R,C error")
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
			return fmt.Errorf("Error CSV Format C[%d],%v", len(record), record)
		}

		if idx >= conf.Board.R {
			return fmt.Errorf("Error CSV Format R[%d]", idx)
		}

		for c := 0; c < conf.Board.C; c++ {
			board[idx][c], err = strconv.Atoi(strings.Trim(record[c], " "))
			if err != nil {
				return fmt.Errorf("Error CSV Data Format [%s]", record[c])
			}
		}
		idx++
	}
	conf.BoardData = board

	gConf = &conf

	return nil
}

func getDefault(name string) ([]byte, error) {

	if name == "" {
		name = getDefaultFilePath()
		if _, err := os.Stat(name); err != nil {
			//作成する
			f, err := os.Create(name)
			if err != nil {
				return nil, xerrors.Errorf("os.Create() error: %w", err)
			}
			defer f.Close()
			_, err = f.Write(defaultFile)
			if err != nil {
				return nil, xerrors.Errorf("file Write() error: %w", err)
			}

			return defaultFile, nil
		}
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, xerrors.Errorf("os.ReadFile() error: %w", err)
	}
	return data, nil
}

const DefaultFileName = ".yuru.xml"

func getDefaultFilePath() string {
	return filepath.Join(getHome(), DefaultFileName)
}

func getHome() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	}
	return os.Getenv(env)
}

package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("fail to open file")
	}

	scaner := bufio.NewScanner(file)

	var lines []string

	for scaner.Scan() {
		lines = append(lines, scaner.Text())
	}

	err = scaner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("fail to read line in  file")
	}

	file.Close()
	return lines, nil

}

func (fm Filemanager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to created file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("faild to convert data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

package utils

import (
	"encoding/json"
	"os"
)

// LoadJSON 加载并解析json文件
func LoadJSON(filename string, v interface{}) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, v)
	return err
}

// SaveJSON 编码并保存json文件
func SaveJSON(filename string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, os.ModeAppend)
	return err
}

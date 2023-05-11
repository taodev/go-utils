package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYAML 加载yaml文件
func LoadYAML(filename string, v interface{}) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bs, v)
	return err
}

// SaveYAML 保存yaml文件
func SaveYAML(filename string, v interface{}) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, os.ModeAppend)
	return err
}

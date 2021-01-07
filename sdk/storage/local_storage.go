package storage

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// LocalStorage 本地存储
type LocalStorage struct{}

var fileName = "storage.yaml"
var content = []byte{}

func init() {
	// 1. 查询配置文件是否存在
	_, err := os.Stat(fileName)
	// 2. 不存在则创建
	if err != nil {
		_, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
	}
	// 3. 存在初始化缓存
	content, err = ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
}

// Get 根据 key 获取 data
func (s *LocalStorage) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("Key cannot be empty")
	}
	m := map[string]interface{}{}
	err := yaml.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	return m[key], nil
}

// Set 根据 key 设置 data
func (s *LocalStorage) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("Key cannot be empty")
	}
	if value == nil {
		return errors.New("Value cannot be empty")
	}
	m := map[string]interface{}{}
	err := yaml.Unmarshal(content, &m)
	if err != nil {
		return err
	}
	m[key] = value
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(fileName, data, 0666); err != nil {
		return err
	}
	content = data
	return nil
}

// Del 根据 key 删除 data
func (s *LocalStorage) Del(key string) error {
	if key == "" {
		return errors.New("Key cannot be empty")
	}
	m := map[string]interface{}{}
	err := yaml.Unmarshal(content, &m)
	if err != nil {
		return err
	}
	delete(m, key)
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(fileName, data, 0666); err != nil {
		return err
	}
	content = data
	return nil
}

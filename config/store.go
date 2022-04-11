package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func NewStore() *Store {
	return &Store{}
}

type Store struct {
}

func (s Store) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s Store) Load(path string, val interface{}) error {
	exist, err := s.PathExists(path)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, val)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}

	return nil
}

func (s Store) Save(path string, val interface{}) error {
	d, err := yaml.Marshal(val)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	_, err = f.Write(d)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil
}

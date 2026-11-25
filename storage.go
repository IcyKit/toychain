package main

import (
	"encoding/json"
	"errors"
	"os"
)

func SaveBlocks(path string, blocks []Block) error {
	if path == "" {
		return errors.New("Empty path error")
	}

	blocksJson, err := json.Marshal(blocks)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, blocksJson, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func LoadBlocks(path string) ([]Block, error) {
	var blocks []Block
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

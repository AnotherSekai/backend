package service

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/AnotherSekai/backend/model"
)

func GetLatestGachas(region string, limit int) (*model.GachaResponse, error) {
	path, err := GetRegionJSONPath(region, "gachas.json")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read gachas.json: %w", err)
	}

	var gachas []model.Gacha
	if err := json.Unmarshal(data, &gachas); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gachas.json: %w", err)
	}

	sort.Slice(gachas, func(i, j int) bool {
		return gachas[i].ID > gachas[j].ID
	})

	if limit > len(gachas) {
		limit = len(gachas)
	}

	return &model.GachaResponse{
		Data: gachas[:limit],
	}, nil
}

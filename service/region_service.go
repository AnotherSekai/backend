package service

import "fmt"

var regionBasePaths = map[string]string{
	"jp": "data/sekai-master-db-diff",
	"en": "data/sekai-master-db-en-diff",
	"cn": "data/sekai-master-db-cn-diff",
	"tc": "data/sekai-master-db-tc-diff",
	"kr": "data/sekai-master-db-kr-diff",
}

func GetRegionJSONPath(region, filename string) (string, error) {
	base, ok := regionBasePaths[region]
	if !ok {
		return "", fmt.Errorf("invalid region: %s", region)
	}
	return base + "/" + filename, nil
}

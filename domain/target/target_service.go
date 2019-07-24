package target

import (
	"encoding/json"
	"errors"

	"github.com/BagusAK95/zaun/common"
)

//TargetService : set target service
type TargetService struct {
	repository TargetRepo
	cache      common.Cache
}

//NewService : instantiate service
func NewService(repo TargetRepo, cache common.Cache) TargetService {
	return TargetService{repo, cache}
}

//GetByName : get list target
func (c *TargetService) GetByName(name string) (target Target, err error) {
	var targets []Target

	cache, errGet := c.cache.Get("zaunTargets")
	if errGet == nil {
		var dat []Target

		err := json.Unmarshal([]byte(cache), &dat)
		if err == nil {
			targets = dat
		} else {
			targets = c.repository.FindAll()
			if targets != nil {
				c.cache.Set("zaunTargets", targets)
			}
		}
	}

	for _, target := range targets {
		if target.Name == name {
			return target, nil
		}
	}

	return target, errors.New("Target not found")
}

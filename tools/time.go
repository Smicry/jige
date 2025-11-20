package tools

import (
	"time"
)

type TimeTool struct{}

func (t *TimeTool) Name() string {
	return "time"
}

func (t *TimeTool) Description() string {
	return "获取当前时间信息"
}

func (t *TimeTool) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	now := time.Now()

	return map[string]interface{}{
		"timestamp":   now.Unix(),
		"datetime":    now.Format("2006-01-02 15:04:05"),
		"date":        now.Format("2006-01-02"),
		"time":        now.Format("15:04:05"),
		"timezone":    now.Location().String(),
		"day_of_week": now.Weekday().String(),
	}, nil
}

package tools

import (
	"fmt"
	"math/rand"
	"time"
)

type WeatherTool struct{}

func (w *WeatherTool) Name() string {
	return "weather"
}

func (w *WeatherTool) Description() string {
	return "查询城市天气信息"
}

func (w *WeatherTool) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	city, ok := input["city"].(string)
	if !ok {
		return nil, fmt.Errorf("缺少城市参数")
	}

	// 模拟天气数据
	rand.Seed(time.Now().UnixNano())
	temperatures := []int{20, 22, 25, 18, 23, 19, 21}
	conditions := []string{"晴朗", "多云", "小雨", "阴天", "晴转多云"}

	return map[string]interface{}{
		"city":        city,
		"temperature": temperatures[rand.Intn(len(temperatures))],
		"condition":   conditions[rand.Intn(len(conditions))],
		"humidity":    rand.Intn(30) + 50, // 50-80%
		"update_time": time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

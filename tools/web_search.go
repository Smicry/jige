package tools

import (
	"fmt"
	"math/rand"
	"time"
)

type WebSearchTool struct{}

func (w *WebSearchTool) Name() string {
	return "web_search"
}

func (w *WebSearchTool) Description() string {
	return "执行网页搜索并返回结果"
}

func (w *WebSearchTool) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	query, ok := input["query"].(string)
	if !ok {
		return nil, fmt.Errorf("缺少搜索查询参数")
	}

	// 模拟搜索结果
	rand.Seed(time.Now().UnixNano())
	results := []map[string]string{
		{
			"title":   fmt.Sprintf("关于 %s 的搜索结果1", query),
			"url":     "https://example.com/result1",
			"snippet": fmt.Sprintf("这是关于 %s 的第一个搜索结果摘要", query),
		},
		{
			"title":   fmt.Sprintf("关于 %s 的搜索结果2", query),
			"url":     "https://example.com/result2",
			"snippet": fmt.Sprintf("这是关于 %s 的第二个搜索结果摘要", query),
		},
	}

	return map[string]interface{}{
		"query":   query,
		"results": results,
		"count":   len(results),
	}, nil
}

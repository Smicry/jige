package agent

import (
	"strings"
)

// Intent 用户意图
type Intent struct {
	Name       string                 `json:"name"`
	Confidence float64                `json:"confidence"`
	Parameters map[string]interface{} `json:"parameters"`
	Reasoning  string                 `json:"reasoning"`
}

// IntentDetector 意图检测器
type IntentDetector struct {
	patterns map[string][]string
}

func NewIntentDetector() *IntentDetector {
	detector := &IntentDetector{
		patterns: make(map[string][]string),
	}

	// 定义意图模式
	detector.patterns["calculator"] = []string{
		"计算", "算一下", "加法", "减法", "乘法", "除法",
		"+", "-", "*", "/", "等于多少", "是多少",
	}

	detector.patterns["weather"] = []string{
		"天气", "气温", "温度", "下雨", "晴天", "多云", "气象",
	}

	detector.patterns["time"] = []string{
		"时间", "几点", "日期", "今天", "现在", "什么时候",
	}

	detector.patterns["web_search"] = []string{
		"搜索", "查找", "查询", "百度", "谷歌", "搜一下", "找找",
	}

	return detector
}

// Detect 检测用户意图
func (id *IntentDetector) Detect(query string) *Intent {
	query = strings.ToLower(query)
	bestIntent := "unknown"
	bestConfidence := 0.0

	for intent, patterns := range id.patterns {
		confidence := id.calculateConfidence(query, patterns)
		if confidence > bestConfidence {
			bestConfidence = confidence
			bestIntent = intent
		}
	}

	parameters := id.extractParameters(bestIntent, query)

	return &Intent{
		Name:       bestIntent,
		Confidence: bestConfidence,
		Parameters: parameters,
	}
}

// Detect 检测用户意图
func (id *IntentDetector) DetectByLLm(query string) *Intent {
	query = strings.ToLower(query)
	bestIntent := "unknown"
	bestConfidence := 0.0

	for intent, patterns := range id.patterns {
		confidence := id.calculateConfidence(query, patterns)
		if confidence > bestConfidence {
			bestConfidence = confidence
			bestIntent = intent
		}
	}

	parameters := id.extractParameters(bestIntent, query)

	return &Intent{
		Name:       bestIntent,
		Confidence: bestConfidence,
		Parameters: parameters,
	}
}

func (id *IntentDetector) calculateConfidence(query string, patterns []string) float64 {
	maxScore := 0.0
	for _, pattern := range patterns {
		if strings.Contains(query, pattern) {
			score := float64(len(pattern)) / float64(len(query))
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func (id *IntentDetector) extractParameters(intentName, query string) map[string]interface{} {
	params := make(map[string]interface{})

	switch intentName {
	case "calculator":
		// 提取数学表达式
		params["expression"] = id.extractExpression(query)
	case "weather":
		// 提取城市名称
		params["city"] = id.extractCity(query)
	case "web_search":
		// 提取搜索查询
		params["query"] = id.extractQuery(query)
	}

	return params
}

func (id *IntentDetector) extractExpression(query string) string {
	// 简单的表达式提取逻辑
	return strings.TrimSpace(query)
}

func (id *IntentDetector) extractCity(query string) string {
	// 简单的城市名称提取
	// 实际应用中可以使用更复杂的地理位置识别
	cities := []string{"北京", "上海", "广州", "深圳", "杭州", "成都", "武汉"}
	for _, city := range cities {
		if strings.Contains(query, city) {
			return city
		}
	}
	return "北京" // 默认城市
}

func (id *IntentDetector) extractQuery(query string) string {
	// 提取搜索关键词
	keywords := []string{"搜索", "查找", "查询", "搜一下"}
	for _, keyword := range keywords {
		if idx := strings.Index(query, keyword); idx != -1 {
			return strings.TrimSpace(query[idx+len(keyword):])
		}
	}
	return query
}

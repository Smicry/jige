package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"jige/prompt"
	"jige/tools"

	"github.com/ollama/ollama/api"
)

// Agent AI代理
type Agent struct {
	intentDetector *IntentDetector
	toolRegistry   *tools.ToolRegistry
	client         *api.Client
}

func NewAgent() (*Agent, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, err
	}
	agent := &Agent{
		intentDetector: NewIntentDetector(),
		toolRegistry:   tools.NewToolRegistry(),
		client:         client,
	}
	agent.registerTools()
	return agent, nil
}

func (a *Agent) registerTools() {
	a.toolRegistry.Register(&tools.CalculatorTool{})
	a.toolRegistry.Register(&tools.WeatherTool{})
	a.toolRegistry.Register(&tools.TimeTool{})
	a.toolRegistry.Register(&tools.WebSearchTool{})
	a.toolRegistry.Register(&tools.UnknownTool{})
}

// Process 处理用户查询
func (a *Agent) Process(query string) (*Response, error) {
	intent, err := a.DetectByLLm(query)
	if err != nil {
		return nil, err
	}
	// fmt.Println(intent.Name)
	// 2. 选择工具
	tool, exists := a.toolRegistry.GetTool(intent.Name)
	if !exists {
		return &Response{
			Query:      query,
			Intent:     intent.Name,
			Confidence: intent.Confidence,
			Result:     "抱歉，我无法处理这个请求",
			Error:      fmt.Sprintf("未找到对应的工具: %s", intent.Name),
		}, nil
	}

	// 3. 执行工具
	result, err := tool.Execute(intent.Parameters)
	if err != nil {
		return &Response{
			Query:      query,
			Intent:     intent.Name,
			Confidence: intent.Confidence,
			Result:     "执行工具时发生错误",
			Error:      err.Error(),
		}, err
	}

	// 4. 格式化响应
	response := a.formatResponse(tool.Name(), result)

	return &Response{
		Query:      query,
		Intent:     intent.Name,
		Confidence: intent.Confidence,
		Result:     response,
		Data:       result,
	}, nil
}

func (a *Agent) formatResponse(toolName string, data map[string]interface{}) string {
	switch toolName {
	case "calculator":
		return fmt.Sprintf("计算结果: %v", data["result"])
	case "weather":
		return fmt.Sprintf("%s天气: %s, 温度: %d°C, 湿度: %d%%",
			data["city"], data["condition"], data["temperature"], data["humidity"])
	case "time":
		return fmt.Sprintf("当前时间: %s", data["datetime"])
	case "web_search":
		return fmt.Sprintf("找到 %d 个搜索结果", data["count"])
	default:
		return fmt.Sprintf("工具执行完成: %v", data)
	}
}

// Response 代理响应
type Response struct {
	Query      string                 `json:"query"`
	Intent     string                 `json:"intent"`
	Confidence float64                `json:"confidence"`
	Result     string                 `json:"result"`
	Data       map[string]interface{} `json:"data,omitempty"`
	Error      string                 `json:"error,omitempty"`
}

// ListTools 列出所有可用工具
func (a *Agent) ListTools() []string {
	return a.toolRegistry.ListTools()
}

func (a *Agent) DetectByLLm(query string) (*Intent, error) {
	p := fmt.Sprintf(prompt.IntentDetector, a.toolRegistry.ToolsDescription(), query)
	req := &api.GenerateRequest{
		Model:  "deepseek-r1:1.5b",
		Prompt: p,
		Stream: new(bool),
	}
	var (
		ctx    = context.Background()
		intent = &Intent{}
	)
	respFunc := func(resp api.GenerateResponse) error {
		rsp := strings.TrimSuffix(strings.TrimPrefix(resp.Response, "```json"), "```")
		fmt.Println(rsp)
		return json.Unmarshal([]byte(rsp), intent)
	}
	err := a.client.Generate(ctx, req, respFunc)
	if err != nil {
		return nil, err
	}
	return intent, nil
}

package agent

import (
	"context"
	"fmt"

	"jige/prompt"
	"jige/tools"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// Agent AI代理
type Agent struct {
	intentDetector *IntentDetector
	toolRegistry   *tools.ToolRegistry
	llm            *ollama.LLM
}

func New() (*Agent, error) {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	if err != nil {
		return nil, err
	}
	agent := &Agent{
		intentDetector: NewIntentDetector(),
		toolRegistry:   tools.NewToolRegistry(),
		llm:            llm,
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
	ctx := context.Background()
	p := fmt.Sprintf(prompt.IntentDetector, a.toolRegistry.ToolsDescription(), query)
	result, err := llms.GenerateFromSinglePrompt(ctx, a.llm, p)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return &Response{}, nil
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

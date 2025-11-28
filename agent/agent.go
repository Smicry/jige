package agent

import (
	"context"
	"fmt"
	
	"jige/localtools"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/tools"
)

type Agent struct {
	agent *agents.OneShotZeroAgent
}

func New() (*Agent, error) {
	llm, err := ollama.New(
		ollama.WithModel(llama3Dot1_8b),
	)
	if err != nil {
		return nil, err
	}
	return &Agent{
		agent: agents.NewOneShotAgent(
			llm,
			[]tools.Tool{
				&localtools.Now{},
			},
		),
	}, nil
}

// Process 处理用户查询
func (a *Agent) Process(query string) (string, error) {
	ctx := context.Background()
	executor := agents.NewExecutor(
		a.agent,
		agents.WithMaxIterations(10),
	)
	result, err := executor.Call(ctx,
		map[string]any{"input": query},
	)
	if err != nil {
		return "", err
	}
	output, ok := result["output"].(string)
	if !ok {
		return "", fmt.Errorf("output missing")
	}
	return output, nil
}

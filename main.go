package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"jige/agent"
)

func main() {
	fmt.Println("🤖 AI Agent 启动中...")

	// 创建AI代理
	aiAgent, err := agent.NewAgent()
	if err != nil {
		fmt.Println("启动代理出错:", err)
		return
	}

	fmt.Println("可用的工具:")
	for _, tool := range aiAgent.ListTools() {
		fmt.Printf("  - %s\n", tool)
	}
	fmt.Println("\n请输入您的查询 (输入 'quit' 退出):")

	// 交互式循环
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n👤 您: ")
		if !scanner.Scan() {
			break
		}

		query := strings.TrimSpace(scanner.Text())
		if query == "quit" || query == "退出" {
			fmt.Println("再见！")
			break
		}

		if query == "" {
			continue
		}

		// 处理查询
		response, err := aiAgent.Process(query)
		if err != nil {
			log.Printf("处理查询时发生错误: %v", err)
			continue
		}

		// 显示响应
		fmt.Printf("🤖 Agent: %s\n", response.Result)
		fmt.Printf("   (意图: %s, 置信度: %.2f)\n", response.Intent, response.Confidence)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("读取输入时发生错误: %v", err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jige/agent"
)

func main() {
	fmt.Println("【几个】：启动中...")
	agt, err := agent.New()
	if err != nil {
		fmt.Printf("【几个】：启动出错，%s\n", err)
		return
	}
	fmt.Println("【几个】：请输入您的查询 (输入 'quit' 退出)。")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("【用户】: ")
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
		response, err := agt.Process(query)
		if err != nil {
			fmt.Printf("【几个】乱了: %s\n", err.Error())
			continue
		}
		fmt.Printf("【几个】： %s\n", response)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("【几个】：我不知道你在说什么。 %s\n", err.Error())
	}
}

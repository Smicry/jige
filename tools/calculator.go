package tools

import (
    "fmt"
    "strconv"
    "strings"
)

type CalculatorTool struct{}

func (c *CalculatorTool) Name() string {
    return "calculator"
}

func (c *CalculatorTool) Description() string {
    return "执行数学计算，支持加减乘除等基本运算"
}

func (c *CalculatorTool) Execute(input map[string]interface{}) (map[string]interface{}, error) {
    expression, ok := input["expression"].(string)
    if !ok {
        return nil, fmt.Errorf("缺少表达式参数")
    }

    result, err := c.evaluateExpression(expression)
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "result":    result,
        "expression": expression,
    }, nil
}

func (c *CalculatorTool) evaluateExpression(expr string) (float64, error) {
    // 简单的表达式求值，实际应用中可以使用更复杂的解析器
    expr = strings.ReplaceAll(expr, " ", "")
    
    // 处理加减乘除
    if strings.Contains(expr, "+") {
        parts := strings.Split(expr, "+")
        if len(parts) != 2 {
            return 0, fmt.Errorf("无效的表达式")
        }
        a, err := strconv.ParseFloat(parts[0], 64)
        if err != nil {
            return 0, err
        }
        b, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            return 0, err
        }
        return a + b, nil
    }
    
    if strings.Contains(expr, "-") {
        parts := strings.Split(expr, "-")
        if len(parts) != 2 {
            return 0, fmt.Errorf("无效的表达式")
        }
        a, err := strconv.ParseFloat(parts[0], 64)
        if err != nil {
            return 0, err
        }
        b, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            return 0, err
        }
        return a - b, nil
    }
    
    if strings.Contains(expr, "*") {
        parts := strings.Split(expr, "*")
        if len(parts) != 2 {
            return 0, fmt.Errorf("无效的表达式")
        }
        a, err := strconv.ParseFloat(parts[0], 64)
        if err != nil {
            return 0, err
        }
        b, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            return 0, err
        }
        return a * b, nil
    }
    
    if strings.Contains(expr, "/") {
        parts := strings.Split(expr, "/")
        if len(parts) != 2 {
            return 0, fmt.Errorf("无效的表达式")
        }
        a, err := strconv.ParseFloat(parts[0], 64)
        if err != nil {
            return 0, err
        }
        b, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            return 0, err
        }
        if b == 0 {
            return 0, fmt.Errorf("除数不能为零")
        }
        return a / b, nil
    }
    
    // 如果是单个数字
    result, err := strconv.ParseFloat(expr, 64)
    if err != nil {
        return 0, fmt.Errorf("无法解析表达式: %s", expr)
    }
    
    return result, nil
}
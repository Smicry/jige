package tools

type UnknownTool struct{}

func (u *UnknownTool) Name() string {
	return "unknown"
}

func (u *UnknownTool) Description() string {
	return "无法判断是否使用工具"
}

func (u *UnknownTool) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

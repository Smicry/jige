package tools

// Tool 工具接口
type Tool interface {
	Name() string
	Description() string
	Execute(input map[string]interface{}) (map[string]interface{}, error)
}

// ToolRegistry 工具注册器
type ToolRegistry struct {
	tools map[string]Tool
}

func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{
		tools: make(map[string]Tool),
	}
}

func (tr *ToolRegistry) Register(tool Tool) {
	tr.tools[tool.Name()] = tool
}

func (tr *ToolRegistry) GetTool(name string) (Tool, bool) {
	tool, exists := tr.tools[name]
	return tool, exists
}

func (tr *ToolRegistry) ListTools() []string {
	var names []string
	for name := range tr.tools {
		names = append(names, name)
	}
	return names
}

func (tr *ToolRegistry) ToolsDescription() string {
	var ds string
	for _, tool := range tr.tools {
		ds += ". " + "工具名称：" + tool.Name() + ", 用途：" + tool.Description() + "\n"
	}
	return ds
}

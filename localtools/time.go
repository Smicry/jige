package localtools

import (
	"context"
	"time"
)

type Now struct {
}

func (w *Now) Name() string {
	return "时间"
}

func (w *Now) Description() string {
	return "用来查询本地时间，无需输入参数，调用会返回RFC3339格式时间。"
}

func (w *Now) Call(_ context.Context, _ string) (string, error) {
	return time.Now().Format(time.RFC3339), nil
}

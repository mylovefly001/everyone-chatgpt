package global

import (
	"context"
	"everyone-chatgpt/core/entity"
)

var (
	RunEnv   string
	RunPort  int
	RootPath string
	Context  context.Context
	Config   entity.Config
)

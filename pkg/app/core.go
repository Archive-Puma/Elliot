package app

import (
	"context"
)

// Core TODO
type Core struct {
	ParentCtx context.Context
}

// NewCore TODO
func NewCore() *Core {
	return &Core{
		ParentCtx: nil,
	}
}

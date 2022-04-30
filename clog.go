package clog

import (
	"context"
	"fmt"

	"github.com/golang/glog"
)

type cLogCtxKey struct {
	name string
}

var (
	ctxKey = cLogCtxKey{"clog"}
	defLog = NewCLog("")
)

type CLog struct {
	rawctx string
	ctx    string
}

func NewCLog(ctx string) *CLog {
	return &CLog{
		rawctx: ctx,
		ctx:    ctx + ": ",
	}
}

func WithCtx(ctx context.Context, logctx string) (context.Context, *CLog) {
	var cl *CLog
	v := ctx.Value(ctxKey)
	if l, ok := v.(*CLog); ok {
		cl = NewCLog(l.rawctx + "/" + logctx)
	} else {
		cl = NewCLog(logctx)
	}
	return context.WithValue(ctx, ctxKey, cl), cl
}
func FromCtx(ctx context.Context) *CLog {
	v := ctx.Value(ctxKey)
	if l, ok := v.(*CLog); ok {
		return l
	} else {
		return defLog
	}
}

func (c *CLog) Infof(format string, args ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(c.ctx+format, args...))
}

func (c *CLog) Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(c.ctx+format, args...))
}

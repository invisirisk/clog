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

/*
NewClog can be used when there is no context available.
*/
func NewCLog(logctx string) *CLog {
	return &CLog{
		rawctx: logctx,
		ctx:    logctx + ": ",
	}
}

/*
WithCtx creates a new clog. If ctx has existing clog, the new clog will inherit from existing logger.
*/
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

/*
FromCtx returns logger from provided context. If context does not have existing logger, returns a default logger with no prefix.
*/
func FromCtx(ctx context.Context) *CLog {
	v := ctx.Value(ctxKey)
	if l, ok := v.(*CLog); ok {
		return l
	} else {
		return defLog
	}
}

/*
WithClog creates a new logger from existing clog.
*/
func WithClog(ctx context.Context, cl *CLog, logctx string) (context.Context, *CLog) {
	cl = NewCLog(cl.rawctx + "/" + logctx)
	return context.WithValue(ctx, ctxKey, cl), cl
}

/*
Infof is similar to [glog.Infof]

The log output will be prepended with logctx
*/
func (c *CLog) Infof(format string, args ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(c.ctx+format, args...))
}

/*
Errorf is similar to [glog.Errorf]

The log output will be prepended with logctx
*/
func (c *CLog) Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(c.ctx+format, args...))
}

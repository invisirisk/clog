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
	defLog = newCLog("")
)

type CLog struct {
	rawctx string
	ctx    string
}

func newCLog(ctx string) *CLog {
	return &CLog{
		rawctx: ctx,
		ctx:    ctx + ": ",
	}
}

/*
WithCtx creates a new clog. If ctx has existing clog, the new clog will inherit from existing logger.
*/
func WithCtx(ctx context.Context, logctx string) (context.Context, *CLog) {
	var cl *CLog
	v := ctx.Value(ctxKey)
	if l, ok := v.(*CLog); ok {
		cl = newCLog(l.rawctx + "/" + logctx)
	} else {
		cl = newCLog(logctx)
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
	cl = newCLog(cl.rawctx + "/" + logctx)
	return context.WithValue(ctx, ctxKey, cl), cl
}

/*
Infof is similar to [github.com/golang/glog.Infof]
*/
func (c *CLog) Infof(format string, args ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf(c.ctx+format, args...))
}

/*
Errorf is similar to [github.com/golang/glog.Errorf]
*/
func (c *CLog) Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(c.ctx+format, args...))
}

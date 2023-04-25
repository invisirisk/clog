/*
pckage clog wraps [glog], with string context. Each log line will be prepended with the string context.

# Usage

Basic Example:

	   	ctx, cl := clog.WithCtx(ctx, "base")
		cl.Infof("this is base log")
		// With log: I0425 15:27:22.513425   38138 clog_test.go:25] base: this is base log

The loggger can be accessed from standard golang [context], allowing eas of use.

	cl := clog.FromCtx(ctx)

The clog package implements for hierarchical logging by creating new logger from existing log context.

	_, cl := WithCtx(ctx, "subrequest")
	cl.Infof("this is another base log")
	// I0425 15:27:22.513502   38138 clog_test.go:10] base/subrequest: this is another base log

# Configuration

All standard [glog] flags can be used with clog
*/
package clog

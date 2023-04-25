package clog

import (
	"context"
	"flag"
)

func processSubRequest(ctx context.Context) {
	_, cl := WithCtx(ctx, "subrequest")
	cl.Infof("this is another base log")
	// I0425 15:27:22.513502   38138 clog_test.go:10] base/subrequest: this is another base log
}

func processRequest(ctx context.Context) {
	cl := FromCtx(ctx)
	cl.Infof("this is another base log")
	// I0425 15:27:22.513500   38138 clog_test.go:15] base: this is another base log
	processSubRequest(ctx)

}

func Example() {
	flag.CommandLine.Parse([]string{"-logtostderr"})
	ctx := context.Background()

	ctx, cl := WithCtx(ctx, "base")
	cl.Infof("this is base log")
	// I0425 15:27:22.513425   38138 clog_test.go:25] base: this is base log
	processRequest(ctx)
}

/*
pckage clog wraps [github.com/golang/glog], with string context. Each log line will be prepended with the string context.

The loggger can be accessed from standard golang [context], allowing eas of use.

The clog package implements for hierarchical logging by creating new logger from existing log context.
*/
package clog

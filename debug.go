// +build debug

package segtree

import "log"

func debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}

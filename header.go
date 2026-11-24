//go:build amd64
// Package spruce provides a highly-optimized buffer pool
package spruce

import "time"

const initialBackoffValue = time.Nanosecond
const maxBackoffValue = time.Millisecond

//go:fix inline
func pool_acquire(*uint64) uint32
//go:fix inline
func pool_release(*uint64, uint32)

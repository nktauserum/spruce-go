//go:build amd64

package models

import "time"

const initialBackoffValue = time.Nanosecond
const maxBackoffValue = time.Millisecond

//go:fix inline
func pool_acquire(*uint64) uint32
//go:fix inline
func pool_release(*uint64, uint32)

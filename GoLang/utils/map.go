package utils

import "sync"

var (
	Store       = sync.Map{}
	ExpireStore = sync.Map{}
)

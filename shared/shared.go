// Package shared is for global SharedState
package shared

import (
	"net/http"
	"sync"
)

// SharedState - global client and mutex
type SharedState struct {
	client *http.Client
	mtx    sync.Mutex
}

// NewSharedState - onstructor
func NewSharedState() *SharedState {
	return &SharedState{client: &http.Client{}, mtx: sync.Mutex{}}
}

func SharedStateSender(s *SharedState) {

}

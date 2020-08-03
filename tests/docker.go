package tests

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

// LaunchIPFSDocker runs a fresh go-ipfs docker image and returns the resource for
// container metadata.
func LaunchIPFSDocker(t require.TestingT) (*dockertest.Resource, func()) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	ipfsDocker, err := pool.Run("ipfs/go-ipfs", "v0.6.0", []string{"IPFS_PROFILE=test"})
	require.NoError(t, err)

	err = ipfsDocker.Expire(180)
	require.NoError(t, err)

	time.Sleep(time.Second * 3)
	return ipfsDocker, func() {
		err = pool.Purge(ipfsDocker)
		require.NoError(t, err)
	}
}

// FlakyT provides retry mechanisms to test.
type FlakyT struct {
	t      *testing.T
	failed bool
	cls    []func()
}

// NewFlakyT creates a new FlakyT.
func NewFlakyT(t *testing.T) *FlakyT {
	return &FlakyT{
		t: t,
	}
}

var _ require.TestingT = (*FlakyT)(nil)

// Errorf registers an error message.
func (ft *FlakyT) Errorf(format string, args ...interface{}) {
	ft.t.Errorf(format, args...)
}

// FailNow indicates to fail the test.
func (ft *FlakyT) FailNow() {
	ft.failed = true
	runtime.Goexit()
}

// Cleanup registers a cleanup function.
func (ft *FlakyT) Cleanup(cls func()) {
	ft.cls = append([]func(){cls}, ft.cls...)
}

var numRetries = 5

// RunFlaky runs a flaky test with retries.
func RunFlaky(t *testing.T, f func(ft *FlakyT)) {
	for i := 0; i < numRetries; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		ft := NewFlakyT(t)
		go func() {
			defer wg.Done()
			f(ft)
		}()
		wg.Wait()
		for _, f := range ft.cls {
			f()
		}
		if !ft.failed {
			return
		}
		ft.t.Logf("test %s attempt %d/%d failed, retrying...", t.Name(), i+1, numRetries)
	}
	t.Fatalf("test failed after %d retries", numRetries)
}

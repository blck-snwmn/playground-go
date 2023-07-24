package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/singleflight"
)

func Test_parallel(t *testing.T) {
	const count = 10
	var g singleflight.Group

	var (
		sg sync.WaitGroup
		// waitsg is used to wait for all goroutines except the first to perform Do.
		waitsg sync.WaitGroup
	)
	waitsg.Add(count - 1)

	var (
		// waitFirst is used to wait for the first goroutine executing Do.
		waitFirst = make(chan struct{})
	)

	go func(i int) {
		v, err, share := g.Do("key1", func() (interface{}, error) {
			waitFirst <- struct{}{}
			waitsg.Wait()

			return i, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, 0, v)
		assert.True(t, share)
	}(0)
	<-waitFirst

	for i := 1; i < count; i++ {
		i := i
		sg.Add(1)
		go func(i int) {
			waitsg.Done()

			defer sg.Done()
			v, err, share := g.Do("key1", func() (interface{}, error) {
				return i, nil
			})
			assert.NoError(t, err)
			assert.Equal(t, 0, v)
			assert.True(t, share)
		}(i)
	}
	sg.Wait()
}
func Test_sequence(t *testing.T) {
	var g singleflight.Group
	{
		v, err, share := g.Do("key1", func() (interface{}, error) {
			return "value1-1", nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "value1-1", v)
		assert.False(t, share)
	}
	{
		v, err, share := g.Do("key2", func() (interface{}, error) {
			return "value1-2", nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "value1-2", v)
		assert.False(t, share)
	}
}

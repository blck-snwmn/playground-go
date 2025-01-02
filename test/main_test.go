package test_test

import "testing"

func TestCleanup(t *testing.T) {
	t.Log("test")
	defer func() {
		t.Log("defer")
	}()
	t.Cleanup(func() {
		t.Log("cleanup")
	})
	t.Fatal("fail in test")
}

func Test_FatalInCleanUp(t *testing.T) {
	t.Log("test")
	defer func() {
		t.Log("defer")
	}()
	t.Cleanup(func() {
		t.Fatal("fail in cleanup")
	})
}

func Test_PanicInCleanUp(t *testing.T) {
	t.Log("test")
	defer func() {
		t.Log("defer")
	}()
	t.Cleanup(func() {
		panic("panic in cleanup")
	})
}

func Test_PanicInTest(t *testing.T) {
	t.Log("test")
	defer func() {
		t.Log("defer")
	}()
	t.Cleanup(func() {
		t.Log("cleanup")
	})
	panic("panic in test")
}

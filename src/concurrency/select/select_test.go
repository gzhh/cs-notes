package main

import (
	"sync"
	"testing"
	"time"
)

// 多渠道的选择
func TestDemo1(t *testing.T) {
	/*
		select {
		case ret := <-retCh1:
			t.Logf("result %s", ret)
		case ret := <-retCh2:
			t.Logf("result %s", ret)
		default:
			t.Error("No one returned")
		}
	*/
}

// 超时控制
func TestDemo2(t *testing.T) {
	/*
		select {
		case ret := <-retCh:
			t.Logf("result %s", ret)
		case <-time.After(time.Second * 1):
			t.Error("time out")
		}
	*/
}

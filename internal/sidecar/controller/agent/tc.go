// +build !linux

package agent

func (ag *agent) setTcClassRate(netLimit uint64) error { return nil }

func (ag *agent) delTcClass() error { return nil }

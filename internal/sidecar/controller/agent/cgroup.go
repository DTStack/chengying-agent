// +build !linux

package agent

type cgroup struct{}

func (ag *agent) installCgroup() error {
	return nil
}

func (ag *agent) unInstallCgroup() error {
	return nil
}

func (ag *agent) updateCgroup(cpuLimit float32, memLimit uint64) error {
	return nil
}

// +build !linux,!windows

package monitor

func getTraffic(classid uint32) (uint64, uint64, error) { return 0, 0, nil }

func setTrafficEnable(pid uint32) error { return nil }

func tcStatistic() {}

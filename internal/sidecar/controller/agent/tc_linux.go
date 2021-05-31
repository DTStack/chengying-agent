package agent

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"easyagent/internal/sidecar/base"
)

var cid = classid{}

type idst struct {
	id   uint32
	free bool
}

type classid struct {
	sync.Mutex

	ids []*idst
}

func allocClassid() uint32 {
	cid.Lock()
	defer cid.Unlock()

	if cid.ids == nil {
		major, _ := strconv.ParseUint(base.Major, 16, 16)
		cid.ids = append(cid.ids, &idst{id: uint32(major<<16) + 1})
		return uint32(major<<16) + 1
	}

	for _, v := range cid.ids {
		if v.free == true {
			v.free = false
			return v.id
		}
	}

	id := cid.ids[len(cid.ids)-1].id + 1
	cid.ids = append(cid.ids, &idst{id: id})
	return id
}

func freeClassid(id uint32) {
	cid.Lock()
	defer cid.Unlock()

	for _, v := range cid.ids {
		if v.id == id {
			v.free = true
		}
	}
}

func (ag *agent) setTcClassRate(netLimit uint64) error {
	if ag.classid == 0 {
		return nil
	}

	if netLimit == 0 {
		return ag.delTcClass()
	}

	hid := base.ClassidToHandle(ag.classid)
	parent := base.Major + ":"
	rate := strconv.FormatUint(netLimit, 10)
	errList := make([]string, 0)
	for _, dev := range base.GetDevs() {
		out, err := exec.Command("tc", "class", "change", "dev", dev, "parent", parent, "classid", hid, "htb", "rate", rate).CombinedOutput()
		if err == nil {
			continue
		}

		out, err = exec.Command("tc", "class", "add", "dev", dev, "parent", parent, "classid", hid, "htb", "rate", rate).CombinedOutput()
		if err != nil {
			if len(out) > 0 {
				errList = append(errList, string(out))
			} else {
				errList = append(errList, err.Error())
			}
		}
	}

	var err error
	if len(errList) > 0 {
		err = errors.New(strings.Join(errList, "\n"))
	}
	return err
}

func (ag *agent) delTcClass() error {
	if ag.classid == 0 {
		return nil
	}

	hid := base.ClassidToHandle(ag.classid)
	parent := base.Major + ":"
	errList := make([]string, 0)
	for _, dev := range base.GetDevs() {
		out, err := exec.Command("tc", "class", "del", "dev", dev, "parent", parent, "classid", hid).CombinedOutput()
		if err != nil {
			if len(out) > 0 {
				if sout := string(out); !strings.Contains(sout, "No such file or directory") {
					errList = append(errList, sout)
				}
			} else {
				errList = append(errList, err.Error())
			}
		}
	}

	var err error
	if len(errList) > 0 {
		err = errors.New(strings.Join(errList, "\n"))
	}
	return err
}

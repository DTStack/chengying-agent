package rpc

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"easyagent/internal/proto"
	"easyagent/internal/server/log"
	"github.com/satori/go.uuid"
)

func init() {
	log.ConfigureLogger("/tmp", 0, 0, 0)
}

func TestRpcService_SendControl(t *testing.T) {
	rs := NewRpcService("", 0)
	sid := uuid.NewV4()
	sc := sidecar{
		sid:    sid,
		ctlCh:  make(chan *proto.ControlResponse),
		stopCh: make(chan struct{}),
	}
	rs.sidecarMap[sid] = sc

	if err := SidecarClient.SendControl(context.Background(), uuid.NewV4(), nil); err != ErrNotFound {
		t.Fatalf("expect error %v, but its %v", ErrNotFound, err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if err := SidecarClient.SendControl(ctx, sid, nil); err != context.DeadlineExceeded {
		t.Fatalf("expect error %v, but its %v", context.DeadlineExceeded, err)
	}

	go func() {
		<-sc.ctlCh
	}()
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err := SidecarClient.SendControl(ctx, sid, nil); err != nil {
		t.Fatal(err)
	}
}

func TestRpcService_SendControlSync(t *testing.T) {
	rs := NewRpcService("", 0)
	sid := uuid.NewV4()
	sc := sidecar{
		sid:    sid,
		ctlCh:  make(chan *proto.ControlResponse),
		stopCh: make(chan struct{}),
	}
	rs.sidecarMap[sid] = sc

	rand.Seed(time.Now().UnixNano())
	ctlResp := &proto.ControlResponse{Seqno: rand.Uint32()}

	go func() {
		<-sc.ctlCh
	}()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if _, err := SidecarClient.SendControlSync(ctx, sid, ctlResp); err != context.DeadlineExceeded {
		t.Fatalf("expect error %v, but its %v", context.DeadlineExceeded, err)
	}

	rand.Seed(time.Now().UnixNano())
	ctlResp.Seqno = rand.Uint32()
	op := &proto.Event_OperationProgress{
		Seqno:   ctlResp.Seqno,
		Failed:  false,
		Message: "test success",
	}
	go func() {
		<-sc.ctlCh

		stopSeqno(ctlResp.Seqno, op)
	}()
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	retOp, err := SidecarClient.SendControlSync(ctx, sid, ctlResp)
	if err != nil {
		t.Fatal(err)
	}
	if retOp.(*proto.Event_OperationProgress).Seqno != op.Seqno {
		t.Fatalf("Seqno not equal %v != %v", op.Seqno, ctlResp.Seqno)
	}
	if retOp.(*proto.Event_OperationProgress).Failed != op.Failed {
		t.Fatalf("Failed not equal %v != %v", retOp.(*proto.Event_OperationProgress).Failed, op.Failed)
	}
	if retOp.(*proto.Event_OperationProgress).Message != op.Message {
		t.Fatalf("Message not equal %v != %v", retOp.(*proto.Event_OperationProgress).Message, op.Message)
	}
}

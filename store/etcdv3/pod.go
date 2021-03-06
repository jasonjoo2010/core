package etcdv3

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/projecteru2/core/scheduler"
	"github.com/projecteru2/core/types"
)

// AddPod add a pod
// save it to etcd
// storage path in etcd is `/pod/info/:podname`
func (m *Mercury) AddPod(ctx context.Context, name, favor, desc string) (*types.Pod, error) {
	key := fmt.Sprintf(podInfoKey, name)
	favor = strings.ToUpper(favor)
	if favor == "" {
		favor = scheduler.MemoryPrior
	} else if favor != scheduler.MemoryPrior && favor != scheduler.CPUPrior {
		return nil, types.NewDetailedErr(types.ErrBadFaver,
			fmt.Sprintf("got bad faver: %s", favor))
	}
	pod := &types.Pod{Name: name, Desc: desc, Favor: favor}

	bytes, err := json.Marshal(pod)
	if err != nil {
		return nil, err
	}
	_, err = m.Put(ctx, key, string(bytes))
	return pod, err
}

// GetPod get a pod from etcd
// storage path in etcd is `/pod/info/:podname`
func (m *Mercury) GetPod(ctx context.Context, name string) (*types.Pod, error) {
	key := fmt.Sprintf(podInfoKey, name)

	ev, err := m.GetOne(ctx, key)
	if err != nil {
		return nil, err
	}

	pod := &types.Pod{}
	if err = json.Unmarshal(ev.Value, pod); err != nil {
		return nil, err
	}
	return pod, err
}

// RemovePod if the pod has no nodes left, otherwise return an error
func (m *Mercury) RemovePod(ctx context.Context, podname string) error {
	key := fmt.Sprintf(podInfoKey, podname)

	ns, err := m.GetNodesByPod(ctx, podname)
	if err != nil {
		return err
	}

	if l := len(ns); l != 0 {
		return types.NewDetailedErr(types.ErrPodHasNodes,
			fmt.Sprintf("pod %s still has %d nodes, delete them first", podname, l))
	}

	_, err = m.Delete(ctx, key)
	return err
}

// GetAllPods get all pods in etcd
// any error will break and return error immediately
// storage path in etcd is `/pod`
func (m *Mercury) GetAllPods(ctx context.Context) ([]*types.Pod, error) {
	resp, err := m.Get(ctx, fmt.Sprintf(podInfoKey, ""), clientv3.WithPrefix())
	if err != nil {
		return []*types.Pod{}, err
	}

	pods := []*types.Pod{}
	for _, ev := range resp.Kvs {
		pod := &types.Pod{}
		if err := json.Unmarshal(ev.Value, pod); err != nil {
			return pods, err
		}
		pods = append(pods, pod)
	}
	return pods, nil
}

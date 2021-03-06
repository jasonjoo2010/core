package complexscheduler

import (
	"fmt"
	"sort"

	"github.com/projecteru2/core/types"
	log "github.com/sirupsen/logrus"
)

// Potassium is a scheduler
type Potassium struct {
	maxshare, sharebase int
}

// New a potassium
func New(config types.Config) (*Potassium, error) {
	return &Potassium{config.Scheduler.MaxShare, config.Scheduler.ShareBase}, nil
}

// MaxIdleNode use for build
func (m *Potassium) MaxIdleNode(nodes []*types.Node) (*types.Node, error) {
	if len(nodes) < 1 {
		return nil, types.ErrInsufficientNodes
	}
	pos := 0
	node := nodes[pos]
	min := float64(node.CPU.Total())/float64(node.InitCPU.Total()) + float64(node.MemCap)/float64(node.InitMemCap)
	for i, node := range nodes {
		idle := float64(node.CPU.Total())/float64(node.InitCPU.Total()) + float64(node.MemCap)/float64(node.InitMemCap)
		if idle < min {
			pos = i
			min = idle
		}
	}
	return nodes[pos], nil
}

// SelectMemoryNodes filter nodes with enough memory
func (m *Potassium) SelectMemoryNodes(nodesInfo []types.NodeInfo, quota float64, memory int64) ([]types.NodeInfo, int, error) {
	log.Debugf("[SelectMemoryNodes] nodesInfo: %v, cpu: %f, memory: %d", nodesInfo, quota, memory)
	if memory <= 0 {
		return nil, 0, types.ErrNegativeMemory
	}

	nodesInfoLength := len(nodesInfo)

	// 筛选出能满足 CPU 需求的
	sort.Slice(nodesInfo, func(i, j int) bool { return len(nodesInfo[i].CPUMap) < len(nodesInfo[j].CPUMap) })
	p := sort.Search(nodesInfoLength, func(i int) bool { return float64(len(nodesInfo[i].CPUMap)) >= quota })
	// p 最大也就是 nodesInfoLength - 1
	if p == nodesInfoLength {
		return nil, 0, types.ErrInsufficientCPU
	}
	nodesInfoLength -= p
	nodesInfo = nodesInfo[p:]

	// 计算是否有足够的内存满足需求
	sort.Slice(nodesInfo, func(i, j int) bool { return nodesInfo[i].MemCap < nodesInfo[j].MemCap })
	p = sort.Search(nodesInfoLength, func(i int) bool { return nodesInfo[i].MemCap >= memory })
	if p == nodesInfoLength {
		return nil, 0, types.ErrInsufficientMEM
	}
	nodesInfoLength -= p
	nodesInfo = nodesInfo[p:]

	// 这里 memCap 一定是大于 memory 的所以不用判断 cap 内容
	volTotal := 0
	for i, nodeInfo := range nodesInfo {
		capacity := int(nodeInfo.MemCap / memory)
		volTotal += capacity
		nodesInfo[i].Capacity = capacity
	}
	return nodesInfo, volTotal, nil
}

// SelectCPUNodes select nodes with enough cpus
func (m *Potassium) SelectCPUNodes(nodesInfo []types.NodeInfo, quota float64, memory int64) ([]types.NodeInfo, map[string][]types.CPUMap, int, error) {
	log.Debugf("[SelectCPUNodes] nodesInfo: %v, cpu: %f memory: %d", nodesInfo, quota, memory)
	if quota <= 0 {
		return nil, nil, 0, types.ErrNegativeQuota
	}
	if len(nodesInfo) == 0 {
		return nil, nil, 0, types.ErrZeroNodes
	}
	return cpuPriorPlan(quota, memory, nodesInfo, m.maxshare, m.sharebase)
}

// CommonDivision deploy containers by their deploy status
// 部署完 N 个后全局尽可能平均
// need 是所需总量，total 是支持部署总量
func (m *Potassium) CommonDivision(nodesInfo []types.NodeInfo, need, total int) ([]types.NodeInfo, error) {
	if total < need {
		return nil, types.NewDetailedErr(types.ErrInsufficientRes,
			fmt.Sprintf("need: %d, vol: %d", need, total))
	}
	return CommunismDivisionPlan(nodesInfo, need)
}

// EachDivision deploy containers by each node
// 容量够的机器每一台部署 N 个
// need 是每台机器所需总量，limit 是限制节点数
func (m *Potassium) EachDivision(nodesInfo []types.NodeInfo, need, limit int) ([]types.NodeInfo, error) {
	return AveragePlan(nodesInfo, need, limit)
}

// FillDivision deploy containers fill nodes by count
// 根据之前部署的策略每一台补充到 N 个，超过 N 个忽略
// need 是每台上限, limit 是限制节点数
func (m *Potassium) FillDivision(nodesInfo []types.NodeInfo, need, limit int) ([]types.NodeInfo, error) {
	return FillPlan(nodesInfo, need, limit)
}

// GlobalDivision deploy containers by their resource costs
// 尽量使得资源消耗平均
// need 是所需总量，total 是支持部署总量
func (m *Potassium) GlobalDivision(nodesInfo []types.NodeInfo, need, total int) ([]types.NodeInfo, error) {
	if total < need {
		return nil, types.NewDetailedErr(types.ErrInsufficientRes,
			fmt.Sprintf("need: %d, vol: %d", need, total))
	}
	return GlobalDivisionPlan(nodesInfo, need)
}

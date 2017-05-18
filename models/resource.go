package models

type Resource struct {
	Id                  int64
	AlgorithmResource   string
	CpuTotalResource    float64
	CpuUsageResource    float64
	CpuUnit             string
	MemoryTotalResource float64
	MemoryUsageResource float64
	MemoryUnit          string
	User                *User
	QuotaNamespace      string
	QuotaName           string
}

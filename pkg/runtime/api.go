package runtime

import (
	"net/http"

	"github.com/pumphouse-p/peek-go/pkg/apiutils"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type CPUProperties struct {
	Num        int      `json:"num"`
	VendorID   string   `json:"vendor_id"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physical_id"`
	CoreID     string   `json:"core_id"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"model_name"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cache_size"`
	Flags      []string `json:"flags"`
}

type CPUStatus struct {
	CPU           []CPUProperties `json:"cpus"`
	CPUUsed       []float64       `json:"cpu_used"`
	CoreUsed      []float64       `json:"core_used"`
	PhysicalCores int             `json:"physical_cores"`
	LogicalCores  int             `json:"logical_cores"`
}

type MemStats struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func initCPUProperties() (cp []CPUProperties) {

	cpuInfo, _ := cpu.Info()

	for _, ci := range cpuInfo {
		core := CPUProperties{}
		core.Num = int(ci.CPU)
		core.VendorID = ci.VendorID
		core.Family = ci.Family
		core.Model = ci.Model
		core.Stepping = ci.Stepping
		core.PhysicalID = ci.PhysicalID
		core.CoreID = ci.CoreID
		core.Cores = ci.Cores
		core.ModelName = ci.ModelName
		core.Mhz = ci.Mhz
		core.CacheSize = ci.CacheSize
		core.Flags = ci.Flags
		cp = append(cp, core)
	}

	return cp
}

func (cs *CPUStatus) Init() {
	cs.CPU = initCPUProperties()
	cs.CPUUsed, _ = cpu.Percent(0, false)
	cs.CoreUsed, _ = cpu.Percent(0, true)
	cs.PhysicalCores, _ = cpu.Counts(false)
	cs.LogicalCores, _ = cpu.Counts(true)
}

func (ms *MemStats) Init() {
	vm, _ := mem.VirtualMemory()
	ms.Total = vm.Total
	ms.Available = vm.Available
	ms.Used = vm.Used
	ms.UsedPercent = vm.UsedPercent
}

type RuntimeStatus struct {
	CPU CPUStatus `json:"cpu"`
	Mem MemStats  `json:"mem"`
}

type Runtime struct{}

func New() *Runtime {
	return &Runtime{}
}

func (rt *Runtime) APIGetCPU(w http.ResponseWriter, r *http.Request) {
	cs := CPUStatus{}
	cs.Init()

	apiutils.ServeJSON(w, cs)
}

func (rt *Runtime) APIGetMem(w http.ResponseWriter, r *http.Request) {
	ms := MemStats{}
	ms.Init()

	apiutils.ServeJSON(w, ms)
}

func (rt *Runtime) APIGetRuntime(w http.ResponseWriter, r *http.Request) {
	rs := RuntimeStatus{}
	rs.CPU = CPUStatus{}
	rs.Mem = MemStats{}

	rs.CPU.Init()
	rs.Mem.Init()

	apiutils.ServeJSON(w, rs)
}

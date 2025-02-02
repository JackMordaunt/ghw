//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"github.com/jackmordaunt/ghw/pkg/baseboard"
	"github.com/jackmordaunt/ghw/pkg/chassis"
	"github.com/jackmordaunt/ghw/pkg/cpu"
	"github.com/jackmordaunt/ghw/pkg/gpu"
	"github.com/jackmordaunt/ghw/pkg/memory"
	"github.com/jackmordaunt/ghw/pkg/option"
	"github.com/jackmordaunt/ghw/pkg/pci"
	pciaddress "github.com/jackmordaunt/ghw/pkg/pci/address"
	"github.com/jackmordaunt/ghw/pkg/product"
	"github.com/jackmordaunt/ghw/pkg/topology"
)

type WithOption = option.Option

var (
	WithChroot      = option.WithChroot
	WithSnapshot    = option.WithSnapshot
	WithAlerter     = option.WithAlerter
	WithNullAlerter = option.WithNullAlerter
	// match the existing environ variable to minimize surprises
	WithDisableWarnings = option.WithNullAlerter
	WithDisableTools    = option.WithDisableTools
	WithPathOverrides   = option.WithPathOverrides
)

type SnapshotOptions = option.SnapshotOptions

type PathOverrides = option.PathOverrides

type CPUInfo = cpu.Info

var (
	CPU = cpu.New
)

type MemoryArea = memory.Area
type MemoryInfo = memory.Info
type MemoryCacheType = memory.CacheType
type MemoryModule = memory.Module

const (
	MEMORY_CACHE_TYPE_UNIFIED     = memory.CACHE_TYPE_UNIFIED
	MEMORY_CACHE_TYPE_INSTRUCTION = memory.CACHE_TYPE_INSTRUCTION
	MEMORY_CACHE_TYPE_DATA        = memory.CACHE_TYPE_DATA
)

var (
	Memory = memory.New
)

type ChassisInfo = chassis.Info

var (
	Chassis = chassis.New
)

type BaseboardInfo = baseboard.Info

var (
	Baseboard = baseboard.New
)

type TopologyInfo = topology.Info
type TopologyNode = topology.Node

var (
	Topology = topology.New
)

type Architecture = topology.Architecture

const (
	ARCHITECTURE_SMP  = topology.ARCHITECTURE_SMP
	ARCHITECTURE_NUMA = topology.ARCHITECTURE_NUMA
)

type PCIInfo = pci.Info
type PCIAddress = pciaddress.Address
type PCIDevice = pci.Device

var (
	PCI                  = pci.New
	PCIAddressFromString = pciaddress.FromString
)

type ProductInfo = product.Info

var (
	Product = product.New
)

type GPUInfo = gpu.Info
type GraphicsCard = gpu.GraphicsCard

var (
	GPU = gpu.New
)

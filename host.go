//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"fmt"

	"github.com/jackmordaunt/ghw/pkg/context"

	"github.com/jackmordaunt/ghw/pkg/baseboard"
	"github.com/jackmordaunt/ghw/pkg/bios"
	"github.com/jackmordaunt/ghw/pkg/block"
	"github.com/jackmordaunt/ghw/pkg/chassis"
	"github.com/jackmordaunt/ghw/pkg/cpu"
	"github.com/jackmordaunt/ghw/pkg/gpu"
	"github.com/jackmordaunt/ghw/pkg/memory"
	"github.com/jackmordaunt/ghw/pkg/net"
	"github.com/jackmordaunt/ghw/pkg/pci"
	"github.com/jackmordaunt/ghw/pkg/product"
	"github.com/jackmordaunt/ghw/pkg/topology"
)

// HostInfo is a wrapper struct containing information about the host system's
// memory, block storage, CPU, etc
type HostInfo struct {
	ctx       *context.Context
	Memory    *memory.Info    `json:"memory"`
	Block     *block.Info     `json:"block"`
	CPU       *cpu.Info       `json:"cpu"`
	Topology  *topology.Info  `json:"topology"`
	Network   *net.Info       `json:"network"`
	GPU       *gpu.Info       `json:"gpu"`
	Chassis   *chassis.Info   `json:"chassis"`
	BIOS      *bios.Info      `json:"bios"`
	Baseboard *baseboard.Info `json:"baseboard"`
	Product   *product.Info   `json:"product"`
	PCI       *pci.Info       `json:"pci"`
}

// Host returns a pointer to a HostInfo struct that contains fields with
// information about the host system's CPU, memory, network devices, etc
func Host(opts ...*WithOption) (*HostInfo, error) {
	ctx := context.New(opts...)

	memInfo, err := memory.New(opts...)
	if err != nil {
		return nil, err
	}
	blockInfo, err := block.New(opts...)
	if err != nil {
		return nil, err
	}
	cpuInfo, err := cpu.New(opts...)
	if err != nil {
		return nil, err
	}
	topologyInfo, err := topology.New(opts...)
	if err != nil {
		return nil, err
	}
	netInfo, err := net.New(opts...)
	if err != nil {
		return nil, err
	}
	gpuInfo, err := gpu.New(opts...)
	if err != nil {
		return nil, err
	}
	chassisInfo, err := chassis.New(opts...)
	if err != nil {
		return nil, err
	}
	biosInfo, err := bios.New(opts...)
	if err != nil {
		return nil, err
	}
	baseboardInfo, err := baseboard.New(opts...)
	if err != nil {
		return nil, err
	}
	productInfo, err := product.New(opts...)
	if err != nil {
		return nil, err
	}
	pciInfo, err := pci.New(opts...)
	if err != nil {
		return nil, err
	}
	return &HostInfo{
		ctx:       ctx,
		CPU:       cpuInfo,
		Memory:    memInfo,
		Block:     blockInfo,
		Topology:  topologyInfo,
		Network:   netInfo,
		GPU:       gpuInfo,
		Chassis:   chassisInfo,
		BIOS:      biosInfo,
		Baseboard: baseboardInfo,
		Product:   productInfo,
		PCI:       pciInfo,
	}, nil
}

// String returns a newline-separated output of the HostInfo's component
// structs' String-ified output
func (info *HostInfo) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		info.Block.String(),
		info.CPU.String(),
		info.GPU.String(),
		info.Memory.String(),
		info.Network.String(),
		info.Topology.String(),
		info.Chassis.String(),
		info.BIOS.String(),
		info.Baseboard.String(),
		info.Product.String(),
		info.PCI.String(),
	)
}

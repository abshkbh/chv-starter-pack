/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VmConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmConfig{}

// VmConfig Virtual machine configuration
type VmConfig struct {
	Cpus *CpusConfig `json:"cpus,omitempty"`
	Memory *MemoryConfig `json:"memory,omitempty"`
	Payload PayloadConfig `json:"payload"`
	RateLimitGroups []RateLimitGroupConfig `json:"rate_limit_groups,omitempty"`
	Disks []DiskConfig `json:"disks,omitempty"`
	Net []NetConfig `json:"net,omitempty"`
	Rng *RngConfig `json:"rng,omitempty"`
	Balloon *BalloonConfig `json:"balloon,omitempty"`
	Fs []FsConfig `json:"fs,omitempty"`
	Pmem []PmemConfig `json:"pmem,omitempty"`
	Serial *ConsoleConfig `json:"serial,omitempty"`
	Console *ConsoleConfig `json:"console,omitempty"`
	DebugConsole *DebugConsoleConfig `json:"debug_console,omitempty"`
	Devices []DeviceConfig `json:"devices,omitempty"`
	Vdpa []VdpaConfig `json:"vdpa,omitempty"`
	Vsock *VsockConfig `json:"vsock,omitempty"`
	SgxEpc []SgxEpcConfig `json:"sgx_epc,omitempty"`
	Numa []NumaConfig `json:"numa,omitempty"`
	Iommu *bool `json:"iommu,omitempty"`
	Watchdog *bool `json:"watchdog,omitempty"`
	Pvpanic *bool `json:"pvpanic,omitempty"`
	PciSegments []PciSegmentConfig `json:"pci_segments,omitempty"`
	Platform *PlatformConfig `json:"platform,omitempty"`
	Tpm *TpmConfig `json:"tpm,omitempty"`
}

type _VmConfig VmConfig

// NewVmConfig instantiates a new VmConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmConfig(payload PayloadConfig) *VmConfig {
	this := VmConfig{}
	this.Payload = payload
	var iommu bool = false
	this.Iommu = &iommu
	var watchdog bool = false
	this.Watchdog = &watchdog
	var pvpanic bool = false
	this.Pvpanic = &pvpanic
	return &this
}

// NewVmConfigWithDefaults instantiates a new VmConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmConfigWithDefaults() *VmConfig {
	this := VmConfig{}
	var iommu bool = false
	this.Iommu = &iommu
	var watchdog bool = false
	this.Watchdog = &watchdog
	var pvpanic bool = false
	this.Pvpanic = &pvpanic
	return &this
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *VmConfig) GetCpus() CpusConfig {
	if o == nil || IsNil(o.Cpus) {
		var ret CpusConfig
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetCpusOk() (*CpusConfig, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *VmConfig) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given CpusConfig and assigns it to the Cpus field.
func (o *VmConfig) SetCpus(v CpusConfig) {
	o.Cpus = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *VmConfig) GetMemory() MemoryConfig {
	if o == nil || IsNil(o.Memory) {
		var ret MemoryConfig
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetMemoryOk() (*MemoryConfig, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *VmConfig) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given MemoryConfig and assigns it to the Memory field.
func (o *VmConfig) SetMemory(v MemoryConfig) {
	o.Memory = &v
}

// GetPayload returns the Payload field value
func (o *VmConfig) GetPayload() PayloadConfig {
	if o == nil {
		var ret PayloadConfig
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *VmConfig) GetPayloadOk() (*PayloadConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *VmConfig) SetPayload(v PayloadConfig) {
	o.Payload = v
}

// GetRateLimitGroups returns the RateLimitGroups field value if set, zero value otherwise.
func (o *VmConfig) GetRateLimitGroups() []RateLimitGroupConfig {
	if o == nil || IsNil(o.RateLimitGroups) {
		var ret []RateLimitGroupConfig
		return ret
	}
	return o.RateLimitGroups
}

// GetRateLimitGroupsOk returns a tuple with the RateLimitGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetRateLimitGroupsOk() ([]RateLimitGroupConfig, bool) {
	if o == nil || IsNil(o.RateLimitGroups) {
		return nil, false
	}
	return o.RateLimitGroups, true
}

// HasRateLimitGroups returns a boolean if a field has been set.
func (o *VmConfig) HasRateLimitGroups() bool {
	if o != nil && !IsNil(o.RateLimitGroups) {
		return true
	}

	return false
}

// SetRateLimitGroups gets a reference to the given []RateLimitGroupConfig and assigns it to the RateLimitGroups field.
func (o *VmConfig) SetRateLimitGroups(v []RateLimitGroupConfig) {
	o.RateLimitGroups = v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *VmConfig) GetDisks() []DiskConfig {
	if o == nil || IsNil(o.Disks) {
		var ret []DiskConfig
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetDisksOk() ([]DiskConfig, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *VmConfig) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []DiskConfig and assigns it to the Disks field.
func (o *VmConfig) SetDisks(v []DiskConfig) {
	o.Disks = v
}

// GetNet returns the Net field value if set, zero value otherwise.
func (o *VmConfig) GetNet() []NetConfig {
	if o == nil || IsNil(o.Net) {
		var ret []NetConfig
		return ret
	}
	return o.Net
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetNetOk() ([]NetConfig, bool) {
	if o == nil || IsNil(o.Net) {
		return nil, false
	}
	return o.Net, true
}

// HasNet returns a boolean if a field has been set.
func (o *VmConfig) HasNet() bool {
	if o != nil && !IsNil(o.Net) {
		return true
	}

	return false
}

// SetNet gets a reference to the given []NetConfig and assigns it to the Net field.
func (o *VmConfig) SetNet(v []NetConfig) {
	o.Net = v
}

// GetRng returns the Rng field value if set, zero value otherwise.
func (o *VmConfig) GetRng() RngConfig {
	if o == nil || IsNil(o.Rng) {
		var ret RngConfig
		return ret
	}
	return *o.Rng
}

// GetRngOk returns a tuple with the Rng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetRngOk() (*RngConfig, bool) {
	if o == nil || IsNil(o.Rng) {
		return nil, false
	}
	return o.Rng, true
}

// HasRng returns a boolean if a field has been set.
func (o *VmConfig) HasRng() bool {
	if o != nil && !IsNil(o.Rng) {
		return true
	}

	return false
}

// SetRng gets a reference to the given RngConfig and assigns it to the Rng field.
func (o *VmConfig) SetRng(v RngConfig) {
	o.Rng = &v
}

// GetBalloon returns the Balloon field value if set, zero value otherwise.
func (o *VmConfig) GetBalloon() BalloonConfig {
	if o == nil || IsNil(o.Balloon) {
		var ret BalloonConfig
		return ret
	}
	return *o.Balloon
}

// GetBalloonOk returns a tuple with the Balloon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetBalloonOk() (*BalloonConfig, bool) {
	if o == nil || IsNil(o.Balloon) {
		return nil, false
	}
	return o.Balloon, true
}

// HasBalloon returns a boolean if a field has been set.
func (o *VmConfig) HasBalloon() bool {
	if o != nil && !IsNil(o.Balloon) {
		return true
	}

	return false
}

// SetBalloon gets a reference to the given BalloonConfig and assigns it to the Balloon field.
func (o *VmConfig) SetBalloon(v BalloonConfig) {
	o.Balloon = &v
}

// GetFs returns the Fs field value if set, zero value otherwise.
func (o *VmConfig) GetFs() []FsConfig {
	if o == nil || IsNil(o.Fs) {
		var ret []FsConfig
		return ret
	}
	return o.Fs
}

// GetFsOk returns a tuple with the Fs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetFsOk() ([]FsConfig, bool) {
	if o == nil || IsNil(o.Fs) {
		return nil, false
	}
	return o.Fs, true
}

// HasFs returns a boolean if a field has been set.
func (o *VmConfig) HasFs() bool {
	if o != nil && !IsNil(o.Fs) {
		return true
	}

	return false
}

// SetFs gets a reference to the given []FsConfig and assigns it to the Fs field.
func (o *VmConfig) SetFs(v []FsConfig) {
	o.Fs = v
}

// GetPmem returns the Pmem field value if set, zero value otherwise.
func (o *VmConfig) GetPmem() []PmemConfig {
	if o == nil || IsNil(o.Pmem) {
		var ret []PmemConfig
		return ret
	}
	return o.Pmem
}

// GetPmemOk returns a tuple with the Pmem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetPmemOk() ([]PmemConfig, bool) {
	if o == nil || IsNil(o.Pmem) {
		return nil, false
	}
	return o.Pmem, true
}

// HasPmem returns a boolean if a field has been set.
func (o *VmConfig) HasPmem() bool {
	if o != nil && !IsNil(o.Pmem) {
		return true
	}

	return false
}

// SetPmem gets a reference to the given []PmemConfig and assigns it to the Pmem field.
func (o *VmConfig) SetPmem(v []PmemConfig) {
	o.Pmem = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *VmConfig) GetSerial() ConsoleConfig {
	if o == nil || IsNil(o.Serial) {
		var ret ConsoleConfig
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetSerialOk() (*ConsoleConfig, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *VmConfig) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given ConsoleConfig and assigns it to the Serial field.
func (o *VmConfig) SetSerial(v ConsoleConfig) {
	o.Serial = &v
}

// GetConsole returns the Console field value if set, zero value otherwise.
func (o *VmConfig) GetConsole() ConsoleConfig {
	if o == nil || IsNil(o.Console) {
		var ret ConsoleConfig
		return ret
	}
	return *o.Console
}

// GetConsoleOk returns a tuple with the Console field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetConsoleOk() (*ConsoleConfig, bool) {
	if o == nil || IsNil(o.Console) {
		return nil, false
	}
	return o.Console, true
}

// HasConsole returns a boolean if a field has been set.
func (o *VmConfig) HasConsole() bool {
	if o != nil && !IsNil(o.Console) {
		return true
	}

	return false
}

// SetConsole gets a reference to the given ConsoleConfig and assigns it to the Console field.
func (o *VmConfig) SetConsole(v ConsoleConfig) {
	o.Console = &v
}

// GetDebugConsole returns the DebugConsole field value if set, zero value otherwise.
func (o *VmConfig) GetDebugConsole() DebugConsoleConfig {
	if o == nil || IsNil(o.DebugConsole) {
		var ret DebugConsoleConfig
		return ret
	}
	return *o.DebugConsole
}

// GetDebugConsoleOk returns a tuple with the DebugConsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetDebugConsoleOk() (*DebugConsoleConfig, bool) {
	if o == nil || IsNil(o.DebugConsole) {
		return nil, false
	}
	return o.DebugConsole, true
}

// HasDebugConsole returns a boolean if a field has been set.
func (o *VmConfig) HasDebugConsole() bool {
	if o != nil && !IsNil(o.DebugConsole) {
		return true
	}

	return false
}

// SetDebugConsole gets a reference to the given DebugConsoleConfig and assigns it to the DebugConsole field.
func (o *VmConfig) SetDebugConsole(v DebugConsoleConfig) {
	o.DebugConsole = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *VmConfig) GetDevices() []DeviceConfig {
	if o == nil || IsNil(o.Devices) {
		var ret []DeviceConfig
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetDevicesOk() ([]DeviceConfig, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *VmConfig) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []DeviceConfig and assigns it to the Devices field.
func (o *VmConfig) SetDevices(v []DeviceConfig) {
	o.Devices = v
}

// GetVdpa returns the Vdpa field value if set, zero value otherwise.
func (o *VmConfig) GetVdpa() []VdpaConfig {
	if o == nil || IsNil(o.Vdpa) {
		var ret []VdpaConfig
		return ret
	}
	return o.Vdpa
}

// GetVdpaOk returns a tuple with the Vdpa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetVdpaOk() ([]VdpaConfig, bool) {
	if o == nil || IsNil(o.Vdpa) {
		return nil, false
	}
	return o.Vdpa, true
}

// HasVdpa returns a boolean if a field has been set.
func (o *VmConfig) HasVdpa() bool {
	if o != nil && !IsNil(o.Vdpa) {
		return true
	}

	return false
}

// SetVdpa gets a reference to the given []VdpaConfig and assigns it to the Vdpa field.
func (o *VmConfig) SetVdpa(v []VdpaConfig) {
	o.Vdpa = v
}

// GetVsock returns the Vsock field value if set, zero value otherwise.
func (o *VmConfig) GetVsock() VsockConfig {
	if o == nil || IsNil(o.Vsock) {
		var ret VsockConfig
		return ret
	}
	return *o.Vsock
}

// GetVsockOk returns a tuple with the Vsock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetVsockOk() (*VsockConfig, bool) {
	if o == nil || IsNil(o.Vsock) {
		return nil, false
	}
	return o.Vsock, true
}

// HasVsock returns a boolean if a field has been set.
func (o *VmConfig) HasVsock() bool {
	if o != nil && !IsNil(o.Vsock) {
		return true
	}

	return false
}

// SetVsock gets a reference to the given VsockConfig and assigns it to the Vsock field.
func (o *VmConfig) SetVsock(v VsockConfig) {
	o.Vsock = &v
}

// GetSgxEpc returns the SgxEpc field value if set, zero value otherwise.
func (o *VmConfig) GetSgxEpc() []SgxEpcConfig {
	if o == nil || IsNil(o.SgxEpc) {
		var ret []SgxEpcConfig
		return ret
	}
	return o.SgxEpc
}

// GetSgxEpcOk returns a tuple with the SgxEpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetSgxEpcOk() ([]SgxEpcConfig, bool) {
	if o == nil || IsNil(o.SgxEpc) {
		return nil, false
	}
	return o.SgxEpc, true
}

// HasSgxEpc returns a boolean if a field has been set.
func (o *VmConfig) HasSgxEpc() bool {
	if o != nil && !IsNil(o.SgxEpc) {
		return true
	}

	return false
}

// SetSgxEpc gets a reference to the given []SgxEpcConfig and assigns it to the SgxEpc field.
func (o *VmConfig) SetSgxEpc(v []SgxEpcConfig) {
	o.SgxEpc = v
}

// GetNuma returns the Numa field value if set, zero value otherwise.
func (o *VmConfig) GetNuma() []NumaConfig {
	if o == nil || IsNil(o.Numa) {
		var ret []NumaConfig
		return ret
	}
	return o.Numa
}

// GetNumaOk returns a tuple with the Numa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetNumaOk() ([]NumaConfig, bool) {
	if o == nil || IsNil(o.Numa) {
		return nil, false
	}
	return o.Numa, true
}

// HasNuma returns a boolean if a field has been set.
func (o *VmConfig) HasNuma() bool {
	if o != nil && !IsNil(o.Numa) {
		return true
	}

	return false
}

// SetNuma gets a reference to the given []NumaConfig and assigns it to the Numa field.
func (o *VmConfig) SetNuma(v []NumaConfig) {
	o.Numa = v
}

// GetIommu returns the Iommu field value if set, zero value otherwise.
func (o *VmConfig) GetIommu() bool {
	if o == nil || IsNil(o.Iommu) {
		var ret bool
		return ret
	}
	return *o.Iommu
}

// GetIommuOk returns a tuple with the Iommu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetIommuOk() (*bool, bool) {
	if o == nil || IsNil(o.Iommu) {
		return nil, false
	}
	return o.Iommu, true
}

// HasIommu returns a boolean if a field has been set.
func (o *VmConfig) HasIommu() bool {
	if o != nil && !IsNil(o.Iommu) {
		return true
	}

	return false
}

// SetIommu gets a reference to the given bool and assigns it to the Iommu field.
func (o *VmConfig) SetIommu(v bool) {
	o.Iommu = &v
}

// GetWatchdog returns the Watchdog field value if set, zero value otherwise.
func (o *VmConfig) GetWatchdog() bool {
	if o == nil || IsNil(o.Watchdog) {
		var ret bool
		return ret
	}
	return *o.Watchdog
}

// GetWatchdogOk returns a tuple with the Watchdog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetWatchdogOk() (*bool, bool) {
	if o == nil || IsNil(o.Watchdog) {
		return nil, false
	}
	return o.Watchdog, true
}

// HasWatchdog returns a boolean if a field has been set.
func (o *VmConfig) HasWatchdog() bool {
	if o != nil && !IsNil(o.Watchdog) {
		return true
	}

	return false
}

// SetWatchdog gets a reference to the given bool and assigns it to the Watchdog field.
func (o *VmConfig) SetWatchdog(v bool) {
	o.Watchdog = &v
}

// GetPvpanic returns the Pvpanic field value if set, zero value otherwise.
func (o *VmConfig) GetPvpanic() bool {
	if o == nil || IsNil(o.Pvpanic) {
		var ret bool
		return ret
	}
	return *o.Pvpanic
}

// GetPvpanicOk returns a tuple with the Pvpanic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetPvpanicOk() (*bool, bool) {
	if o == nil || IsNil(o.Pvpanic) {
		return nil, false
	}
	return o.Pvpanic, true
}

// HasPvpanic returns a boolean if a field has been set.
func (o *VmConfig) HasPvpanic() bool {
	if o != nil && !IsNil(o.Pvpanic) {
		return true
	}

	return false
}

// SetPvpanic gets a reference to the given bool and assigns it to the Pvpanic field.
func (o *VmConfig) SetPvpanic(v bool) {
	o.Pvpanic = &v
}

// GetPciSegments returns the PciSegments field value if set, zero value otherwise.
func (o *VmConfig) GetPciSegments() []PciSegmentConfig {
	if o == nil || IsNil(o.PciSegments) {
		var ret []PciSegmentConfig
		return ret
	}
	return o.PciSegments
}

// GetPciSegmentsOk returns a tuple with the PciSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetPciSegmentsOk() ([]PciSegmentConfig, bool) {
	if o == nil || IsNil(o.PciSegments) {
		return nil, false
	}
	return o.PciSegments, true
}

// HasPciSegments returns a boolean if a field has been set.
func (o *VmConfig) HasPciSegments() bool {
	if o != nil && !IsNil(o.PciSegments) {
		return true
	}

	return false
}

// SetPciSegments gets a reference to the given []PciSegmentConfig and assigns it to the PciSegments field.
func (o *VmConfig) SetPciSegments(v []PciSegmentConfig) {
	o.PciSegments = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *VmConfig) GetPlatform() PlatformConfig {
	if o == nil || IsNil(o.Platform) {
		var ret PlatformConfig
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetPlatformOk() (*PlatformConfig, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *VmConfig) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformConfig and assigns it to the Platform field.
func (o *VmConfig) SetPlatform(v PlatformConfig) {
	o.Platform = &v
}

// GetTpm returns the Tpm field value if set, zero value otherwise.
func (o *VmConfig) GetTpm() TpmConfig {
	if o == nil || IsNil(o.Tpm) {
		var ret TpmConfig
		return ret
	}
	return *o.Tpm
}

// GetTpmOk returns a tuple with the Tpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConfig) GetTpmOk() (*TpmConfig, bool) {
	if o == nil || IsNil(o.Tpm) {
		return nil, false
	}
	return o.Tpm, true
}

// HasTpm returns a boolean if a field has been set.
func (o *VmConfig) HasTpm() bool {
	if o != nil && !IsNil(o.Tpm) {
		return true
	}

	return false
}

// SetTpm gets a reference to the given TpmConfig and assigns it to the Tpm field.
func (o *VmConfig) SetTpm(v TpmConfig) {
	o.Tpm = &v
}

func (o VmConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	toSerialize["payload"] = o.Payload
	if !IsNil(o.RateLimitGroups) {
		toSerialize["rate_limit_groups"] = o.RateLimitGroups
	}
	if !IsNil(o.Disks) {
		toSerialize["disks"] = o.Disks
	}
	if !IsNil(o.Net) {
		toSerialize["net"] = o.Net
	}
	if !IsNil(o.Rng) {
		toSerialize["rng"] = o.Rng
	}
	if !IsNil(o.Balloon) {
		toSerialize["balloon"] = o.Balloon
	}
	if !IsNil(o.Fs) {
		toSerialize["fs"] = o.Fs
	}
	if !IsNil(o.Pmem) {
		toSerialize["pmem"] = o.Pmem
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Console) {
		toSerialize["console"] = o.Console
	}
	if !IsNil(o.DebugConsole) {
		toSerialize["debug_console"] = o.DebugConsole
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Vdpa) {
		toSerialize["vdpa"] = o.Vdpa
	}
	if !IsNil(o.Vsock) {
		toSerialize["vsock"] = o.Vsock
	}
	if !IsNil(o.SgxEpc) {
		toSerialize["sgx_epc"] = o.SgxEpc
	}
	if !IsNil(o.Numa) {
		toSerialize["numa"] = o.Numa
	}
	if !IsNil(o.Iommu) {
		toSerialize["iommu"] = o.Iommu
	}
	if !IsNil(o.Watchdog) {
		toSerialize["watchdog"] = o.Watchdog
	}
	if !IsNil(o.Pvpanic) {
		toSerialize["pvpanic"] = o.Pvpanic
	}
	if !IsNil(o.PciSegments) {
		toSerialize["pci_segments"] = o.PciSegments
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Tpm) {
		toSerialize["tpm"] = o.Tpm
	}
	return toSerialize, nil
}

func (o *VmConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payload",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVmConfig := _VmConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmConfig)

	if err != nil {
		return err
	}

	*o = VmConfig(varVmConfig)

	return err
}

type NullableVmConfig struct {
	value *VmConfig
	isSet bool
}

func (v NullableVmConfig) Get() *VmConfig {
	return v.value
}

func (v *NullableVmConfig) Set(val *VmConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVmConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVmConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmConfig(val *VmConfig) *NullableVmConfig {
	return &NullableVmConfig{value: val, isSet: true}
}

func (v NullableVmConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


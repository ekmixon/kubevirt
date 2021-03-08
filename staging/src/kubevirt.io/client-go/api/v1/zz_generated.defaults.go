// +build !ignore_autogenerated

/*
Copyright 2021 The KubeVirt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&VirtualMachine{}, func(obj interface{}) { SetObjectDefaults_VirtualMachine(obj.(*VirtualMachine)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachineInstance{}, func(obj interface{}) { SetObjectDefaults_VirtualMachineInstance(obj.(*VirtualMachineInstance)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachineInstanceList{}, func(obj interface{}) { SetObjectDefaults_VirtualMachineInstanceList(obj.(*VirtualMachineInstanceList)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachineInstancePreset{}, func(obj interface{}) {
		SetObjectDefaults_VirtualMachineInstancePreset(obj.(*VirtualMachineInstancePreset))
	})
	scheme.AddTypeDefaultingFunc(&VirtualMachineInstancePresetList{}, func(obj interface{}) {
		SetObjectDefaults_VirtualMachineInstancePresetList(obj.(*VirtualMachineInstancePresetList))
	})
	scheme.AddTypeDefaultingFunc(&VirtualMachineInstanceReplicaSet{}, func(obj interface{}) {
		SetObjectDefaults_VirtualMachineInstanceReplicaSet(obj.(*VirtualMachineInstanceReplicaSet))
	})
	scheme.AddTypeDefaultingFunc(&VirtualMachineInstanceReplicaSetList{}, func(obj interface{}) {
		SetObjectDefaults_VirtualMachineInstanceReplicaSetList(obj.(*VirtualMachineInstanceReplicaSetList))
	})
	scheme.AddTypeDefaultingFunc(&VirtualMachineList{}, func(obj interface{}) { SetObjectDefaults_VirtualMachineList(obj.(*VirtualMachineList)) })
	return nil
}

func SetObjectDefaults_VirtualMachine(in *VirtualMachine) {
	if in.Spec.Template != nil {
		if in.Spec.Template.Spec.Domain.Firmware != nil {
			SetDefaults_Firmware(in.Spec.Template.Spec.Domain.Firmware)
		}
		if in.Spec.Template.Spec.Domain.Clock != nil {
			if in.Spec.Template.Spec.Domain.Clock.Timer != nil {
				if in.Spec.Template.Spec.Domain.Clock.Timer.HPET != nil {
					SetDefaults_HPETTimer(in.Spec.Template.Spec.Domain.Clock.Timer.HPET)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.KVM != nil {
					SetDefaults_KVMTimer(in.Spec.Template.Spec.Domain.Clock.Timer.KVM)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.PIT != nil {
					SetDefaults_PITTimer(in.Spec.Template.Spec.Domain.Clock.Timer.PIT)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.RTC != nil {
					SetDefaults_RTCTimer(in.Spec.Template.Spec.Domain.Clock.Timer.RTC)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.Hyperv != nil {
					SetDefaults_HypervTimer(in.Spec.Template.Spec.Domain.Clock.Timer.Hyperv)
				}
			}
		}
		if in.Spec.Template.Spec.Domain.Features != nil {
			SetDefaults_FeatureState(&in.Spec.Template.Spec.Domain.Features.ACPI)
			if in.Spec.Template.Spec.Domain.Features.APIC != nil {
				SetDefaults_FeatureAPIC(in.Spec.Template.Spec.Domain.Features.APIC)
			}
			if in.Spec.Template.Spec.Domain.Features.Hyperv != nil {
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Relaxed != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Relaxed)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VAPIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.VAPIC)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Spinlocks != nil {
					SetDefaults_FeatureSpinlocks(in.Spec.Template.Spec.Domain.Features.Hyperv.Spinlocks)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VPIndex != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.VPIndex)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Runtime != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Runtime)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNIC)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
					SetDefaults_SyNICTimer(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer)
					if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer.Direct != nil {
						SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer.Direct)
					}
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Reset != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Reset)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VendorID != nil {
					SetDefaults_FeatureVendorID(in.Spec.Template.Spec.Domain.Features.Hyperv.VendorID)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Frequencies != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Frequencies)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Reenlightenment != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Reenlightenment)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.TLBFlush != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.TLBFlush)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.IPI != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.IPI)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.EVMCS != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.EVMCS)
				}
			}
			if in.Spec.Template.Spec.Domain.Features.SMM != nil {
				SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.SMM)
			}
			if in.Spec.Template.Spec.Domain.Features.Pvspinlock != nil {
				SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Pvspinlock)
			}
		}
		for i := range in.Spec.Template.Spec.Domain.Devices.Disks {
			a := &in.Spec.Template.Spec.Domain.Devices.Disks[i]
			SetDefaults_DiskDevice(&a.DiskDevice)
			if a.DiskDevice.Floppy != nil {
				SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
			}
			if a.DiskDevice.CDRom != nil {
				SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
			}
		}
		if in.Spec.Template.Spec.Domain.Devices.Watchdog != nil {
			SetDefaults_Watchdog(in.Spec.Template.Spec.Domain.Devices.Watchdog)
			if in.Spec.Template.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
				SetDefaults_I6300ESBWatchdog(in.Spec.Template.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
			}
		}
	}
	for i := range in.Status.VolumeRequests {
		a := &in.Status.VolumeRequests[i]
		if a.AddVolumeOptions != nil {
			if a.AddVolumeOptions.Disk != nil {
				SetDefaults_DiskDevice(&a.AddVolumeOptions.Disk.DiskDevice)
				if a.AddVolumeOptions.Disk.DiskDevice.Floppy != nil {
					SetDefaults_FloppyTarget(a.AddVolumeOptions.Disk.DiskDevice.Floppy)
				}
				if a.AddVolumeOptions.Disk.DiskDevice.CDRom != nil {
					SetDefaults_CDRomTarget(a.AddVolumeOptions.Disk.DiskDevice.CDRom)
				}
			}
		}
	}
}

func SetObjectDefaults_VirtualMachineInstance(in *VirtualMachineInstance) {
	SetDefaults_VirtualMachineInstance(in)
	if in.Spec.Domain.Firmware != nil {
		SetDefaults_Firmware(in.Spec.Domain.Firmware)
	}
	if in.Spec.Domain.Clock != nil {
		if in.Spec.Domain.Clock.Timer != nil {
			if in.Spec.Domain.Clock.Timer.HPET != nil {
				SetDefaults_HPETTimer(in.Spec.Domain.Clock.Timer.HPET)
			}
			if in.Spec.Domain.Clock.Timer.KVM != nil {
				SetDefaults_KVMTimer(in.Spec.Domain.Clock.Timer.KVM)
			}
			if in.Spec.Domain.Clock.Timer.PIT != nil {
				SetDefaults_PITTimer(in.Spec.Domain.Clock.Timer.PIT)
			}
			if in.Spec.Domain.Clock.Timer.RTC != nil {
				SetDefaults_RTCTimer(in.Spec.Domain.Clock.Timer.RTC)
			}
			if in.Spec.Domain.Clock.Timer.Hyperv != nil {
				SetDefaults_HypervTimer(in.Spec.Domain.Clock.Timer.Hyperv)
			}
		}
	}
	if in.Spec.Domain.Features != nil {
		SetDefaults_FeatureState(&in.Spec.Domain.Features.ACPI)
		if in.Spec.Domain.Features.APIC != nil {
			SetDefaults_FeatureAPIC(in.Spec.Domain.Features.APIC)
		}
		if in.Spec.Domain.Features.Hyperv != nil {
			if in.Spec.Domain.Features.Hyperv.Relaxed != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Relaxed)
			}
			if in.Spec.Domain.Features.Hyperv.VAPIC != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VAPIC)
			}
			if in.Spec.Domain.Features.Hyperv.Spinlocks != nil {
				SetDefaults_FeatureSpinlocks(in.Spec.Domain.Features.Hyperv.Spinlocks)
			}
			if in.Spec.Domain.Features.Hyperv.VPIndex != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VPIndex)
			}
			if in.Spec.Domain.Features.Hyperv.Runtime != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Runtime)
			}
			if in.Spec.Domain.Features.Hyperv.SyNIC != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNIC)
			}
			if in.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
				SetDefaults_SyNICTimer(in.Spec.Domain.Features.Hyperv.SyNICTimer)
				if in.Spec.Domain.Features.Hyperv.SyNICTimer.Direct != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNICTimer.Direct)
				}
			}
			if in.Spec.Domain.Features.Hyperv.Reset != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Reset)
			}
			if in.Spec.Domain.Features.Hyperv.VendorID != nil {
				SetDefaults_FeatureVendorID(in.Spec.Domain.Features.Hyperv.VendorID)
			}
			if in.Spec.Domain.Features.Hyperv.Frequencies != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Frequencies)
			}
			if in.Spec.Domain.Features.Hyperv.Reenlightenment != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Reenlightenment)
			}
			if in.Spec.Domain.Features.Hyperv.TLBFlush != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.TLBFlush)
			}
			if in.Spec.Domain.Features.Hyperv.IPI != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.IPI)
			}
			if in.Spec.Domain.Features.Hyperv.EVMCS != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.EVMCS)
			}
		}
		if in.Spec.Domain.Features.SMM != nil {
			SetDefaults_FeatureState(in.Spec.Domain.Features.SMM)
		}
		if in.Spec.Domain.Features.Pvspinlock != nil {
			SetDefaults_FeatureState(in.Spec.Domain.Features.Pvspinlock)
		}
	}
	for i := range in.Spec.Domain.Devices.Disks {
		a := &in.Spec.Domain.Devices.Disks[i]
		SetDefaults_DiskDevice(&a.DiskDevice)
		if a.DiskDevice.Floppy != nil {
			SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
		}
		if a.DiskDevice.CDRom != nil {
			SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
		}
	}
	if in.Spec.Domain.Devices.Watchdog != nil {
		SetDefaults_Watchdog(in.Spec.Domain.Devices.Watchdog)
		if in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
			SetDefaults_I6300ESBWatchdog(in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
		}
	}
}

func SetObjectDefaults_VirtualMachineInstanceList(in *VirtualMachineInstanceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachineInstance(a)
	}
}

func SetObjectDefaults_VirtualMachineInstancePreset(in *VirtualMachineInstancePreset) {
	if in.Spec.Domain != nil {
		if in.Spec.Domain.Firmware != nil {
			SetDefaults_Firmware(in.Spec.Domain.Firmware)
		}
		if in.Spec.Domain.Clock != nil {
			if in.Spec.Domain.Clock.Timer != nil {
				if in.Spec.Domain.Clock.Timer.HPET != nil {
					SetDefaults_HPETTimer(in.Spec.Domain.Clock.Timer.HPET)
				}
				if in.Spec.Domain.Clock.Timer.KVM != nil {
					SetDefaults_KVMTimer(in.Spec.Domain.Clock.Timer.KVM)
				}
				if in.Spec.Domain.Clock.Timer.PIT != nil {
					SetDefaults_PITTimer(in.Spec.Domain.Clock.Timer.PIT)
				}
				if in.Spec.Domain.Clock.Timer.RTC != nil {
					SetDefaults_RTCTimer(in.Spec.Domain.Clock.Timer.RTC)
				}
				if in.Spec.Domain.Clock.Timer.Hyperv != nil {
					SetDefaults_HypervTimer(in.Spec.Domain.Clock.Timer.Hyperv)
				}
			}
		}
		if in.Spec.Domain.Features != nil {
			SetDefaults_FeatureState(&in.Spec.Domain.Features.ACPI)
			if in.Spec.Domain.Features.APIC != nil {
				SetDefaults_FeatureAPIC(in.Spec.Domain.Features.APIC)
			}
			if in.Spec.Domain.Features.Hyperv != nil {
				if in.Spec.Domain.Features.Hyperv.Relaxed != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Relaxed)
				}
				if in.Spec.Domain.Features.Hyperv.VAPIC != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VAPIC)
				}
				if in.Spec.Domain.Features.Hyperv.Spinlocks != nil {
					SetDefaults_FeatureSpinlocks(in.Spec.Domain.Features.Hyperv.Spinlocks)
				}
				if in.Spec.Domain.Features.Hyperv.VPIndex != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VPIndex)
				}
				if in.Spec.Domain.Features.Hyperv.Runtime != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Runtime)
				}
				if in.Spec.Domain.Features.Hyperv.SyNIC != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNIC)
				}
				if in.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
					SetDefaults_SyNICTimer(in.Spec.Domain.Features.Hyperv.SyNICTimer)
					if in.Spec.Domain.Features.Hyperv.SyNICTimer.Direct != nil {
						SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNICTimer.Direct)
					}
				}
				if in.Spec.Domain.Features.Hyperv.Reset != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Reset)
				}
				if in.Spec.Domain.Features.Hyperv.VendorID != nil {
					SetDefaults_FeatureVendorID(in.Spec.Domain.Features.Hyperv.VendorID)
				}
				if in.Spec.Domain.Features.Hyperv.Frequencies != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Frequencies)
				}
				if in.Spec.Domain.Features.Hyperv.Reenlightenment != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Reenlightenment)
				}
				if in.Spec.Domain.Features.Hyperv.TLBFlush != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.TLBFlush)
				}
				if in.Spec.Domain.Features.Hyperv.IPI != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.IPI)
				}
				if in.Spec.Domain.Features.Hyperv.EVMCS != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.EVMCS)
				}
			}
			if in.Spec.Domain.Features.SMM != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.SMM)
			}
			if in.Spec.Domain.Features.Pvspinlock != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Pvspinlock)
			}
		}
		for i := range in.Spec.Domain.Devices.Disks {
			a := &in.Spec.Domain.Devices.Disks[i]
			SetDefaults_DiskDevice(&a.DiskDevice)
			if a.DiskDevice.Floppy != nil {
				SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
			}
			if a.DiskDevice.CDRom != nil {
				SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
			}
		}
		if in.Spec.Domain.Devices.Watchdog != nil {
			SetDefaults_Watchdog(in.Spec.Domain.Devices.Watchdog)
			if in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
				SetDefaults_I6300ESBWatchdog(in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
			}
		}
	}
}

func SetObjectDefaults_VirtualMachineInstancePresetList(in *VirtualMachineInstancePresetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachineInstancePreset(a)
	}
}

func SetObjectDefaults_VirtualMachineInstanceReplicaSet(in *VirtualMachineInstanceReplicaSet) {
	if in.Spec.Template != nil {
		if in.Spec.Template.Spec.Domain.Firmware != nil {
			SetDefaults_Firmware(in.Spec.Template.Spec.Domain.Firmware)
		}
		if in.Spec.Template.Spec.Domain.Clock != nil {
			if in.Spec.Template.Spec.Domain.Clock.Timer != nil {
				if in.Spec.Template.Spec.Domain.Clock.Timer.HPET != nil {
					SetDefaults_HPETTimer(in.Spec.Template.Spec.Domain.Clock.Timer.HPET)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.KVM != nil {
					SetDefaults_KVMTimer(in.Spec.Template.Spec.Domain.Clock.Timer.KVM)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.PIT != nil {
					SetDefaults_PITTimer(in.Spec.Template.Spec.Domain.Clock.Timer.PIT)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.RTC != nil {
					SetDefaults_RTCTimer(in.Spec.Template.Spec.Domain.Clock.Timer.RTC)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.Hyperv != nil {
					SetDefaults_HypervTimer(in.Spec.Template.Spec.Domain.Clock.Timer.Hyperv)
				}
			}
		}
		if in.Spec.Template.Spec.Domain.Features != nil {
			SetDefaults_FeatureState(&in.Spec.Template.Spec.Domain.Features.ACPI)
			if in.Spec.Template.Spec.Domain.Features.APIC != nil {
				SetDefaults_FeatureAPIC(in.Spec.Template.Spec.Domain.Features.APIC)
			}
			if in.Spec.Template.Spec.Domain.Features.Hyperv != nil {
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Relaxed != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Relaxed)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VAPIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.VAPIC)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Spinlocks != nil {
					SetDefaults_FeatureSpinlocks(in.Spec.Template.Spec.Domain.Features.Hyperv.Spinlocks)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VPIndex != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.VPIndex)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Runtime != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Runtime)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNIC)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
					SetDefaults_SyNICTimer(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer)
					if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer.Direct != nil {
						SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer.Direct)
					}
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Reset != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Reset)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VendorID != nil {
					SetDefaults_FeatureVendorID(in.Spec.Template.Spec.Domain.Features.Hyperv.VendorID)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Frequencies != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Frequencies)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Reenlightenment != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Reenlightenment)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.TLBFlush != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.TLBFlush)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.IPI != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.IPI)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.EVMCS != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.EVMCS)
				}
			}
			if in.Spec.Template.Spec.Domain.Features.SMM != nil {
				SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.SMM)
			}
			if in.Spec.Template.Spec.Domain.Features.Pvspinlock != nil {
				SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Pvspinlock)
			}
		}
		for i := range in.Spec.Template.Spec.Domain.Devices.Disks {
			a := &in.Spec.Template.Spec.Domain.Devices.Disks[i]
			SetDefaults_DiskDevice(&a.DiskDevice)
			if a.DiskDevice.Floppy != nil {
				SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
			}
			if a.DiskDevice.CDRom != nil {
				SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
			}
		}
		if in.Spec.Template.Spec.Domain.Devices.Watchdog != nil {
			SetDefaults_Watchdog(in.Spec.Template.Spec.Domain.Devices.Watchdog)
			if in.Spec.Template.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
				SetDefaults_I6300ESBWatchdog(in.Spec.Template.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
			}
		}
	}
}

func SetObjectDefaults_VirtualMachineInstanceReplicaSetList(in *VirtualMachineInstanceReplicaSetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachineInstanceReplicaSet(a)
	}
}

func SetObjectDefaults_VirtualMachineList(in *VirtualMachineList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachine(a)
	}
}

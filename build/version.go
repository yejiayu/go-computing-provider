package build

var CurrentCommit string

const BuildVersion = "0.4.4"

const UBITaskImageIntelCpu = "filswan/ubi-worker-cpu-intel:v2.0"
const UBITaskImageIntelGpu = "filswan/ubi-worker-gpu-intel:v2.0"
const UBITaskImageAmdCpu = "filswan/ubi-worker-cpu-amd:v2.0"
const UBITaskImageAmdGpu = "filswan/ubi-worker-gpu-amd:v2.0"

const PRIVATE_UBUNTU_2204 = "filswan/vm-ubuntu-22.04:v1"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}

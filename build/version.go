package build

var CurrentCommit string

const BuildVersion = "0.4.2"

const UBITaskImageVersion = "filswan/ubi-worker:v1.0"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}

package config

import (
	"os"
	"strings"
)

func isPodNameOverriden() bool {
	return os.Getenv("HTTP_HEALTHCHECK_SIDECAR_TARGET_POD_OVERRIDE") != ""
}

func TargetsSpecificContainer() bool {
	return TargettedContainer() != ""
}

func TargettedContainer() string {
	return os.Getenv(("HTTP_HEALTHCHECK_SIDECAR_TARGET_CONTAINER"))
}

func PodName() string {
	if isPodNameOverriden() {
		return os.Getenv("HTTP_HEALTHCHECK_SIDECAR_TARGET_POD_OVERRIDE")
	} else {
		return os.Getenv("HOSTNAME")
	}
}

func IsStatusOverriden() bool {
	return os.Getenv("HTTP_HEALTHCHECK_SIDECAR_STATE_OVERRIDE") != ""
}

func GetStatusOverrideState() bool {
	return strings.ToLower(os.Getenv("HTTP_HEALTHCHECK_SIDECAR_STATE_OVERRIDE")) == "true"
}

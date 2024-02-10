package health

import (
	"github.com/bryopsida/http-healthcheck-sidecar/config"
	"github.com/bryopsida/http-healthcheck-sidecar/kubernetes"
)

func IsHealthy() bool {
	if config.TargetsSpecificContainer() {
		return kubernetes.IsPodContainerHealthy(config.PodName(), config.TargettedContainer())
	} else {
		return kubernetes.IsPodHealthy(config.PodName())
	}
}

READY_STATE ?= true

image:
	docker build . -t ghcr.io/bryopsida/http-healthcheck-sidecar:local

run:
	docker run --env="HTTP_HEALTHCHECK_SIDECAR_STATE_OVERRIDE=$(READY_STATE)" -p 3000:3000 ghcr.io/bryopsida/http-healthcheck-sidecar:local

test:
	skaffold build -q > build_result.json
	skaffold deploy -a build_result.json
	skaffold verify -a build_result.json

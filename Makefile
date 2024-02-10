image:
	docker build . -t ghcr.io/bryopsida/http-healthcheck-sidecar:local

run:
	docker run -p 3000:3000 ghcr.io/bryopsida/http-healthcheck-sidecar:local

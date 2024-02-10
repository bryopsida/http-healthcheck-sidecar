# HTTP Healthcheck Sidecar

## What is this?

This solves the problem of how do you expose a HTTP based health check for a service that doesn't provide a health check in it's own protocol.
Perhaps it's a UDP service etc. It does this by using the kubernetes client to query its pod information and check the specified container status, or if one is not provided, it queries the overall container ready states of the pod.

## But doesn't kubernetes handles this for us with load balancers?

Mostly yes, however, in cases where you might be using a NodePort service and something like HAProxy or another load balancer that has no integration or awareness of kubernetes state, you need to provide it another mechanism to check the status of the service on that node, that's the gap this sidecar fills.

## How would I use this in my manifests?

You would attach this as sidecar on your pod, and create a new kubernetes NodePort service with the parameters needed to expose this to your upstream load balancer.


## Environment Variables

### HTTP_HEALTHCHECK_SIDECAR_TARGET_CONTAINER

This is an optional string value that when provided, is used to determine if the service is healthy or not. Be sure to have a good healthcheck inside that container, such as a exec probe.
If this is not provided the sidecar will check if all containers are ready or not and use that to determine a healthy (http status code 200) or non healthy (http status code 500)

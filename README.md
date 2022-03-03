## Validation
Validation of the Go code and Kubernetes manifests is handled via GitHub Actions:

 - validate-code.yaml
	 - Runs when a PR is raised that changes the Go source code. Runs `go build` against the code to ensure that everything is ok.
 - validate-manifests.yaml
	 - Runs when a PR is raised that changes any of the Kubernetes manifests. Uses `kubeval` to ensure that all resources are sane.

## Build and Publish
Docker builds/publishing is also handled via GitHub actions:

 - build-and-push.yaml
	 - Runs when a tag matching the pattern `v*` is pushed. Builds the Docker image and pushes it to Docker Hub.

## Deployment
Kubernetes manifests are deployed via ArgoCD. The pipeline definition can be found in `argocd-application.yaml`

## TODO

 - Better error handling in the Go code
 - Split out the code into a separate repository? Could end up with a bit of a race condition (where ArgoCD tries to deploy a new image before it's been built) if you update the Kubernetes manifests and push a new tag to the repo in quick succession.

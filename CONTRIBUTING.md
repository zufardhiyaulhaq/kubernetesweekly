## Development
For small things like fixing typos in documentation, you can [make edits through GitHub](https://help.github.com/articles/editing-files-in-another-user-s-repository/), which will handle forking and making a pull request (PR) for you. For anything bigger or more complex, you'll probably want to set up a development environment on your machine, a quick procedure for which is as folows:

### Setup your machine
Prerequisites:
- make
- [Go 1.13+](https://golang.org/doc/install)

- export this variable
```
export GITHUB_TOKEN="GITHUB_TOKEN"
export GITHUB_ORGANIZATION="GITHUB_USERNAME/ORGANIZATION"
export GITHUB_REPOSITORY="GITHUB_REPOSITORY"
export GITHUB_REPOSITORY_PATH="GITHUB_REPOSITORY_PATH"
export GITHUB_BRANCH="BRANCH"

export COMMUNITY="COMMUNITY_NAME"
export TAGS="TAGS"
export NAMESPACE="CRD_NAMESPACE"
export IMAGE="IMAGE_FOR_WEEKLY"
```
for example
```
export GITHUB_TOKEN="token"
export GITHUB_ORGANIZATION="zufardhiyaulhaq"
export GITHUB_REPOSITORY="community-ops"
export GITHUB_REPOSITORY_PATH="./manifest/kubernetes-community/"
export GITHUB_BRANCH="master"

export COMMUNITY="Kubernetes Indonesia Community"
export TAGS="weekly,kubernetes"
export NAMESPACE="kubernetes-community"
export IMAGE="https://raw.githubusercontent.com/cncf/artwork/master/other/cncf/horizontal/color/cncf-color.png"
```
- Lint, test, build, and run
```
make lint
make test
make build
make run
```

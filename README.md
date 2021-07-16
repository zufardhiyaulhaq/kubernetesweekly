# kubernetesweekly
[![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org)
[![Github main branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/kubernetesweekly/Main)](https://github.com/zufardhiyaulhaq/kubernetesweekly/actions/workflows/main.yml)
[![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/kubernetesweekly)](https://github.com/zufardhiyaulhaq/kubernetesweekly/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/kubernetesweekly)](https://github.com/zufardhiyaulhaq/kubernetesweekly/pulls)

Get data from kubeweekly and create Weekly CRDs based on community-operator & push to datastore

### Installing the charts
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/kubernetesweekly --name-template kubernetesweekly -f values.yaml
```

### Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| community | string | `"Cloud Native Indonesia Community"` |  |
| cronSchedule | string | `"0 8 * * 6"` |  |
| github.branch | string | `"master"` |  |
| github.organization | string | `"zufardhiyaulhaq"` |  |
| github.repository | string | `"community-ops"` |  |
| github.repository_path | string | `"./manifest/kubernetes-community/"` |  |
| github.token | string | `"your_token"` |  |
| image.name | string | `"kubernetesweekly"` |  |
| image.repository | string | `"zufardhiyaulhaq/kubernetesweekly"` |  |
| image.tag | string | `"v1.0.0"` |  |
| image_url | string | `"https://raw.githubusercontent.com/cncf/artwork/master/other/cncf/horizontal/color/cncf-color.png"` |  |
| jobHistoryLimit | int | `1` |  |
| namespace | string | `"kubernetes-community"` |  |
| tags | string | `"weekly,kubernetes"` |  |

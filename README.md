# kubernetesweekly

Get data from kubeweekly and create Weekly CRDs based on community-operator and push to git datastore

![Version: 2.1.0](https://img.shields.io/badge/Version-2.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.1.0](https://img.shields.io/badge/AppVersion-2.1.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github main branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/kubernetesweekly/Main)](https://github.com/zufardhiyaulhaq/kubernetesweekly/actions/workflows/main.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/kubernetesweekly)](https://github.com/zufardhiyaulhaq/kubernetesweekly/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/kubernetesweekly)](https://github.com/zufardhiyaulhaq/kubernetesweekly/pulls)

## Installing the Chart

To install the chart with the release name `my-kubernetesweekly` and secret `foo`:

```console
kubectl apply -f secret.yaml

helm repo add kubernetesweekly https://zufardhiyaulhaq.com/kubernetesweekly/charts/releases/
helm install my-kubernetesweekly kubernetesweekly/kubernetesweekly --values values.yaml
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cronSchedule | string | `"0 8 * * 6"` |  |
| image.repository | string | `"zufardhiyaulhaq/kubernetesweekly"` |  |
| image.tag | string | `"v2.1.0"` |  |
| secret | string | `""` |  |
| weekly.community | string | `"Cloud Native Indonesia Community"` |  |
| weekly.image_url | string | `"https://raw.githubusercontent.com/cncf/artwork/master/other/cncf/horizontal/color/cncf-color.png"` |  |
| weekly.namespace | string | `"kubernetes-community"` |  |
| weekly.tags | string | `"weekly,kubernetes"` |  |


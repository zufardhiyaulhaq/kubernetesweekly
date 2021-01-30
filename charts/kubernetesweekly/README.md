# kubernetesweekly charts
Helm chart for kubernetesweeklys

### Installing the charts
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/kubernetesweekly --name-template kubernetesweekly -f values.yaml
```

### Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| community | string | `"Kubernetes Indonesia Community"` |  |
| cronSchedule | string | `"0 8 * * 6"` |  |
| github.branch | string | `"master"` |  |
| github.organization | string | `"zufardhiyaulhaq"` |  |
| github.repository | string | `"community-ops"` |  |
| github.repository_path | string | `"./manifest/kubernetes-community/"` |  |
| github.token | string | `"your_token"` |  |
| image.name | string | `"kubernetesweekly"` |  |
| image.repository | string | `"zufardhiyaulhaq/kubernetesweekly"` |  |
| image.tag | string | `"0.0.1"` |  |
| image_url | string | `"https://raw.githubusercontent.com/cncf/artwork/master/other/cncf/horizontal/color/cncf-color.png"` |  |
| jobHistoryLimit | int | `1` |  |
| namespace | string | `"kubernetes-community"` |  |
| tags | string | `"weekly,kubernetes"` |  |

check & modify values.yaml for details

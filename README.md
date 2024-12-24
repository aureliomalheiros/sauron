# Study prometheus

![sauron-observability](img/sauron.png)

This repository provides a study environment for Prometheus monitoring, transitioning from Vagrant-based setups to containerized environments using Docker and Kubernetes. It includes configurations and manifests to explore Prometheus features in modern cloud-native infrastructures.

## Prerequisites

![Vagrant](https://img.shields.io/badge/-vagrant-3371e3?style=for-the-badge&logo=vagrant&logoColor=white)

### Project Structure:

```bash
prometheus
├── README.md
└── vagrant
    ├── config
    │   └── prometheus.yml
    ├── README.md
    ├── scripts
    │   └── startup-prometheus.sh
    └── Vagrantfile
```

- [Vagrant](vagrant/README.md)
  
#### Reference

- [Prometheus](https://prometheus.io/docs/introduction/overview/)

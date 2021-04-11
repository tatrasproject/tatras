# ![Test Image 1](img/tatras-small.svg) Tatras


Table of Contents
-----------------

* [Summary](#summary)
* [Installation](#installation)
* [Usage](#usage)
  * [Local](#local)
  * [Docker](#docker)
  * [Kubernetes](#kubernetes)
* [Contributor](#contributor)


## Summary

Tatras is a cloud native CD tool built to deploy single tenant applications from a single portal. It relies on Kubernetes and ArgoCD. ArgoCD and Flux can be used to deploy a Helm chart to one or many clusters using a single values.yaml. Tatras allows you to quickly deploy multiple instances of the same Helm chart, passing in custom values.yaml files as needed. 

Tatras is recommended for cases when multiple parties require access to a single tenant application. This could be for deploying a PostgreSQL server for multiple teams, or by a SaaS company deploying a product for each customer. 


## Installation

Pull down repo, set up environment, and deploy
```bash
echo TBD
```

## Usage

### Local

```bash
./dev/build.sh
```

## Contributor



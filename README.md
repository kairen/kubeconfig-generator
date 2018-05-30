[![Docker Build Status](https://img.shields.io/docker/build/inwinstack/kubeconfig-generator.svg)](https://hub.docker.com/r/inwinstack/kubeconfig-generator/)
# Kubeconfig Generator
A server to generate kubeconfig of the user by LDAP query token.

# Quick Start
In this first, modified the `deploy/ldap-generator-server-dp.yml` file to match our LDAP and Kubernetes API server endpoint:
```yml
# container args
spec:
  template:
    spec:
      containers:
      - name: kubeconfig-generator-server
        image: inwinstack/kubeconfig-generator:v0.1.0
        args:
        - serve
        - --kube-apiserver-endpoint=https://192.16.35.11:6443
        - --ldap-address=192.16.35.20:389
```

And then apply to Kubernetes cluster:
```sh
$ kubectl apply -f deploy/
```
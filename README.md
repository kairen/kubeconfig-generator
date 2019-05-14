[![Build Status](https://travis-ci.org/kubedev/kubeconfig-generator.svg?branch=master)](https://travis-ci.org/kubedev/kubeconfig-generator) [![Docker Pulls](https://img.shields.io/docker/pulls/kubedev/kubeconfig-generator.svg)](https://hub.docker.com/r/kubedev/kubeconfig-generator/)

# Kubeconfig Generator
Kubeconfig Generator is a tool to generate kubeconfig for the auth webhook.

Now support webhook as below:
* [x] LDAP Webhook.
* [ ] Keystone Webhook.

## Building from Source
Clone into your go path under `$GOPATH/src`:
```sh
$ git clone https://github.com/kubedev/kubeconfig-generator.git $GOPATH/src/github.com/kubedev/kubeconfig-generator.git
$ cd $GOPATH/src/github.com/kubedev/kubeconfig-generator.git
$ make
```

## Quick Start
In this first, modified the `deploy/deployment.yml` file to match our LDAP and Kubernetes API server endpoint:
```yml
# container args
spec:
  template:
    spec:
      containers:
      - name: kubeconfig-generator
        image: kubedev/kubeconfig-generator:v0.1.0
        args:
        - serve
        - --kube-apiserver-endpoint=https://192.16.35.11:6443
        - --ldap-address=192.16.35.20:389
        - --ldap-dc=dc=k8s,dc=com
        - --user-search-base=ou=People,dc=k8s,dc=com
        - --user-name-attribute=givenName
        - --user-token-arttribute=kubernetesToken
```

And then apply to Kubernetes cluster:
```sh
$ kubectl apply -f deploy/
```

To generate the config using kgctl:
```sh
$ kgctl ldap --url http://172.22.132.40:32400 \
    --dn "uid=user1,ou=People,dc=k8s,dc=com" \
    --password "user1" \
    -o test.conf
# output
Generate the Kubernetes config to `test.conf`.

$ export KUBECONFIG=test.conf
$ kubectl -n user1 get po
```
> Or access `Web-based UI`.

![web-ui](snapshots/home.png)

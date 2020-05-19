# Kubernetes KMS Plugin for Self-Defending KMS

This repository contains an implementation of [Kubernetes KMS Plugin] for
[Self-Defending KMS].

## General Usage

At a high level you need to follow these steps to encrypt Kubernetes secrets
with key(s) stored in Self-Defending KMS:

1. Setup encryption key and app in Self-Defending KMS
2. Create a configuration file for the KMS plugin
3. Deploy the KMS plugin on the Kubernetes master nodes
4. Setup Kubernetes encryption configuration

### 1. Setup encryption key and app in Self-Defending KMS

Create an AES key in Self-Defending KMS. This key will be used by the KMS
plugin to encrypt/decrypt Kubernetes secrets.

Create an app with access to the encryption key in Self-Defending KMS. Make
sure the app is allowed to perform encrypt and decrypt operations on the
encryption key. Use the API Key authentication method for this app.

### 2. Create a configuration file for the KMS plugin

The KMS plugin needs the following configuration values:

- Self-Defending KMS endpoint URL, e.g. https://sdkms.fortanix.com
- API Key
- Name or UUID of the encryption key
- Unix domain socket path for its gRPC endpoint

This plugin allows for rotation of the encryption key in Self-Defending KMS
without the need to change Kubernetes encryption configuration. You would need
to specify the encryption key **by name** to enable this feature. If you don't
want to use this feature, then you should specify the encryption key by its
UUID instead.

The KMS plugin expects its configuration in JSON format. Here is an example `k8s-sdkms-plugin.json`:

```json
{
  "sdkms_endpoint": "https://sdkms.fortanix.com",
  "api_key": "N2Q3MGRiZWMtMGMyMC00ZTRjLTk5YjktMmFkYz...",
  "key_name": "Kubernetes Secret Encryption Key",
  "socket_file": "/var/run/kms-plugin/socket"
}
```

if you don't need to rotate the encryption key, replace the entry
`key_name` with `key_id` in the config file. Here is an example:

```js
{
  // ...
  "key_id": "4b3d13a8-2c3a-47dc-8779-311dad6843a2",
  // ...
}
```

### 3. Deploy the KMS plugin on the Kubernetes master nodes

The KMS plugin needs to run on all master nodes to be able to communicate with
`kube-apiserver` through a Unix domain socket. You can either use [Static Pods]
or use Docker directly. Regardless of the method you use, the static pod or the
Docker container needs:

- Plugin configuration file, e.g. `/etc/fortanix/k8s-sdkms-plugin.json`
- A shared directory to store the Unix domain socket, e.g. `/var/run/kms-plugin`

#### Delpoy the plugin using Docker

A prebuilt Docker image containing this KMS plugin is available on Docker Hub
at `fortanix/k8s-sdkms-plugin:$TAG` where `TAG` refers to a published git tag
of this repo.

Pull the Docker image on Kubernetes master nodes:

```
$ docker pull "fortanix/k8s-sdkms-plugin:0.1.0"
```

Then start the container:

```
$ docker run -d -v /var/run/kms-plugin:/var/run/kms-plugin \
    -v /etc/fortanix/k8s-sdkms-plugin.json:/etc/fortanix/k8s-sdkms-plugin.json \
    --name sdkms-plugin \
    fortanix/k8s-sdkms-plugin:0.1.0
```

#### Deploy the plugin using Static Pods

Use the following Pod specification to create the static pods:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: k8s-sdkms-plugin
spec:
  containers:
  - name: k8s-sdkms-plugin
    image: fortanix/k8s-sdkms-plugin:0.1.0
    volumeMounts:
    - name: socket-dir
      mountPath: /var/run/kms-plugin
    - name: config
      mountPath: /etc/fortanix/k8s-sdkms-plugin.json
  volumes:
  - name: socket-dir
    hostPath:
      path: /var/run/kms-plugin
      type: DirectoryOrCreate
  - name: config
    hostPath:
      path: /etc/fortanix/k8s-sdkms-plugin.json
      type: File
```

### 4. Setup Kubernetes encryption configuration

Update the `kube-apiserver`'s encryption configuration to enable the KMS plugin
for encrypting/decrypting Kubernetes secrets. Here is an example:

```yaml
apiVersion: apiserver.config.k8s.io/v1
kind: EncryptionConfiguration
resources:
- resources:
  - secrets
  providers:
  - kms:
      name: sdkms-plugin
      endpoint: unix:///var/run/kms-plugin/socket
      cachesize: 100
      timeout: 3s
  - identity: {}
```

Note that `kube-apiserver` needs to be able to access the shared socket file
through `/var/run/kms-plugin/socket`. The details of how to ensure the socket
is accessible is highly specific to the type of Kubernetes deployment you are
using.

Before changing the `kube-apiserver` encryption configuration make sure you
understand the various provider types and how this configuration is used by
Kubernetes to encrypt and decrypt secrets. Please refer to
[Encrypting Secret Data at Rest] for more information.

If you are using [Rancher] to deploy your Kubernetes cluster, you can follow
the instructions in [rancher-guide](./rancher-guide.md) to apply the encryption
configuration.

# Contributing

We gratefully accept bug reports and contributions from the community.
By participating in this community, you agree to abide by [Code of Conduct](./CODE_OF_CONDUCT.md).
All contributions are covered under the Developer's Certificate of Origin (DCO).

## Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
have the right to submit it under the open source license
indicated in the file; or

(b) The contribution is based upon previous work that, to the best
of my knowledge, is covered under an appropriate open source
license and I have the right under that license to submit that
work with modifications, whether created in whole or in part
by me, under the same open source license (unless I am
permitted to submit under a different license), as indicated
in the file; or

(c) The contribution was provided directly to me by some other
person who certified (a), (b) or (c) and I have not modified
it.

(d) I understand and agree that this project and the contribution
are public and that a record of the contribution (including all
personal information I submit with it, including my sign-off) is
maintained indefinitely and may be redistributed consistent with
this project or the open source license(s) involved.

# License

This project is primarily distributed under the terms of the Mozilla Public License (MPL) 2.0, see [LICENSE](./LICENSE) for details.

[Kubernetes KMS Plugin]: https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/
[Self-Defending KMS]: https://fortanix.com/products/sdkms/
[Static Pods]: https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/
[Encrypting Secret Data at Rest]: https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/
[Rancher]: https://rancher.com/

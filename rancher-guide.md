# Deploying on Rancher

The instructions here are complimentary to those in the [README](./README.md)
and only apply to Kubernetes clusters deployed using [Rancher].

Please start by reading the [README](./README.md) first.

## KMS Plugin Deployment

Use Docker to deploy the KMS plugin as there is no easy way to define Static
Pods in Rancher.

Make sure to verify that the KMS plugin is running on all control plane nodes
and check its logs to ensure it has initialized correctly:

```
$ docker logs sdkms-plugin
```

## Kubernetes Setup

Rancher makes it easier to setup encryption configuration for Kubernetes
through its cluster settings. [Encrypting Secret Data at Rest] provides details
of how that works.

In order to change the encryption configuration, follow the steps in
[Edit Cluster Settings] to edit the cluster settings in YAML format.

### Important note

Before you update your cluster settings to use the KMS provider, make sure
that encryption keys are not being managed by Rancher.

If your cluster configuration has the following settings:

```yaml
  services:
    kube-api:
      secrets_encryption_config:
        enabled: true
```

Then **you should [disable encryption] first.**

Otherwise your Kubernetes secrets will not be decryptable and you might lose
your rancher-managed encryption keys.

### Enabling the KMS provider in cluster settings

Change your cluster settings to look like the following:

```yaml
  services:
    # ...
    kube-api:
      # ...
      extra_binds:
        - "/var/run/kms-plugin/:/var/run/kms-plugin/"
      secrets_encryption_config:
        enabled: true
        custom_config:
          apiVersion: apiserver.config.k8s.io/v1
          kind: EncryptionConfiguration
          resources:
            - resources:
              - secrets
              providers:
              - kms:
                  name: sdkms-encryption-provider
                  endpoint: unix:///var/run/kms-plugin/socket
                  cachesize: 100
                  timeout: 3s
              - identity: {}
```

The `extra_bind` will make sure that `kube-apiserver` has access to the Unix
domain socket created by the KMS plugin.

Once you apply the changes, Rancher will deploy the proper encryption
configuration settings to all control plane nodes and restart all
`kube-apiserver`s. Then it will rewrite all secrets so they get encrypted
with the new settings.

## Learn More

- How Kubernetes uses the encryption providers: [Kubernetes docs]
- How Rancher applies `secrets_encryption_config`: [Encrypting Secret Data at Rest]



[Rancher]: https://rancher.com/
[Encrypting Secret Data at Rest]: https://rancher.com/docs/rke/latest/en/config-options/secrets-encryption/
[Edit Cluster Settings]: https://rancher.com/docs/rancher/v2.x/en/cluster-admin/editing-clusters/
[disable encryption]: https://rancher.com/docs/rke/latest/en/config-options/secrets-encryption/#disable-encryption
[Kubernetes docs]: https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/

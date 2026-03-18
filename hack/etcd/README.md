# Proxima etcd

This directory contains a dedicated 3-node etcd cluster for remote debugging in Kubernetes.

The debug profile uses `emptyDir` instead of network storage because the cluster default `shared-nfs` StorageClass is not suitable for etcd WAL files.

## Deploy

```bash
kubectl apply -k hack/etcd/kustomize/overlays/develop
```

## Verify

```bash
kubectl -n proxima-etcd get pods
kubectl -n proxima-etcd get pvc
kubectl -n proxima-etcd exec proxima-etcd-0 -- \
  /usr/local/bin/etcdctl --endpoints=http://127.0.0.1:2379 endpoint status -w table
```

## Local tunnel for debugging

Run this on the cluster master:

```bash
kubectl -n proxima-etcd port-forward svc/proxima-etcd-client 2379:2379
```

Then from your local machine:

```bash
ssh -p 60011 -L 2379:127.0.0.1:2379 root@nps.nobugs.net.cn
```

After that, local clients can connect to `127.0.0.1:2379`.
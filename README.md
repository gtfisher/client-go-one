go example that talks to Kubernetes to list pods in kube-system namespace using client-go

Should work with

If you have minikube running and kubctl configured to talk to it, it should run with

```
go run main.go
```

And get output including
```
etcd-minikube
 kube-apiserver-minikube
 kube-controller-manager-minikube
 kube-proxy-6gzqz
 kube-scheduler-minikube
 storage-provisioner
 ```
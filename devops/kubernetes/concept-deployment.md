## Deployment

通过声明的方式在 Kubernetes 中部署和更新应用

deployment → replicaSet → pod

升级方式

- 更新
- 回滚
- 终止

查看

`kubectl get deployments`

删除

`kubectl delete deployments {name}`

参考

[https://linuxhandbook.com/kubectl-delete-deployment/](https://linuxhandbook.com/kubectl-delete-deployment/)


### Q&A
1. What is the difference between a pod and a deployment?
- [https://stackoverflow.com/questions/41325087/what-is-the-difference-between-a-pod-and-a-deployment#:~:text=Pods - runs one or more,right number of them exist](https://stackoverflow.com/questions/41325087/what-is-the-difference-between-a-pod-and-a-deployment#:~:text=Pods%20-%20runs%20one%20or%20more,right%20number%20of%20them%20exist)

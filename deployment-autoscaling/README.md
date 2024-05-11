### deployment autoscaling

``` shell
# create a deployment
kubectl create -f nginx-deployment.yaml


# get pods
kubectl get pod

# NAME                                READY   STATUS    RESTARTS   AGE2
# nginx-deployment-77d8468669-477nl   1/1     Running   0          67s
# nginx-deployment-77d8468669-76scr   1/1     Running   0          67s

# delete pod
# e.g. kubectl delete pod nginx-deployment-77d8468669-477nl   
kubectl delete pod <pod-name>


# then check the pod again
kubectl get pod


# 아래 결과에서 nginx-deployment-77d8468669-hxxn9가 새로 생성된 모습을 확인할 수 있다.
# NAME                                READY   STATUS    RESTARTS   AGE2
# nginx-deployment-77d8468669-76scr   1/1     Running   0          113s
# nginx-deployment-77d8468669-hxxn9   1/1     Running   0          6s

```
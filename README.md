# eks-workshop-x-ray-sample-back

The example [AWS X-Ray](https://aws.amazon.com/xray/) instrumented back-end service in the [EKS Workshop](https://eksworkshop.com/)

**Command reference**

Deploy
```
kubectl apply -f x-ray-sample-back-k8s.yml
```

Delete
```
kubectl delete deployment x-ray-sample-back-k8s

kubectl delete service x-ray-sample-back-k8s
```


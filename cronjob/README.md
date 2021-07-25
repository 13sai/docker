## 创建cronJob

kubectl apply -f cronjob/cronJob.yaml 

## 查看

kubectl get cronjob

## 执行后等待

kubectl get jobs --watch


## 查看pods

kubectl get pods

## 打印日志

kubectl logs hello-1624860600-w7vkc
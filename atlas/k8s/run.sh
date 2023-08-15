# k8s ui
kubectl -n kubernetes-dashboard create token admin-user

kubectl proxy


## docker registry
docker run -d -p 5000:5000 --restart=always --name registry registry

docker push 192.168.2.11:5000/atlas-admin-ui:latest

docker image push  192.168.2.11:5000/atlas-admin-ui:latest

#k8s hitting service
kubectl port-forward atlas 9811:9811


## Delete the NuoDB Database, NuoDB Insights, and YCSB sample application

Run the following commands,
```
# delete the nuodb CE nuodb.lic file
kubectl delete -n $OPERATOR_NAMESPACE configmap nuodb-lic-configmap

# delete cluster-admin permissions from the nuodb-operator service account
kubectl delete -f nuodb-operator/deploy/cluster-op-admin.yaml

kubectl delete pod/insights-client -n $OPERATOR_NAMESPACE

kubectl delete crd elasticsearches.elasticsearch.k8s.elastic.co &
kubectl delete crd kibanas.kibana.k8s.elastic.co &
kubectl delete crd grafanadatasources.integreatly.org &
kubectl delete crd grafanas.integreatly.org &

kubectl delete statefulset insights-server-release-logstash 
kubectl delete service insights-server-release-logstash 
kubectl delete deployment grafana-operator

# delete statement for a database named `db`. Adjust if your db is named differently.
kubectl delete nuodb nuodb-db

kubectl delete nuodbycsbwl nuodbycsbwl
kubectl delete nuodbinsightsserver insightsserver &

kubectl delete -n $OPERATOR_NAMESPACE pvc --all
```
Then run,

`kubectl edit nuodbinsightsserver`

and remove the below line using the vi editor and save the file.

` - ECK.finalizers.nuodbinsightsserver.nuodb.com`

If the local-disk storage class was used, then delete the NuoDB Storage Manager(SM) disk storage and storage class
```
ssh -i ~/Documents/cluster.pem $JUMP_HOST
ssh -i ~/.ssh/cluster.pem core@ip-n-n-n-n.ec2.internal  'rm -rf /mnt/local-storage/disk0/*'

kubectl delete -f nuodb-operator/build/etc/charts/nuodb-helm/local-disk-class.yaml
```
**NOTE:** If planning to restart another NuoDB database after deleting one, please wait a few minutes to ensure all Kubernetes objects have been terminated and removed. You can check by running `kubectl get all` to ensure only your NuoDB Operator is running before initialing a new database deployment.


## Delete the NuoDB Operator

### Red Hat OpenShift v4.x
From the OpenShift WebUI, locate the OperatorHub under the Catalog left-bar selection. Select the NuoDB Operator and click the `Uninstall` button.

### Google Public Cloud GKE Kubernetes
From the Google Cloud Platform dashboard, locate the Applications selection on the left tool bar in the Google Kubernetes management section. Select the NuoDB Operator application and click the `Delete` button.

### Amazon EKS, Azure AKS, Google Anthos GKE, or Open Source K8S

Run the following commands,
```
kubectl delete crd nuodbinsightsservers.nuodb.com
kubectl delete crd nuodbycsbwls.nuodb.com
kubectl delete crd nuodbs.nuodb.com

kubectl delete -f nuodb-csv.yaml

kubectl delete role grafana-operator
kubectl delete role nuodb-operator
kubectl delete rolebinding grafana-operator
kubectl delete rolebinding nuodb-operator
kubectl delete serviceaccount grafana-operator
kubectl delete serviceaccount nuodb-operator

kubectl delete -f nuodb-operator/deploy/cluster_role_binding.yaml
kubectl delete -n $OPERATOR_NAMESPACE -f nuodb-operator/deploy/cluster_role.yaml
kubectl delete -n $OPERATOR_NAMESPACE -f nuodb-operator/deploy/operatorGroup.yaml
```

_**NOTE:** The instruction prerequisite steps on this page use the Kubernetes `kubectl` command (for portability reasons across Kubernetes environments). For Red Hat OpenShift environments, the `kubectl` command is an alias that points to the OpenShift `oc` client program._

# NuoDB Operator Prerequisites

## 1. Provision a Kubernetes cluster

Create a Kubernetes cluster and connect to the cluster. 
In our verification tests, we regularly verify the samples workloads outlined on this page using the following configuration:
* 4 worker nodes, each with with 4 CPUs and 16 GB of RAM (e.g. AWS m5.xlarge instance type)
* 5 GB disk for Admin pods
* 20 GB disk for Storage Manager(SM) pods

Please use this as a guideline for a minimal configuration when you create your cluster. To run larger SQL workloads using the included YCSB sample application, adjust node CPU and Memory upwards as required. To determine resources used, monitor your NuoDB database process resource consumption using the NuoDB Insights visual monitoring tool. 

## 2. Create environment variables
```
export OPERATOR_NAMESPACE=nuodb
export NUODB_OPERATOR_VERSION=2.0.3           --confirm you set the correction NuoDB Operator version here.
```

## 3. Create the "nuodb" project/namespace

`kubectl create namespace $OPERATOR_NAMESPACE`

## 4. Download the latest NuoDB Operator release zip file from the [Releases](https://github.com/nuodb/nuodb-operator/releases) tab and unzip 
For example, in your home or working directory, run:

`unzip nuodb-operator-<latest version>.zip`

# NuoDB Database Prerequisites

## 5. Cluster Node Labeling
Label the cluster nodes you want to use to run NuoDB pods.

 `kubectl  label node <node name> nuodb.com/zone=nuodb`

_**NOTE:** The label value, in this example "nuodb", can be any value._

Next, label one of these nodes as your storage node that will host the NuoDB Storage Manager (SM) pod. If using Local storage, ensure there is sufficient disk space on this node. To create this label run:

`kubectl  label node <yourStorageNodeDNSName> nuodb.com/node-type=storage`

Once your cluster nodes are labeled for NuoDB use, run the following `kubectl get nodes` command to confirm nodes are labeled properly. The display output should look similar to the below
```
kubectl get nodes -l nuodb.com/zone -L nuodb.com/zone,nuodb.com/node-type
NAME                           STATUS   ROLES    AGE   VERSION             ZONE    NODE-TYPE
ip-10-0-141-113.ec2.internal   Ready    worker   15d   v1.15.11            nuodb   storage
ip-10-0-152-147.ec2.internal   Ready    worker   15d   v1.15.11            nuodb   
ip-10-0-162-73.ec2.internal    Ready    worker   15d   v1.15.11            nuodb   
ip-10-0-184-233.ec2.internal   Ready    worker   15d   v1.15.11            nuodb   
ip-10-0-206-8.ec2.internal     Ready    worker   15d   v1.15.11            nuodb 
```

## 6. Optionally Use Cluster Node Local Storage

NuoDB supports cloud platform storage (e.g. AWS EBS, GCP PD, Azure Disk, etc), 3rd-party CSI storage (e.g. Portworx, OpenEBS, Linbit, etc.), and the use of local storage via Hostpath. If planning to use any of these storage types, configure them prior to deploying your database. 

### To Setup Cluster Node local storage (using HOSTPATH): 
Configure the local storage permissions on each cluster node to enable hosting storage for either the NuoDB Admin or the Storage Manager (SM) pods.
**NOTE:** When using the local disk storage option only 1 Admin pod is supported.

```
sudo mkdir -p /mnt/local-storage/disk0
sudo chmod -R 777 /mnt/local-storage/
sudo chcon -R unconfined_u:object_r:svirt_sandbox_file_t:s0 /mnt/local-storage
sudo chown -R root:root /mnt/local-storage
```
Create the Kubernetes storage class "local-disk" and persistent volume

 `kubectl create -f nuodb-operator/build/etc/charts/nuodb-helm/local-disk-class.yaml`
 
## 7. Apply a NuoDB License File

Each time a NuoDB Admin pod starts it will load a Kubernetes configmap that contains the current NuoDB license level information and places its contents in the /etc/nuodb/nuodb.lic file. When a request is made to either start a NuoDB Transaction Engine (TE) or Storage Manager (SM) process, the NuoDB Admin will check the license file contents to ensure the process request is authorized.

### Obtain your Community Edition (CE) License File
If you have not already obtained your NuoDB Community Edition product evaluation license file, please visit the [Get NuoDB Community Edition](http://nuodb.com/get-community-edition) webpage to promptly receive your license file by email.

### Apply your (CE) license file

`kubectl create configmap nuodb-lic-configmap -n $OPERATOR_NAMESPACE --from-file=nuodb.lic`

**NOTE**: Applying a NuoDB Enterprise Edition (EE) is done in the same manner. Obtain your NuoDB EE license file from your NuoDB Sales or Support representative and copy the file to `nuodb.lic`

### Upgrade your license to an Enterprise Edition (EE) license on an already running database
```
kubectl delete configmap nuodb-lic-configmap -n $OPERATOR_NAMESPACE
kubectl create configmap nuodb-lic-configmap -n $OPERATOR_NAMESPACE --from-file=nuodb.lic
```
Then, delete a NuoDB Admin pod, and once the Admin pod has been restarted, connect to the new pod and run,

`nuocmd set license --license-file /etc/nuodb/nuodb.lic`

This command will propagate the new NuoDB EE license throughout the Admin tier (remaining pods).  

**NOTE:** The file name specified in the above commands must be nuodb.lic

To check the effective NuoDB license and confirm license level, run

`nuocmd --show-json get effective-license`

 
## 8. If Using Red Hat OpenShift 

### To permit the pulling of the NuoDB database and operator container images, create the Kubernetes image pull secret

This secret will be used to pull the NuoDB Operator and NuoDB container images from the Red Hat Container
Catalog (RHCC). Enter your Red Hat login credentials for the --docker-username and --docker-password values.

```
kubectl  create secret docker-registry pull-secret \
   --docker-username='yourUserName' --docker-password='yourPassword' \
   --docker-email='yourEmailAddr'  --docker-server='registry.connect.redhat.com'
 ```
**NOTE:** If using Quay.io (or other supported public repo) to pull the NuoDB container images, a Kubernetes secret is not required because the NuoDB repository is public. For example, to pull the image from quay.io, run at the command prompt, docker pull quay.io/nuodb/nuodb-operator.

### Disable Linux Transparent Huge Pages (THP). Run the following required command to create a security context constraint which will allow the Operator to disable THP during Operator deployment.
```
kubectl create -n $OPERATOR_NAMESPACE -f nuodb-operator/deploy/thp-scc.yaml
```
### Run the following oc admin policy commands,
```
oc adm policy add-scc-to-user privileged system:serviceaccount:$OPERATOR_NAMESPACE:nuodb-operator
oc adm policy add-scc-to-user privileged system:serviceaccount:elastic-system:elastic-operator
oc adm policy add-scc-to-user privileged system:serviceaccount:$OPERATOR_NAMESPACE:insights-server-release-logstash
```
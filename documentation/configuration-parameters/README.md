## Admin Domain Parameters


**adminCount** - Number of admin service pods. Requires 1 cluster node available for each Admin Service

&ensp; `adminCount: 3`


**adminStorageSize** - Admin service log volume size (GB)

&ensp; `adminStorageSize: 5G`


**adminStorageClass** - Admin persistent storage class name

&ensp; `adminStorageClass: glusterfs-storage`


**apiServer** - Load balancer service URL. hostname:port (or LB address) for nuocmd and nuodocker process to connect to.

&ensp; `apiServer: https://domain:8888`

**insightsEnabled** - (Deprecated) This variable was previously used to enable hosted Insights which provided an option to send and store performance data to a secure AWS cloud portal that was only accessible by using your Subscription ID. This option is no longer available, and therefore this variable should remain false.

&ensp; `insightsEnabled: false`


## Database Parameters

**storageMode** - Run NuoDB using persistent durable storage "persistent" or volatile storage "ephemeral". Must be set to one of those two values.

&ensp; `storageMode: persistent`


**dbName** - NuoDB Database name. must consist of lowercase alphanumeric characters '[a-z0-9]+' 

&ensp; `dbName: test`


**dbUser** - Name of Database user

&ensp; `dbUser: dba`


**dbPassword** - Database password

&ensp; `dbPassword: secret`


**smMemory** - SM memory (in GB)

&ensp; `smMemory: 2Gi`


**smCpu** - SM CPU cores to request

&ensp; `smCpu: "1"`


**smStorageSize** - Storage manager (SM) volume size (GB)

&ensp; `smStorageSize: 20G`


**smStorageClass** - SM persistent storage class name

&ensp; `smStorageClass: local-disk`


**engineOptions** - Additional "nuodb" engine options Format: …

&ensp; `engineOptions: ""`


**teCount** - Number of transaction engines (TE) nodes. Limit is 3 in CE version of NuoDB

&ensp; `teCount: 1`


**teMemory** - TE memory (in GB)

&ensp; `teMemory: 2Gi`


**teCpu** - TE CPU cores to request

&ensp; `teCpu: "1"`


**container** - NuoDB fully qualified image name (FQIN) for the Docker image to use

Below are examples that pull the NuoDB container image from Red Hat (RHCC), Google Cloud Platform Marketplace, AWS Marketplace, and DockerHub.

```
container: registry.connect.redhat.com/nuodb/nuodb-ce:latest
container: marketplace.gcr.io/nuodb/nuodb:latest
container: 117940112483.dkr.ecr.us-east-1.amazonaws.com/d893f8e5-fe12-4e43-b792-8cb98ffc11c0/cg-1228790192/docker.io/nuodb/nuodb-ce:2.0.3-2-latest
```


## NuoDB Insights Parameters

**elasticVersion** - Version of ElasticSearch

&ensp; `elasticVersion: 7.3.0`

**elasticNodeCount** - Number of nodes in the ElasticSearch Cluster

&ensp; `elasticNodeCount: 1`

**kibanaVersion** - Version of Kibana

&ensp; `kibanaVersion: 7.3.0`

**kibanaNodeCount** - Version of Kibana

&ensp; `kibanaNodeCount: 1`

**storageClass** - Kubernetes Persistent Storage Class

&ensp; `storageClass: ""`



## YCSB Sample Application Workload Parameters


**dbName** - The NuoDB database to run the ycsb application

&ensp; `dbName: db`


**ycsbWorkloadCount** - The number of ycsb pods to run

&ensp; `ycsbWorkloadCount: 0`
  

**ycsbLoadName** - YCSB workload pod name

&ensp; `ycsbLoadName: ycsb-load`


**ycsbWorkload** - Sample SQL activity workload. Valid values are a-f. Each letter determines a different mix of read and update workload percentage generated. a= 50/50, b=95/5, c=100 read. Refer to YCSB documentation for more detail.

&ensp; `ycsbWorkload: b`


**ycsbLbPolicy** - YCSB load-balancer policy. Name of an existing load-balancer policy, that has already been created using the 'nuocmd set load-balancer' command.

&ensp; `ycsbLbPolicy: ""`


**ycsbNoOfProcesses** - Number of YCSB processes. Number of concurrent YCSB processes that will be started in each YCSB pod. Each YCSB process makes a connection to the Database.

&ensp; `ycsbNoOfProcesses: 2`


**ycsbNoOfRows** - YCSB number of rows inserted into the ycsb table

&ensp; `ycsbNoOfRows: 10000`


**ycsbNoOfIterations** - YCSB number of iterations. After the last iteration finishes the benchmark will stop. A value of 0 will run the benchmark continuously.

&ensp; `ycsbNoOfIterations: 0`


**ycsbOpsPerIteration** - Number of YCSB SQL operations to perform in each iteration. This value controls the number of SQL operations performed in each benchmark iteration. Increasing this value increases the run-time of each iteration, and also reduces the frequency at which new connections are made during the sample workload run period. 

&ensp; `ycsbOpsPerIteration: 10000`


**ycsbMaxDelay** - YCSB maximum workload delay in milliseconds (Default is 4 minutes)

&ensp; `ycsbMaxDelay: 240000`


**ycsbDbSchema** - YCSB Database schema. Default schema to use to resolve tables, views, etc.

&ensp; `ycsbDbSchema: User1`


**ycsbContainer** - YCSB fully qualified image name (FQIN) for the ycsb docker image to use. See examples below pulling the image from dockerhub and the AWS Marketplace.

```
ycsbContainer: nuodb/ycsb:latest
ycsbContainer: 117940112483.dkr.ecr.us-east-1.amazonaws.com/d893f8e5-fe12-4e43-b792-8cb98ffc11c0/cg-1228790192/docker.io/nuodb/ycsb:2.0.3-2-latest
```

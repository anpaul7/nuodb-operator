// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v2alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.Nuodb":                     schema_pkg_apis_nuodb_v2alpha1_Nuodb(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServer":       schema_pkg_apis_nuodb_v2alpha1_NuodbInsightsServer(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServerSpec":   schema_pkg_apis_nuodb_v2alpha1_NuodbInsightsServerSpec(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServerStatus": schema_pkg_apis_nuodb_v2alpha1_NuodbInsightsServerStatus(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbSpec":                 schema_pkg_apis_nuodb_v2alpha1_NuodbSpec(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbStatus":               schema_pkg_apis_nuodb_v2alpha1_NuodbStatus(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWl":               schema_pkg_apis_nuodb_v2alpha1_NuodbYcsbWl(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWlSpec":           schema_pkg_apis_nuodb_v2alpha1_NuodbYcsbWlSpec(ref),
		"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWlStatus":         schema_pkg_apis_nuodb_v2alpha1_NuodbYcsbWlStatus(ref),
	}
}

func schema_pkg_apis_nuodb_v2alpha1_Nuodb(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Nuodb is the Schema for the nuodbs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbSpec", "github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbInsightsServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbInsightsServer is the Schema for the nuodbinsightsservers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServerSpec", "github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbInsightsServerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbInsightsServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbInsightsServerSpec defines the desired state of NuodbInsightsServer",
				Properties: map[string]spec.Schema{
					"elasticVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "ElasticSearch Version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"elasticNodeCount": {
						SchemaProps: spec.SchemaProps{
							Description: "ElasticSearch Node Count",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"kibanaVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "Kibana Version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kibanaNodeCount": {
						SchemaProps: spec.SchemaProps{
							Description: "Kibana Node Count",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"storageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "Persistent Storage Class for internal components.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"elasticVersion", "elasticNodeCount", "kibanaVersion", "kibanaNodeCount", "storageClass"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbInsightsServerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbInsightsServerStatus defines the observed state of NuodbInsightsServer",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbSpec defines the desired state of Nuodb",
				Properties: map[string]spec.Schema{
					"adminCount": {
						SchemaProps: spec.SchemaProps{
							Description: "adminCount Number of admin service pods. Requires 1 server available for each Admin Service example: adminCount: 1",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"adminStorageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "adminStorageClass Admin persistent storage class name example: adminStorageClass: glusterfs-storage",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"adminStorageSize": {
						SchemaProps: spec.SchemaProps{
							Description: "adminStorageSize Admin service log volume size (GB) example: adminStorageSize: 5G",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiServer": {
						SchemaProps: spec.SchemaProps{
							Description: "apiServer Load balancer service URL.  hostname:port (or LB address) for nuoadmin process to connect to. Example: apiServer: https://domain:8888",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"container": {
						SchemaProps: spec.SchemaProps{
							Description: "container NuoDB fully qualified image name (FQIN) for the Docker image to use container: \"registry.connect.redhat.com/nuodb/nuodb-ce:latest\" Example: container: nuodb/nuodb-ce:latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dbAvailability": {
						SchemaProps: spec.SchemaProps{
							Description: "dbAvailability\n\nAvailability requirement for this database.  Values are:\n\nsingle-instance - The operator manages a single instance of each NuoDB\n                  component (Admin, SM, TE).  In the event that one\n                  component goes down or becomes unavailable, the\n                  operator will automatically replace the failed component.\n                  This is the most resource friendly option for applications\n                  that can tolerate a brief outage.\n\nmultiple-instance - The operator manages multiple instances of each\n                    NuoDB component (Admin, SM, TE).\n\nhigh-availability - The operator will maximize performance and reliability.\n\nmanual - The operator will enforce custom provided Admin, SM, TE instance counts.\n\nThe default is: \"high-availability\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dbName": {
						SchemaProps: spec.SchemaProps{
							Description: "dbName NuoDB Database name.  must consist of lowercase alphanumeric characters '[a-z0-9]+' example: dbName: test",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dbPassword": {
						SchemaProps: spec.SchemaProps{
							Description: "dbPassword Database password Example: dbPassword: secret",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dbUser": {
						SchemaProps: spec.SchemaProps{
							Description: "dbUser Name of Database user example: dbUser: dba",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"engineOptions": {
						SchemaProps: spec.SchemaProps{
							Description: "engineOptions Additional \"nuodb\" engine options Format: <option> <value> <option> <value> ... Example: engineOptions: \"\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"insightsEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "insightsEnabled Use to control Insights Opt In.  Insights provides database monitoring.  Set to \"true\" to activate or \"false\" to deactivate example: insightsEnabled: false",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"smCount": {
						SchemaProps: spec.SchemaProps{
							Description: "smCount Number of SM service pods. Requires 1 SM available for each NuoDB Database example: smCount: 1",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"smCpu": {
						SchemaProps: spec.SchemaProps{
							Description: "smCpu SM CPU cores to request example: smCpu: 1",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"smMemory": {
						SchemaProps: spec.SchemaProps{
							Description: "smMemory SM memory example: smMemory: 2Gi",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"smStorageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "smStorageClass SM persistent storage class name Example: smStorageClass: local-disk",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"smStorageSize": {
						SchemaProps: spec.SchemaProps{
							Description: "smStorageSize Storage manager (SM) volume size (GB) Example: smStorageSize: 20G",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"storageMode": {
						SchemaProps: spec.SchemaProps{
							Description: "storageMode Run NuoDB using a persistent, local, disk volume \"persistent\" or volatile storage \"ephemeral\".  Must be set to one of those values. example: storageMode: persistent",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"teCount": {
						SchemaProps: spec.SchemaProps{
							Description: "teCount Number of transaction engines (TE) nodes. Limit is 3 in CE version of NuoDB Example: teCount: 1",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"teCpu": {
						SchemaProps: spec.SchemaProps{
							Description: "teCpu TE CPU cores to request Example: teCpu: 1",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"teMemory": {
						SchemaProps: spec.SchemaProps{
							Description: "teMemory TE memory Example: teMemory: 2Gi",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"adminCount", "adminStorageClass", "adminStorageSize", "apiServer", "container", "dbAvailability", "dbName", "dbPassword", "dbUser", "engineOptions", "insightsEnabled", "smCount", "smCpu", "smMemory", "smStorageClass", "smStorageSize", "storageMode", "teCount", "teCpu", "teMemory"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbStatus defines the observed state of Nuodb",
				Properties: map[string]spec.Schema{
					"adminReadyCount": {
						SchemaProps: spec.SchemaProps{
							Description: "Admin Node Ready Count",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"smReadyCount": {
						SchemaProps: spec.SchemaProps{
							Description: "SM Node Ready Count",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"teReadyCount": {
						SchemaProps: spec.SchemaProps{
							Description: "TE Node Ready Count",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"adminHealth": {
						SchemaProps: spec.SchemaProps{
							Description: "AdminHealth of the NuoDB Domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"smHealth": {
						SchemaProps: spec.SchemaProps{
							Description: "SM Health of the NuoDB Domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"teHealth": {
						SchemaProps: spec.SchemaProps{
							Description: "TE Health of the NuoDB Domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"domainHealth": {
						SchemaProps: spec.SchemaProps{
							Description: "DomainHealth of the NuoDB Domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Orchestration phase of the NuoDB Domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"controllerVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "ControllerVersion is the version of the controller that last updated the NuoDB Domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbYcsbWl(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbYcsbWl is the Schema for the nuodbycsbwls API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWlSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWlStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWlSpec", "github.com/nuodb/nuodb-operator/pkg/apis/nuodb/v2alpha1.NuodbYcsbWlStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbYcsbWlSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbYcsbWlSpec defines the desired state of NuodbYcsbWl",
				Properties: map[string]spec.Schema{
					"dbName": {
						SchemaProps: spec.SchemaProps{
							Description: "dbName NuoDB Database name.  must consist of lowercase alphanumeric characters '[a-z0-9]+' example: dbName: test",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ycsbWorkloadCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycsbLoadName": {
						SchemaProps: spec.SchemaProps{
							Description: "ycsbLoadName",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ycsbWorkload": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ycsbLbPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ycsbNoOfProcesses": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycsbNoOfRows": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycsbNoOfIterations": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycsbOpsPerIteration": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycsbMaxDelay": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"ycsbDbSchema": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ycsbContainer": {
						SchemaProps: spec.SchemaProps{
							Description: "container NuoDB YCSB fully qualified image name (FQIN) for the Docker image to use. Example: container: nuodb/ycsb:latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"dbName", "ycsbWorkloadCount", "ycsbLoadName", "ycsbWorkload", "ycsbLbPolicy", "ycsbNoOfProcesses", "ycsbNoOfRows", "ycsbNoOfIterations", "ycsbOpsPerIteration", "ycsbMaxDelay", "ycsbDbSchema", "ycsbContainer"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nuodb_v2alpha1_NuodbYcsbWlStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NuodbYcsbWlStatus defines the observed state of NuodbYcsbWl",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

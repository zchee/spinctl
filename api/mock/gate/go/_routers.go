/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gatemock

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"//",
		Index,
	},

	Route{
		"InstanceTypesUsingGET",
		strings.ToUpper("Get"),
		"//instanceTypes",
		InstanceTypesUsingGET,
	},

	Route{
		"SubnetsUsingGET",
		strings.ToUpper("Get"),
		"//subnets",
		SubnetsUsingGET,
	},

	Route{
		"VpcsUsingGET",
		strings.ToUpper("Get"),
		"//vpcs",
		VpcsUsingGET,
	},

	Route{
		"CancelPipelineUsingPUT",
		strings.ToUpper("Put"),
		"//applications/{application}/pipelines/{id}/cancel",
		CancelPipelineUsingPUT,
	},

	Route{
		"CancelTaskUsingPUT",
		strings.ToUpper("Put"),
		"//applications/{application}/tasks/{id}/cancel",
		CancelTaskUsingPUT,
	},

	Route{
		"GetAllApplicationsUsingGET",
		strings.ToUpper("Get"),
		"//applications",
		GetAllApplicationsUsingGET,
	},

	Route{
		"GetApplicationHistoryUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/history",
		GetApplicationHistoryUsingGET,
	},

	Route{
		"GetApplicationUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}",
		GetApplicationUsingGET,
	},

	Route{
		"GetPipelineConfigUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/pipelineConfigs/{pipelineName}",
		GetPipelineConfigUsingGET,
	},

	Route{
		"GetPipelineConfigsForApplicationUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/pipelineConfigs",
		GetPipelineConfigsForApplicationUsingGET,
	},

	Route{
		"GetPipelinesUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/pipelines",
		GetPipelinesUsingGET,
	},

	Route{
		"GetStrategyConfigUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/strategyConfigs/{strategyName}",
		GetStrategyConfigUsingGET,
	},

	Route{
		"GetStrategyConfigsForApplicationUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/strategyConfigs",
		GetStrategyConfigsForApplicationUsingGET,
	},

	Route{
		"GetTaskDetailsUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/tasks/{id}/details/{taskDetailsId}",
		GetTaskDetailsUsingGET,
	},

	Route{
		"GetTaskUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/tasks/{id}",
		GetTaskUsingGET,
	},

	Route{
		"GetTasksUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/tasks",
		GetTasksUsingGET,
	},

	Route{
		"InvokePipelineConfigUsingPOST",
		strings.ToUpper("Post"),
		"//applications/{application}/pipelineConfigs/{pipelineName}",
		InvokePipelineConfigUsingPOST,
	},

	Route{
		"TaskUsingPOST",
		strings.ToUpper("Post"),
		"//applications/{application}/tasks",
		TaskUsingPOST,
	},

	Route{
		"AllUsingGET",
		strings.ToUpper("Get"),
		"//artifacts/credentials",
		AllUsingGET,
	},

	Route{
		"FindByPrincipalAndAfterAndTypeUsingGET",
		strings.ToUpper("Get"),
		"//auditevents",
		FindByPrincipalAndAfterAndTypeUsingGET,
	},

	Route{
		"FindByPrincipalAndAfterAndTypeUsingGET1",
		strings.ToUpper("Get"),
		"//auditevents.json",
		FindByPrincipalAndAfterAndTypeUsingGET1,
	},

	Route{
		"GetServiceAccountsUsingGET",
		strings.ToUpper("Get"),
		"//auth/user/serviceAccounts",
		GetServiceAccountsUsingGET,
	},

	Route{
		"LoggedOutUsingGET",
		strings.ToUpper("Get"),
		"//auth/loggedOut",
		LoggedOutUsingGET,
	},

	Route{
		"RedirectUsingGET",
		strings.ToUpper("Get"),
		"//auth/redirect",
		RedirectUsingGET,
	},

	Route{
		"SyncUsingPOST",
		strings.ToUpper("Post"),
		"//auth/roles/sync",
		SyncUsingPOST,
	},

	Route{
		"UserUsingGET",
		strings.ToUpper("Get"),
		"//auth/user",
		UserUsingGET,
	},

	Route{
		"BakeOptionsUsingGET",
		strings.ToUpper("Get"),
		"//bakery/options/{cloudProvider}",
		BakeOptionsUsingGET,
	},

	Route{
		"BakeOptionsUsingGET1",
		strings.ToUpper("Get"),
		"//bakery/options",
		BakeOptionsUsingGET1,
	},

	Route{
		"LookupLogsUsingGET",
		strings.ToUpper("Get"),
		"//bakery/logs/{region}/{statusId}",
		LookupLogsUsingGET,
	},

	Route{
		"GetBuildMastersUsingGET",
		strings.ToUpper("Get"),
		"//v2/builds",
		GetBuildMastersUsingGET,
	},

	Route{
		"GetBuildUsingGET",
		strings.ToUpper("Get"),
		"//v2/builds/{buildMaster}/build/{number}/**",
		GetBuildUsingGET,
	},

	Route{
		"GetBuildsUsingGET",
		strings.ToUpper("Get"),
		"//v2/builds/{buildMaster}/builds/**",
		GetBuildsUsingGET,
	},

	Route{
		"GetJobConfigUsingGET",
		strings.ToUpper("Get"),
		"//v2/builds/{buildMaster}/jobs/**",
		GetJobConfigUsingGET,
	},

	Route{
		"GetJobsForBuildMasterUsingGET",
		strings.ToUpper("Get"),
		"//v2/builds/{buildMaster}/jobs",
		GetJobsForBuildMasterUsingGET,
	},

	Route{
		"GetClusterLoadBalancersUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}/{clusterName}/{type}/loadBalancers",
		GetClusterLoadBalancersUsingGET,
	},

	Route{
		"GetClustersUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}/{clusterName}",
		GetClustersUsingGET,
	},

	Route{
		"GetClustersUsingGET1",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}",
		GetClustersUsingGET1,
	},

	Route{
		"GetClustersUsingGET2",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters",
		GetClustersUsingGET2,
	},

	Route{
		"GetScalingActivitiesUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}/scalingActivities",
		GetScalingActivitiesUsingGET,
	},

	Route{
		"GetServerGroupsUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}",
		GetServerGroupsUsingGET,
	},

	Route{
		"GetServerGroupsUsingGET1",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}/{clusterName}/serverGroups",
		GetServerGroupsUsingGET1,
	},

	Route{
		"GetTargetServerGroupUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target}",
		GetTargetServerGroupUsingGET,
	},

	Route{
		"GetAccountUsingGET",
		strings.ToUpper("Get"),
		"//credentials/{account}",
		GetAccountUsingGET,
	},

	Route{
		"GetAccountsUsingGET",
		strings.ToUpper("Get"),
		"//credentials",
		GetAccountsUsingGET,
	},

	Route{
		"GetLatestExecutionsByConfigIdsUsingGET",
		strings.ToUpper("Get"),
		"//executions",
		GetLatestExecutionsByConfigIdsUsingGET,
	},

	Route{
		"SearchForPipelineExecutionsByTriggerUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/executions/search",
		SearchForPipelineExecutionsByTriggerUsingGET,
	},

	Route{
		"AllByAccountUsingGET",
		strings.ToUpper("Get"),
		"//firewalls/{account}",
		AllByAccountUsingGET,
	},

	Route{
		"AllUsingGET1",
		strings.ToUpper("Get"),
		"//firewalls",
		AllUsingGET1,
	},

	Route{
		"GetSecurityGroupUsingGET",
		strings.ToUpper("Get"),
		"//firewalls/{account}/{region}/{name}",
		GetSecurityGroupUsingGET,
	},

	Route{
		"FindImagesUsingGET",
		strings.ToUpper("Get"),
		"//images/find",
		FindImagesUsingGET,
	},

	Route{
		"FindTagsUsingGET",
		strings.ToUpper("Get"),
		"//images/tags",
		FindTagsUsingGET,
	},

	Route{
		"GetImageDetailsUsingGET",
		strings.ToUpper("Get"),
		"//images/{account}/{region}/{imageId}",
		GetImageDetailsUsingGET,
	},

	Route{
		"GetConsoleOutputUsingGET",
		strings.ToUpper("Get"),
		"//instances/{account}/{region}/{instanceId}/console",
		GetConsoleOutputUsingGET,
	},

	Route{
		"GetInstanceDetailsUsingGET",
		strings.ToUpper("Get"),
		"//instances/{account}/{region}/{instanceId}",
		GetInstanceDetailsUsingGET,
	},

	Route{
		"GetJobUsingGET",
		strings.ToUpper("Get"),
		"//applications/{applicationName}/jobs/{account}/{region}/{name}",
		GetJobUsingGET,
	},

	Route{
		"GetJobsUsingGET",
		strings.ToUpper("Get"),
		"//applications/{applicationName}/jobs",
		GetJobsUsingGET,
	},

	Route{
		"GetAllUsingGET",
		strings.ToUpper("Get"),
		"//loadBalancers",
		GetAllUsingGET,
	},

	Route{
		"GetApplicationLoadBalancersUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/loadBalancers",
		GetApplicationLoadBalancersUsingGET,
	},

	Route{
		"GetLoadBalancerDetailsUsingGET",
		strings.ToUpper("Get"),
		"//loadBalancers/{account}/{region}/{name}",
		GetLoadBalancerDetailsUsingGET,
	},

	Route{
		"GetLoadBalancerUsingGET",
		strings.ToUpper("Get"),
		"//loadBalancers/{name}",
		GetLoadBalancerUsingGET,
	},

	Route{
		"AllByCloudProviderUsingGET",
		strings.ToUpper("Get"),
		"//networks/{cloudProvider}",
		AllByCloudProviderUsingGET,
	},

	Route{
		"AllUsingGET2",
		strings.ToUpper("Get"),
		"//networks",
		AllUsingGET2,
	},

	Route{
		"CancelPipelineUsingPUT1",
		strings.ToUpper("Put"),
		"//pipelines/{id}/cancel",
		CancelPipelineUsingPUT1,
	},

	Route{
		"DeletePipelineUsingDELETE",
		strings.ToUpper("Delete"),
		"//pipelines/{application}/{pipelineName}",
		DeletePipelineUsingDELETE,
	},

	Route{
		"DeletePipelineUsingDELETE1",
		strings.ToUpper("Delete"),
		"//pipelines/{id}",
		DeletePipelineUsingDELETE1,
	},

	Route{
		"EvaluateExpressionForExecutionUsingDELETE",
		strings.ToUpper("Delete"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingDELETE,
	},

	Route{
		"EvaluateExpressionForExecutionUsingGET",
		strings.ToUpper("Get"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingGET,
	},

	Route{
		"EvaluateExpressionForExecutionUsingHEAD",
		strings.ToUpper("Head"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingHEAD,
	},

	Route{
		"EvaluateExpressionForExecutionUsingOPTIONS",
		strings.ToUpper("Options"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingOPTIONS,
	},

	Route{
		"EvaluateExpressionForExecutionUsingPATCH",
		strings.ToUpper("Patch"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingPATCH,
	},

	Route{
		"EvaluateExpressionForExecutionUsingPOST",
		strings.ToUpper("Post"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingPOST,
	},

	Route{
		"EvaluateExpressionForExecutionUsingPUT",
		strings.ToUpper("Put"),
		"//pipelines/{id}/evaluateExpression",
		EvaluateExpressionForExecutionUsingPUT,
	},

	Route{
		"GetPipelineLogsUsingGET",
		strings.ToUpper("Get"),
		"//pipelines/{id}/logs",
		GetPipelineLogsUsingGET,
	},

	Route{
		"GetPipelineUsingGET",
		strings.ToUpper("Get"),
		"//pipelines/{id}",
		GetPipelineUsingGET,
	},

	Route{
		"InvokePipelineConfigUsingPOST1",
		strings.ToUpper("Post"),
		"//pipelines/{application}/{pipelineNameOrId}",
		InvokePipelineConfigUsingPOST1,
	},

	Route{
		"PausePipelineUsingPUT",
		strings.ToUpper("Put"),
		"//pipelines/{id}/pause",
		PausePipelineUsingPUT,
	},

	Route{
		"RenamePipelineUsingPOST",
		strings.ToUpper("Post"),
		"//pipelines/move",
		RenamePipelineUsingPOST,
	},

	Route{
		"RestartStageUsingPUT",
		strings.ToUpper("Put"),
		"//pipelines/{id}/stages/{stageId}/restart",
		RestartStageUsingPUT,
	},

	Route{
		"ResumePipelineUsingPUT",
		strings.ToUpper("Put"),
		"//pipelines/{id}/resume",
		ResumePipelineUsingPUT,
	},

	Route{
		"SavePipelineUsingPOST",
		strings.ToUpper("Post"),
		"//pipelines",
		SavePipelineUsingPOST,
	},

	Route{
		"StartUsingPOST",
		strings.ToUpper("Post"),
		"//pipelines/start",
		StartUsingPOST,
	},

	Route{
		"UpdatePipelineUsingPUT",
		strings.ToUpper("Put"),
		"//pipelines/{id}",
		UpdatePipelineUsingPUT,
	},

	Route{
		"UpdateStageUsingPATCH",
		strings.ToUpper("Patch"),
		"//pipelines/{id}/stages/{stageId}",
		UpdateStageUsingPATCH,
	},

	Route{
		"AllPipelinesForProjectUsingGET",
		strings.ToUpper("Get"),
		"//projects/{id}/pipelines",
		AllPipelinesForProjectUsingGET,
	},

	Route{
		"SearchUsingGET",
		strings.ToUpper("Get"),
		"//search",
		SearchUsingGET,
	},

	Route{
		"AllByAccountUsingGET1",
		strings.ToUpper("Get"),
		"//securityGroups/{account}",
		AllByAccountUsingGET1,
	},

	Route{
		"AllUsingGET3",
		strings.ToUpper("Get"),
		"//securityGroups",
		AllUsingGET3,
	},

	Route{
		"GetSecurityGroupUsingGET1",
		strings.ToUpper("Get"),
		"//securityGroups/{account}/{region}/{name}",
		GetSecurityGroupUsingGET1,
	},

	Route{
		"GetServerGroupDetailsUsingGET",
		strings.ToUpper("Get"),
		"//applications/{applicationName}/serverGroups/{account}/{region}/{serverGroupName}",
		GetServerGroupDetailsUsingGET,
	},

	Route{
		"GetServerGroupsForApplicationUsingGET",
		strings.ToUpper("Get"),
		"//applications/{applicationName}/serverGroups",
		GetServerGroupsForApplicationUsingGET,
	},

	Route{
		"GetServerGroupManagersForApplicationUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/serverGroupManagers",
		GetServerGroupManagersForApplicationUsingGET,
	},

	Route{
		"GetCurrentSnapshotUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/snapshots/{account}",
		GetCurrentSnapshotUsingGET,
	},

	Route{
		"GetSnapshotHistoryUsingGET",
		strings.ToUpper("Get"),
		"//applications/{application}/snapshots/{account}/history",
		GetSnapshotHistoryUsingGET,
	},

	Route{
		"AllByCloudProviderUsingGET1",
		strings.ToUpper("Get"),
		"//subnets/{cloudProvider}",
		AllByCloudProviderUsingGET1,
	},

	Route{
		"CancelTaskUsingPUT1",
		strings.ToUpper("Put"),
		"//tasks/{id}/cancel",
		CancelTaskUsingPUT1,
	},

	Route{
		"CancelTasksUsingPUT",
		strings.ToUpper("Put"),
		"//tasks/cancel",
		CancelTasksUsingPUT,
	},

	Route{
		"DeleteTaskUsingDELETE",
		strings.ToUpper("Delete"),
		"//tasks/{id}",
		DeleteTaskUsingDELETE,
	},

	Route{
		"GetTaskDetailsUsingGET1",
		strings.ToUpper("Get"),
		"//tasks/{id}/details/{taskDetailsId}",
		GetTaskDetailsUsingGET1,
	},

	Route{
		"GetTaskUsingGET1",
		strings.ToUpper("Get"),
		"//tasks/{id}",
		GetTaskUsingGET1,
	},

	Route{
		"TaskUsingPOST1",
		strings.ToUpper("Post"),
		"//tasks",
		TaskUsingPOST1,
	},

	Route{
		"CreateCanaryConfigUsingPOST",
		strings.ToUpper("Post"),
		"//v2/canaryConfig",
		CreateCanaryConfigUsingPOST,
	},

	Route{
		"DeleteCanaryConfigUsingDELETE",
		strings.ToUpper("Delete"),
		"//v2/canaryConfig/{id}",
		DeleteCanaryConfigUsingDELETE,
	},

	Route{
		"GetCanaryConfigUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaryConfig/{id}",
		GetCanaryConfigUsingGET,
	},

	Route{
		"GetCanaryConfigsUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaryConfig",
		GetCanaryConfigsUsingGET,
	},

	Route{
		"UpdateCanaryConfigUsingPUT",
		strings.ToUpper("Put"),
		"//v2/canaryConfig/{id}",
		UpdateCanaryConfigUsingPUT,
	},

	Route{
		"GetCanaryResultUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaries/canary/{canaryConfigId}/{canaryExecutionId}",
		GetCanaryResultUsingGET,
	},

	Route{
		"GetCanaryResultsByApplicationUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaries/{application}/executions",
		GetCanaryResultsByApplicationUsingGET,
	},

	Route{
		"GetMetricSetPairListUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaries/metricSetPairList/{metricSetPairListId}",
		GetMetricSetPairListUsingGET,
	},

	Route{
		"ListCredentialsUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaries/credentials",
		ListCredentialsUsingGET,
	},

	Route{
		"ListJudgesUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaries/judges",
		ListJudgesUsingGET,
	},

	Route{
		"ListMetricsServiceMetadataUsingGET",
		strings.ToUpper("Get"),
		"//v2/canaries/metadata/metricsService",
		ListMetricsServiceMetadataUsingGET,
	},

	Route{
		"PreconfiguredWebhooksUsingGET",
		strings.ToUpper("Get"),
		"//webhooks/preconfigured",
		PreconfiguredWebhooksUsingGET,
	},

	Route{
		"WebhooksUsingPOST",
		strings.ToUpper("Post"),
		"//webhooks/{type}/{source}",
		WebhooksUsingPOST,
	},
}

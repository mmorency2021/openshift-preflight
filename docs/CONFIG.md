# Configuring Preflight

The following configurables are available for the `preflight` tool.

## Common Configuration

|Variable|Kind|Doc|Required or Optional|Default|
|--|--|--|--|--|
|`PFLT_LOGLEVEL`|env|The verbosity of the preflight tool itself. Ex. warn, debug, trace, info, error|optional|[warn](https://github.com/redhat-openshift-ecosystem/openshift-preflight/blob/main/cmd/defaults.go#L6)|
|`PFLT_LOGFILE`|env|Where the execution logfile will be written.|optional|[preflight.log](https://github.com/redhat-openshift-ecosystem/openshift-preflight/blob/main/cmd/defaults.go#L5)|
|`PFLT_ARTIFACTS`|env|Where check-specific artifacts will be written.|optional|[artifacts/](https://github.com/redhat-openshift-ecosystem/openshift-preflight/blob/main/cmd/defaults.go#L7)|


## Operator Policy Configuration

These configurables are specific to cases where `preflight check operator ...`
is called.

|Variable|Kind|Doc|Required or Optional|Default|
|--|--|--|--|--|
|`KUBECONFIG`|env|The operator policy must interact with a Kubernetes cluster for checks such as `DeployableByOLM` and running [OperatorSDK Scorecard](https://sdk.operatorframework.io/docs/testing-operators/scorecard/).|required|-|
|`PFLT_NAMESPACE`|env|The namespace to use when running [OperatorSDK Scorecard](https://sdk.operatorframework.io/docs/testing-operators/scorecard/)|optional|[default](https://github.com/redhat-openshift-ecosystem/openshift-preflight/blob/main/cmd/defaults.go#L8)|
|`PFLT_SERVICEACCOUNT`|env|The service account to use when running [OperatorSDK Scorecard](https://sdk.operatorframework.io/docs/testing-operators/scorecard/)|optional|[default](https://github.com/redhat-openshift-ecosystem/openshift-preflight/blob/main/cmd/defaults.go#L9)|
|`PFLT_INDEXIMAGE`|env|The index image to use when testing that an operator is `DeployableByOLM`|optional|-|

## Container Policy Configuration

There are no configurables specific to the container policy (i.e. when running `preflight check container ...`).

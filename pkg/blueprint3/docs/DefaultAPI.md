# \DefaultAPI

All URIs are relative to *https://osbuild.org/wip/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateBlueprint**](DefaultAPI.md#ValidateBlueprint) | **Post** /validate_blueprint | Validate blueprint



## ValidateBlueprint

> Blueprint ValidateBlueprint(ctx).Blueprint(blueprint).Execute()

Validate blueprint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	blueprint := *openapiclient.NewBlueprint("Name_example") // Blueprint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ValidateBlueprint(context.Background()).Blueprint(blueprint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ValidateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateBlueprint`: Blueprint
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ValidateBlueprint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprint** | [**Blueprint**](Blueprint.md) |  | 

### Return type

[**Blueprint**](Blueprint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


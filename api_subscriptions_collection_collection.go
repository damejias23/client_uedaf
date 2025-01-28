/*
Nuedaf_EventExposure

UEDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uedaf

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// SubscriptionsCollectionCollectionApiService SubscriptionsCollectionCollectionApi service
type SubscriptionsCollectionCollectionApiService service

type ApiCreateSubscriptionRequest struct {
	ctx context.Context
	ApiService *SubscriptionsCollectionCollectionApiService
	uedafCreateEventSubscription *UedafCreateEventSubscription
}

func (r ApiCreateSubscriptionRequest) UedafCreateEventSubscription(uedafCreateEventSubscription UedafCreateEventSubscription) ApiCreateSubscriptionRequest {
	r.uedafCreateEventSubscription = &uedafCreateEventSubscription
	return r
}

func (r ApiCreateSubscriptionRequest) Execute() (*UedafCreatedEventSubscription, *http.Response, error) {
	return r.ApiService.CreateSubscriptionExecute(r)
}

/*
CreateSubscription Nuedaf_EventExposure Subscribe service Operation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateSubscriptionRequest
*/
func (a *SubscriptionsCollectionCollectionApiService) CreateSubscription(ctx context.Context) ApiCreateSubscriptionRequest {
	return ApiCreateSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UedafCreatedEventSubscription
func (a *SubscriptionsCollectionCollectionApiService) CreateSubscriptionExecute(r ApiCreateSubscriptionRequest) (*UedafCreatedEventSubscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UedafCreatedEventSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsCollectionCollectionApiService.CreateSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uedafCreateEventSubscription == nil {
		return localVarReturnValue, nil, reportError("uedafCreateEventSubscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uedafCreateEventSubscription
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

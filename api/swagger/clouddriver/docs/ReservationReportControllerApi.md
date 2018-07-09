# \ReservationReportControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReservationReportsByNameUsingGET**](ReservationReportControllerApi.md#GetReservationReportsByNameUsingGET) | **Get** /reports/reservation/{name} | getReservationReportsByName
[**GetReservationReportsUsingGET**](ReservationReportControllerApi.md#GetReservationReportsUsingGET) | **Get** /reports/reservation | getReservationReports


# **GetReservationReportsByNameUsingGET**
> []ReservationReport GetReservationReportsByNameUsingGET(ctx, name, filters)
getReservationReportsByName

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| name | 
  **filters** | **string**| filters | 

### Return type

[**[]ReservationReport**](ReservationReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReservationReportsUsingGET**
> []ReservationReport GetReservationReportsUsingGET(ctx, filters)
getReservationReports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **filters** | **string**| filters | 

### Return type

[**[]ReservationReport**](ReservationReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ReferenceDataAPIService ReferenceDataAPI service
type ReferenceDataAPIService service

type ReferenceDataAPIGetBondsRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	symbol     *string
	exchange   *string
	country    *string
	format     *FormatEnum
	delimiter  *string
	showPlan   *bool
	page       *int64
	outputsize *int64
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetBondsRequest) Symbol(symbol string) ReferenceDataAPIGetBondsRequest {
	r.symbol = &symbol
	return r
}

// Filter by exchange name
func (r ReferenceDataAPIGetBondsRequest) Exchange(exchange string) ReferenceDataAPIGetBondsRequest {
	r.exchange = &exchange
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ReferenceDataAPIGetBondsRequest) Country(country string) ReferenceDataAPIGetBondsRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetBondsRequest) Format(format FormatEnum) ReferenceDataAPIGetBondsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetBondsRequest) Delimiter(delimiter string) ReferenceDataAPIGetBondsRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ReferenceDataAPIGetBondsRequest) ShowPlan(showPlan bool) ReferenceDataAPIGetBondsRequest {
	r.showPlan = &showPlan
	return r
}

// Page number of the results to fetch
func (r ReferenceDataAPIGetBondsRequest) Page(page int64) ReferenceDataAPIGetBondsRequest {
	r.page = &page
	return r
}

// Determines the number of data points returned in the output
func (r ReferenceDataAPIGetBondsRequest) Outputsize(outputsize int64) ReferenceDataAPIGetBondsRequest {
	r.outputsize = &outputsize
	return r
}

func (r ReferenceDataAPIGetBondsRequest) Execute() (*GetBonds200Response, *http.Response, error) {
	return r.ApiService.GetBondsExecute(r)
}

/*
GetBonds Fixed income

The fixed income endpoint provides a daily updated list of available bonds. It returns an array containing detailed information about each bond, including identifiers, names, and other relevant attributes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetBondsRequest
*/
func (a *ReferenceDataAPIService) GetBonds(ctx context.Context) ReferenceDataAPIGetBondsRequest {
	return ReferenceDataAPIGetBondsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBonds200Response
func (a *ReferenceDataAPIService) GetBondsExecute(r ReferenceDataAPIGetBondsRequest) (*GetBonds200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBonds200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetBonds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bonds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
		r.showPlan = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int64 = 1
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
		r.page = &defaultValue
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
		var defaultValue int64 = 5000
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
		r.outputsize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetCommoditiesRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	symbol     *string
	category   *string
	format     *FormatEnum
	delimiter  *string
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetCommoditiesRequest) Symbol(symbol string) ReferenceDataAPIGetCommoditiesRequest {
	r.symbol = &symbol
	return r
}

// Filter by category of commodity
func (r ReferenceDataAPIGetCommoditiesRequest) Category(category string) ReferenceDataAPIGetCommoditiesRequest {
	r.category = &category
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetCommoditiesRequest) Format(format FormatEnum) ReferenceDataAPIGetCommoditiesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetCommoditiesRequest) Delimiter(delimiter string) ReferenceDataAPIGetCommoditiesRequest {
	r.delimiter = &delimiter
	return r
}

func (r ReferenceDataAPIGetCommoditiesRequest) Execute() (*GetCommodities200Response, *http.Response, error) {
	return r.ApiService.GetCommoditiesExecute(r)
}

/*
GetCommodities Commodities

The commodities endpoint provides a daily updated list of available commodity pairs, across precious metals, livestock, softs, grains, etc.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetCommoditiesRequest
*/
func (a *ReferenceDataAPIService) GetCommodities(ctx context.Context) ReferenceDataAPIGetCommoditiesRequest {
	return ReferenceDataAPIGetCommoditiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCommodities200Response
func (a *ReferenceDataAPIService) GetCommoditiesExecute(r ReferenceDataAPIGetCommoditiesRequest) (*GetCommodities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCommodities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCommodities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commodities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetCountriesRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
}

func (r ReferenceDataAPIGetCountriesRequest) Execute() (*GetCountries200Response, *http.Response, error) {
	return r.ApiService.GetCountriesExecute(r)
}

/*
GetCountries Countries

The countries endpoint provides a comprehensive list of countries, including their ISO codes, official names, capitals, and currencies. This data is essential for applications requiring accurate country information for tasks such as localization, currency conversion, or geographic analysis.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetCountriesRequest
*/
func (a *ReferenceDataAPIService) GetCountries(ctx context.Context) ReferenceDataAPIGetCountriesRequest {
	return ReferenceDataAPIGetCountriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCountries200Response
func (a *ReferenceDataAPIService) GetCountriesExecute(r ReferenceDataAPIGetCountriesRequest) (*GetCountries200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCountries200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCountries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/countries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetCrossListingsRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	symbol     *string
	exchange   *string
	micCode    *string
	country    *string
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetCrossListingsRequest) Symbol(symbol string) ReferenceDataAPIGetCrossListingsRequest {
	r.symbol = &symbol
	return r
}

// Exchange where instrument is traded
func (r ReferenceDataAPIGetCrossListingsRequest) Exchange(exchange string) ReferenceDataAPIGetCrossListingsRequest {
	r.exchange = &exchange
	return r
}

// Market identifier code (MIC) under ISO 10383 standard
func (r ReferenceDataAPIGetCrossListingsRequest) MicCode(micCode string) ReferenceDataAPIGetCrossListingsRequest {
	r.micCode = &micCode
	return r
}

// Country to which stock exchange belongs to
func (r ReferenceDataAPIGetCrossListingsRequest) Country(country string) ReferenceDataAPIGetCrossListingsRequest {
	r.country = &country
	return r
}

func (r ReferenceDataAPIGetCrossListingsRequest) Execute() (*GetCrossListings200Response, *http.Response, error) {
	return r.ApiService.GetCrossListingsExecute(r)
}

/*
GetCrossListings Cross listings

The cross_listings endpoint provides a daily updated list of cross-listed symbols for a specified financial instrument. Cross-listed symbols represent the same security available on multiple exchanges. This endpoint is useful for identifying all the exchanges where a particular security is traded, allowing users to access comprehensive trading information across different markets.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetCrossListingsRequest
*/
func (a *ReferenceDataAPIService) GetCrossListings(ctx context.Context) ReferenceDataAPIGetCrossListingsRequest {
	return ReferenceDataAPIGetCrossListingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCrossListings200Response
func (a *ReferenceDataAPIService) GetCrossListingsExecute(r ReferenceDataAPIGetCrossListingsRequest) (*GetCrossListings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCrossListings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCrossListings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cross_listings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetCryptocurrenciesRequest struct {
	ctx           context.Context
	ApiService    *ReferenceDataAPIService
	symbol        *string
	exchange      *string
	currencyBase  *string
	currencyQuote *string
	format        *FormatEnum
	delimiter     *string
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetCryptocurrenciesRequest) Symbol(symbol string) ReferenceDataAPIGetCryptocurrenciesRequest {
	r.symbol = &symbol
	return r
}

// Filter by exchange name. E.g. &#x60;Binance&#x60;, &#x60;Coinbase&#x60;, etc.
func (r ReferenceDataAPIGetCryptocurrenciesRequest) Exchange(exchange string) ReferenceDataAPIGetCryptocurrenciesRequest {
	r.exchange = &exchange
	return r
}

// Filter by currency base
func (r ReferenceDataAPIGetCryptocurrenciesRequest) CurrencyBase(currencyBase string) ReferenceDataAPIGetCryptocurrenciesRequest {
	r.currencyBase = &currencyBase
	return r
}

// Filter by currency quote
func (r ReferenceDataAPIGetCryptocurrenciesRequest) CurrencyQuote(currencyQuote string) ReferenceDataAPIGetCryptocurrenciesRequest {
	r.currencyQuote = &currencyQuote
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetCryptocurrenciesRequest) Format(format FormatEnum) ReferenceDataAPIGetCryptocurrenciesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetCryptocurrenciesRequest) Delimiter(delimiter string) ReferenceDataAPIGetCryptocurrenciesRequest {
	r.delimiter = &delimiter
	return r
}

func (r ReferenceDataAPIGetCryptocurrenciesRequest) Execute() (*GetCryptocurrencies200Response, *http.Response, error) {
	return r.ApiService.GetCryptocurrenciesExecute(r)
}

/*
GetCryptocurrencies Cryptocurrency pairs

The cryptocurrencies endpoint provides a daily updated list of all available cryptos. It returns an array containing detailed information about each cryptocurrency, including its symbol, name, and other relevant identifiers. This endpoint is useful for retrieving a comprehensive catalog of cryptocurrencies for applications that require up-to-date market listings or need to display available crypto assets to users.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetCryptocurrenciesRequest
*/
func (a *ReferenceDataAPIService) GetCryptocurrencies(ctx context.Context) ReferenceDataAPIGetCryptocurrenciesRequest {
	return ReferenceDataAPIGetCryptocurrenciesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCryptocurrencies200Response
func (a *ReferenceDataAPIService) GetCryptocurrenciesExecute(r ReferenceDataAPIGetCryptocurrenciesRequest) (*GetCryptocurrencies200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCryptocurrencies200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCryptocurrencies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cryptocurrencies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.currencyBase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_base", r.currencyBase, "form", "")
	}
	if r.currencyQuote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_quote", r.currencyQuote, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetCryptocurrencyExchangesRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	format     *FormatEnum
	delimiter  *string
}

// The format of the response data
func (r ReferenceDataAPIGetCryptocurrencyExchangesRequest) Format(format FormatEnum) ReferenceDataAPIGetCryptocurrencyExchangesRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r ReferenceDataAPIGetCryptocurrencyExchangesRequest) Delimiter(delimiter string) ReferenceDataAPIGetCryptocurrencyExchangesRequest {
	r.delimiter = &delimiter
	return r
}

func (r ReferenceDataAPIGetCryptocurrencyExchangesRequest) Execute() (*GetCryptocurrencyExchanges200Response, *http.Response, error) {
	return r.ApiService.GetCryptocurrencyExchangesExecute(r)
}

/*
GetCryptocurrencyExchanges Cryptocurrency exchanges

The cryptocurrency exchanges endpoint provides a daily updated list of available cryptocurrency exchanges. It returns an array containing details about each exchange, such as exchange names and identifiers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetCryptocurrencyExchangesRequest
*/
func (a *ReferenceDataAPIService) GetCryptocurrencyExchanges(ctx context.Context) ReferenceDataAPIGetCryptocurrencyExchangesRequest {
	return ReferenceDataAPIGetCryptocurrencyExchangesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCryptocurrencyExchanges200Response
func (a *ReferenceDataAPIService) GetCryptocurrencyExchangesExecute(r ReferenceDataAPIGetCryptocurrencyExchangesRequest) (*GetCryptocurrencyExchanges200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCryptocurrencyExchanges200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCryptocurrencyExchanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cryptocurrency_exchanges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetEarliestTimestampRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	interval   *IntervalEnum
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	timezone   *string
}

// Interval between two consecutive points in time series.
func (r ReferenceDataAPIGetEarliestTimestampRequest) Interval(interval IntervalEnum) ReferenceDataAPIGetEarliestTimestampRequest {
	r.interval = &interval
	return r
}

// Symbol ticker of the instrument.
func (r ReferenceDataAPIGetEarliestTimestampRequest) Symbol(symbol string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r ReferenceDataAPIGetEarliestTimestampRequest) Figi(figi string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetEarliestTimestampRequest) Isin(isin string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetEarliestTimestampRequest) Cusip(cusip string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded.
func (r ReferenceDataAPIGetEarliestTimestampRequest) Exchange(exchange string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard.
func (r ReferenceDataAPIGetEarliestTimestampRequest) MicCode(micCode string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.micCode = &micCode
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.&lt;/p&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ReferenceDataAPIGetEarliestTimestampRequest) Timezone(timezone string) ReferenceDataAPIGetEarliestTimestampRequest {
	r.timezone = &timezone
	return r
}

func (r ReferenceDataAPIGetEarliestTimestampRequest) Execute() (*GetEarliestTimestamp200Response, *http.Response, error) {
	return r.ApiService.GetEarliestTimestampExecute(r)
}

/*
GetEarliestTimestamp Earliest timestamp

The earliest_timestamp endpoint provides the earliest available date and time for a specified financial instrument at a given data interval. This endpoint is useful for determining the starting point of historical data availability for various assets, such as stocks or currencies, allowing users to understand the time range covered by the data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetEarliestTimestampRequest
*/
func (a *ReferenceDataAPIService) GetEarliestTimestamp(ctx context.Context) ReferenceDataAPIGetEarliestTimestampRequest {
	return ReferenceDataAPIGetEarliestTimestampRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEarliestTimestamp200Response
func (a *ReferenceDataAPIService) GetEarliestTimestampExecute(r ReferenceDataAPIGetEarliestTimestampRequest) (*GetEarliestTimestamp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEarliestTimestamp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetEarliestTimestamp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/earliest_timestamp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	} else {
		var defaultValue string = "Exchange"
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", defaultValue, "form", "")
		r.timezone = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetEtfRequest struct {
	ctx             context.Context
	ApiService      *ReferenceDataAPIService
	symbol          *string
	figi            *string
	isin            *string
	cusip           *string
	cik             *string
	exchange        *string
	micCode         *string
	country         *string
	format          *FormatEnum
	delimiter       *string
	showPlan        *bool
	includeDelisted *bool
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetEtfRequest) Symbol(symbol string) ReferenceDataAPIGetEtfRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r ReferenceDataAPIGetEtfRequest) Figi(figi string) ReferenceDataAPIGetEtfRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetEtfRequest) Isin(isin string) ReferenceDataAPIGetEtfRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetEtfRequest) Cusip(cusip string) ReferenceDataAPIGetEtfRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ReferenceDataAPIGetEtfRequest) Cik(cik string) ReferenceDataAPIGetEtfRequest {
	r.cik = &cik
	return r
}

// Filter by exchange name
func (r ReferenceDataAPIGetEtfRequest) Exchange(exchange string) ReferenceDataAPIGetEtfRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ReferenceDataAPIGetEtfRequest) MicCode(micCode string) ReferenceDataAPIGetEtfRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ReferenceDataAPIGetEtfRequest) Country(country string) ReferenceDataAPIGetEtfRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetEtfRequest) Format(format FormatEnum) ReferenceDataAPIGetEtfRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetEtfRequest) Delimiter(delimiter string) ReferenceDataAPIGetEtfRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ReferenceDataAPIGetEtfRequest) ShowPlan(showPlan bool) ReferenceDataAPIGetEtfRequest {
	r.showPlan = &showPlan
	return r
}

// Include delisted identifiers
func (r ReferenceDataAPIGetEtfRequest) IncludeDelisted(includeDelisted bool) ReferenceDataAPIGetEtfRequest {
	r.includeDelisted = &includeDelisted
	return r
}

func (r ReferenceDataAPIGetEtfRequest) Execute() (*GetEtf200Response, *http.Response, error) {
	return r.ApiService.GetEtfExecute(r)
}

/*
GetEtf ETFs

The ETFs endpoint provides a daily updated list of all available Exchange-Traded Funds. It returns an array containing detailed information about each ETF, including its symbol, name, and other relevant identifiers. This endpoint is useful for retrieving a comprehensive catalog of ETFs for portfolio management, investment tracking, or financial analysis.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetEtfRequest
*/
func (a *ReferenceDataAPIService) GetEtf(ctx context.Context) ReferenceDataAPIGetEtfRequest {
	return ReferenceDataAPIGetEtfRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEtf200Response
func (a *ReferenceDataAPIService) GetEtfExecute(r ReferenceDataAPIGetEtfRequest) (*GetEtf200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEtf200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetEtf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/etfs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
		r.showPlan = &defaultValue
	}
	if r.includeDelisted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", r.includeDelisted, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", defaultValue, "form", "")
		r.includeDelisted = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetExchangeScheduleRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	micName    *string
	micCode    *string
	country    *string
	date       *string
}

// Filter by exchange name
func (r ReferenceDataAPIGetExchangeScheduleRequest) MicName(micName string) ReferenceDataAPIGetExchangeScheduleRequest {
	r.micName = &micName
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ReferenceDataAPIGetExchangeScheduleRequest) MicCode(micCode string) ReferenceDataAPIGetExchangeScheduleRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ReferenceDataAPIGetExchangeScheduleRequest) Country(country string) ReferenceDataAPIGetExchangeScheduleRequest {
	r.country = &country
	return r
}

// &lt;p&gt; If a date is provided, the API returns the schedule for the specified date; otherwise, it returns the default (common) schedule. &lt;/p&gt; The date can be specified in one of the following formats: &lt;ul&gt; &lt;li&gt;An exact date (e.g., &lt;code&gt;2021-10-27&lt;/code&gt;)&lt;/li&gt; &lt;li&gt;A human-readable keyword: &lt;code&gt;today&lt;/code&gt; or &lt;code&gt;yesterday&lt;/code&gt;&lt;/li&gt; &lt;li&gt;A full datetime string in UTC (e.g., &lt;code&gt;2025-04-11T20:00:00&lt;/code&gt;) to retrieve the schedule corresponding to the day in the specified time.&lt;/li&gt; &lt;/ul&gt; When using a datetime value, the resulting schedule will correspond to the local calendar day at the specified time. For example, &lt;code&gt;2025-04-11T20:00:00 UTC&lt;/code&gt; corresponds to: &lt;ul&gt; &lt;li&gt;&lt;code&gt;2025-04-11&lt;/code&gt; in the &lt;code&gt;America/New_York&lt;/code&gt; timezone&lt;/li&gt; &lt;li&gt;&lt;code&gt;2025-04-12&lt;/code&gt; in the &lt;code&gt;Australia/Sydney&lt;/code&gt; timezone&lt;/li&gt; &lt;/ul&gt;
func (r ReferenceDataAPIGetExchangeScheduleRequest) Date(date string) ReferenceDataAPIGetExchangeScheduleRequest {
	r.date = &date
	return r
}

func (r ReferenceDataAPIGetExchangeScheduleRequest) Execute() (*GetExchangeSchedule200Response, *http.Response, error) {
	return r.ApiService.GetExchangeScheduleExecute(r)
}

/*
GetExchangeSchedule Exchanges schedule

The exchanges schedule endpoint provides detailed information about various stock exchanges, including their trading hours and operational days. This data is essential for users who need to know when specific exchanges are open for trading, allowing them to plan their activities around the availability of these markets.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetExchangeScheduleRequest
*/
func (a *ReferenceDataAPIService) GetExchangeSchedule(ctx context.Context) ReferenceDataAPIGetExchangeScheduleRequest {
	return ReferenceDataAPIGetExchangeScheduleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetExchangeSchedule200Response
func (a *ReferenceDataAPIService) GetExchangeScheduleExecute(r ReferenceDataAPIGetExchangeScheduleRequest) (*GetExchangeSchedule200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetExchangeSchedule200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetExchangeSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchange_schedule"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.micName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_name", r.micName, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetExchangesRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	type_      *TypeEnum
	name       *string
	code       *string
	country    *string
	format     *FormatEnum
	delimiter  *string
	showPlan   *bool
}

// The asset class to which the instrument belongs
func (r ReferenceDataAPIGetExchangesRequest) Type_(type_ TypeEnum) ReferenceDataAPIGetExchangesRequest {
	r.type_ = &type_
	return r
}

// Filter by exchange name
func (r ReferenceDataAPIGetExchangesRequest) Name(name string) ReferenceDataAPIGetExchangesRequest {
	r.name = &name
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ReferenceDataAPIGetExchangesRequest) Code(code string) ReferenceDataAPIGetExchangesRequest {
	r.code = &code
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ReferenceDataAPIGetExchangesRequest) Country(country string) ReferenceDataAPIGetExchangesRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetExchangesRequest) Format(format FormatEnum) ReferenceDataAPIGetExchangesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetExchangesRequest) Delimiter(delimiter string) ReferenceDataAPIGetExchangesRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ReferenceDataAPIGetExchangesRequest) ShowPlan(showPlan bool) ReferenceDataAPIGetExchangesRequest {
	r.showPlan = &showPlan
	return r
}

func (r ReferenceDataAPIGetExchangesRequest) Execute() (*GetExchanges200Response, *http.Response, error) {
	return r.ApiService.GetExchangesExecute(r)
}

/*
GetExchanges Exchanges

The exchanges endpoint provides a comprehensive list of all available equity exchanges. It returns an array containing detailed information about each exchange, such as exchange code, name, country, and timezone. This data is updated daily.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetExchangesRequest
*/
func (a *ReferenceDataAPIService) GetExchanges(ctx context.Context) ReferenceDataAPIGetExchangesRequest {
	return ReferenceDataAPIGetExchangesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetExchanges200Response
func (a *ReferenceDataAPIService) GetExchangesExecute(r ReferenceDataAPIGetExchangesRequest) (*GetExchanges200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetExchanges200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetExchanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchanges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
		r.showPlan = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetForexPairsRequest struct {
	ctx           context.Context
	ApiService    *ReferenceDataAPIService
	symbol        *string
	currencyBase  *string
	currencyQuote *string
	format        *FormatEnum
	delimiter     *string
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetForexPairsRequest) Symbol(symbol string) ReferenceDataAPIGetForexPairsRequest {
	r.symbol = &symbol
	return r
}

// Filter by currency base
func (r ReferenceDataAPIGetForexPairsRequest) CurrencyBase(currencyBase string) ReferenceDataAPIGetForexPairsRequest {
	r.currencyBase = &currencyBase
	return r
}

// Filter by currency quote
func (r ReferenceDataAPIGetForexPairsRequest) CurrencyQuote(currencyQuote string) ReferenceDataAPIGetForexPairsRequest {
	r.currencyQuote = &currencyQuote
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetForexPairsRequest) Format(format FormatEnum) ReferenceDataAPIGetForexPairsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetForexPairsRequest) Delimiter(delimiter string) ReferenceDataAPIGetForexPairsRequest {
	r.delimiter = &delimiter
	return r
}

func (r ReferenceDataAPIGetForexPairsRequest) Execute() (*GetForexPairs200Response, *http.Response, error) {
	return r.ApiService.GetForexPairsExecute(r)
}

/*
GetForexPairs Forex pairs

The forex pairs endpoint provides a comprehensive list of all available foreign exchange currency pairs. It returns an array of forex pairs, which is updated daily.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetForexPairsRequest
*/
func (a *ReferenceDataAPIService) GetForexPairs(ctx context.Context) ReferenceDataAPIGetForexPairsRequest {
	return ReferenceDataAPIGetForexPairsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetForexPairs200Response
func (a *ReferenceDataAPIService) GetForexPairsExecute(r ReferenceDataAPIGetForexPairsRequest) (*GetForexPairs200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetForexPairs200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetForexPairs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/forex_pairs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.currencyBase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_base", r.currencyBase, "form", "")
	}
	if r.currencyQuote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_quote", r.currencyQuote, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetFundsRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	cik        *string
	exchange   *string
	country    *string
	format     *FormatEnum
	delimiter  *string
	showPlan   *bool
	page       *int64
	outputsize *int64
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetFundsRequest) Symbol(symbol string) ReferenceDataAPIGetFundsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r ReferenceDataAPIGetFundsRequest) Figi(figi string) ReferenceDataAPIGetFundsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetFundsRequest) Isin(isin string) ReferenceDataAPIGetFundsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetFundsRequest) Cusip(cusip string) ReferenceDataAPIGetFundsRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ReferenceDataAPIGetFundsRequest) Cik(cik string) ReferenceDataAPIGetFundsRequest {
	r.cik = &cik
	return r
}

// Filter by exchange name
func (r ReferenceDataAPIGetFundsRequest) Exchange(exchange string) ReferenceDataAPIGetFundsRequest {
	r.exchange = &exchange
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ReferenceDataAPIGetFundsRequest) Country(country string) ReferenceDataAPIGetFundsRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetFundsRequest) Format(format FormatEnum) ReferenceDataAPIGetFundsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetFundsRequest) Delimiter(delimiter string) ReferenceDataAPIGetFundsRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ReferenceDataAPIGetFundsRequest) ShowPlan(showPlan bool) ReferenceDataAPIGetFundsRequest {
	r.showPlan = &showPlan
	return r
}

// Page number of the results to fetch
func (r ReferenceDataAPIGetFundsRequest) Page(page int64) ReferenceDataAPIGetFundsRequest {
	r.page = &page
	return r
}

// Determines the number of data points returned in the output
func (r ReferenceDataAPIGetFundsRequest) Outputsize(outputsize int64) ReferenceDataAPIGetFundsRequest {
	r.outputsize = &outputsize
	return r
}

func (r ReferenceDataAPIGetFundsRequest) Execute() (*GetFunds200Response, *http.Response, error) {
	return r.ApiService.GetFundsExecute(r)
}

/*
GetFunds Funds

The funds endpoint provides a daily updated list of available investment funds. It returns an array containing detailed information about each fund, including identifiers, names, and other relevant attributes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetFundsRequest
*/
func (a *ReferenceDataAPIService) GetFunds(ctx context.Context) ReferenceDataAPIGetFundsRequest {
	return ReferenceDataAPIGetFundsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFunds200Response
func (a *ReferenceDataAPIService) GetFundsExecute(r ReferenceDataAPIGetFundsRequest) (*GetFunds200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetFunds200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetFunds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/funds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
		r.showPlan = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int64 = 1
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
		r.page = &defaultValue
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
		var defaultValue int64 = 5000
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
		r.outputsize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetInstrumentTypeRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
}

func (r ReferenceDataAPIGetInstrumentTypeRequest) Execute() (*GetInstrumentType200Response, *http.Response, error) {
	return r.ApiService.GetInstrumentTypeExecute(r)
}

/*
GetInstrumentType Instrument type

The instrument type endpoint lists all available financial instrument types, such as stocks, ETFs, and cryptos. This information is essential for users to identify and categorize different financial instruments when accessing or analyzing market data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetInstrumentTypeRequest
*/
func (a *ReferenceDataAPIService) GetInstrumentType(ctx context.Context) ReferenceDataAPIGetInstrumentTypeRequest {
	return ReferenceDataAPIGetInstrumentTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetInstrumentType200Response
func (a *ReferenceDataAPIService) GetInstrumentTypeExecute(r ReferenceDataAPIGetInstrumentTypeRequest) (*GetInstrumentType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetInstrumentType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetInstrumentType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/instrument_type"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetIntervalsRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
}

func (r ReferenceDataAPIGetIntervalsRequest) Execute() (*GetIntervals200Response, *http.Response, error) {
	return r.ApiService.GetIntervalsExecute(r)
}

/*
GetIntervals Intervals List

The intervals endpoint provides a list of supported time intervals that can be used for querying financial data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetIntervalsRequest
*/
func (a *ReferenceDataAPIService) GetIntervals(ctx context.Context) ReferenceDataAPIGetIntervalsRequest {
	return ReferenceDataAPIGetIntervalsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetIntervals200Response
func (a *ReferenceDataAPIService) GetIntervalsExecute(r ReferenceDataAPIGetIntervalsRequest) (*GetIntervals200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetIntervals200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetIntervals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intervals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetMarketStateRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	exchange   *string
	code       *string
	country    *string
}

// The exchange name where the instrument is traded.
func (r ReferenceDataAPIGetMarketStateRequest) Exchange(exchange string) ReferenceDataAPIGetMarketStateRequest {
	r.exchange = &exchange
	return r
}

// The Market Identifier Code (MIC) of the exchange where the instrument is traded.
func (r ReferenceDataAPIGetMarketStateRequest) Code(code string) ReferenceDataAPIGetMarketStateRequest {
	r.code = &code
	return r
}

// The country where the exchange is located. Takes country name or alpha code.
func (r ReferenceDataAPIGetMarketStateRequest) Country(country string) ReferenceDataAPIGetMarketStateRequest {
	r.country = &country
	return r
}

func (r ReferenceDataAPIGetMarketStateRequest) Execute() ([]MarketStateResponseItem, *http.Response, error) {
	return r.ApiService.GetMarketStateExecute(r)
}

/*
GetMarketState Market state

The market state endpoint provides real-time information on the operational status of all available stock exchanges. It returns data on whether each exchange is currently open or closed, along with the time remaining until the next opening or closing. This endpoint is useful for users who need to monitor exchange hours and plan their trading activities accordingly.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetMarketStateRequest
*/
func (a *ReferenceDataAPIService) GetMarketState(ctx context.Context) ReferenceDataAPIGetMarketStateRequest {
	return ReferenceDataAPIGetMarketStateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []MarketStateResponseItem
func (a *ReferenceDataAPIService) GetMarketStateExecute(r ReferenceDataAPIGetMarketStateRequest) ([]MarketStateResponseItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []MarketStateResponseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetMarketState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_state"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetStocksRequest struct {
	ctx             context.Context
	ApiService      *ReferenceDataAPIService
	symbol          *string
	figi            *string
	isin            *string
	cusip           *string
	cik             *string
	exchange        *string
	micCode         *string
	country         *string
	type_           *TypeEnum
	format          *FormatEnum
	delimiter       *string
	showPlan        *bool
	includeDelisted *bool
}

// The ticker symbol of an instrument for which data is requested
func (r ReferenceDataAPIGetStocksRequest) Symbol(symbol string) ReferenceDataAPIGetStocksRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r ReferenceDataAPIGetStocksRequest) Figi(figi string) ReferenceDataAPIGetStocksRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetStocksRequest) Isin(isin string) ReferenceDataAPIGetStocksRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r ReferenceDataAPIGetStocksRequest) Cusip(cusip string) ReferenceDataAPIGetStocksRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ReferenceDataAPIGetStocksRequest) Cik(cik string) ReferenceDataAPIGetStocksRequest {
	r.cik = &cik
	return r
}

// Filter by exchange name
func (r ReferenceDataAPIGetStocksRequest) Exchange(exchange string) ReferenceDataAPIGetStocksRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ReferenceDataAPIGetStocksRequest) MicCode(micCode string) ReferenceDataAPIGetStocksRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ReferenceDataAPIGetStocksRequest) Country(country string) ReferenceDataAPIGetStocksRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r ReferenceDataAPIGetStocksRequest) Type_(type_ TypeEnum) ReferenceDataAPIGetStocksRequest {
	r.type_ = &type_
	return r
}

// The format of the response data
func (r ReferenceDataAPIGetStocksRequest) Format(format FormatEnum) ReferenceDataAPIGetStocksRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ReferenceDataAPIGetStocksRequest) Delimiter(delimiter string) ReferenceDataAPIGetStocksRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ReferenceDataAPIGetStocksRequest) ShowPlan(showPlan bool) ReferenceDataAPIGetStocksRequest {
	r.showPlan = &showPlan
	return r
}

// Include delisted identifiers
func (r ReferenceDataAPIGetStocksRequest) IncludeDelisted(includeDelisted bool) ReferenceDataAPIGetStocksRequest {
	r.includeDelisted = &includeDelisted
	return r
}

func (r ReferenceDataAPIGetStocksRequest) Execute() (*GetStocks200Response, *http.Response, error) {
	return r.ApiService.GetStocksExecute(r)
}

/*
GetStocks Stocks

The stocks endpoint provides a daily updated list of all available stock symbols. It returns an array containing the symbols, which can be used to identify and access specific stock data across various services. This endpoint is essential for users needing to retrieve the latest stock symbol information for further data requests or integration into financial applications.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetStocksRequest
*/
func (a *ReferenceDataAPIService) GetStocks(ctx context.Context) ReferenceDataAPIGetStocksRequest {
	return ReferenceDataAPIGetStocksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetStocks200Response
func (a *ReferenceDataAPIService) GetStocksExecute(r ReferenceDataAPIGetStocksRequest) (*GetStocks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetStocks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetStocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stocks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue FormatEnum = "JSON"
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
		r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
		var defaultValue string = ";"
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
		r.delimiter = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
		r.showPlan = &defaultValue
	}
	if r.includeDelisted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", r.includeDelisted, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", defaultValue, "form", "")
		r.includeDelisted = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetSymbolSearchRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
	symbol     *string
	outputsize *int64
	showPlan   *bool
}

// Symbol to search. Supports: &lt;ul&gt; &lt;li&gt;Ticker symbol of instrument.&lt;/li&gt; &lt;li&gt;International securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section&lt;/li&gt; &lt;li&gt;The FIGI (Financial Instrument Global Identifier) parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.&lt;/li&gt; &lt;li&gt;Composite FIGI parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.&lt;/li&gt; &lt;li&gt;Share Class FIGI parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.&lt;/li&gt; &lt;/ul&gt;
func (r ReferenceDataAPIGetSymbolSearchRequest) Symbol(symbol string) ReferenceDataAPIGetSymbolSearchRequest {
	r.symbol = &symbol
	return r
}

// Number of matches in response. Max &lt;code&gt;120&lt;/code&gt;
func (r ReferenceDataAPIGetSymbolSearchRequest) Outputsize(outputsize int64) ReferenceDataAPIGetSymbolSearchRequest {
	r.outputsize = &outputsize
	return r
}

// Adds info on which plan symbol is available.
func (r ReferenceDataAPIGetSymbolSearchRequest) ShowPlan(showPlan bool) ReferenceDataAPIGetSymbolSearchRequest {
	r.showPlan = &showPlan
	return r
}

func (r ReferenceDataAPIGetSymbolSearchRequest) Execute() (*GetSymbolSearch200Response, *http.Response, error) {
	return r.ApiService.GetSymbolSearchExecute(r)
}

/*
GetSymbolSearch Symbol search

The symbol search endpoint allows users to find financial instruments by name or symbol. It returns a list of matching symbols, ordered by relevance, with the most relevant instrument first. This is useful for quickly locating specific stocks, ETFs, or other financial instruments when only partial information is available.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetSymbolSearchRequest
*/
func (a *ReferenceDataAPIService) GetSymbolSearch(ctx context.Context) ReferenceDataAPIGetSymbolSearchRequest {
	return ReferenceDataAPIGetSymbolSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSymbolSearch200Response
func (a *ReferenceDataAPIService) GetSymbolSearchExecute(r ReferenceDataAPIGetSymbolSearchRequest) (*GetSymbolSearch200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSymbolSearch200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetSymbolSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/symbol_search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
		var defaultValue int64 = 30
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
		r.outputsize = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
		r.showPlan = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ReferenceDataAPIGetTechnicalIndicatorsRequest struct {
	ctx        context.Context
	ApiService *ReferenceDataAPIService
}

func (r ReferenceDataAPIGetTechnicalIndicatorsRequest) Execute() (*GetTechnicalIndicators200Response, *http.Response, error) {
	return r.ApiService.GetTechnicalIndicatorsExecute(r)
}

/*
GetTechnicalIndicators Technical indicators

The technical indicators endpoint provides a comprehensive list of available technical indicators, each represented as an object. This endpoint is useful for developers looking to integrate a variety of technical analysis tools into their applications, allowing for streamlined access to indicator data without needing to manually configure each one.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReferenceDataAPIGetTechnicalIndicatorsRequest
*/
func (a *ReferenceDataAPIService) GetTechnicalIndicators(ctx context.Context) ReferenceDataAPIGetTechnicalIndicatorsRequest {
	return ReferenceDataAPIGetTechnicalIndicatorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTechnicalIndicators200Response
func (a *ReferenceDataAPIService) GetTechnicalIndicatorsExecute(r ReferenceDataAPIGetTechnicalIndicatorsRequest) (*GetTechnicalIndicators200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTechnicalIndicators200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetTechnicalIndicators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/technical_indicators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiBadRequestErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiUnauthorizedErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiForbiddenErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ApiParameterTooLongErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiTooManyRequestsErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiInternalServerErrorResponseBody
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

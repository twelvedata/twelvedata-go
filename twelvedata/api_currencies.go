// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CurrenciesAPIService CurrenciesAPI service
type CurrenciesAPIService service

type CurrenciesAPIGetCurrencyConversionRequest struct {
	ctx        context.Context
	ApiService *CurrenciesAPIService
	symbol     *string
	amount     *float64
	date       *string
	format     *FormatEnum
	delimiter  *string
	dp         *int64
	timezone   *string
}

// The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct
func (r CurrenciesAPIGetCurrencyConversionRequest) Symbol(symbol string) CurrenciesAPIGetCurrencyConversionRequest {
	r.symbol = &symbol
	return r
}

// Amount of base currency to be converted into quote currency. Supports values in the range from &#x60;0&#x60; and above
func (r CurrenciesAPIGetCurrencyConversionRequest) Amount(amount float64) CurrenciesAPIGetCurrencyConversionRequest {
	r.amount = &amount
	return r
}

// If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone
func (r CurrenciesAPIGetCurrencyConversionRequest) Date(date string) CurrenciesAPIGetCurrencyConversionRequest {
	r.date = &date
	return r
}

// Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60;
func (r CurrenciesAPIGetCurrencyConversionRequest) Format(format FormatEnum) CurrenciesAPIGetCurrencyConversionRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60;
func (r CurrenciesAPIGetCurrencyConversionRequest) Delimiter(delimiter string) CurrenciesAPIGetCurrencyConversionRequest {
	r.delimiter = &delimiter
	return r
}

// The number of decimal places for the data
func (r CurrenciesAPIGetCurrencyConversionRequest) Dp(dp int64) CurrenciesAPIGetCurrencyConversionRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r CurrenciesAPIGetCurrencyConversionRequest) Timezone(timezone string) CurrenciesAPIGetCurrencyConversionRequest {
	r.timezone = &timezone
	return r
}

func (r CurrenciesAPIGetCurrencyConversionRequest) Execute() (*GetCurrencyConversion200Response, *http.Response, error) {
	return r.ApiService.GetCurrencyConversionExecute(r)
}

/*
GetCurrencyConversion Currency conversion

The currency conversion endpoint provides real-time exchange rates and calculates the converted amount for specified currency pairs, including both forex and cryptocurrencies. This endpoint is useful for obtaining up-to-date conversion values between two currencies, facilitating tasks such as financial reporting, e-commerce transactions, and travel budgeting.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CurrenciesAPIGetCurrencyConversionRequest
*/
func (a *CurrenciesAPIService) GetCurrencyConversion(ctx context.Context) CurrenciesAPIGetCurrencyConversionRequest {
	return CurrenciesAPIGetCurrencyConversionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCurrencyConversion200Response
func (a *CurrenciesAPIService) GetCurrencyConversionExecute(r CurrenciesAPIGetCurrencyConversionRequest) (*GetCurrencyConversion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCurrencyConversion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrenciesAPIService.GetCurrencyConversion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/currency_conversion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}
	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
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

type CurrenciesAPIGetExchangeRateRequest struct {
	ctx        context.Context
	ApiService *CurrenciesAPIService
	symbol     *string
	date       *string
	format     *FormatEnum
	delimiter  *string
	dp         *int64
	timezone   *string
}

// The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct
func (r CurrenciesAPIGetExchangeRateRequest) Symbol(symbol string) CurrenciesAPIGetExchangeRateRequest {
	r.symbol = &symbol
	return r
}

// If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone
func (r CurrenciesAPIGetExchangeRateRequest) Date(date string) CurrenciesAPIGetExchangeRateRequest {
	r.date = &date
	return r
}

// Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60;
func (r CurrenciesAPIGetExchangeRateRequest) Format(format FormatEnum) CurrenciesAPIGetExchangeRateRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60;
func (r CurrenciesAPIGetExchangeRateRequest) Delimiter(delimiter string) CurrenciesAPIGetExchangeRateRequest {
	r.delimiter = &delimiter
	return r
}

// The number of decimal places for the data
func (r CurrenciesAPIGetExchangeRateRequest) Dp(dp int64) CurrenciesAPIGetExchangeRateRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r CurrenciesAPIGetExchangeRateRequest) Timezone(timezone string) CurrenciesAPIGetExchangeRateRequest {
	r.timezone = &timezone
	return r
}

func (r CurrenciesAPIGetExchangeRateRequest) Execute() (*GetExchangeRate200Response, *http.Response, error) {
	return r.ApiService.GetExchangeRateExecute(r)
}

/*
GetExchangeRate Exchange rate

The exchange rate endpoint provides real-time exchange rates for specified currency pairs, including both forex and cryptocurrency. It returns the current exchange rate value between two currencies, allowing users to quickly access up-to-date conversion rates for financial transactions or market analysis.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CurrenciesAPIGetExchangeRateRequest
*/
func (a *CurrenciesAPIService) GetExchangeRate(ctx context.Context) CurrenciesAPIGetExchangeRateRequest {
	return CurrenciesAPIGetExchangeRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetExchangeRate200Response
func (a *CurrenciesAPIService) GetExchangeRateExecute(r CurrenciesAPIGetExchangeRateRequest) (*GetExchangeRate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetExchangeRate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CurrenciesAPIService.GetExchangeRate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchange_rate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
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

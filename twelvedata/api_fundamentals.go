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
	"strings"
)

// FundamentalsAPIService FundamentalsAPI service
type FundamentalsAPIService service

type FundamentalsAPIGetBalanceSheetRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	period     *PeriodEnum
	startDate  *string
	endDate    *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetBalanceSheetRequest) Symbol(symbol string) FundamentalsAPIGetBalanceSheetRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetBalanceSheetRequest) Figi(figi string) FundamentalsAPIGetBalanceSheetRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetBalanceSheetRequest) Isin(isin string) FundamentalsAPIGetBalanceSheetRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetBalanceSheetRequest) Cusip(cusip string) FundamentalsAPIGetBalanceSheetRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetBalanceSheetRequest) Exchange(exchange string) FundamentalsAPIGetBalanceSheetRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetBalanceSheetRequest) MicCode(micCode string) FundamentalsAPIGetBalanceSheetRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetBalanceSheetRequest) Country(country string) FundamentalsAPIGetBalanceSheetRequest {
	r.country = &country
	return r
}

// The reporting period for the balane sheet data
func (r FundamentalsAPIGetBalanceSheetRequest) Period(period PeriodEnum) FundamentalsAPIGetBalanceSheetRequest {
	r.period = &period
	return r
}

// Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetBalanceSheetRequest) StartDate(startDate string) FundamentalsAPIGetBalanceSheetRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetBalanceSheetRequest) EndDate(endDate string) FundamentalsAPIGetBalanceSheetRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r FundamentalsAPIGetBalanceSheetRequest) Outputsize(outputsize int64) FundamentalsAPIGetBalanceSheetRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetBalanceSheetRequest) Execute() (*GetBalanceSheet200Response, *http.Response, error) {
	return r.ApiService.GetBalanceSheetExecute(r)
}

/*
GetBalanceSheet Balance sheet

The balance sheet endpoint provides a detailed financial statement for a company, outlining its assets, liabilities, and shareholders' equity. This endpoint returns structured data that includes current and non-current assets, total liabilities, and equity figures, enabling users to assess a company's financial health and stability.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetBalanceSheetRequest
*/
func (a *FundamentalsAPIService) GetBalanceSheet(ctx context.Context) FundamentalsAPIGetBalanceSheetRequest {
	return FundamentalsAPIGetBalanceSheetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBalanceSheet200Response
func (a *FundamentalsAPIService) GetBalanceSheetExecute(r FundamentalsAPIGetBalanceSheetRequest) (*GetBalanceSheet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBalanceSheet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetBalanceSheet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/balance_sheet"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetBalanceSheetConsolidatedRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	period     *PeriodEnum
	startDate  *string
	endDate    *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Symbol(symbol string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Figi(figi string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Isin(isin string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Cusip(cusip string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Exchange(exchange string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) MicCode(micCode string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Country(country string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.country = &country
	return r
}

// The reporting period for the balance sheet data.
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Period(period PeriodEnum) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.period = &period
	return r
}

// Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) StartDate(startDate string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) EndDate(endDate string) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Outputsize(outputsize int64) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetBalanceSheetConsolidatedRequest) Execute() (*GetBalanceSheetConsolidated200Response, *http.Response, error) {
	return r.ApiService.GetBalanceSheetConsolidatedExecute(r)
}

/*
GetBalanceSheetConsolidated Balance sheet consolidated

The balance sheet consolidated endpoint provides a detailed overview of a company's raw balance sheet, including a comprehensive summary of its assets, liabilities, and shareholders' equity. This endpoint is useful for retrieving financial data that reflects the overall financial position of a company, allowing users to access critical information about its financial health and structure.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetBalanceSheetConsolidatedRequest
*/
func (a *FundamentalsAPIService) GetBalanceSheetConsolidated(ctx context.Context) FundamentalsAPIGetBalanceSheetConsolidatedRequest {
	return FundamentalsAPIGetBalanceSheetConsolidatedRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBalanceSheetConsolidated200Response
func (a *FundamentalsAPIService) GetBalanceSheetConsolidatedExecute(r FundamentalsAPIGetBalanceSheetConsolidatedRequest) (*GetBalanceSheetConsolidated200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBalanceSheetConsolidated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetBalanceSheetConsolidated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/balance_sheet/consolidated"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetCashFlowRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	period     *PeriodEnum
	startDate  *string
	endDate    *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetCashFlowRequest) Symbol(symbol string) FundamentalsAPIGetCashFlowRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetCashFlowRequest) Figi(figi string) FundamentalsAPIGetCashFlowRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetCashFlowRequest) Isin(isin string) FundamentalsAPIGetCashFlowRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetCashFlowRequest) Cusip(cusip string) FundamentalsAPIGetCashFlowRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetCashFlowRequest) Exchange(exchange string) FundamentalsAPIGetCashFlowRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetCashFlowRequest) MicCode(micCode string) FundamentalsAPIGetCashFlowRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetCashFlowRequest) Country(country string) FundamentalsAPIGetCashFlowRequest {
	r.country = &country
	return r
}

// The reporting period for the cash flow statements
func (r FundamentalsAPIGetCashFlowRequest) Period(period PeriodEnum) FundamentalsAPIGetCashFlowRequest {
	r.period = &period
	return r
}

// Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetCashFlowRequest) StartDate(startDate string) FundamentalsAPIGetCashFlowRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetCashFlowRequest) EndDate(endDate string) FundamentalsAPIGetCashFlowRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r FundamentalsAPIGetCashFlowRequest) Outputsize(outputsize int64) FundamentalsAPIGetCashFlowRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetCashFlowRequest) Execute() (*GetCashFlow200Response, *http.Response, error) {
	return r.ApiService.GetCashFlowExecute(r)
}

/*
GetCashFlow Cash flow

The cash flow endpoint provides detailed information on a company's cash flow activities, including the net cash and cash equivalents moving in and out of the business. This data includes operating, investing, and financing cash flows, offering a comprehensive view of the company's liquidity and financial health.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetCashFlowRequest
*/
func (a *FundamentalsAPIService) GetCashFlow(ctx context.Context) FundamentalsAPIGetCashFlowRequest {
	return FundamentalsAPIGetCashFlowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCashFlow200Response
func (a *FundamentalsAPIService) GetCashFlowExecute(r FundamentalsAPIGetCashFlowRequest) (*GetCashFlow200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCashFlow200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetCashFlow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cash_flow"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetCashFlowConsolidatedRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	period     *PeriodEnum
	startDate  *string
	endDate    *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Symbol(symbol string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Figi(figi string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Isin(isin string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Cusip(cusip string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Exchange(exchange string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) MicCode(micCode string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Country(country string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.country = &country
	return r
}

// The reporting period for the cash flow statements
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Period(period PeriodEnum) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.period = &period
	return r
}

// Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) StartDate(startDate string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) EndDate(endDate string) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Outputsize(outputsize int64) FundamentalsAPIGetCashFlowConsolidatedRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetCashFlowConsolidatedRequest) Execute() (*GetCashFlowConsolidated200Response, *http.Response, error) {
	return r.ApiService.GetCashFlowConsolidatedExecute(r)
}

/*
GetCashFlowConsolidated Cash flow consolidated

The cash flow consolidated endpoint provides raw data on a company's consolidated cash flow, including the net cash and cash equivalents moving in and out of the business. It returns information on operating, investing, and financing activities, helping users track liquidity and financial health over a specified period.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetCashFlowConsolidatedRequest
*/
func (a *FundamentalsAPIService) GetCashFlowConsolidated(ctx context.Context) FundamentalsAPIGetCashFlowConsolidatedRequest {
	return FundamentalsAPIGetCashFlowConsolidatedRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCashFlowConsolidated200Response
func (a *FundamentalsAPIService) GetCashFlowConsolidatedExecute(r FundamentalsAPIGetCashFlowConsolidatedRequest) (*GetCashFlowConsolidated200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCashFlowConsolidated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetCashFlowConsolidated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cash_flow/consolidated"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetDividendsRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	range_     *RangeEnum
	startDate  *string
	endDate    *string
	adjust     *bool
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetDividendsRequest) Symbol(symbol string) FundamentalsAPIGetDividendsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetDividendsRequest) Figi(figi string) FundamentalsAPIGetDividendsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetDividendsRequest) Isin(isin string) FundamentalsAPIGetDividendsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetDividendsRequest) Cusip(cusip string) FundamentalsAPIGetDividendsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetDividendsRequest) Exchange(exchange string) FundamentalsAPIGetDividendsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetDividendsRequest) MicCode(micCode string) FundamentalsAPIGetDividendsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetDividendsRequest) Country(country string) FundamentalsAPIGetDividendsRequest {
	r.country = &country
	return r
}

// Specifies the time range for which to retrieve dividend data. Accepts values such as &#x60;last&#x60; (most recent dividend), &#x60;next&#x60; (upcoming dividend), &#x60;1m&#x60; - &#x60;5y&#x60; for respective periods, or &#x60;full&#x60; for all available data. If provided together with &#x60;start_date&#x60; and/or &#x60;end_date&#x60;, this parameter takes precedence.
func (r FundamentalsAPIGetDividendsRequest) Range_(range_ RangeEnum) FundamentalsAPIGetDividendsRequest {
	r.range_ = &range_
	return r
}

// Start date for the dividend data query. Only dividends with dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;. If provided together with &#x60;range&#x60; parameter, &#x60;range&#x60; will take precedence.
func (r FundamentalsAPIGetDividendsRequest) StartDate(startDate string) FundamentalsAPIGetDividendsRequest {
	r.startDate = &startDate
	return r
}

// End date for the dividend data query. Only dividends with dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;. If provided together with &#x60;range&#x60; parameter, &#x60;range&#x60; will take precedence.
func (r FundamentalsAPIGetDividendsRequest) EndDate(endDate string) FundamentalsAPIGetDividendsRequest {
	r.endDate = &endDate
	return r
}

// Specifies if there should be an adjustment
func (r FundamentalsAPIGetDividendsRequest) Adjust(adjust bool) FundamentalsAPIGetDividendsRequest {
	r.adjust = &adjust
	return r
}

func (r FundamentalsAPIGetDividendsRequest) Execute() (*GetDividends200Response, *http.Response, error) {
	return r.ApiService.GetDividendsExecute(r)
}

/*
GetDividends Dividends

The dividends endpoint provides historical dividend data for a specified stock, in many cases covering over a decade. It returns information on dividend payouts, including the amount and frequency. This endpoint is ideal for users tracking dividend histories or evaluating the income potential of stocks.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetDividendsRequest
*/
func (a *FundamentalsAPIService) GetDividends(ctx context.Context) FundamentalsAPIGetDividendsRequest {
	return FundamentalsAPIGetDividendsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDividends200Response
func (a *FundamentalsAPIService) GetDividendsExecute(r FundamentalsAPIGetDividendsRequest) (*GetDividends200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDividends200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetDividends")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dividends"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.range_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range", r.range_, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.adjust != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adjust", r.adjust, "form", "")
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

type FundamentalsAPIGetDividendsCalendarRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	startDate  *string
	endDate    *string
	outputsize *int64
	page       *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetDividendsCalendarRequest) Symbol(symbol string) FundamentalsAPIGetDividendsCalendarRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetDividendsCalendarRequest) Figi(figi string) FundamentalsAPIGetDividendsCalendarRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetDividendsCalendarRequest) Isin(isin string) FundamentalsAPIGetDividendsCalendarRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetDividendsCalendarRequest) Cusip(cusip string) FundamentalsAPIGetDividendsCalendarRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetDividendsCalendarRequest) Exchange(exchange string) FundamentalsAPIGetDividendsCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetDividendsCalendarRequest) MicCode(micCode string) FundamentalsAPIGetDividendsCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetDividendsCalendarRequest) Country(country string) FundamentalsAPIGetDividendsCalendarRequest {
	r.country = &country
	return r
}

// Start date for the dividends calendar query. Only dividends with ex-dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetDividendsCalendarRequest) StartDate(startDate string) FundamentalsAPIGetDividendsCalendarRequest {
	r.startDate = &startDate
	return r
}

// End date for the dividends calendar query. Only dividends with ex-dates on or before this date will be returned. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetDividendsCalendarRequest) EndDate(endDate string) FundamentalsAPIGetDividendsCalendarRequest {
	r.endDate = &endDate
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;500&#x60;. Default &#x60;100&#x60; when no date parameters are set, otherwise set to maximum
func (r FundamentalsAPIGetDividendsCalendarRequest) Outputsize(outputsize int64) FundamentalsAPIGetDividendsCalendarRequest {
	r.outputsize = &outputsize
	return r
}

// Page number
func (r FundamentalsAPIGetDividendsCalendarRequest) Page(page int64) FundamentalsAPIGetDividendsCalendarRequest {
	r.page = &page
	return r
}

func (r FundamentalsAPIGetDividendsCalendarRequest) Execute() ([]DividendsCalendarItem, *http.Response, error) {
	return r.ApiService.GetDividendsCalendarExecute(r)
}

/*
GetDividendsCalendar Dividends calendar

The dividends calendar endpoint provides a detailed schedule of upcoming and past dividend events for specified date ranges. By using the `start_date` and `end_date` parameters, users can retrieve a list of companies issuing dividends, including the ex-dividend date and dividend amount. This endpoint is ideal for tracking dividend payouts and planning investment strategies based on dividend schedules.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetDividendsCalendarRequest
*/
func (a *FundamentalsAPIService) GetDividendsCalendar(ctx context.Context) FundamentalsAPIGetDividendsCalendarRequest {
	return FundamentalsAPIGetDividendsCalendarRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []DividendsCalendarItem
func (a *FundamentalsAPIService) GetDividendsCalendarExecute(r FundamentalsAPIGetDividendsCalendarRequest) ([]DividendsCalendarItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DividendsCalendarItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetDividendsCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dividends_calendar"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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

type FundamentalsAPIGetEarningsRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	type_      *TypeEnum
	period     *PeriodEarningsEnum
	outputsize *int64
	format     *FormatEnum
	delimiter  *string
	startDate  *string
	endDate    *string
	dp         *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetEarningsRequest) Symbol(symbol string) FundamentalsAPIGetEarningsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetEarningsRequest) Figi(figi string) FundamentalsAPIGetEarningsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetEarningsRequest) Isin(isin string) FundamentalsAPIGetEarningsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetEarningsRequest) Cusip(cusip string) FundamentalsAPIGetEarningsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetEarningsRequest) Exchange(exchange string) FundamentalsAPIGetEarningsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetEarningsRequest) MicCode(micCode string) FundamentalsAPIGetEarningsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetEarningsRequest) Country(country string) FundamentalsAPIGetEarningsRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r FundamentalsAPIGetEarningsRequest) Type_(type_ TypeEnum) FundamentalsAPIGetEarningsRequest {
	r.type_ = &type_
	return r
}

// Type of earning, returns only 1 record. When is not empty, dates and outputsize parameters are ignored
func (r FundamentalsAPIGetEarningsRequest) Period(period PeriodEarningsEnum) FundamentalsAPIGetEarningsRequest {
	r.period = &period
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;1000&#x60;. Default &#x60;10&#x60; when no date parameters are set, otherwise set to maximum
func (r FundamentalsAPIGetEarningsRequest) Outputsize(outputsize int64) FundamentalsAPIGetEarningsRequest {
	r.outputsize = &outputsize
	return r
}

// The format of the response data
func (r FundamentalsAPIGetEarningsRequest) Format(format FormatEnum) FundamentalsAPIGetEarningsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r FundamentalsAPIGetEarningsRequest) Delimiter(delimiter string) FundamentalsAPIGetEarningsRequest {
	r.delimiter = &delimiter
	return r
}

// The date from which the data is requested. The date format is &#x60;YYYY-MM-DD&#x60;.
func (r FundamentalsAPIGetEarningsRequest) StartDate(startDate string) FundamentalsAPIGetEarningsRequest {
	r.startDate = &startDate
	return r
}

// The date to which the data is requested. The date format is &#x60;YYYY-MM-DD&#x60;.
func (r FundamentalsAPIGetEarningsRequest) EndDate(endDate string) FundamentalsAPIGetEarningsRequest {
	r.endDate = &endDate
	return r
}

// The number of decimal places in the response data. Should be in range [0,11] inclusive
func (r FundamentalsAPIGetEarningsRequest) Dp(dp int64) FundamentalsAPIGetEarningsRequest {
	r.dp = &dp
	return r
}

func (r FundamentalsAPIGetEarningsRequest) Execute() (*GetEarnings200Response, *http.Response, error) {
	return r.ApiService.GetEarningsExecute(r)
}

/*
GetEarnings Earnings

The earnings endpoint provides comprehensive earnings data for a specified company, including both the estimated and actual Earnings Per Share (EPS) figures. This endpoint delivers historical earnings information, allowing users to track a company's financial performance over time.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetEarningsRequest
*/
func (a *FundamentalsAPIService) GetEarnings(ctx context.Context) FundamentalsAPIGetEarningsRequest {
	return FundamentalsAPIGetEarningsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEarnings200Response
func (a *FundamentalsAPIService) GetEarningsExecute(r FundamentalsAPIGetEarningsRequest) (*GetEarnings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEarnings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetEarnings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/earnings"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
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

type FundamentalsAPIGetEarningsCalendarRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	exchange   *string
	micCode    *string
	country    *string
	format     *FormatEnum
	delimiter  *string
	startDate  *string
	endDate    *string
	dp         *int64
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetEarningsCalendarRequest) Exchange(exchange string) FundamentalsAPIGetEarningsCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetEarningsCalendarRequest) MicCode(micCode string) FundamentalsAPIGetEarningsCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetEarningsCalendarRequest) Country(country string) FundamentalsAPIGetEarningsCalendarRequest {
	r.country = &country
	return r
}

// Value can be JSON or CSV
func (r FundamentalsAPIGetEarningsCalendarRequest) Format(format FormatEnum) FundamentalsAPIGetEarningsCalendarRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r FundamentalsAPIGetEarningsCalendarRequest) Delimiter(delimiter string) FundamentalsAPIGetEarningsCalendarRequest {
	r.delimiter = &delimiter
	return r
}

// Can be used separately and together with end_date. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;
func (r FundamentalsAPIGetEarningsCalendarRequest) StartDate(startDate string) FundamentalsAPIGetEarningsCalendarRequest {
	r.startDate = &startDate
	return r
}

// Can be used separately and together with start_date. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;
func (r FundamentalsAPIGetEarningsCalendarRequest) EndDate(endDate string) FundamentalsAPIGetEarningsCalendarRequest {
	r.endDate = &endDate
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive
func (r FundamentalsAPIGetEarningsCalendarRequest) Dp(dp int64) FundamentalsAPIGetEarningsCalendarRequest {
	r.dp = &dp
	return r
}

func (r FundamentalsAPIGetEarningsCalendarRequest) Execute() (*GetEarningsCalendar200Response, *http.Response, error) {
	return r.ApiService.GetEarningsCalendarExecute(r)
}

/*
GetEarningsCalendar Earnings calendar

The earnings calendar endpoint provides a schedule of company earnings announcements for a specified date range. By default, it returns earnings data for the current day. Users can customize the date range using the `start_date` and `end_date` parameters to retrieve earnings information for specific periods. This endpoint is useful for tracking upcoming earnings reports and planning around key financial announcements.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetEarningsCalendarRequest
*/
func (a *FundamentalsAPIService) GetEarningsCalendar(ctx context.Context) FundamentalsAPIGetEarningsCalendarRequest {
	return FundamentalsAPIGetEarningsCalendarRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEarningsCalendar200Response
func (a *FundamentalsAPIService) GetEarningsCalendarExecute(r FundamentalsAPIGetEarningsCalendarRequest) (*GetEarningsCalendar200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEarningsCalendar200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetEarningsCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/earnings_calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
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

type FundamentalsAPIGetIncomeStatementRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	period     *PeriodEnum
	startDate  *string
	endDate    *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetIncomeStatementRequest) Symbol(symbol string) FundamentalsAPIGetIncomeStatementRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetIncomeStatementRequest) Figi(figi string) FundamentalsAPIGetIncomeStatementRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetIncomeStatementRequest) Isin(isin string) FundamentalsAPIGetIncomeStatementRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetIncomeStatementRequest) Cusip(cusip string) FundamentalsAPIGetIncomeStatementRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetIncomeStatementRequest) Exchange(exchange string) FundamentalsAPIGetIncomeStatementRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetIncomeStatementRequest) MicCode(micCode string) FundamentalsAPIGetIncomeStatementRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetIncomeStatementRequest) Country(country string) FundamentalsAPIGetIncomeStatementRequest {
	r.country = &country
	return r
}

// The reporting period for the income statement data
func (r FundamentalsAPIGetIncomeStatementRequest) Period(period PeriodEnum) FundamentalsAPIGetIncomeStatementRequest {
	r.period = &period
	return r
}

// Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetIncomeStatementRequest) StartDate(startDate string) FundamentalsAPIGetIncomeStatementRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetIncomeStatementRequest) EndDate(endDate string) FundamentalsAPIGetIncomeStatementRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r FundamentalsAPIGetIncomeStatementRequest) Outputsize(outputsize int64) FundamentalsAPIGetIncomeStatementRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetIncomeStatementRequest) Execute() (*GetIncomeStatement200Response, *http.Response, error) {
	return r.ApiService.GetIncomeStatementExecute(r)
}

/*
GetIncomeStatement Income statement

The income statement endpoint provides detailed financial data on a company's income statement, including revenues, expenses, and net income for specified periods, either annually or quarterly. This endpoint is essential for retrieving comprehensive financial performance metrics of a company, allowing users to access historical and current financial results.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetIncomeStatementRequest
*/
func (a *FundamentalsAPIService) GetIncomeStatement(ctx context.Context) FundamentalsAPIGetIncomeStatementRequest {
	return FundamentalsAPIGetIncomeStatementRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetIncomeStatement200Response
func (a *FundamentalsAPIService) GetIncomeStatementExecute(r FundamentalsAPIGetIncomeStatementRequest) (*GetIncomeStatement200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetIncomeStatement200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetIncomeStatement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/income_statement"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetIncomeStatementConsolidatedRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	period     *PeriodEnum
	startDate  *string
	endDate    *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Symbol(symbol string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Figi(figi string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Isin(isin string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Cusip(cusip string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Exchange(exchange string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) MicCode(micCode string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Country(country string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.country = &country
	return r
}

// The reporting period for the income statement data
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Period(period PeriodEnum) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.period = &period
	return r
}

// Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) StartDate(startDate string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) EndDate(endDate string) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Outputsize(outputsize int64) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetIncomeStatementConsolidatedRequest) Execute() (*GetIncomeStatementConsolidated200Response, *http.Response, error) {
	return r.ApiService.GetIncomeStatementConsolidatedExecute(r)
}

/*
GetIncomeStatementConsolidated Income statement consolidated

The income statement consolidated endpoint provides a company's raw income statement, detailing revenue, expenses, and net income for specified periods, either annually or quarterly. This data is essential for evaluating a company's financial performance over time, allowing users to access comprehensive financial results in a structured format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetIncomeStatementConsolidatedRequest
*/
func (a *FundamentalsAPIService) GetIncomeStatementConsolidated(ctx context.Context) FundamentalsAPIGetIncomeStatementConsolidatedRequest {
	return FundamentalsAPIGetIncomeStatementConsolidatedRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetIncomeStatementConsolidated200Response
func (a *FundamentalsAPIService) GetIncomeStatementConsolidatedExecute(r FundamentalsAPIGetIncomeStatementConsolidatedRequest) (*GetIncomeStatementConsolidated200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetIncomeStatementConsolidated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetIncomeStatementConsolidated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/income_statement/consolidated"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetIpoCalendarRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	exchange   *string
	micCode    *string
	country    *string
	startDate  *string
	endDate    *string
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetIpoCalendarRequest) Exchange(exchange string) FundamentalsAPIGetIpoCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetIpoCalendarRequest) MicCode(micCode string) FundamentalsAPIGetIpoCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetIpoCalendarRequest) Country(country string) FundamentalsAPIGetIpoCalendarRequest {
	r.country = &country
	return r
}

// The earliest IPO date to include in the results. Format: &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetIpoCalendarRequest) StartDate(startDate string) FundamentalsAPIGetIpoCalendarRequest {
	r.startDate = &startDate
	return r
}

// The latest IPO date to include in the results. Format: &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetIpoCalendarRequest) EndDate(endDate string) FundamentalsAPIGetIpoCalendarRequest {
	r.endDate = &endDate
	return r
}

func (r FundamentalsAPIGetIpoCalendarRequest) Execute() (*map[string][]GetIpoCalendar200ResponseValueInner, *http.Response, error) {
	return r.ApiService.GetIpoCalendarExecute(r)
}

/*
GetIpoCalendar IPO calendar

The IPO Calendar endpoint provides detailed information on initial public offerings (IPOs), including those that have occurred in the past, are happening today, or are scheduled for the future. Users can access data such as company names, IPO dates, and offering details, allowing them to track and monitor IPO activity efficiently.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetIpoCalendarRequest
*/
func (a *FundamentalsAPIService) GetIpoCalendar(ctx context.Context) FundamentalsAPIGetIpoCalendarRequest {
	return FundamentalsAPIGetIpoCalendarRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string][]GetIpoCalendar200ResponseValueInner
func (a *FundamentalsAPIService) GetIpoCalendarExecute(r FundamentalsAPIGetIpoCalendarRequest) (*map[string][]GetIpoCalendar200ResponseValueInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *map[string][]GetIpoCalendar200ResponseValueInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetIpoCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ipo_calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
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

type FundamentalsAPIGetKeyExecutivesRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetKeyExecutivesRequest) Symbol(symbol string) FundamentalsAPIGetKeyExecutivesRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetKeyExecutivesRequest) Figi(figi string) FundamentalsAPIGetKeyExecutivesRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetKeyExecutivesRequest) Isin(isin string) FundamentalsAPIGetKeyExecutivesRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetKeyExecutivesRequest) Cusip(cusip string) FundamentalsAPIGetKeyExecutivesRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetKeyExecutivesRequest) Exchange(exchange string) FundamentalsAPIGetKeyExecutivesRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetKeyExecutivesRequest) MicCode(micCode string) FundamentalsAPIGetKeyExecutivesRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetKeyExecutivesRequest) Country(country string) FundamentalsAPIGetKeyExecutivesRequest {
	r.country = &country
	return r
}

func (r FundamentalsAPIGetKeyExecutivesRequest) Execute() (*GetKeyExecutives200Response, *http.Response, error) {
	return r.ApiService.GetKeyExecutivesExecute(r)
}

/*
GetKeyExecutives Key executives

The key executives endpoint provides detailed information about a company's key executives identified by a specific stock symbol. It returns data such as names, titles, and roles of the executives, which can be useful for understanding the leadership structure of the company.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetKeyExecutivesRequest
*/
func (a *FundamentalsAPIService) GetKeyExecutives(ctx context.Context) FundamentalsAPIGetKeyExecutivesRequest {
	return FundamentalsAPIGetKeyExecutivesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetKeyExecutives200Response
func (a *FundamentalsAPIService) GetKeyExecutivesExecute(r FundamentalsAPIGetKeyExecutivesRequest) (*GetKeyExecutives200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetKeyExecutives200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetKeyExecutives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key_executives"

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

type FundamentalsAPIGetLastChangesRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	endpoint   EndpointEnum
	startDate  *string
	symbol     *string
	exchange   *string
	micCode    *string
	country    *string
	page       *int64
	outputsize *int64
}

// The starting date and time for data selection, in &#x60;2006-01-02T15:04:05&#x60; format
func (r FundamentalsAPIGetLastChangesRequest) StartDate(startDate string) FundamentalsAPIGetLastChangesRequest {
	r.startDate = &startDate
	return r
}

// Filter by symbol
func (r FundamentalsAPIGetLastChangesRequest) Symbol(symbol string) FundamentalsAPIGetLastChangesRequest {
	r.symbol = &symbol
	return r
}

// Filter by exchange name
func (r FundamentalsAPIGetLastChangesRequest) Exchange(exchange string) FundamentalsAPIGetLastChangesRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetLastChangesRequest) MicCode(micCode string) FundamentalsAPIGetLastChangesRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetLastChangesRequest) Country(country string) FundamentalsAPIGetLastChangesRequest {
	r.country = &country
	return r
}

// Page number
func (r FundamentalsAPIGetLastChangesRequest) Page(page int64) FundamentalsAPIGetLastChangesRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r FundamentalsAPIGetLastChangesRequest) Outputsize(outputsize int64) FundamentalsAPIGetLastChangesRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetLastChangesRequest) Execute() (*GetLastChanges200Response, *http.Response, error) {
	return r.ApiService.GetLastChangesExecute(r)
}

/*
GetLastChanges Last changes

The last change endpoint provides the most recent updates to fundamental data for a specified dataset. It returns a timestamp indicating when the data was last modified, allowing users to efficiently manage API requests by only fetching new data when changes occur. This helps optimize data retrieval and reduce unnecessary API credit usage.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param endpoint Endpoint name
	@return FundamentalsAPIGetLastChangesRequest
*/
func (a *FundamentalsAPIService) GetLastChanges(ctx context.Context, endpoint EndpointEnum) FundamentalsAPIGetLastChangesRequest {
	return FundamentalsAPIGetLastChangesRequest{
		ApiService: a,
		ctx:        ctx,
		endpoint:   endpoint,
	}
}

// Execute executes the request
//
//	@return GetLastChanges200Response
func (a *FundamentalsAPIService) GetLastChangesExecute(r FundamentalsAPIGetLastChangesRequest) (*GetLastChanges200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLastChanges200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetLastChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/last_change/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", url.PathEscape(parameterValueToString(r.endpoint, "endpoint")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
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
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetLogoRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	exchange   *string
	micCode    *string
	country    *string
}

// The ticker symbol of an instrument for which data is requested, e.g., &#x60;AAPL&#x60;, &#x60;BTC/USD&#x60;, &#x60;EUR/USD&#x60;.
func (r FundamentalsAPIGetLogoRequest) Symbol(symbol string) FundamentalsAPIGetLogoRequest {
	r.symbol = &symbol
	return r
}

// The exchange name where the instrument is traded, e.g., &#x60;NASDAQ&#x60;, &#x60;NSE&#x60;
func (r FundamentalsAPIGetLogoRequest) Exchange(exchange string) FundamentalsAPIGetLogoRequest {
	r.exchange = &exchange
	return r
}

// The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., &#x60;XNAS&#x60;, &#x60;XLON&#x60;
func (r FundamentalsAPIGetLogoRequest) MicCode(micCode string) FundamentalsAPIGetLogoRequest {
	r.micCode = &micCode
	return r
}

// The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetLogoRequest) Country(country string) FundamentalsAPIGetLogoRequest {
	r.country = &country
	return r
}

func (r FundamentalsAPIGetLogoRequest) Execute() (*GetLogo200Response, *http.Response, error) {
	return r.ApiService.GetLogoExecute(r)
}

/*
GetLogo Logo

The logo endpoint provides the official logo image for a specified company, cryptocurrency, or forex pair. This endpoint is useful for integrating visual branding elements into financial applications, websites, or reports, ensuring that users can easily identify and associate the correct logo with the respective financial asset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetLogoRequest
*/
func (a *FundamentalsAPIService) GetLogo(ctx context.Context) FundamentalsAPIGetLogoRequest {
	return FundamentalsAPIGetLogoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLogo200Response
func (a *FundamentalsAPIService) GetLogoExecute(r FundamentalsAPIGetLogoRequest) (*GetLogo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLogo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetLogo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logo"

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

type FundamentalsAPIGetMarketCapRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	startDate  *string
	endDate    *string
	page       *int64
	outputsize *int64
}

// Filter by symbol
func (r FundamentalsAPIGetMarketCapRequest) Symbol(symbol string) FundamentalsAPIGetMarketCapRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetMarketCapRequest) Figi(figi string) FundamentalsAPIGetMarketCapRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetMarketCapRequest) Isin(isin string) FundamentalsAPIGetMarketCapRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetMarketCapRequest) Cusip(cusip string) FundamentalsAPIGetMarketCapRequest {
	r.cusip = &cusip
	return r
}

// Filter by exchange name
func (r FundamentalsAPIGetMarketCapRequest) Exchange(exchange string) FundamentalsAPIGetMarketCapRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetMarketCapRequest) MicCode(micCode string) FundamentalsAPIGetMarketCapRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetMarketCapRequest) Country(country string) FundamentalsAPIGetMarketCapRequest {
	r.country = &country
	return r
}

// Start date for market capitalization data retrieval. Data will be returned from this date onwards. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetMarketCapRequest) StartDate(startDate string) FundamentalsAPIGetMarketCapRequest {
	r.startDate = &startDate
	return r
}

// End date for market capitalization data retrieval. Data will be returned up to and including this date. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetMarketCapRequest) EndDate(endDate string) FundamentalsAPIGetMarketCapRequest {
	r.endDate = &endDate
	return r
}

// Page number
func (r FundamentalsAPIGetMarketCapRequest) Page(page int64) FundamentalsAPIGetMarketCapRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r FundamentalsAPIGetMarketCapRequest) Outputsize(outputsize int64) FundamentalsAPIGetMarketCapRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIGetMarketCapRequest) Execute() (*GetMarketCap200Response, *http.Response, error) {
	return r.ApiService.GetMarketCapExecute(r)
}

/*
GetMarketCap Market capitalization

The Market Capitalization History endpoint provides historical data on a company's market capitalization over a specified time period. It returns a time series of market cap values, allowing users to track changes in a company's market value.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetMarketCapRequest
*/
func (a *FundamentalsAPIService) GetMarketCap(ctx context.Context) FundamentalsAPIGetMarketCapRequest {
	return FundamentalsAPIGetMarketCapRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMarketCap200Response
func (a *FundamentalsAPIService) GetMarketCapExecute(r FundamentalsAPIGetMarketCapRequest) (*GetMarketCap200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMarketCap200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetMarketCap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_cap"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

type FundamentalsAPIGetProfileRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetProfileRequest) Symbol(symbol string) FundamentalsAPIGetProfileRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetProfileRequest) Figi(figi string) FundamentalsAPIGetProfileRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetProfileRequest) Isin(isin string) FundamentalsAPIGetProfileRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetProfileRequest) Cusip(cusip string) FundamentalsAPIGetProfileRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetProfileRequest) Exchange(exchange string) FundamentalsAPIGetProfileRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetProfileRequest) MicCode(micCode string) FundamentalsAPIGetProfileRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetProfileRequest) Country(country string) FundamentalsAPIGetProfileRequest {
	r.country = &country
	return r
}

func (r FundamentalsAPIGetProfileRequest) Execute() (*GetProfile200Response, *http.Response, error) {
	return r.ApiService.GetProfileExecute(r)
}

/*
GetProfile Profile

The profile endpoint provides detailed company information, including the company's name, industry, sector, CEO, headquarters location, and market capitalization. This data is useful for obtaining a comprehensive overview of a company's business and financial standing.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetProfileRequest
*/
func (a *FundamentalsAPIService) GetProfile(ctx context.Context) FundamentalsAPIGetProfileRequest {
	return FundamentalsAPIGetProfileRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetProfile200Response
func (a *FundamentalsAPIService) GetProfileExecute(r FundamentalsAPIGetProfileRequest) (*GetProfile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetProfile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profile"

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

type FundamentalsAPIGetSplitsRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	range_     *RangeSplitsEnum
	startDate  *string
	endDate    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetSplitsRequest) Symbol(symbol string) FundamentalsAPIGetSplitsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetSplitsRequest) Figi(figi string) FundamentalsAPIGetSplitsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetSplitsRequest) Isin(isin string) FundamentalsAPIGetSplitsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetSplitsRequest) Cusip(cusip string) FundamentalsAPIGetSplitsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetSplitsRequest) Exchange(exchange string) FundamentalsAPIGetSplitsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetSplitsRequest) MicCode(micCode string) FundamentalsAPIGetSplitsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetSplitsRequest) Country(country string) FundamentalsAPIGetSplitsRequest {
	r.country = &country
	return r
}

// Range of data to be returned
func (r FundamentalsAPIGetSplitsRequest) Range_(range_ RangeSplitsEnum) FundamentalsAPIGetSplitsRequest {
	r.range_ = &range_
	return r
}

// The starting date for data selection. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetSplitsRequest) StartDate(startDate string) FundamentalsAPIGetSplitsRequest {
	r.startDate = &startDate
	return r
}

// The ending date for data selection. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetSplitsRequest) EndDate(endDate string) FundamentalsAPIGetSplitsRequest {
	r.endDate = &endDate
	return r
}

func (r FundamentalsAPIGetSplitsRequest) Execute() (*GetSplits200Response, *http.Response, error) {
	return r.ApiService.GetSplitsExecute(r)
}

/*
GetSplits Splits

The splits endpoint provides historical data on stock split events for a specified company. It returns details including the date of each split and the corresponding split factor.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetSplitsRequest
*/
func (a *FundamentalsAPIService) GetSplits(ctx context.Context) FundamentalsAPIGetSplitsRequest {
	return FundamentalsAPIGetSplitsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSplits200Response
func (a *FundamentalsAPIService) GetSplitsExecute(r FundamentalsAPIGetSplitsRequest) (*GetSplits200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSplits200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetSplits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/splits"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.range_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range", r.range_, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
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

type FundamentalsAPIGetSplitsCalendarRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	startDate  *string
	endDate    *string
	outputsize *int64
	page       *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetSplitsCalendarRequest) Symbol(symbol string) FundamentalsAPIGetSplitsCalendarRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetSplitsCalendarRequest) Figi(figi string) FundamentalsAPIGetSplitsCalendarRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetSplitsCalendarRequest) Isin(isin string) FundamentalsAPIGetSplitsCalendarRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetSplitsCalendarRequest) Cusip(cusip string) FundamentalsAPIGetSplitsCalendarRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetSplitsCalendarRequest) Exchange(exchange string) FundamentalsAPIGetSplitsCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetSplitsCalendarRequest) MicCode(micCode string) FundamentalsAPIGetSplitsCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetSplitsCalendarRequest) Country(country string) FundamentalsAPIGetSplitsCalendarRequest {
	r.country = &country
	return r
}

// The starting date (inclusive) for filtering split events in the calendar. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetSplitsCalendarRequest) StartDate(startDate string) FundamentalsAPIGetSplitsCalendarRequest {
	r.startDate = &startDate
	return r
}

// The ending date (inclusive) for filtering split events in the calendar. Format &#x60;2006-01-02&#x60;
func (r FundamentalsAPIGetSplitsCalendarRequest) EndDate(endDate string) FundamentalsAPIGetSplitsCalendarRequest {
	r.endDate = &endDate
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;500&#x60;. Default &#x60;100&#x60; when no date parameters are set, otherwise set to maximum
func (r FundamentalsAPIGetSplitsCalendarRequest) Outputsize(outputsize int64) FundamentalsAPIGetSplitsCalendarRequest {
	r.outputsize = &outputsize
	return r
}

// Page number
func (r FundamentalsAPIGetSplitsCalendarRequest) Page(page string) FundamentalsAPIGetSplitsCalendarRequest {
	r.page = &page
	return r
}

func (r FundamentalsAPIGetSplitsCalendarRequest) Execute() ([]SplitsCalendarResponseItem, *http.Response, error) {
	return r.ApiService.GetSplitsCalendarExecute(r)
}

/*
GetSplitsCalendar Splits calendar

The splits calendar endpoint provides a detailed calendar of stock split events within a specified date range. By setting the `start_date` and `end_date` parameters, users can retrieve a list of upcoming or past stock splits, including the company name, split ratio, and effective date. This endpoint is useful for tracking changes in stock structure and planning investment strategies around these events.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetSplitsCalendarRequest
*/
func (a *FundamentalsAPIService) GetSplitsCalendar(ctx context.Context) FundamentalsAPIGetSplitsCalendarRequest {
	return FundamentalsAPIGetSplitsCalendarRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []SplitsCalendarResponseItem
func (a *FundamentalsAPIService) GetSplitsCalendarExecute(r FundamentalsAPIGetSplitsCalendarRequest) ([]SplitsCalendarResponseItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []SplitsCalendarResponseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetSplitsCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/splits_calendar"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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

type FundamentalsAPIGetStatisticsRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIGetStatisticsRequest) Symbol(symbol string) FundamentalsAPIGetStatisticsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIGetStatisticsRequest) Figi(figi string) FundamentalsAPIGetStatisticsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetStatisticsRequest) Isin(isin string) FundamentalsAPIGetStatisticsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIGetStatisticsRequest) Cusip(cusip string) FundamentalsAPIGetStatisticsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIGetStatisticsRequest) Exchange(exchange string) FundamentalsAPIGetStatisticsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIGetStatisticsRequest) MicCode(micCode string) FundamentalsAPIGetStatisticsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r FundamentalsAPIGetStatisticsRequest) Country(country string) FundamentalsAPIGetStatisticsRequest {
	r.country = &country
	return r
}

func (r FundamentalsAPIGetStatisticsRequest) Execute() (*GetStatistics200Response, *http.Response, error) {
	return r.ApiService.GetStatisticsExecute(r)
}

/*
GetStatistics Statistics

The statistics endpoint provides a comprehensive snapshot of a company's key financial statistics, including valuation metrics, revenue figures, profit margins, and other essential financial data. This endpoint is ideal for users seeking detailed insights into a company's financial health and performance metrics.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIGetStatisticsRequest
*/
func (a *FundamentalsAPIService) GetStatistics(ctx context.Context) FundamentalsAPIGetStatisticsRequest {
	return FundamentalsAPIGetStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetStatistics200Response
func (a *FundamentalsAPIService) GetStatisticsExecute(r FundamentalsAPIGetStatisticsRequest) (*GetStatistics200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetStatistics200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/statistics"

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

type FundamentalsAPIPressReleasesListParametersRequest struct {
	ctx        context.Context
	ApiService *FundamentalsAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	startDate  *string
	endDate    *string
	timezone   *string
	language   *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r FundamentalsAPIPressReleasesListParametersRequest) Symbol(symbol string) FundamentalsAPIPressReleasesListParametersRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r FundamentalsAPIPressReleasesListParametersRequest) Figi(figi string) FundamentalsAPIPressReleasesListParametersRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIPressReleasesListParametersRequest) Isin(isin string) FundamentalsAPIPressReleasesListParametersRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r FundamentalsAPIPressReleasesListParametersRequest) Cusip(cusip string) FundamentalsAPIPressReleasesListParametersRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r FundamentalsAPIPressReleasesListParametersRequest) Exchange(exchange string) FundamentalsAPIPressReleasesListParametersRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r FundamentalsAPIPressReleasesListParametersRequest) MicCode(micCode string) FundamentalsAPIPressReleasesListParametersRequest {
	r.micCode = &micCode
	return r
}

// Begin date for filtering items. Returns press releases with release date on or after this date. Format &#x60;2025-12-24T02:07:00&#x60;
func (r FundamentalsAPIPressReleasesListParametersRequest) StartDate(startDate string) FundamentalsAPIPressReleasesListParametersRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering items. Returns press releases with release date on or before this date. Format &#x60;2025-12-24T02:07:00&#x60;
func (r FundamentalsAPIPressReleasesListParametersRequest) EndDate(endDate string) FundamentalsAPIPressReleasesListParametersRequest {
	r.endDate = &endDate
	return r
}

// Time zone for date filtering. Default is the identifier time zone.
func (r FundamentalsAPIPressReleasesListParametersRequest) Timezone(timezone string) FundamentalsAPIPressReleasesListParametersRequest {
	r.timezone = &timezone
	return r
}

// Comma-separated list of languages to filter press releases by language.
func (r FundamentalsAPIPressReleasesListParametersRequest) Language(language string) FundamentalsAPIPressReleasesListParametersRequest {
	r.language = &language
	return r
}

// Number of latest press releases returned. Only used if no data range is specified. Maximum value is &#x60;10&#x60;.  type: number
func (r FundamentalsAPIPressReleasesListParametersRequest) Outputsize(outputsize int64) FundamentalsAPIPressReleasesListParametersRequest {
	r.outputsize = &outputsize
	return r
}

func (r FundamentalsAPIPressReleasesListParametersRequest) Execute() (*PressReleasesListParameters200Response, *http.Response, error) {
	return r.ApiService.PressReleasesListParametersExecute(r)
}

/*
PressReleasesListParameters Press releases

The press releases endpoint offers structured, real-time access to official company press releases and corporate announcements from public entities across global markets.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FundamentalsAPIPressReleasesListParametersRequest
*/
func (a *FundamentalsAPIService) PressReleasesListParameters(ctx context.Context) FundamentalsAPIPressReleasesListParametersRequest {
	return FundamentalsAPIPressReleasesListParametersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PressReleasesListParameters200Response
func (a *FundamentalsAPIService) PressReleasesListParametersExecute(r FundamentalsAPIPressReleasesListParametersRequest) (*PressReleasesListParameters200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PressReleasesListParameters200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.PressReleasesListParameters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/press_releases"

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
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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

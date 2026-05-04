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

// RegulatoryAPIService RegulatoryAPI service
type RegulatoryAPIService service

type RegulatoryAPIGetDirectHoldersRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r RegulatoryAPIGetDirectHoldersRequest) Symbol(symbol string) RegulatoryAPIGetDirectHoldersRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r RegulatoryAPIGetDirectHoldersRequest) Figi(figi string) RegulatoryAPIGetDirectHoldersRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetDirectHoldersRequest) Isin(isin string) RegulatoryAPIGetDirectHoldersRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetDirectHoldersRequest) Cusip(cusip string) RegulatoryAPIGetDirectHoldersRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r RegulatoryAPIGetDirectHoldersRequest) Exchange(exchange string) RegulatoryAPIGetDirectHoldersRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r RegulatoryAPIGetDirectHoldersRequest) MicCode(micCode string) RegulatoryAPIGetDirectHoldersRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r RegulatoryAPIGetDirectHoldersRequest) Country(country string) RegulatoryAPIGetDirectHoldersRequest {
	r.country = &country
	return r
}

func (r RegulatoryAPIGetDirectHoldersRequest) Execute() (*GetDirectHolders200Response, *http.Response, error) {
	return r.ApiService.GetDirectHoldersExecute(r)
}

/*
GetDirectHolders Direct holders

The direct holders endpoint provides detailed information about the number of shares directly held by individuals or entities as recorded in a company's official share registry. This data is essential for understanding the distribution of stock ownership within a company, helping users identify major shareholders and assess shareholder concentration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegulatoryAPIGetDirectHoldersRequest
*/
func (a *RegulatoryAPIService) GetDirectHolders(ctx context.Context) RegulatoryAPIGetDirectHoldersRequest {
	return RegulatoryAPIGetDirectHoldersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDirectHolders200Response
func (a *RegulatoryAPIService) GetDirectHoldersExecute(r RegulatoryAPIGetDirectHoldersRequest) (*GetDirectHolders200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDirectHolders200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetDirectHolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/direct_holders"

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

type RegulatoryAPIGetEdgarFilingsArchiveRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	formType   *string
	filledFrom *string
	filledTo   *string
	page       *int64
	pageSize   *int64
}

// The ticker symbol of an instrument for which data is requested
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Symbol(symbol string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Figi(figi string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Isin(isin string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Cusip(cusip string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.cusip = &cusip
	return r
}

// Filter by exchange name
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Exchange(exchange string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) MicCode(micCode string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Country(country string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.country = &country
	return r
}

// Filter by form types, example &#x60;8-K&#x60;, &#x60;EX-1.1&#x60;
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) FormType(formType string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.formType = &formType
	return r
}

// Filter by filled time from
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) FilledFrom(filledFrom string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.filledFrom = &filledFrom
	return r
}

// Filter by filled time to
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) FilledTo(filledTo string) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.filledTo = &filledTo
	return r
}

// Page number
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Page(page int64) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) PageSize(pageSize int64) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	r.pageSize = &pageSize
	return r
}

func (r RegulatoryAPIGetEdgarFilingsArchiveRequest) Execute() (*GetEdgarFilingsArchive200Response, *http.Response, error) {
	return r.ApiService.GetEdgarFilingsArchiveExecute(r)
}

/*
GetEdgarFilingsArchive EDGAR fillings

The EDGAR fillings endpoint provides access to a comprehensive collection of financial documents submitted to the SEC, including real-time and historical forms, filings, and exhibits. Users can retrieve detailed information about company disclosures, financial statements, and regulatory submissions, enabling them to access essential compliance and financial data directly from the SEC's EDGAR system.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegulatoryAPIGetEdgarFilingsArchiveRequest
*/
func (a *RegulatoryAPIService) GetEdgarFilingsArchive(ctx context.Context) RegulatoryAPIGetEdgarFilingsArchiveRequest {
	return RegulatoryAPIGetEdgarFilingsArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEdgarFilingsArchive200Response
func (a *RegulatoryAPIService) GetEdgarFilingsArchiveExecute(r RegulatoryAPIGetEdgarFilingsArchiveRequest) (*GetEdgarFilingsArchive200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEdgarFilingsArchive200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetEdgarFilingsArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edgar_filings/archive"

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
	if r.formType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "form_type", r.formType, "form", "")
	}
	if r.filledFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filled_from", r.filledFrom, "form", "")
	}
	if r.filledTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filled_to", r.filledTo, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
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

type RegulatoryAPIGetFundHoldersRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r RegulatoryAPIGetFundHoldersRequest) Symbol(symbol string) RegulatoryAPIGetFundHoldersRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r RegulatoryAPIGetFundHoldersRequest) Figi(figi string) RegulatoryAPIGetFundHoldersRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetFundHoldersRequest) Isin(isin string) RegulatoryAPIGetFundHoldersRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetFundHoldersRequest) Cusip(cusip string) RegulatoryAPIGetFundHoldersRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r RegulatoryAPIGetFundHoldersRequest) Exchange(exchange string) RegulatoryAPIGetFundHoldersRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r RegulatoryAPIGetFundHoldersRequest) MicCode(micCode string) RegulatoryAPIGetFundHoldersRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r RegulatoryAPIGetFundHoldersRequest) Country(country string) RegulatoryAPIGetFundHoldersRequest {
	r.country = &country
	return r
}

func (r RegulatoryAPIGetFundHoldersRequest) Execute() (*GetFundHolders200Response, *http.Response, error) {
	return r.ApiService.GetFundHoldersExecute(r)
}

/*
GetFundHolders Fund holders

The fund holders endpoint provides detailed information about the proportion of a company's stock that is owned by mutual fund holders. It returns data on the number of shares held, the percentage of total shares outstanding, and the names of the mutual funds involved. This endpoint is useful for users looking to understand mutual fund investment in a specific company.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegulatoryAPIGetFundHoldersRequest
*/
func (a *RegulatoryAPIService) GetFundHolders(ctx context.Context) RegulatoryAPIGetFundHoldersRequest {
	return RegulatoryAPIGetFundHoldersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFundHolders200Response
func (a *RegulatoryAPIService) GetFundHoldersExecute(r RegulatoryAPIGetFundHoldersRequest) (*GetFundHolders200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetFundHolders200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetFundHolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fund_holders"

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

type RegulatoryAPIGetInsiderTransactionsRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// The ticker symbol of an instrument for which data is requested, e.g., &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r RegulatoryAPIGetInsiderTransactionsRequest) Symbol(symbol string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r RegulatoryAPIGetInsiderTransactionsRequest) Figi(figi string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetInsiderTransactionsRequest) Isin(isin string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetInsiderTransactionsRequest) Cusip(cusip string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded, e.g., &#x60;Nasdaq&#x60;, &#x60;NSE&#x60;
func (r RegulatoryAPIGetInsiderTransactionsRequest) Exchange(exchange string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r RegulatoryAPIGetInsiderTransactionsRequest) MicCode(micCode string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., United States or US.
func (r RegulatoryAPIGetInsiderTransactionsRequest) Country(country string) RegulatoryAPIGetInsiderTransactionsRequest {
	r.country = &country
	return r
}

func (r RegulatoryAPIGetInsiderTransactionsRequest) Execute() (*GetInsiderTransactions200Response, *http.Response, error) {
	return r.ApiService.GetInsiderTransactionsExecute(r)
}

/*
GetInsiderTransactions Insider transaction

The insider transaction endpoint provides detailed data on trades executed by company insiders, such as executives and directors. It returns information including the insider's name, their role, the transaction type, the number of shares, the transaction date, and the price per share. This endpoint is useful for tracking insider activity and understanding potential insider sentiment towards a company's stock.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegulatoryAPIGetInsiderTransactionsRequest
*/
func (a *RegulatoryAPIService) GetInsiderTransactions(ctx context.Context) RegulatoryAPIGetInsiderTransactionsRequest {
	return RegulatoryAPIGetInsiderTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetInsiderTransactions200Response
func (a *RegulatoryAPIService) GetInsiderTransactionsExecute(r RegulatoryAPIGetInsiderTransactionsRequest) (*GetInsiderTransactions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetInsiderTransactions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetInsiderTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/insider_transactions"

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

type RegulatoryAPIGetInstitutionalHoldersRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r RegulatoryAPIGetInstitutionalHoldersRequest) Symbol(symbol string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r RegulatoryAPIGetInstitutionalHoldersRequest) Figi(figi string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetInstitutionalHoldersRequest) Isin(isin string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetInstitutionalHoldersRequest) Cusip(cusip string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r RegulatoryAPIGetInstitutionalHoldersRequest) Exchange(exchange string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r RegulatoryAPIGetInstitutionalHoldersRequest) MicCode(micCode string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r RegulatoryAPIGetInstitutionalHoldersRequest) Country(country string) RegulatoryAPIGetInstitutionalHoldersRequest {
	r.country = &country
	return r
}

func (r RegulatoryAPIGetInstitutionalHoldersRequest) Execute() (*GetInstitutionalHolders200Response, *http.Response, error) {
	return r.ApiService.GetInstitutionalHoldersExecute(r)
}

/*
GetInstitutionalHolders Institutional holders

The institutional holders endpoint provides detailed information on the percentage and amount of a company's stock owned by institutional investors, such as pension funds, insurance companies, and investment firms. This data is essential for understanding the influence and involvement of large entities in a company's ownership structure.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegulatoryAPIGetInstitutionalHoldersRequest
*/
func (a *RegulatoryAPIService) GetInstitutionalHolders(ctx context.Context) RegulatoryAPIGetInstitutionalHoldersRequest {
	return RegulatoryAPIGetInstitutionalHoldersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetInstitutionalHolders200Response
func (a *RegulatoryAPIService) GetInstitutionalHoldersExecute(r RegulatoryAPIGetInstitutionalHoldersRequest) (*GetInstitutionalHolders200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetInstitutionalHolders200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetInstitutionalHolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/institutional_holders"

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

type RegulatoryAPIGetSourceSanctionedEntitiesRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	source     SourceEnum
}

func (r RegulatoryAPIGetSourceSanctionedEntitiesRequest) Execute() (*GetSourceSanctionedEntities200Response, *http.Response, error) {
	return r.ApiService.GetSourceSanctionedEntitiesExecute(r)
}

/*
GetSourceSanctionedEntities Sanctioned entities

The sanctions entities endpoint provides a comprehensive list of entities sanctioned by a specified authority, such as OFAC, UK, EU, or AU. Users can retrieve detailed information about individuals, organizations, and other entities subject to sanctions from the chosen source, facilitating compliance and risk management processes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param source Sanctions source
	@return RegulatoryAPIGetSourceSanctionedEntitiesRequest
*/
func (a *RegulatoryAPIService) GetSourceSanctionedEntities(ctx context.Context, source SourceEnum) RegulatoryAPIGetSourceSanctionedEntitiesRequest {
	return RegulatoryAPIGetSourceSanctionedEntitiesRequest{
		ApiService: a,
		ctx:        ctx,
		source:     source,
	}
}

// Execute executes the request
//
//	@return GetSourceSanctionedEntities200Response
func (a *RegulatoryAPIService) GetSourceSanctionedEntitiesExecute(r RegulatoryAPIGetSourceSanctionedEntitiesRequest) (*GetSourceSanctionedEntities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSourceSanctionedEntities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetSourceSanctionedEntities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sanctions/{source}"
	localVarPath = strings.Replace(localVarPath, "{"+"source"+"}", url.PathEscape(parameterValueToString(r.source, "source")), -1)

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

type RegulatoryAPIGetTaxInfoRequest struct {
	ctx        context.Context
	ApiService *RegulatoryAPIService
	symbol     *string
	isin       *string
	figi       *string
	cusip      *string
	exchange   *string
	micCode    *string
}

// The ticker symbol of an instrument for which data is requested, e.g., &#x60;SKYQ&#x60;, &#x60;AIRE&#x60;, &#x60;ALM:BME&#x60;, &#x60;HSI:HKEX&#x60;.
func (r RegulatoryAPIGetTaxInfoRequest) Symbol(symbol string) RegulatoryAPIGetTaxInfoRequest {
	r.symbol = &symbol
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetTaxInfoRequest) Isin(isin string) RegulatoryAPIGetTaxInfoRequest {
	r.isin = &isin
	return r
}

// The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r RegulatoryAPIGetTaxInfoRequest) Figi(figi string) RegulatoryAPIGetTaxInfoRequest {
	r.figi = &figi
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r RegulatoryAPIGetTaxInfoRequest) Cusip(cusip string) RegulatoryAPIGetTaxInfoRequest {
	r.cusip = &cusip
	return r
}

// The exchange name where the instrument is traded, e.g., &#x60;Nasdaq&#x60;, &#x60;Euronext&#x60;
func (r RegulatoryAPIGetTaxInfoRequest) Exchange(exchange string) RegulatoryAPIGetTaxInfoRequest {
	r.exchange = &exchange
	return r
}

// The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., &#x60;XNAS&#x60;, &#x60;XLON&#x60;
func (r RegulatoryAPIGetTaxInfoRequest) MicCode(micCode string) RegulatoryAPIGetTaxInfoRequest {
	r.micCode = &micCode
	return r
}

func (r RegulatoryAPIGetTaxInfoRequest) Execute() (*GetTaxInfo200Response, *http.Response, error) {
	return r.ApiService.GetTaxInfoExecute(r)
}

/*
GetTaxInfo Tax information

The tax information endpoint provides detailed tax-related data for a specified financial instrument, including applicable tax rates and relevant tax codes. This information is essential for users needing to understand the tax implications associated with trading or investing in specific instruments.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegulatoryAPIGetTaxInfoRequest
*/
func (a *RegulatoryAPIService) GetTaxInfo(ctx context.Context) RegulatoryAPIGetTaxInfoRequest {
	return RegulatoryAPIGetTaxInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTaxInfo200Response
func (a *RegulatoryAPIService) GetTaxInfoExecute(r RegulatoryAPIGetTaxInfoRequest) (*GetTaxInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTaxInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegulatoryAPIService.GetTaxInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
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

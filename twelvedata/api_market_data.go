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

// MarketDataAPIService MarketDataAPI service
type MarketDataAPIService service

type MarketDataAPIGetEodRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	type_      *TypeEnum
	date       *string
	prepost    *bool
	dp         *int64
}

// Symbol ticker of the instrument
func (r MarketDataAPIGetEodRequest) Symbol(symbol string) MarketDataAPIGetEodRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r MarketDataAPIGetEodRequest) Figi(figi string) MarketDataAPIGetEodRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetEodRequest) Isin(isin string) MarketDataAPIGetEodRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetEodRequest) Cusip(cusip string) MarketDataAPIGetEodRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r MarketDataAPIGetEodRequest) Exchange(exchange string) MarketDataAPIGetEodRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r MarketDataAPIGetEodRequest) MicCode(micCode string) MarketDataAPIGetEodRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r MarketDataAPIGetEodRequest) Country(country string) MarketDataAPIGetEodRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r MarketDataAPIGetEodRequest) Type_(type_ TypeEnum) MarketDataAPIGetEodRequest {
	r.type_ = &type_
	return r
}

// If not null, then return data from a specific date
func (r MarketDataAPIGetEodRequest) Date(date string) MarketDataAPIGetEodRequest {
	r.date = &date
	return r
}

// Parameter is optional. Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume
func (r MarketDataAPIGetEodRequest) Prepost(prepost bool) MarketDataAPIGetEodRequest {
	r.prepost = &prepost
	return r
}

// Specifies the number of decimal places for floating values Should be in range [0,11] inclusive
func (r MarketDataAPIGetEodRequest) Dp(dp int64) MarketDataAPIGetEodRequest {
	r.dp = &dp
	return r
}

func (r MarketDataAPIGetEodRequest) Execute() (*GetEod200Response, *http.Response, error) {
	return r.ApiService.GetEodExecute(r)
}

/*
GetEod End of day price

The End of Day (EOD) Prices endpoint provides the closing price and other relevant metadata for a financial instrument at the end of a trading day. This endpoint is useful for retrieving daily historical data for stocks, ETFs, or other securities, allowing users to track performance over time and compare daily market movements.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MarketDataAPIGetEodRequest
*/
func (a *MarketDataAPIService) GetEod(ctx context.Context) MarketDataAPIGetEodRequest {
	return MarketDataAPIGetEodRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEod200Response
func (a *MarketDataAPIService) GetEodExecute(r MarketDataAPIGetEodRequest) (*GetEod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetEod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eod"

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
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
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

type MarketDataAPIGetMarketMoversRequest struct {
	ctx              context.Context
	ApiService       *MarketDataAPIService
	market           MarketEnum
	direction        *DirectionEnum
	outputsize       *int64
	country          *string
	priceGreaterThan *string
	dp               *string
}

// Specifies direction of the snapshot gainers or losers
func (r MarketDataAPIGetMarketMoversRequest) Direction(direction DirectionEnum) MarketDataAPIGetMarketMoversRequest {
	r.direction = &direction
	return r
}

// Specifies the size of the snapshot. Can be in a range from &#x60;1&#x60; to &#x60;50&#x60;
func (r MarketDataAPIGetMarketMoversRequest) Outputsize(outputsize int64) MarketDataAPIGetMarketMoversRequest {
	r.outputsize = &outputsize
	return r
}

// Country of the snapshot, applicable to non-currencies only. Takes country name or alpha code
func (r MarketDataAPIGetMarketMoversRequest) Country(country string) MarketDataAPIGetMarketMoversRequest {
	r.country = &country
	return r
}

// Takes values with price grater than specified value
func (r MarketDataAPIGetMarketMoversRequest) PriceGreaterThan(priceGreaterThan string) MarketDataAPIGetMarketMoversRequest {
	r.priceGreaterThan = &priceGreaterThan
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive
func (r MarketDataAPIGetMarketMoversRequest) Dp(dp string) MarketDataAPIGetMarketMoversRequest {
	r.dp = &dp
	return r
}

func (r MarketDataAPIGetMarketMoversRequest) Execute() (*MarketMoversResponseBody, *http.Response, error) {
	return r.ApiService.GetMarketMoversExecute(r)
}

/*
GetMarketMovers Market movers

The market movers endpoint provides a ranked list of the top-gaining and losing assets for the current trading day. It returns detailed data on the highest percentage price increases and decreases since the previous day's close. This endpoint supports international equities, forex, and cryptocurrencies, enabling users to quickly identify significant market movements across various asset classes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param market Maket type
	@return MarketDataAPIGetMarketMoversRequest
*/
func (a *MarketDataAPIService) GetMarketMovers(ctx context.Context, market MarketEnum) MarketDataAPIGetMarketMoversRequest {
	return MarketDataAPIGetMarketMoversRequest{
		ApiService: a,
		ctx:        ctx,
		market:     market,
	}
}

// Execute executes the request
//
//	@return MarketMoversResponseBody
func (a *MarketDataAPIService) GetMarketMoversExecute(r MarketDataAPIGetMarketMoversRequest) (*MarketMoversResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MarketMoversResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetMarketMovers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_movers/{market}"
	localVarPath = strings.Replace(localVarPath, "{"+"market"+"}", url.PathEscape(parameterValueToString(r.market, "market")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.priceGreaterThan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "price_greater_than", r.priceGreaterThan, "form", "")
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

type MarketDataAPIGetPriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	figi       *string
	isin       *string
	cusip      *string
	exchange   *string
	micCode    *string
	country    *string
	type_      *TypeEnum
	format     *FormatEnum
	delimiter  *string
	prepost    *bool
	dp         *int64
}

// Symbol ticker of the instrument
func (r MarketDataAPIGetPriceRequest) Symbol(symbol string) MarketDataAPIGetPriceRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r MarketDataAPIGetPriceRequest) Figi(figi string) MarketDataAPIGetPriceRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetPriceRequest) Isin(isin string) MarketDataAPIGetPriceRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetPriceRequest) Cusip(cusip string) MarketDataAPIGetPriceRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r MarketDataAPIGetPriceRequest) Exchange(exchange string) MarketDataAPIGetPriceRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r MarketDataAPIGetPriceRequest) MicCode(micCode string) MarketDataAPIGetPriceRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r MarketDataAPIGetPriceRequest) Country(country string) MarketDataAPIGetPriceRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r MarketDataAPIGetPriceRequest) Type_(type_ TypeEnum) MarketDataAPIGetPriceRequest {
	r.type_ = &type_
	return r
}

// Value can be JSON or CSV
func (r MarketDataAPIGetPriceRequest) Format(format FormatEnum) MarketDataAPIGetPriceRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r MarketDataAPIGetPriceRequest) Delimiter(delimiter string) MarketDataAPIGetPriceRequest {
	r.delimiter = &delimiter
	return r
}

// Parameter is optional. Only for Pro or Venture, and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume.
func (r MarketDataAPIGetPriceRequest) Prepost(prepost bool) MarketDataAPIGetPriceRequest {
	r.prepost = &prepost
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive
func (r MarketDataAPIGetPriceRequest) Dp(dp int64) MarketDataAPIGetPriceRequest {
	r.dp = &dp
	return r
}

func (r MarketDataAPIGetPriceRequest) Execute() (*GetPrice200Response, *http.Response, error) {
	return r.ApiService.GetPriceExecute(r)
}

/*
GetPrice Latest price

The latest price endpoint provides the latest market price for a specified financial instrument. It returns a single data point representing the current (or the most recently available) trading price.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MarketDataAPIGetPriceRequest
*/
func (a *MarketDataAPIService) GetPrice(ctx context.Context) MarketDataAPIGetPriceRequest {
	return MarketDataAPIGetPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPrice200Response
func (a *MarketDataAPIService) GetPriceExecute(r MarketDataAPIGetPriceRequest) (*GetPrice200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPrice200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetPrice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price"

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
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
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

type MarketDataAPIGetQuoteRequest struct {
	ctx              context.Context
	ApiService       *MarketDataAPIService
	symbol           *string
	figi             *string
	isin             *string
	cusip            *string
	interval         *IntervalEnum
	exchange         *string
	micCode          *string
	country          *string
	volumeTimePeriod *int64
	type_            *TypeEnum
	format           *FormatEnum
	delimiter        *string
	prepost          *bool
	eod              *bool
	rollingPeriod    *int64
	dp               *int64
	timezone         *string
}

// Symbol ticker of the instrument
func (r MarketDataAPIGetQuoteRequest) Symbol(symbol string) MarketDataAPIGetQuoteRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r MarketDataAPIGetQuoteRequest) Figi(figi string) MarketDataAPIGetQuoteRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetQuoteRequest) Isin(isin string) MarketDataAPIGetQuoteRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetQuoteRequest) Cusip(cusip string) MarketDataAPIGetQuoteRequest {
	r.cusip = &cusip
	return r
}

// Interval of the quote
func (r MarketDataAPIGetQuoteRequest) Interval(interval IntervalEnum) MarketDataAPIGetQuoteRequest {
	r.interval = &interval
	return r
}

// Exchange where instrument is traded
func (r MarketDataAPIGetQuoteRequest) Exchange(exchange string) MarketDataAPIGetQuoteRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r MarketDataAPIGetQuoteRequest) MicCode(micCode string) MarketDataAPIGetQuoteRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r MarketDataAPIGetQuoteRequest) Country(country string) MarketDataAPIGetQuoteRequest {
	r.country = &country
	return r
}

// Number of periods for Average Volume
func (r MarketDataAPIGetQuoteRequest) VolumeTimePeriod(volumeTimePeriod int64) MarketDataAPIGetQuoteRequest {
	r.volumeTimePeriod = &volumeTimePeriod
	return r
}

// The asset class to which the instrument belongs
func (r MarketDataAPIGetQuoteRequest) Type_(type_ TypeEnum) MarketDataAPIGetQuoteRequest {
	r.type_ = &type_
	return r
}

// Value can be JSON or CSV Default JSON
func (r MarketDataAPIGetQuoteRequest) Format(format FormatEnum) MarketDataAPIGetQuoteRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r MarketDataAPIGetQuoteRequest) Delimiter(delimiter string) MarketDataAPIGetQuoteRequest {
	r.delimiter = &delimiter
	return r
}

// Parameter is optional. Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume.
func (r MarketDataAPIGetQuoteRequest) Prepost(prepost bool) MarketDataAPIGetQuoteRequest {
	r.prepost = &prepost
	return r
}

// If true, then return data for closed day
func (r MarketDataAPIGetQuoteRequest) Eod(eod bool) MarketDataAPIGetQuoteRequest {
	r.eod = &eod
	return r
}

// Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168].
func (r MarketDataAPIGetQuoteRequest) RollingPeriod(rollingPeriod int64) MarketDataAPIGetQuoteRequest {
	r.rollingPeriod = &rollingPeriod
	return r
}

// Specifies the number of decimal places for floating values Should be in range [0,11] inclusive
func (r MarketDataAPIGetQuoteRequest) Dp(dp int64) MarketDataAPIGetQuoteRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.&lt;/p&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r MarketDataAPIGetQuoteRequest) Timezone(timezone string) MarketDataAPIGetQuoteRequest {
	r.timezone = &timezone
	return r
}

func (r MarketDataAPIGetQuoteRequest) Execute() (*GetQuote200Response, *http.Response, error) {
	return r.ApiService.GetQuoteExecute(r)
}

/*
GetQuote Quote

The quote endpoint provides real-time data for a selected financial instrument, returning essential information such as the latest price, open, high, low, close, volume, and price change. This endpoint is ideal for users needing up-to-date market data to track price movements and trading activity for specific stocks, ETFs, or other securities.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MarketDataAPIGetQuoteRequest
*/
func (a *MarketDataAPIService) GetQuote(ctx context.Context) MarketDataAPIGetQuoteRequest {
	return MarketDataAPIGetQuoteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetQuote200Response
func (a *MarketDataAPIService) GetQuoteExecute(r MarketDataAPIGetQuoteRequest) (*GetQuote200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetQuote200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetQuote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/quote"

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
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
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
	if r.volumeTimePeriod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume_time_period", r.volumeTimePeriod, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
	}
	if r.eod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eod", r.eod, "form", "")
	}
	if r.rollingPeriod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rolling_period", r.rollingPeriod, "form", "")
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

type MarketDataAPIGetTimeSeriesRequest struct {
	ctx           context.Context
	ApiService    *MarketDataAPIService
	interval      *IntervalEnum
	symbol        *string
	isin          *string
	figi          *string
	cusip         *string
	outputsize    *int64
	exchange      *string
	micCode       *string
	country       *string
	type_         *TypeEnum
	timezone      *string
	startDate     *string
	endDate       *string
	date          *string
	order         *OrderEnum
	prepost       *bool
	format        *FormatEnum
	delimiter     *string
	dp            *int64
	previousClose *bool
	adjust        *AdjustEnum
}

// Interval between two consecutive points in time series
func (r MarketDataAPIGetTimeSeriesRequest) Interval(interval IntervalEnum) MarketDataAPIGetTimeSeriesRequest {
	r.interval = &interval
	return r
}

// Symbol ticker of the instrument. E.g. &#x60;AAPL&#x60;, &#x60;EUR/USD&#x60;, &#x60;ETH/BTC&#x60;, ...
func (r MarketDataAPIGetTimeSeriesRequest) Symbol(symbol string) MarketDataAPIGetTimeSeriesRequest {
	r.symbol = &symbol
	return r
}

// Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetTimeSeriesRequest) Isin(isin string) MarketDataAPIGetTimeSeriesRequest {
	r.isin = &isin
	return r
}

// The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.
func (r MarketDataAPIGetTimeSeriesRequest) Figi(figi string) MarketDataAPIGetTimeSeriesRequest {
	r.figi = &figi
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section
func (r MarketDataAPIGetTimeSeriesRequest) Cusip(cusip string) MarketDataAPIGetTimeSeriesRequest {
	r.cusip = &cusip
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;5000&#x60;. Default &#x60;30&#x60; when no date parameters are set, otherwise set to maximum
func (r MarketDataAPIGetTimeSeriesRequest) Outputsize(outputsize int64) MarketDataAPIGetTimeSeriesRequest {
	r.outputsize = &outputsize
	return r
}

// Exchange where instrument is traded
func (r MarketDataAPIGetTimeSeriesRequest) Exchange(exchange string) MarketDataAPIGetTimeSeriesRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r MarketDataAPIGetTimeSeriesRequest) MicCode(micCode string) MarketDataAPIGetTimeSeriesRequest {
	r.micCode = &micCode
	return r
}

// The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r MarketDataAPIGetTimeSeriesRequest) Country(country string) MarketDataAPIGetTimeSeriesRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r MarketDataAPIGetTimeSeriesRequest) Type_(type_ TypeEnum) MarketDataAPIGetTimeSeriesRequest {
	r.type_ = &type_
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.&lt;/p&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r MarketDataAPIGetTimeSeriesRequest) Timezone(timezone string) MarketDataAPIGetTimeSeriesRequest {
	r.timezone = &timezone
	return r
}

// Can be used separately and together with &#x60;end_date&#x60;. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;  Default location: &lt;ul&gt; &lt;li&gt;Forex and Cryptocurrencies - &lt;code&gt;UTC&lt;/code&gt;&lt;/li&gt; &lt;li&gt;Stocks - where exchange is located (e.g. for AAPL it will be &lt;code&gt;America/New_York&lt;/code&gt;)&lt;/li&gt; &lt;/ul&gt; Both parameters take into account if &lt;code&gt;timezone&lt;/code&gt; parameter is provided.&lt;br/&gt; If &lt;code&gt;timezone&lt;/code&gt; is given then, &lt;code&gt;start_date&lt;/code&gt; and &lt;code&gt;end_date&lt;/code&gt; will be used in the specified location  Examples: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;&amp;symbol&#x3D;AAPL&amp;start_date&#x3D;2019-08-09T15:50:00&amp;…&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 New York time up to current date&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;&amp;symbol&#x3D;EUR/USD&amp;timezone&#x3D;Asia/Singapore&amp;start_date&#x3D;2019-08-09T15:50:00&amp;…&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date&lt;/li&gt; &lt;li&gt;3. &lt;code&gt;&amp;symbol&#x3D;ETH/BTC&amp;timezone&#x3D;Europe/Zurich&amp;start_date&#x3D;2019-08-09T15:50:00&amp;end_date&#x3D;2019-08-09T15:55:00&amp;...&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00&lt;/li&gt; &lt;/ul&gt;
func (r MarketDataAPIGetTimeSeriesRequest) StartDate(startDate string) MarketDataAPIGetTimeSeriesRequest {
	r.startDate = &startDate
	return r
}

// The ending date and time for data selection, see &#x60;start_date&#x60; description for details.
func (r MarketDataAPIGetTimeSeriesRequest) EndDate(endDate string) MarketDataAPIGetTimeSeriesRequest {
	r.endDate = &endDate
	return r
}

// Specifies the exact date to get the data for. Could be the exact date, e.g. &#x60;2021-10-27&#x60;, or in human language &#x60;today&#x60; or &#x60;yesterday&#x60;
func (r MarketDataAPIGetTimeSeriesRequest) Date(date string) MarketDataAPIGetTimeSeriesRequest {
	r.date = &date
	return r
}

// Sorting order of the output
func (r MarketDataAPIGetTimeSeriesRequest) Order(order OrderEnum) MarketDataAPIGetTimeSeriesRequest {
	r.order = &order
	return r
}

// Returns quotes that include pre-market and post-market data. Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume
func (r MarketDataAPIGetTimeSeriesRequest) Prepost(prepost bool) MarketDataAPIGetTimeSeriesRequest {
	r.prepost = &prepost
	return r
}

// The format of the response data
func (r MarketDataAPIGetTimeSeriesRequest) Format(format FormatEnum) MarketDataAPIGetTimeSeriesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r MarketDataAPIGetTimeSeriesRequest) Delimiter(delimiter string) MarketDataAPIGetTimeSeriesRequest {
	r.delimiter = &delimiter
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided
func (r MarketDataAPIGetTimeSeriesRequest) Dp(dp int64) MarketDataAPIGetTimeSeriesRequest {
	r.dp = &dp
	return r
}

// A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object
func (r MarketDataAPIGetTimeSeriesRequest) PreviousClose(previousClose bool) MarketDataAPIGetTimeSeriesRequest {
	r.previousClose = &previousClose
	return r
}

// Adjusting mode for prices
func (r MarketDataAPIGetTimeSeriesRequest) Adjust(adjust AdjustEnum) MarketDataAPIGetTimeSeriesRequest {
	r.adjust = &adjust
	return r
}

func (r MarketDataAPIGetTimeSeriesRequest) Execute() (*GetTimeSeries200Response, *http.Response, error) {
	return r.ApiService.GetTimeSeriesExecute(r)
}

/*
GetTimeSeries Time series

The time series endpoint provides detailed historical data for a specified financial instrument. It returns two main components: metadata, which includes essential information about the instrument, and a time series dataset. The time series consists of chronological entries with Open, High, Low, and Close prices, and for applicable instruments, it also includes trading volume. This endpoint is ideal for retrieving comprehensive historical price data for analysis or visualization purposes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MarketDataAPIGetTimeSeriesRequest
*/
func (a *MarketDataAPIService) GetTimeSeries(ctx context.Context) MarketDataAPIGetTimeSeriesRequest {
	return MarketDataAPIGetTimeSeriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTimeSeries200Response
func (a *MarketDataAPIService) GetTimeSeriesExecute(r MarketDataAPIGetTimeSeriesRequest) (*GetTimeSeries200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTimeSeries200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetTimeSeries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time_series"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
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
	if r.previousClose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "previous_close", r.previousClose, "form", "")
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

type MarketDataAPIGetTimeSeriesCrossRequest struct {
	ctx           context.Context
	ApiService    *MarketDataAPIService
	base          *string
	quote         *string
	interval      *IntervalEnum
	baseType      *string
	baseExchange  *string
	baseMicCode   *string
	quoteType     *string
	quoteExchange *string
	quoteMicCode  *string
	outputsize    *int64
	format        *FormatEnum
	delimiter     *string
	prepost       *bool
	startDate     *string
	endDate       *string
	adjust        *bool
	dp            *int64
	timezone      *string
}

// Base currency symbol
func (r MarketDataAPIGetTimeSeriesCrossRequest) Base(base string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.base = &base
	return r
}

// Quote currency symbol
func (r MarketDataAPIGetTimeSeriesCrossRequest) Quote(quote string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.quote = &quote
	return r
}

// Interval between two consecutive points in time series
func (r MarketDataAPIGetTimeSeriesCrossRequest) Interval(interval IntervalEnum) MarketDataAPIGetTimeSeriesCrossRequest {
	r.interval = &interval
	return r
}

// Base instrument type according to the &#x60;/instrument_type&#x60; endpoint
func (r MarketDataAPIGetTimeSeriesCrossRequest) BaseType(baseType string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.baseType = &baseType
	return r
}

// Base exchange
func (r MarketDataAPIGetTimeSeriesCrossRequest) BaseExchange(baseExchange string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.baseExchange = &baseExchange
	return r
}

// Base MIC code
func (r MarketDataAPIGetTimeSeriesCrossRequest) BaseMicCode(baseMicCode string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.baseMicCode = &baseMicCode
	return r
}

// Quote instrument type according to the &#x60;/instrument_type&#x60; endpoint
func (r MarketDataAPIGetTimeSeriesCrossRequest) QuoteType(quoteType string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.quoteType = &quoteType
	return r
}

// Quote exchange
func (r MarketDataAPIGetTimeSeriesCrossRequest) QuoteExchange(quoteExchange string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.quoteExchange = &quoteExchange
	return r
}

// Quote MIC code
func (r MarketDataAPIGetTimeSeriesCrossRequest) QuoteMicCode(quoteMicCode string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.quoteMicCode = &quoteMicCode
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;5000&#x60;. Default &#x60;30&#x60; when no date parameters are set, otherwise set to maximum
func (r MarketDataAPIGetTimeSeriesCrossRequest) Outputsize(outputsize int64) MarketDataAPIGetTimeSeriesCrossRequest {
	r.outputsize = &outputsize
	return r
}

// Format of the response data
func (r MarketDataAPIGetTimeSeriesCrossRequest) Format(format FormatEnum) MarketDataAPIGetTimeSeriesCrossRequest {
	r.format = &format
	return r
}

// Delimiter used in CSV file
func (r MarketDataAPIGetTimeSeriesCrossRequest) Delimiter(delimiter string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.delimiter = &delimiter
	return r
}

// Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume.
func (r MarketDataAPIGetTimeSeriesCrossRequest) Prepost(prepost bool) MarketDataAPIGetTimeSeriesCrossRequest {
	r.prepost = &prepost
	return r
}

// Start date for the time series data
func (r MarketDataAPIGetTimeSeriesCrossRequest) StartDate(startDate string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.startDate = &startDate
	return r
}

// End date for the time series data
func (r MarketDataAPIGetTimeSeriesCrossRequest) EndDate(endDate string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.endDate = &endDate
	return r
}

// Specifies if there should be an adjustment
func (r MarketDataAPIGetTimeSeriesCrossRequest) Adjust(adjust bool) MarketDataAPIGetTimeSeriesCrossRequest {
	r.adjust = &adjust
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive.
func (r MarketDataAPIGetTimeSeriesCrossRequest) Dp(dp int64) MarketDataAPIGetTimeSeriesCrossRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r MarketDataAPIGetTimeSeriesCrossRequest) Timezone(timezone string) MarketDataAPIGetTimeSeriesCrossRequest {
	r.timezone = &timezone
	return r
}

func (r MarketDataAPIGetTimeSeriesCrossRequest) Execute() (*GetTimeSeriesCross200Response, *http.Response, error) {
	return r.ApiService.GetTimeSeriesCrossExecute(r)
}

/*
GetTimeSeriesCross Time series cross

The Time Series Cross endpoint calculates and returns historical cross-rate data for exotic forex pairs, cryptocurrencies, or stocks (e.g., Apple Inc. price in Indian Rupees) on the fly. It provides metadata about the requested symbol and a time series array with Open, High, Low, and Close prices, sorted descending by time, enabling analysis of price history and market trends.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MarketDataAPIGetTimeSeriesCrossRequest
*/
func (a *MarketDataAPIService) GetTimeSeriesCross(ctx context.Context) MarketDataAPIGetTimeSeriesCrossRequest {
	return MarketDataAPIGetTimeSeriesCrossRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTimeSeriesCross200Response
func (a *MarketDataAPIService) GetTimeSeriesCrossExecute(r MarketDataAPIGetTimeSeriesCrossRequest) (*GetTimeSeriesCross200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTimeSeriesCross200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetTimeSeriesCross")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time_series/cross"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.base == nil {
		return localVarReturnValue, nil, reportError("base is required and must be specified")
	}
	if r.quote == nil {
		return localVarReturnValue, nil, reportError("quote is required and must be specified")
	}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "base", r.base, "form", "")
	if r.baseType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_type", r.baseType, "form", "")
	}
	if r.baseExchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_exchange", r.baseExchange, "form", "")
	}
	if r.baseMicCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_mic_code", r.baseMicCode, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "quote", r.quote, "form", "")
	if r.quoteType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_type", r.quoteType, "form", "")
	}
	if r.quoteExchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_exchange", r.quoteExchange, "form", "")
	}
	if r.quoteMicCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_mic_code", r.quoteMicCode, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
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

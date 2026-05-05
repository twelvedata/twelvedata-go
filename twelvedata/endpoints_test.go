//go:build integration

package twelvedata

import (
	"context"
	"testing"
	"time"
)

const (
	symbolStock      = "AAPL"
	symbolEtf        = "SPY"
	symbolForex      = "EUR/USD"
	symbolCrypto     = "BTC/USD"
	symbolMutualFund = "VFINX"
	startDate        = "2025-01-01"
	endDate          = "2025-01-31"
	timezone         = "UTC"
	currencyAmount   = float64(100)
	outputsize       = int64(10)
	delayMs          = 100 * time.Millisecond
)

func setupClient(t *testing.T) (*APIClient, context.Context) {
	t.Helper()
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { time.Sleep(delayMs) })
	return NewAPIClient(cfg), context.Background()
}

// --- MarketDataAPI ---

func TestMarketDataGetTimeSeries(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MarketDataAPI.GetTimeSeries(ctx).
		Symbol(symbolStock).
		Interval(INTERVALENUM__1DAY).
		Outputsize(outputsize).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

func TestMarketDataGetExchangeRate(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.CurrenciesAPI.GetExchangeRate(ctx).
		Symbol(symbolForex).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Rate == 0 {
		t.Fatal("expected non-zero rate")
	}
	t.Logf("%+v", resp)
}

func TestMarketDataGetPrice(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MarketDataAPI.GetPrice(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Price == "" {
		t.Fatal("expected non-empty price")
	}
	t.Logf("%+v", resp)
}

func TestMarketDataGetQuote(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MarketDataAPI.GetQuote(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Symbol == "" {
		t.Fatal("expected non-empty symbol")
	}
	t.Logf("%+v", resp)
}

func TestMarketDataGetEod(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MarketDataAPI.GetEod(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Close == "" {
		t.Fatal("expected non-empty close")
	}
	t.Logf("%+v", resp)
}

func TestMarketDataGetCurrencyConversion(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.CurrenciesAPI.GetCurrencyConversion(ctx).
		Symbol(symbolForex).
		Amount(currencyAmount).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Rate == 0 {
		t.Fatal("expected non-zero rate")
	}
	t.Logf("%+v", resp)
}

func TestMarketDataGetMarketMovers(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MarketDataAPI.GetMarketMovers(ctx, MARKETENUM_STOCKS).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

func TestMarketDataGetTimeSeriesCross(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MarketDataAPI.GetTimeSeriesCross(ctx).
		Base("EUR").
		Quote("USD").
		Interval(INTERVALENUM__1DAY).
		Outputsize(outputsize).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

// --- FundamentalsAPI ---

func TestFundamentalsGetIpoCalendar(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetIpoCalendar(ctx).
		StartDate(startDate).
		EndDate(endDate).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil || len(*resp) == 0 {
		t.Fatal("expected non-empty response data")
	}
	t.Logf("%+v", *resp)
}

func TestFundamentalsGetEarningsCalendar(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetEarningsCalendar(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if resp.Earnings == nil {
		t.Fatal("expected response data")
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetDividendsCalendar(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetDividendsCalendar(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetSplitsCalendar(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetSplitsCalendar(ctx).
		StartDate(startDate).
		EndDate(endDate).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetEarnings(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetEarnings(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Earnings) == 0 {
		t.Fatal("expected earnings data")
	}
	t.Logf("%+v", resp.Earnings[0])
}

func TestFundamentalsGetDividends(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetDividends(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetSplits(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetSplits(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetBalanceSheet(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetBalanceSheet(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetBalanceSheetConsolidated(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetBalanceSheetConsolidated(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetCashFlow(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetCashFlow(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetCashFlowConsolidated(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetCashFlowConsolidated(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetIncomeStatement(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetIncomeStatement(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetIncomeStatementConsolidated(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetIncomeStatementConsolidated(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetProfile(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetProfile(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Symbol == "" {
		t.Fatal("expected non-empty symbol")
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetStatistics(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetStatistics(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp.Statistics)
}

func TestFundamentalsGetMarketCap(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetMarketCap(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetKeyExecutives(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetKeyExecutives(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetLogo(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetLogo(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsPressReleasesListParameters(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.PressReleasesListParameters(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp)
}

func TestFundamentalsGetLastChanges(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.FundamentalsAPI.GetLastChanges(ctx, ENDPOINTENUM_PROFILE).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

// --- AnalysisAPI ---

func TestAnalysisGetAnalystRatingsLight(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetAnalystRatingsLight(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp)
}

func TestAnalysisGetAnalystRatingsUsEquities(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetAnalystRatingsUsEquities(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp)
}

func TestAnalysisGetPriceTarget(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetPriceTarget(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.PriceTarget)
}

func TestAnalysisGetRecommendations(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetRecommendations(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Trends)
}

func TestAnalysisGetEarningsEstimate(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetEarningsEstimate(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.EarningsEstimate) == 0 {
		t.Fatal("expected earnings_estimate data")
	}
	t.Logf("%+v", resp.EarningsEstimate[0])
}

func TestAnalysisGetEpsRevisions(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetEpsRevisions(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.EpsRevision) == 0 {
		t.Fatal("expected eps_revision data")
	}
	t.Logf("%+v", resp.EpsRevision[0])
}

func TestAnalysisGetEpsTrend(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetEpsTrend(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.EpsTrend) == 0 {
		t.Fatal("expected eps_trend data")
	}
	t.Logf("%+v", resp.EpsTrend[0])
}

func TestAnalysisGetGrowthEstimates(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetGrowthEstimates(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if resp.GrowthEstimates == nil {
		t.Fatal("expected growth_estimates data")
	}
	t.Logf("%+v", resp.GrowthEstimates)
}

func TestAnalysisGetRevenueEstimate(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AnalysisAPI.GetRevenueEstimate(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.RevenueEstimate) == 0 {
		t.Fatal("expected revenue_estimate data")
	}
	t.Logf("%+v", resp.RevenueEstimate[0])
}

// --- EtfsAPI ---

func TestEtfsGetETFsList(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsList(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Result)
}

func TestEtfsGetETFsType(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsType(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
	t.Logf("%+v", resp.Result)
}

func TestEtfsGetETFsFamily(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsFamily(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
	t.Logf("%+v", resp.Result)
}

func TestEtfsGetETFsWorld(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsWorld(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Etf)
}

func TestEtfsGetETFsWorldSummary(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsWorldSummary(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Etf)
}

func TestEtfsGetETFsWorldComposition(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsWorldComposition(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Etf)
}

func TestEtfsGetETFsWorldPerformance(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsWorldPerformance(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Etf)
}

func TestEtfsGetETFsWorldRisk(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.EtfsAPI.GetETFsWorldRisk(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Etf)
}

// --- MutualFundsAPI ---

func TestMutualFundsGetMutualFundsList(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsList(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Result)
}

func TestMutualFundsGetMutualFundsType(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsType(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
	t.Logf("%+v", resp.Result)
}

func TestMutualFundsGetMutualFundsFamily(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsFamily(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
	t.Logf("%+v", resp.Result)
}

func TestMutualFundsGetMutualFundsWorld(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorld(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldSummary(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldSummary(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldComposition(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldComposition(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldPerformance(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldPerformance(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldRisk(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldRisk(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldRatings(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldRatings(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldPurchaseInfo(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldPurchaseInfo(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

func TestMutualFundsGetMutualFundsWorldSustainability(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.MutualFundsAPI.GetMutualFundsWorldSustainability(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.MutualFund)
}

// --- ReferenceDataAPI ---

func TestReferenceDataGetStocks(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetStocks(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetCryptocurrencies(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetCryptocurrencies(ctx).
		Symbol(symbolCrypto).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetCryptocurrencyExchanges(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetCryptocurrencyExchanges(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetForexPairs(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetForexPairs(ctx).
		Symbol(symbolForex).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetExchanges(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetExchanges(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetExchangeSchedule(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetExchangeSchedule(ctx).
		MicCode("XNYS").
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetCountries(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetCountries(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetIntervals(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetIntervals(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data)
}

func TestReferenceDataGetMarketState(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetMarketState(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp) == 0 {
		t.Fatal("expected non-empty array")
	}
	t.Logf("%+v", resp[0])
}

func TestReferenceDataGetEarliestTimestamp(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetEarliestTimestamp(ctx).
		Symbol(symbolStock).
		Interval(INTERVALENUM__1DAY).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Datetime == "" {
		t.Fatal("expected non-empty datetime")
	}
	t.Logf("%+v", resp)
}

func TestReferenceDataGetSymbolSearch(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetSymbolSearch(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetInstrumentType(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetInstrumentType(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Result) == 0 {
		t.Fatal("expected non-empty result")
	}
	t.Logf("%+v", resp.Result)
}

func TestReferenceDataGetCrossListings(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetCrossListings(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result == nil {
		t.Fatal("expected non-nil result")
	}
	t.Logf("%+v", resp.Result)
}

func TestReferenceDataGetBonds(t *testing.T) {
	// TODO: uncomment this when the API response content-type is fixed
	t.Skip("skipping GetBonds test due to incorrect API response content-type")
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetBonds(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Result)
}

func TestReferenceDataGetCommodities(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetCommodities(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetEtf(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetEtf(ctx).
		Symbol(symbolEtf).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", resp.Data[0])
}

func TestReferenceDataGetFunds(t *testing.T) {
	// TODO: uncomment this when the API response content-type is fixed
	t.Skip("skipping GetFunds test due to incorrect API response content-type")
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetFunds(ctx).
		Symbol(symbolMutualFund).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Result)
}

func TestReferenceDataGetTechnicalIndicators(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.ReferenceDataAPI.GetTechnicalIndicators(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Fatal("expected non-empty data")
	}
	t.Logf("%+v", len(resp.Data))
}

// --- RegulatoryAPI ---

func TestRegulatoryGetDirectHolders(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetDirectHolders(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestRegulatoryGetFundHolders(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetFundHolders(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestRegulatoryGetInstitutionalHolders(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetInstitutionalHolders(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestRegulatoryGetInsiderTransactions(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetInsiderTransactions(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

func TestRegulatoryGetTaxInfo(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetTaxInfo(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	t.Logf("%+v", resp.Data)
}

func TestRegulatoryGetSourceSanctionedEntities(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetSourceSanctionedEntities(ctx, SOURCEENUM_OFAC).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Sanctions) == 0 {
		t.Fatal("expected non-empty sanctions")
	}
	t.Logf("%+v", resp.Sanctions[0])
}

func TestRegulatoryGetEdgarFilingsArchive(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.RegulatoryAPI.GetEdgarFilingsArchive(ctx).
		Symbol(symbolStock).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}

// --- TechnicalIndicatorAPI (RSI + MACD only) ---

func TestTechnicalIndicatorGetTimeSeriesRsi(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.TechnicalIndicatorAPI.GetTimeSeriesRsi(ctx).
		Symbol(symbolStock).
		Interval(INTERVALENUM__1DAY).
		Outputsize(outputsize).
		Timezone(timezone).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

func TestTechnicalIndicatorGetTimeSeriesMacd(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.TechnicalIndicatorAPI.GetTimeSeriesMacd(ctx).
		Symbol(symbolStock).
		Interval(INTERVALENUM__1DAY).
		Outputsize(outputsize).
		Timezone(timezone).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "ok" {
		t.Fatalf("status=%q", resp.Status)
	}
	if len(resp.Values) == 0 {
		t.Fatal("empty values")
	}
	t.Logf("%+v", resp.Values[0])
}

// --- AdvancedAPI ---

func TestAdvancedGetApiUsage(t *testing.T) {
	client, ctx := setupClient(t)
	resp, _, err := client.AdvancedAPI.GetApiUsage(ctx).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Timestamp == "" {
		t.Fatal("expected non-empty timestamp")
	}
	t.Logf("%+v", resp)
}

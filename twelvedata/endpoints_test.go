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

func TestEndpoints(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	client := NewAPIClient(cfg)
	ctx := context.Background()

	runTest := func(name string, fn func(t *testing.T)) {
		t.Run(name, fn)
		time.Sleep(delayMs)
	}

	// --- MarketDataAPI ---

	runTest("MarketData.GetTimeSeries", func(t *testing.T) {
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
	})

	runTest("MarketData.GetExchangeRate", func(t *testing.T) {
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
	})

	runTest("MarketData.GetPrice", func(t *testing.T) {
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
	})

	runTest("MarketData.GetQuote", func(t *testing.T) {
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
	})

	runTest("MarketData.GetEod", func(t *testing.T) {
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
	})

	runTest("MarketData.GetCurrencyConversion", func(t *testing.T) {
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
	})

	runTest("MarketData.GetMarketMovers", func(t *testing.T) {
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
	})

	runTest("MarketData.GetTimeSeriesCross", func(t *testing.T) {
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
	})

	// --- FundamentalsAPI ---

	runTest("Fundamentals.GetIpoCalendar", func(t *testing.T) {
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
	})

	runTest("Fundamentals.GetEarningsCalendar", func(t *testing.T) {
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
	})

	runTest("Fundamentals.GetDividendsCalendar", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetDividendsCalendar(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetSplitsCalendar", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetSplitsCalendar(ctx).
			StartDate(startDate).
			EndDate(endDate).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetEarnings", func(t *testing.T) {
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
	})

	runTest("Fundamentals.GetDividends", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetDividends(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetSplits", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetSplits(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetBalanceSheet", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetBalanceSheet(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetBalanceSheetConsolidated", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetBalanceSheetConsolidated(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetCashFlow", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetCashFlow(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetCashFlowConsolidated", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetCashFlowConsolidated(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetIncomeStatement", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetIncomeStatement(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetIncomeStatementConsolidated", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetIncomeStatementConsolidated(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetProfile", func(t *testing.T) {
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
	})

	runTest("Fundamentals.GetStatistics", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetStatistics(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp.Statistics)
	})

	runTest("Fundamentals.GetMarketCap", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetMarketCap(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetKeyExecutives", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetKeyExecutives(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.GetLogo", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetLogo(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Fundamentals.PressReleasesListParameters", func(t *testing.T) {
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
	})

	runTest("Fundamentals.GetLastChanges", func(t *testing.T) {
		resp, _, err := client.FundamentalsAPI.GetLastChanges(ctx, ENDPOINTENUM_PROFILE).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	// --- AnalysisAPI ---

	runTest("Analysis.GetAnalystRatingsLight", func(t *testing.T) {
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
	})

	runTest("Analysis.GetAnalystRatingsUsEquities", func(t *testing.T) {
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
	})

	runTest("Analysis.GetPriceTarget", func(t *testing.T) {
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
	})

	runTest("Analysis.GetRecommendations", func(t *testing.T) {
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
	})

	runTest("Analysis.GetEarningsEstimate", func(t *testing.T) {
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
	})

	runTest("Analysis.GetEpsRevisions", func(t *testing.T) {
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
	})

	runTest("Analysis.GetEpsTrend", func(t *testing.T) {
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
	})

	runTest("Analysis.GetGrowthEstimates", func(t *testing.T) {
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
	})

	runTest("Analysis.GetRevenueEstimate", func(t *testing.T) {
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
	})

	// --- EtfsAPI ---

	runTest("Etfs.GetETFsList", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsType", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsFamily", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsWorld", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsWorldSummary", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsWorldComposition", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsWorldPerformance", func(t *testing.T) {
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
	})

	runTest("Etfs.GetETFsWorldRisk", func(t *testing.T) {
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
	})

	// --- MutualFundsAPI ---

	runTest("MutualFunds.GetMutualFundsList", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsType", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsFamily", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorld", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldSummary", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldComposition", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldPerformance", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldRisk", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldRatings", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldPurchaseInfo", func(t *testing.T) {
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
	})

	runTest("MutualFunds.GetMutualFundsWorldSustainability", func(t *testing.T) {
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
	})

	// --- ReferenceDataAPI ---

	runTest("ReferenceData.GetStocks", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetCryptocurrencies", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetCryptocurrencyExchanges", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetForexPairs", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetExchanges", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetExchangeSchedule", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetCountries", func(t *testing.T) {
		resp, _, err := client.ReferenceDataAPI.GetCountries(ctx).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.Data) == 0 {
			t.Fatal("expected non-empty data")
		}
		t.Logf("%+v", resp.Data[0])
	})

	runTest("ReferenceData.GetIntervals", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetMarketState", func(t *testing.T) {
		resp, _, err := client.ReferenceDataAPI.GetMarketState(ctx).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		if len(resp) == 0 {
			t.Fatal("expected non-empty array")
		}
		t.Logf("%+v", resp[0])
	})

	runTest("ReferenceData.GetEarliestTimestamp", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetSymbolSearch", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetInstrumentType", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetCrossListings", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetBonds", func(t *testing.T) {
		// TODO: uncomment this when the API response content-type is fixed
		t.Skip("skipping GetBonds test due to incorrect API response content-type")
		resp, _, err := client.ReferenceDataAPI.GetBonds(ctx).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		if resp.Status != "ok" {
			t.Fatalf("status=%q", resp.Status)
		}
		t.Logf("%+v", resp.Result)
	})

	runTest("ReferenceData.GetCommodities", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetEtf", func(t *testing.T) {
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
	})

	runTest("ReferenceData.GetFunds", func(t *testing.T) {
		// TODO: uncomment this when the API response content-type is fixed
		t.Skip("skipping GetFunds test due to incorrect API response content-type")
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
	})

	runTest("ReferenceData.GetTechnicalIndicators", func(t *testing.T) {
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
	})

	// --- RegulatoryAPI ---

	runTest("Regulatory.GetDirectHolders", func(t *testing.T) {
		resp, _, err := client.RegulatoryAPI.GetDirectHolders(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Regulatory.GetFundHolders", func(t *testing.T) {
		resp, _, err := client.RegulatoryAPI.GetFundHolders(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Regulatory.GetInstitutionalHolders", func(t *testing.T) {
		resp, _, err := client.RegulatoryAPI.GetInstitutionalHolders(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Regulatory.GetInsiderTransactions", func(t *testing.T) {
		resp, _, err := client.RegulatoryAPI.GetInsiderTransactions(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	runTest("Regulatory.GetTaxInfo", func(t *testing.T) {
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
	})

	runTest("Regulatory.GetSourceSanctionedEntities", func(t *testing.T) {
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
	})

	runTest("Regulatory.GetEdgarFilingsArchive", func(t *testing.T) {
		resp, _, err := client.RegulatoryAPI.GetEdgarFilingsArchive(ctx).
			Symbol(symbolStock).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", resp)
	})

	// --- TechnicalIndicatorAPI (RSI + MACD only) ---

	runTest("TechnicalIndicator.GetTimeSeriesRsi", func(t *testing.T) {
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
	})

	runTest("TechnicalIndicator.GetTimeSeriesMacd", func(t *testing.T) {
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
	})

	// --- AdvancedAPI ---

	runTest("Advanced.GetApiUsage", func(t *testing.T) {
		resp, _, err := client.AdvancedAPI.GetApiUsage(ctx).
			Execute()
		if err != nil {
			t.Fatal(err)
		}
		if resp.Timestamp == "" {
			t.Fatal("expected non-empty timestamp")
		}
		t.Logf("%+v", resp)
	})
}

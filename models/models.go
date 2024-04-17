package models

type AccountNumbers []struct {
	AccountNumber string `json:"accountNumber"`
	AccountHash   string `json:"accountHash"`
}

type SecuritiesAccounts []struct {
	SecuritiesAccount SecuritiesAccount
}

type SecuritiesAccount struct {
	Type                    string            `json:"type"`
	AccountNumber           string            `json:"accountNumber"`
	RoundTrips              int               `json:"roundTrips"`
	IsDayTrader             bool              `json:"isDayTrader"`
	IsClosingOnlyRestricted bool              `json:"isClosingOnlyRestricted"`
	PfcbFlag                bool              `json:"pfcbFlag"`
	Positions               []Position        `json:"positions"`
	InitialBalances         InitialBalances   `json:"initialBalances"`
	CurrentBalances         CurrentBalances   `json:"currentBalances"`
	ProjectedBalances       ProjectedBalances `json:"projectedBalances"`
}

type Position struct {
	ShortQuantity                  float64            `json:"shortQuantity"`
	AveragePrice                   float64            `json:"averagePrice"`
	CurrentDayProfitLoss           float64            `json:"currentDayProfitLoss"`
	CurrentDayProfitLossPercentage float64            `json:"currentDayProfitLossPercentage"`
	LongQuantity                   float64            `json:"longQuantity"`
	SettledLongQuantity            float64            `json:"settledLongQuantity"`
	SettledShortQuantity           float64            `json:"settledShortQuantity"`
	AgedQuantity                   float64            `json:"agedQuantity"`
	Instrument                     PositionInstrument `json:"instrument"`
	MarketValue                    float64            `json:"marketValue"`
	MaintenanceRequirement         float64            `json:"maintenanceRequirement"`
	AverageLongPrice               float64            `json:"averageLongPrice"`
	AverageShortPrice              float64            `json:"averageShortPrice"`
	TaxLotAverageLongPrice         float64            `json:"taxLotAverageLongPrice"`
	TaxLotAverageShortPrice        float64            `json:"taxLotAverageShortPrice"`
	LongOpenProfitLoss             float64            `json:"longOpenProfitLoss"`
	ShortOpenProfitLoss            float64            `json:"shortOpenProfitLoss"`
	PreviousSessionLongQuantity    float64            `json:"previousSessionLongQuantity"`
	PreviousSessionShortQuantity   float64            `json:"previousSessionShortQuantity"`
	CurrentDayCost                 float64            `json:"currentDayCost"`
}

type PositionInstrument struct {
	AssetType        string  `json:"assetType"`
	Cusip            string  `json:"cusip"`
	Symbol           string  `json:"symbol"`
	Description      string  `json:"description"`
	NetChange        float64 `json:"netChange"`
	Type             string  `json:"type"`
	PutCall          string  `json:"putCall"`
	UnderlyingSymbol string  `json:"underlyingSymbol"`
}

type InitialBalances struct {
	AccruedInterest                  float64 `json:"accruedInterest"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BondValue                        float64 `json:"bondValue"`
	BuyingPower                      float64 `json:"buyingPower"`
	CashBalance                      float64 `json:"cashBalance"`
	CashAvailableForTrading          float64 `json:"cashAvailableForTrading"`
	CashReceipts                     float64 `json:"cashReceipts"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower"`
	DayTradingBuyingPowerCall        float64 `json:"dayTradingBuyingPowerCall"`
	DayTradingEquityCall             float64 `json:"dayTradingEquityCall"`
	Equity                           float64 `json:"equity"`
	EquityPercentage                 float64 `json:"equityPercentage"`
	LiquidationValue                 float64 `json:"liquidationValue"`
	LongMarginValue                  float64 `json:"longMarginValue"`
	LongOptionMarketValue            float64 `json:"longOptionMarketValue"`
	LongStockValue                   float64 `json:"longStockValue"`
	MaintenanceCall                  float64 `json:"maintenanceCall"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement"`
	Margin                           float64 `json:"margin"`
	MarginEquity                     float64 `json:"marginEquity"`
	MoneyMarketFund                  float64 `json:"moneyMarketFund"`
	MutualFundValue                  float64 `json:"mutualFundValue"`
	RegTCall                         float64 `json:"regTCall"`
	ShortMarginValue                 float64 `json:"shortMarginValue"`
	ShortOptionMarketValue           float64 `json:"shortOptionMarketValue"`
	ShortStockValue                  float64 `json:"shortStockValue"`
	TotalCash                        float64 `json:"totalCash"`
	IsInCall                         bool    `json:"isInCall"`
	PendingDeposits                  float64 `json:"pendingDeposits"`
	MarginBalance                    float64 `json:"marginBalance"`
	ShortBalance                     float64 `json:"shortBalance"`
	AccountValue                     float64 `json:"accountValue"`
}

type CurrentBalances struct {
	AccruedInterest                  float64 `json:"accruedInterest"`
	CashBalance                      float64 `json:"cashBalance"`
	CashReceipts                     float64 `json:"cashReceipts"`
	LongOptionMarketValue            float64 `json:"longOptionMarketValue"`
	LiquidationValue                 float64 `json:"liquidationValue"`
	LongMarketValue                  float64 `json:"longMarketValue"`
	MoneyMarketFund                  float64 `json:"moneyMarketFund"`
	Savings                          float64 `json:"savings"`
	ShortMarketValue                 float64 `json:"shortMarketValue"`
	PendingDeposits                  float64 `json:"pendingDeposits"`
	MutualFundValue                  float64 `json:"mutualFundValue"`
	BondValue                        float64 `json:"bondValue"`
	ShortOptionMarketValue           float64 `json:"shortOptionMarketValue"`
	AvailableFunds                   float64 `json:"availableFunds"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BuyingPower                      float64 `json:"buyingPower"`
	BuyingPowerNonMarginableTrade    float64 `json:"buyingPowerNonMarginableTrade"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower"`
	Equity                           float64 `json:"equity"`
	EquityPercentage                 float64 `json:"equityPercentage"`
	LongMarginValue                  float64 `json:"longMarginValue"`
	MaintenanceCall                  float64 `json:"maintenanceCall"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement"`
	MarginBalance                    float64 `json:"marginBalance"`
	RegTCall                         float64 `json:"regTCall"`
	ShortBalance                     float64 `json:"shortBalance"`
	ShortMarginValue                 float64 `json:"shortMarginValue"`
	Sma                              float64 `json:"sma"`
}

type ProjectedBalances struct {
	AvailableFunds                   float64 `json:"availableFunds"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BuyingPower                      float64 `json:"buyingPower"`
	DayTradingBuyingPower            float64 `json:"dayTradingBuyingPower"`
	DayTradingBuyingPowerCall        float64 `json:"dayTradingBuyingPowerCall"`
	MaintenanceCall                  float64 `json:"maintenanceCall"`
	RegTCall                         float64 `json:"regTCall"`
	IsInCall                         bool    `json:"isInCall"`
	StockBuyingPower                 float64 `json:"stockBuyingPower"`
}

type Transactions []struct {
	ActivityID     int64          `json:"activityId"`
	Time           string         `json:"time"`
	AccountNumber  string         `json:"accountNumber"`
	Type           string         `json:"type"`
	Status         string         `json:"status"`
	SubAccount     string         `json:"subAccount"`
	TradeDate      string         `json:"tradeDate"`
	PositionID     int64          `json:"positionId"`
	OrderID        int64          `json:"orderId,omitempty"`
	NetAmount      float64        `json:"netAmount"`
	TransferItems  []TransferItem `json:"transferItems"`
	Description    string         `json:"description,omitempty"`
	SettlementDate string         `json:"settlementDate,omitempty"`
}
type TransactionInstrument struct {
	AssetType               string               `json:"assetType"`
	Status                  string               `json:"status"`
	Symbol                  string               `json:"symbol"`
	Description             string               `json:"description"`
	InstrumentID            int64                `json:"instrumentId"`
	ClosingPrice            float64              `json:"closingPrice"`
	ExpirationDate          string               `json:"expirationDate"`
	OptionDeliverables      []OptionDeliverables `json:"optionDeliverables"`
	OptionPremiumMultiplier int64                `json:"optionPremiumMultiplier"`
	PutCall                 string               `json:"putCall"`
	StrikePrice             float64              `json:"strikePrice"`
	Type                    string               `json:"type"`
	UnderlyingSymbol        string               `json:"underlyingSymbol"`
	UnderlyingCusip         string               `json:"underlyingCusip"`
}
type Deliverable struct {
	AssetType    string  `json:"assetType"`
	Status       string  `json:"status"`
	Symbol       string  `json:"symbol"`
	InstrumentID int64   `json:"instrumentId"`
	ClosingPrice float64 `json:"closingPrice"`
	Type         string  `json:"type"`
}
type OptionDeliverables struct {
	RootSymbol        string      `json:"rootSymbol"`
	StrikePercent     int         `json:"strikePercent"`
	DeliverableNumber int         `json:"deliverableNumber"`
	DeliverableUnits  int         `json:"deliverableUnits"`
	Deliverable       Deliverable `json:"deliverable"`
}

type TransferItem struct {
	Instrument     TransactionInstrument `json:"instrument,omitempty"`
	Amount         float64               `json:"amount"`
	Cost           float64               `json:"cost"`
	FeeType        string                `json:"feeType,omitempty"`
	Price          float64               `json:"price,omitempty"`
	PositionEffect string                `json:"positionEffect,omitempty"`
}

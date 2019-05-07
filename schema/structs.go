package schema

import (
	"github.com/niftynei/glightning/glightning"
)

//feerates
type FeeRateEstimate_ql struct {
	Style                      FeeRateStyle_ql
	Details                    *glightning.FeeRateDetails
	OnchainEstimate            *OnchainEstimate_ql
	Warning                    string
}
type FeeRateStyle_ql string
const (
	SatPerKiloByte FeeRateStyle_ql = "perkb"
	SatPerKiloSipa FeeRateStyle_ql = "perkw"
)
type OnchainEstimate_ql struct {
	OpeningChannelSatoshis     string
	MutualCloseSatoshis        string
	UnilateralCloseSatoshis    string
}
//feerates ^^

//listchannels
type Channel_ql struct {
	Source              string `json:"source"`
	Destination         string `json:"destination"`
	ShortChannelId      string `json:"short_channel_id"`
	IsPublic            bool   `json:"public"`
	Satoshis            string `json:"satoshis"`     //uint64
	MessageFlags        uint   `json:"message_flags"`
	ChannelFlags        uint   `json:"channel_flags"`
	IsActive            bool   `json:"active"`
	LastUpdate          uint   `json:"last_update"`
	BaseFeeMillisatoshi string `json:"base_fee_millisatoshi"` //uint64
	FeePerMillionth     string `json:"fee_per_millionth"`     //uint64
	Delay               uint   `json:"delay"`
}
//listchannels ^^

//listinvoices
type Invoice_ql struct {
	PaymentHash     string `json:"payment_hash"`
	ExpiresAt       string `json:"expires_at"` // uint64
	Bolt11          string `json:"bolt11"`
	WarningOffline  string `json:"warning_offline"`
	WarningCapacity string `json:"warning_capacity"`
	Label           string `json:"label"`
	Status          string `json:"status"`
	Description     string `json:"description"`
}
//listinvoices ^^

//getinfo
type NodeInfo_ql struct {
	Id                         string            `json:"id"`
	Alias                      string            `json:"alias"`
	Color                      string            `json:"color"`
	PeerCount                  int               `json:"num_peers"`
	PendingChannelCount        int               `json:"num_pending_channels"`
	ActiveChannelCount         int               `json:"num_active_channels"`
	InactiveChannelCount       int               `json:"num_inactive_channels"`
	Addresses                  []glightning.Address         `json:"address"`
	Binding                    []glightning.AddressInternal `json:"binding"`
	Version                    string            `json:"version"`
	Blockheight                int               `json:"blockheight"`
	Network                    string            `json:"network"`
	FeesCollectedMilliSatoshis string            `json:"msatoshi_fees_collected"` //graphql protocol cannot handle uint64, so turn in into string
}
//getinfo ^^

//listpeers
type Peer_ql struct {
	Id             string        `json:"id"`
	Connected      bool          `json:"connected"`
	NetAddresses   []string      `json:"netaddr"`
	GlobalFeatures string        `json:"globalfeatures"`
	LocalFeatures  string        `json:"localfeatures"`
	Channels       []PeerChannel_ql `json:"channels"`
	Logs           []glightning.Log         `json:"log,omitempty"`
}
type PeerChannel_ql struct {
	State                            string            `json:"state"`
	ScratchTxId                      string            `json:"scratch_txid"`
	Owner                            string            `json:"owner"`
	ShortChannelId                   string            `json:"short_channel_id"`
	ChannelDirection                 int               `json:"direction"`
	ChannelId                        string            `json:"channel_id"`
	FundingTxId                      string            `json:"funding_txid"`
	Funding                          string            `json:"funding"`
	Status                           []string          `json:"status"`
	Private                          bool              `json:"private"`
	FundingAllocations               []FundingAllocations_ql  `json:"funding_allocation_msat"`
	MilliSatoshiToUs                 string            `json:"msatoshi_to_us"`
	MilliSatoshiToUsMin              string            `json:"msatoshi_to_us_min"`
	MilliSatoshiToUsMax              string            `json:"msatoshi_to_us_max"`
	MilliSatoshiTotal                string            `json:"msatoshi_total"`
	DustLimitSatoshi                 string            `json:"dust_limit_satoshis"`
	MaxHtlcValueInFlightMilliSatoshi string            `json:"max_htlc_value_in_flight_msat"`
	TheirChannelReserveSatoshi       string            `json:"their_channel_reserve_satoshis"`
	OurChannelReserveSatoshi         string            `json:"our_channel_reserve_satoshis"`
	SpendableMilliSatoshi            string            `json:"spendable_msatoshi"`
	HtlcMinMilliSatoshi              string            `json:"htlc_minimum_msat"`
	TheirToSelfDelay                 uint              `json:"their_to_self_delay"`
	OurToSelfDelay                   uint              `json:"our_to_self_delay"`
	MaxAcceptedHtlcs                 uint              `json:"max_accepted_htlcs"`
	InPaymentsOffered                string            `json:"in_payments_offered"`
	InMilliSatoshiOffered            string            `json:"in_msatoshi_offered"`
	InPaymentsFulfilled              string            `json:"in_payments_fulfilled"`
	InMilliSatoshiFulfilled          string            `json:"in_msatoshi_fulfilled"`
	OutPaymentsOffered               string            `json:"out_payments_offered"`
	OutMilliSatoshiOffered           string            `json:"out_msatoshi_offered"`
	OutPaymentsFulfilled             string            `json:"out_payments_fulfilled"`
	OutMilliSatoshiFulfilled         string            `json:"out_msatoshi_fulfilled"`
	Htlcs                            []Htlc_ql         `json:"htlcs"`
}
type FundingAllocations_ql struct {
        id string
	msat string
}
type Htlc_ql struct {
	Direction    string `json:"direction"`
	Id           string `json:"id"`
	MilliSatoshi string `json:"msatoshi"` // uint64
	Expiry       string `json:"expiry"`   // uint64
	PaymentHash  string `json:"payment_hash"`
	State        string `json:"state"`
}
//listpeers ^^


//pay
type PaymentSuccess_ql struct {
	PaymentFields PaymentFields_ql
	GetRouteTries int          `json:"getroute_tries"`
	SendPayTries  int          `json:"sendpay_tries"`
	Route         []RouteHop_ql   `json:"route"`
	Failures      []PayFailure_ql `json:"failures"`
}
type RouteHop_ql struct {
	Id              string
	ShortChannelId  string
	MilliSatoshi    string  // uint64
	Delay           uint
}
type PayFailure_ql struct {
	Message       string     `json:"message"`
	Type          string     `json:"type"`
	OnionReply    string     `json:"onionreply"`
	ErringIndex   int        `json:"erring_index"`
	FailCode      int        `json:"failcode"`
	ErringNode    string     `json:"erring_node"`
	ErringChannel string     `json:"erring_channel"`
	ChannelUpdate string     `json:"channel_update"`
	Route         []RouteHop_ql `json:"route"`
}
type PaymentFields_ql struct {
	Id               string `json:"id"`             //uint64
	PaymentHash      string `json:"payment_hash"`
	Destination      string `json:"destination"`
	MilliSatoshi     string `json:"msatoshi"`       //uint64
	MilliSatoshiSent string `json:"msatoshi_sent"`  //uint64
	CreatedAt        string `json:"created_at"`     //uint64
	Status           string `json:"status"`
	PaymentPreimage  string `json:"payment_preimage"`
	Description      string `json:"description"`
}
//pay ^^



//waitanyinvoice
type CompletedInvoice_ql struct {
	Label                string `json:"label"`
	Bolt11               string `json:"bolt11"`
	PaymentHash          string `json:"payment_hash"`
	Status               string `json:"status"`
	Description          string `json:"description"`
	PayIndex             int    `json:"pay_index"`
	MilliSatoshi         string `json:"msatoshi"` //uint64
	MilliSatoshiReceived string `json:"msatoshi_received"` //uint64
	PaidAt               string `json:"paid_at"` //uint64
	ExpiresAt            string `json:"expires_at"` //uint64
}
//waitanyinvoice ^^




type PayRequest_ql struct {
	Bolt11         string
	MilliSatoshi   string  // uint64
	Desc           string
	RiskFactor     float32
	MaxFeePercent  float32
	RetryFor       uint
	MaxDelay       uint
	ExemptFee      bool
}


package model

type Account struct {
	Domain string `json:"domain"`
}

type Advertisable struct {
	Eid string `json:"eid"`
}

type AttributesAccountSource struct {
	Accounts []*Account `json:"accounts"`
}

type CompanyAttributes struct {
	Industry string `json:"industry"`
}

type TalAccountSource struct {
	Accounts []*Account `json:"accounts"`
}

type IntentAccountSource struct {
	Accounts []*Account `json:"accounts"`
}

type DatFilters struct {
	CompanySize float64 `json:"companySize"`
}

type IntentFilters struct {
	Topics []string `json:"topics"`
}

type IntentSurgeRecords struct {
	Topic string `json:"topic"`
}

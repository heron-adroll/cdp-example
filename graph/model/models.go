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
	FirstParty FirstPartyCompanyAttributes
	ThirdParty ThirdPartyCompanyAttributes
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

type FirstPartyCompanyAttributes struct {
	Industry string `json:"industry"`
}

type ThirdPartyCompanyAttributes struct {
	Employees float64 `json:"employees"`
}

package countries

type Country struct {
	CountrySummary
	States []State
}

type CountrySummary struct {
	CallingCode string
	CountryKey  string
	Label       string
}

type State struct {
	Label    string
	StateKey string
}

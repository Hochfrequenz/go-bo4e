package waehrungscode

// Waehrungscode is an enum for currencies
//go:generate stringer --type Waehrungscode
//go:generate jsonenums --type Waehrungscode
type Waehrungscode int

const (
	// AFN is the code of Afghani
	AFN Waehrungscode = iota + 1
	ALL               // ALL is the code of Lek
	AMD               // AMD is the code of Armenian Dram
	ANG               // ANG is the code of Netherlands Antillean Guilder
	AOA               // AOA is the code of Kwanza
	ARS               // ARS is the code of Argentine Peso
	AUD               // AUD is the code of Australian Dollar
	AWG               // AWG is the code of Aruban Florin
	AZN               // AZN is the code of Azerbaijanian Manat
	BAM               // BAM is the code of Convertible Mark
	BBD               // BBD is the code of Barbados Dollar
	BDT               // BDT is the code of Taka
	BGN               // BGN is the code of Bulgarian Lev
	BHD               // BHD is the code of Bahraini Dinar
	BIF               // BIF is the code of Burundi Franc
	BMD               // BMD is the code of Bermudian Dollar
	BND               // BND is the code of Brunei Dollar
	BOB               // BOB is the code of Boliviano
	BOV               // BOV is the code of Mvdol
	BRL               // BRL is the code of Brazilian Real
	BSD               // BSD is the code of Bahamian Dollar
	BTN               // BTN is the code of Ngultrum
	BWP               // BWP is the code of Pula
	BYN               // BYN is the code of Belarusian Ruble
	BYR               // BYR is the code of Belarusian Ruble
	BZD               // BZD is the code of Belize Dollar
	CAD               // CAD is the code of Canadian Dollar
	CDF               // CDF is the code of Congolese Franc
	CHE               // CHE is the code of WIR Euro
	CHF               // CHF is the code of Swiss Franc
	CHW               // CHW is the code of WIR Franc
	CLF               // CLF is the code of Unidad de Fomento
	CLP               // CLP is the code of Chilean Peso
	CNY               // CNY is the code of Yuan Renminbi
	COP               // COP is the code of Colombian Peso
	COU               // COU is the code of Unidad de Valor Real
	CRC               // CRC is the code of Costa Rican Colon
	CUC               // CUC is the code of Peso Convertible
	CUP               // CUP is the code of Cuban Peso
	CVE               // CVE is the code of Cape Verde Escudo
	CZK               // CZK is the code of Czech Koruna
	DJF               // DJF is the code of Djibouti Franc
	DKK               // DKK is the code of Danish Krone
	DOP               // DOP is the code of Dominican Peso
	DZD               // DZD is the code of Algerian Dinar
	EGP               // EGP is the code of Egyptian Pound
	ERN               // ERN is the code of Nakfa
	ETB               // ETB is the code of Ethiopian Birr
	EUR               // EUR is the code of Euro
	FJD               // FJD is the code of Fiji Dollar
	FKP               // FKP is the code of Falkland Islands Pound
	GBP               // GBP is the code of Pound Sterling
	GEL               // GEL is the code of Lari
	GHS               // GHS is the code of Ghana Cedi
	GIP               // GIP is the code of Gibraltar Pound
	GMD               // GMD is the code of Dalasi
	GNF               // GNF is the code of Guinea Franc
	GTQ               // GTQ is the code of Quetzal
	GYD               // GYD is the code of Guyana Dollar
	HKD               // HKD is the code of Hong Kong Dollar
	HNL               // HNL is the code of Lempira
	HRK               // HRK is the code of Croatian Kuna
	HTG               // HTG is the code of Gourde
	HUF               // HUF is the code of Forint
	IDR               // IDR is the code of Rupiah
	ILS               // ILS is the code of New Israeli Sheqel
	INR               // INR is the code of Indian Rupee
	IQD               // IQD is the code of Iraqi Dinar
	IRR               // IRR is the code of Iranian Rial
	ISK               // ISK is the code of Iceland Krona
	JMD               // JMD is the code of Jamaican Dollar
	JOD               // JOD is the code of Jordanian Dinar
	JPY               // JPY is the code of Yen
	KES               // KES is the code of Kenyan Shilling
	KGS               // KGS is the code of Som
	KHR               // KHR is the code of Riel
	KMF               // KMF is the code of Comoro Franc
	KPW               // KPW is the code of North Korean Won
	KRW               // KRW is the code of Won
	KWD               // KWD is the code of Kuwaiti Dinar
	KYD               // KYD is the code of Cayman Islands Dollar
	KZT               // KZT is the code of Tenge
	LAK               // LAK is the code of Kip
	LBP               // LBP is the code of Lebanese Pound
	LKR               // LKR is the code of Sri Lanka Rupee
	LRD               // LRD is the code of Liberian Dollar
	LSL               // LSL is the code of Loti
	LTL               // LTL is the code of Lithuanian Litas
	LYD               // LYD is the code of Libyan Dinar
	MAD               // MAD is the code of Moroccan Dirham
	MDL               // MDL is the code of Moldovan Leu
	MGA               // MGA is the code of Malagasy Ariary
	MKD               // MKD is the code of Denar
	MMK               // MMK is the code of Kyat
	MNT               // MNT is the code of Tugrik
	MOP               // MOP is the code of Pataca
	MRO               // MRO is the code of Ouguiya
	MUR               // MUR is the code of Mauritius Rupee
	MVR               // MVR is the code of Rufiyaa
	MWK               // MWK is the code of Kwacha
	MXN               // MXN is the code of Mexican Peso
	MXV               // MXV is the code of Mexican Unidad de Inversion (UDI)
	MYR               // MYR is the code of Malaysian Ringgit
	MZN               // MZN is the code of Mozambique Metical
	NAD               // NAD is the code of Namibia Dollar
	NGN               // NGN is the code of Naira
	NIO               // NIO is the code of Cordoba Oro
	NOK               // NOK is the code of Norwegian Krone
	NPR               // NPR is the code of Nepalese Rupee
	NZD               // NZD is the code of New Zealand Dollar
	OMR               // OMR is the code of Rial Omani
	PAB               // PAB is the code of Balboa
	PEN               // PEN is the code of Nuevo Sol
	PGK               // PGK is the code of Kina
	PHP               // PHP is the code of Philippine Peso
	PKR               // PKR is the code of Pakistan Rupee
	PLN               // PLN is the code of Zloty
	PYG               // PYG is the code of Guarani
	QAR               // QAR is the code of Qatari Rial
	RON               // RON is the code of New Romanian Leu
	RSD               // RSD is the code of Serbian Dinar
	RUB               // RUB is the code of Russian Ruble
	RUR               // RUR is the code of Russian Ruble
	RWF               // RWF is the code of Rwanda Franc
	SAR               // SAR is the code of Saudi Riyal
	SBD               // SBD is the code of Solomon Islands Dollar
	SCR               // SCR is the code of Seychelles Rupee
	SDG               // SDG is the code of Sudanese Pound
	SEK               // SEK is the code of Swedish Krona
	SGD               // SGD is the code of Singapore Dollar
	SHP               // SHP is the code of Saint Helena Pound
	SLL               // SLL is the code of Leone
	SOS               // SOS is the code of Somali Shilling
	SRD               // SRD is the code of Surinam Dollar
	SSP               // SSP is the code of South Sudanese Pound
	STD               // STD is the code of Dobra
	SVC               // SVC is the code of El Salvador Colon
	SYP               // SYP is the code of Syrian Pound
	SZL               // SZL is the code of Lilangeni
	THB               // THB is the code of Baht
	TJS               // TJS is the code of Somoni
	TMT               // TMT is the code of Turkmenistan New Manat
	TND               // TND is the code of Tunisian Dinar
	TOP               // TOP is the code of PaÊ»anga
	TRY               // TRY is the code of Turkish Lira
	TTD               // TTD is the code of Trinidad and Tobago Dollar
	TWD               // TWD is the code of New Taiwan Dollar
	TZS               // TZS is the code of Tanzanian Shilling
	UAH               // UAH is the code of Hryvnia
	UGX               // UGX is the code of Uganda Shilling
	USD               // USD is the code of US Dollar
	USN               // USN is the code of US Dollar (Next day)
	USS               // USS is the code of US Dollar (Same day)
	UYI               // UYI is the code of Uruguay Peso en Unidades Indexadas (URUIURUI)
	UYU               // UYU is the code of Peso Uruguayo
	UZS               // UZS is the code of Uzbekistan Sum
	VEF               // VEF is the code of Bolivar
	VND               // VND is the code of Dong
	VUV               // VUV is the code of Vatu
	WST               // WST is the code of Tala
	XAF               // XAF is the code of CFA Franc BEAC
	XAG               // XAG is the code of Silver
	XAU               // XAU is the code of Gold
	XBA               // XBA is the code of Bond Markets Unit European Composite Unit (EURCO)
	XBB               // XBB is the code of Bond Markets Unit European Monetary Unit (E.M.U.-6)
	XBC               // XBC is the code of Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	XBD               // XBD is the code of Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	XCD               // XCD is the code of East Caribbean Dollar
	XDR               // XDR is the code of SDR (Special Drawing Right)
	XOF               // XOF is the code of CFA Franc BCEAO
	XPD               // XPD is the code of Palladium
	XPF               // XPF is the code of CFP Franc
	XPT               // XPT is the code of Platinum
	XSU               // XSU is the code of Sucre
	XTS               // XTS is the code of Codes specifically reserved for testing purposes
	XUA               // XUA is the code of ADB Unit of Account
	XXX               // XXX is the code of The codes assigned for transactions where no currency ist involved
	YER               // YER is the code of Yemeni Rial
	ZAR               // ZAR is the code of Rand
	ZMW               // ZMW is the code of Zambian Kwacha
	ZWL               // ZWL is the code of Zimbabwe Dollar
)

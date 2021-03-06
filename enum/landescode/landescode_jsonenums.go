// Code generated by jsonenums -type=Landescode; DO NOT EDIT.

package landescode

import (
	"encoding/json"
	"fmt"
)

var (
	_LandescodeNameToValue = map[string]Landescode{
		"AC": AC,
		"AD": AD,
		"AE": AE,
		"AF": AF,
		"AG": AG,
		"AI": AI,
		"AL": AL,
		"AM": AM,
		"AN": AN,
		"AO": AO,
		"AQ": AQ,
		"AR": AR,
		"AS": AS,
		"AT": AT,
		"AU": AU,
		"AW": AW,
		"AX": AX,
		"AZ": AZ,
		"BA": BA,
		"BB": BB,
		"BD": BD,
		"BE": BE,
		"BF": BF,
		"BG": BG,
		"BH": BH,
		"BI": BI,
		"BJ": BJ,
		"BL": BL,
		"BM": BM,
		"BN": BN,
		"BO": BO,
		"BQ": BQ,
		"BR": BR,
		"BS": BS,
		"BT": BT,
		"BU": BU,
		"BV": BV,
		"BW": BW,
		"BY": BY,
		"BZ": BZ,
		"CA": CA,
		"CC": CC,
		"CD": CD,
		"CF": CF,
		"CG": CG,
		"CH": CH,
		"CI": CI,
		"CK": CK,
		"CL": CL,
		"CM": CM,
		"CN": CN,
		"CO": CO,
		"CP": CP,
		"CR": CR,
		"CS": CS,
		"CU": CU,
		"CV": CV,
		"CW": CW,
		"CX": CX,
		"CY": CY,
		"CZ": CZ,
		"DE": DE,
		"DG": DG,
		"DJ": DJ,
		"DK": DK,
		"DM": DM,
		"DO": DO,
		"DZ": DZ,
		"EA": EA,
		"EC": EC,
		"EE": EE,
		"EG": EG,
		"EH": EH,
		"ER": ER,
		"ES": ES,
		"ET": ET,
		"EU": EU,
		"FI": FI,
		"FJ": FJ,
		"FK": FK,
		"FM": FM,
		"FO": FO,
		"FR": FR,
		"FX": FX,
		"GA": GA,
		"GB": GB,
		"GD": GD,
		"GE": GE,
		"GF": GF,
		"GG": GG,
		"GH": GH,
		"GI": GI,
		"GL": GL,
		"GM": GM,
		"GN": GN,
		"GP": GP,
		"GQ": GQ,
		"GR": GR,
		"GS": GS,
		"GT": GT,
		"GU": GU,
		"GW": GW,
		"GY": GY,
		"HK": HK,
		"HM": HM,
		"HN": HN,
		"HR": HR,
		"HT": HT,
		"HU": HU,
		"IC": IC,
		"ID": ID,
		"IE": IE,
		"IL": IL,
		"IM": IM,
		"IN": IN,
		"IO": IO,
		"IQ": IQ,
		"IR": IR,
		"IS": IS,
		"IT": IT,
		"JE": JE,
		"JM": JM,
		"JO": JO,
		"JP": JP,
		"KE": KE,
		"KG": KG,
		"KH": KH,
		"KI": KI,
		"KM": KM,
		"KN": KN,
		"KP": KP,
		"KR": KR,
		"KW": KW,
		"KY": KY,
		"KZ": KZ,
		"LA": LA,
		"LB": LB,
		"LC": LC,
		"LI": LI,
		"LK": LK,
		"LR": LR,
		"LS": LS,
		"LT": LT,
		"LU": LU,
		"LV": LV,
		"LY": LY,
		"MA": MA,
		"MC": MC,
		"MD": MD,
		"ME": ME,
		"MF": MF,
		"MG": MG,
		"MH": MH,
		"MK": MK,
		"ML": ML,
		"MM": MM,
		"MN": MN,
		"MO": MO,
		"MP": MP,
		"MQ": MQ,
		"MR": MR,
		"MS": MS,
		"MT": MT,
		"MU": MU,
		"MV": MV,
		"MW": MW,
		"MX": MX,
		"MY": MY,
		"MZ": MZ,
		"NA": NA,
		"NC": NC,
		"NE": NE,
		"NF": NF,
		"NG": NG,
		"NI": NI,
		"NL": NL,
		"NO": NO,
		"NP": NP,
		"NR": NR,
		"NT": NT,
		"NU": NU,
		"NZ": NZ,
		"OM": OM,
		"PA": PA,
		"PE": PE,
		"PF": PF,
		"PG": PG,
		"PH": PH,
		"PK": PK,
		"PL": PL,
		"PM": PM,
		"PN": PN,
		"PR": PR,
		"PS": PS,
		"PT": PT,
		"PW": PW,
		"PY": PY,
		"QA": QA,
		"RE": RE,
		"RO": RO,
		"RS": RS,
		"RU": RU,
		"RW": RW,
		"SA": SA,
		"SB": SB,
		"SC": SC,
		"SD": SD,
		"SE": SE,
		"SF": SF,
		"SG": SG,
		"SH": SH,
		"SI": SI,
		"SJ": SJ,
		"SK": SK,
		"SL": SL,
		"SM": SM,
		"SN": SN,
		"SO": SO,
		"SR": SR,
		"SS": SS,
		"ST": ST,
		"SU": SU,
		"SV": SV,
		"SX": SX,
		"SY": SY,
		"SZ": SZ,
		"TA": TA,
		"TC": TC,
		"TD": TD,
		"TF": TF,
		"TG": TG,
		"TJ": TJ,
		"TK": TK,
		"TL": TL,
		"TM": TM,
		"TN": TN,
		"TO": TO,
		"TP": TP,
		"TR": TR,
		"TT": TT,
		"TV": TV,
		"TW": TW,
		"TZ": TZ,
		"UA": UA,
		"UG": UG,
		"UK": UK,
		"UM": UM,
		"US": US,
		"UY": UY,
		"UZ": UZ,
		"VA": VA,
		"VC": VC,
		"VE": VE,
		"VG": VG,
		"VI": VI,
		"VN": VN,
		"VU": VU,
		"WF": WF,
		"WS": WS,
		"XK": XK,
		"YE": YE,
		"YT": YT,
		"YU": YU,
		"ZA": ZA,
		"ZM": ZM,
		"ZR": ZR,
		"ZW": ZW,
	}

	_LandescodeValueToName = map[Landescode]string{
		AC: "AC",
		AD: "AD",
		AE: "AE",
		AF: "AF",
		AG: "AG",
		AI: "AI",
		AL: "AL",
		AM: "AM",
		AN: "AN",
		AO: "AO",
		AQ: "AQ",
		AR: "AR",
		AS: "AS",
		AT: "AT",
		AU: "AU",
		AW: "AW",
		AX: "AX",
		AZ: "AZ",
		BA: "BA",
		BB: "BB",
		BD: "BD",
		BE: "BE",
		BF: "BF",
		BG: "BG",
		BH: "BH",
		BI: "BI",
		BJ: "BJ",
		BL: "BL",
		BM: "BM",
		BN: "BN",
		BO: "BO",
		BQ: "BQ",
		BR: "BR",
		BS: "BS",
		BT: "BT",
		BU: "BU",
		BV: "BV",
		BW: "BW",
		BY: "BY",
		BZ: "BZ",
		CA: "CA",
		CC: "CC",
		CD: "CD",
		CF: "CF",
		CG: "CG",
		CH: "CH",
		CI: "CI",
		CK: "CK",
		CL: "CL",
		CM: "CM",
		CN: "CN",
		CO: "CO",
		CP: "CP",
		CR: "CR",
		CS: "CS",
		CU: "CU",
		CV: "CV",
		CW: "CW",
		CX: "CX",
		CY: "CY",
		CZ: "CZ",
		DE: "DE",
		DG: "DG",
		DJ: "DJ",
		DK: "DK",
		DM: "DM",
		DO: "DO",
		DZ: "DZ",
		EA: "EA",
		EC: "EC",
		EE: "EE",
		EG: "EG",
		EH: "EH",
		ER: "ER",
		ES: "ES",
		ET: "ET",
		EU: "EU",
		FI: "FI",
		FJ: "FJ",
		FK: "FK",
		FM: "FM",
		FO: "FO",
		FR: "FR",
		FX: "FX",
		GA: "GA",
		GB: "GB",
		GD: "GD",
		GE: "GE",
		GF: "GF",
		GG: "GG",
		GH: "GH",
		GI: "GI",
		GL: "GL",
		GM: "GM",
		GN: "GN",
		GP: "GP",
		GQ: "GQ",
		GR: "GR",
		GS: "GS",
		GT: "GT",
		GU: "GU",
		GW: "GW",
		GY: "GY",
		HK: "HK",
		HM: "HM",
		HN: "HN",
		HR: "HR",
		HT: "HT",
		HU: "HU",
		IC: "IC",
		ID: "ID",
		IE: "IE",
		IL: "IL",
		IM: "IM",
		IN: "IN",
		IO: "IO",
		IQ: "IQ",
		IR: "IR",
		IS: "IS",
		IT: "IT",
		JE: "JE",
		JM: "JM",
		JO: "JO",
		JP: "JP",
		KE: "KE",
		KG: "KG",
		KH: "KH",
		KI: "KI",
		KM: "KM",
		KN: "KN",
		KP: "KP",
		KR: "KR",
		KW: "KW",
		KY: "KY",
		KZ: "KZ",
		LA: "LA",
		LB: "LB",
		LC: "LC",
		LI: "LI",
		LK: "LK",
		LR: "LR",
		LS: "LS",
		LT: "LT",
		LU: "LU",
		LV: "LV",
		LY: "LY",
		MA: "MA",
		MC: "MC",
		MD: "MD",
		ME: "ME",
		MF: "MF",
		MG: "MG",
		MH: "MH",
		MK: "MK",
		ML: "ML",
		MM: "MM",
		MN: "MN",
		MO: "MO",
		MP: "MP",
		MQ: "MQ",
		MR: "MR",
		MS: "MS",
		MT: "MT",
		MU: "MU",
		MV: "MV",
		MW: "MW",
		MX: "MX",
		MY: "MY",
		MZ: "MZ",
		NA: "NA",
		NC: "NC",
		NE: "NE",
		NF: "NF",
		NG: "NG",
		NI: "NI",
		NL: "NL",
		NO: "NO",
		NP: "NP",
		NR: "NR",
		NT: "NT",
		NU: "NU",
		NZ: "NZ",
		OM: "OM",
		PA: "PA",
		PE: "PE",
		PF: "PF",
		PG: "PG",
		PH: "PH",
		PK: "PK",
		PL: "PL",
		PM: "PM",
		PN: "PN",
		PR: "PR",
		PS: "PS",
		PT: "PT",
		PW: "PW",
		PY: "PY",
		QA: "QA",
		RE: "RE",
		RO: "RO",
		RS: "RS",
		RU: "RU",
		RW: "RW",
		SA: "SA",
		SB: "SB",
		SC: "SC",
		SD: "SD",
		SE: "SE",
		SF: "SF",
		SG: "SG",
		SH: "SH",
		SI: "SI",
		SJ: "SJ",
		SK: "SK",
		SL: "SL",
		SM: "SM",
		SN: "SN",
		SO: "SO",
		SR: "SR",
		SS: "SS",
		ST: "ST",
		SU: "SU",
		SV: "SV",
		SX: "SX",
		SY: "SY",
		SZ: "SZ",
		TA: "TA",
		TC: "TC",
		TD: "TD",
		TF: "TF",
		TG: "TG",
		TJ: "TJ",
		TK: "TK",
		TL: "TL",
		TM: "TM",
		TN: "TN",
		TO: "TO",
		TP: "TP",
		TR: "TR",
		TT: "TT",
		TV: "TV",
		TW: "TW",
		TZ: "TZ",
		UA: "UA",
		UG: "UG",
		UK: "UK",
		UM: "UM",
		US: "US",
		UY: "UY",
		UZ: "UZ",
		VA: "VA",
		VC: "VC",
		VE: "VE",
		VG: "VG",
		VI: "VI",
		VN: "VN",
		VU: "VU",
		WF: "WF",
		WS: "WS",
		XK: "XK",
		YE: "YE",
		YT: "YT",
		YU: "YU",
		ZA: "ZA",
		ZM: "ZM",
		ZR: "ZR",
		ZW: "ZW",
	}
)

func init() {
	var v Landescode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_LandescodeNameToValue = map[string]Landescode{
			interface{}(AC).(fmt.Stringer).String(): AC,
			interface{}(AD).(fmt.Stringer).String(): AD,
			interface{}(AE).(fmt.Stringer).String(): AE,
			interface{}(AF).(fmt.Stringer).String(): AF,
			interface{}(AG).(fmt.Stringer).String(): AG,
			interface{}(AI).(fmt.Stringer).String(): AI,
			interface{}(AL).(fmt.Stringer).String(): AL,
			interface{}(AM).(fmt.Stringer).String(): AM,
			interface{}(AN).(fmt.Stringer).String(): AN,
			interface{}(AO).(fmt.Stringer).String(): AO,
			interface{}(AQ).(fmt.Stringer).String(): AQ,
			interface{}(AR).(fmt.Stringer).String(): AR,
			interface{}(AS).(fmt.Stringer).String(): AS,
			interface{}(AT).(fmt.Stringer).String(): AT,
			interface{}(AU).(fmt.Stringer).String(): AU,
			interface{}(AW).(fmt.Stringer).String(): AW,
			interface{}(AX).(fmt.Stringer).String(): AX,
			interface{}(AZ).(fmt.Stringer).String(): AZ,
			interface{}(BA).(fmt.Stringer).String(): BA,
			interface{}(BB).(fmt.Stringer).String(): BB,
			interface{}(BD).(fmt.Stringer).String(): BD,
			interface{}(BE).(fmt.Stringer).String(): BE,
			interface{}(BF).(fmt.Stringer).String(): BF,
			interface{}(BG).(fmt.Stringer).String(): BG,
			interface{}(BH).(fmt.Stringer).String(): BH,
			interface{}(BI).(fmt.Stringer).String(): BI,
			interface{}(BJ).(fmt.Stringer).String(): BJ,
			interface{}(BL).(fmt.Stringer).String(): BL,
			interface{}(BM).(fmt.Stringer).String(): BM,
			interface{}(BN).(fmt.Stringer).String(): BN,
			interface{}(BO).(fmt.Stringer).String(): BO,
			interface{}(BQ).(fmt.Stringer).String(): BQ,
			interface{}(BR).(fmt.Stringer).String(): BR,
			interface{}(BS).(fmt.Stringer).String(): BS,
			interface{}(BT).(fmt.Stringer).String(): BT,
			interface{}(BU).(fmt.Stringer).String(): BU,
			interface{}(BV).(fmt.Stringer).String(): BV,
			interface{}(BW).(fmt.Stringer).String(): BW,
			interface{}(BY).(fmt.Stringer).String(): BY,
			interface{}(BZ).(fmt.Stringer).String(): BZ,
			interface{}(CA).(fmt.Stringer).String(): CA,
			interface{}(CC).(fmt.Stringer).String(): CC,
			interface{}(CD).(fmt.Stringer).String(): CD,
			interface{}(CF).(fmt.Stringer).String(): CF,
			interface{}(CG).(fmt.Stringer).String(): CG,
			interface{}(CH).(fmt.Stringer).String(): CH,
			interface{}(CI).(fmt.Stringer).String(): CI,
			interface{}(CK).(fmt.Stringer).String(): CK,
			interface{}(CL).(fmt.Stringer).String(): CL,
			interface{}(CM).(fmt.Stringer).String(): CM,
			interface{}(CN).(fmt.Stringer).String(): CN,
			interface{}(CO).(fmt.Stringer).String(): CO,
			interface{}(CP).(fmt.Stringer).String(): CP,
			interface{}(CR).(fmt.Stringer).String(): CR,
			interface{}(CS).(fmt.Stringer).String(): CS,
			interface{}(CU).(fmt.Stringer).String(): CU,
			interface{}(CV).(fmt.Stringer).String(): CV,
			interface{}(CW).(fmt.Stringer).String(): CW,
			interface{}(CX).(fmt.Stringer).String(): CX,
			interface{}(CY).(fmt.Stringer).String(): CY,
			interface{}(CZ).(fmt.Stringer).String(): CZ,
			interface{}(DE).(fmt.Stringer).String(): DE,
			interface{}(DG).(fmt.Stringer).String(): DG,
			interface{}(DJ).(fmt.Stringer).String(): DJ,
			interface{}(DK).(fmt.Stringer).String(): DK,
			interface{}(DM).(fmt.Stringer).String(): DM,
			interface{}(DO).(fmt.Stringer).String(): DO,
			interface{}(DZ).(fmt.Stringer).String(): DZ,
			interface{}(EA).(fmt.Stringer).String(): EA,
			interface{}(EC).(fmt.Stringer).String(): EC,
			interface{}(EE).(fmt.Stringer).String(): EE,
			interface{}(EG).(fmt.Stringer).String(): EG,
			interface{}(EH).(fmt.Stringer).String(): EH,
			interface{}(ER).(fmt.Stringer).String(): ER,
			interface{}(ES).(fmt.Stringer).String(): ES,
			interface{}(ET).(fmt.Stringer).String(): ET,
			interface{}(EU).(fmt.Stringer).String(): EU,
			interface{}(FI).(fmt.Stringer).String(): FI,
			interface{}(FJ).(fmt.Stringer).String(): FJ,
			interface{}(FK).(fmt.Stringer).String(): FK,
			interface{}(FM).(fmt.Stringer).String(): FM,
			interface{}(FO).(fmt.Stringer).String(): FO,
			interface{}(FR).(fmt.Stringer).String(): FR,
			interface{}(FX).(fmt.Stringer).String(): FX,
			interface{}(GA).(fmt.Stringer).String(): GA,
			interface{}(GB).(fmt.Stringer).String(): GB,
			interface{}(GD).(fmt.Stringer).String(): GD,
			interface{}(GE).(fmt.Stringer).String(): GE,
			interface{}(GF).(fmt.Stringer).String(): GF,
			interface{}(GG).(fmt.Stringer).String(): GG,
			interface{}(GH).(fmt.Stringer).String(): GH,
			interface{}(GI).(fmt.Stringer).String(): GI,
			interface{}(GL).(fmt.Stringer).String(): GL,
			interface{}(GM).(fmt.Stringer).String(): GM,
			interface{}(GN).(fmt.Stringer).String(): GN,
			interface{}(GP).(fmt.Stringer).String(): GP,
			interface{}(GQ).(fmt.Stringer).String(): GQ,
			interface{}(GR).(fmt.Stringer).String(): GR,
			interface{}(GS).(fmt.Stringer).String(): GS,
			interface{}(GT).(fmt.Stringer).String(): GT,
			interface{}(GU).(fmt.Stringer).String(): GU,
			interface{}(GW).(fmt.Stringer).String(): GW,
			interface{}(GY).(fmt.Stringer).String(): GY,
			interface{}(HK).(fmt.Stringer).String(): HK,
			interface{}(HM).(fmt.Stringer).String(): HM,
			interface{}(HN).(fmt.Stringer).String(): HN,
			interface{}(HR).(fmt.Stringer).String(): HR,
			interface{}(HT).(fmt.Stringer).String(): HT,
			interface{}(HU).(fmt.Stringer).String(): HU,
			interface{}(IC).(fmt.Stringer).String(): IC,
			interface{}(ID).(fmt.Stringer).String(): ID,
			interface{}(IE).(fmt.Stringer).String(): IE,
			interface{}(IL).(fmt.Stringer).String(): IL,
			interface{}(IM).(fmt.Stringer).String(): IM,
			interface{}(IN).(fmt.Stringer).String(): IN,
			interface{}(IO).(fmt.Stringer).String(): IO,
			interface{}(IQ).(fmt.Stringer).String(): IQ,
			interface{}(IR).(fmt.Stringer).String(): IR,
			interface{}(IS).(fmt.Stringer).String(): IS,
			interface{}(IT).(fmt.Stringer).String(): IT,
			interface{}(JE).(fmt.Stringer).String(): JE,
			interface{}(JM).(fmt.Stringer).String(): JM,
			interface{}(JO).(fmt.Stringer).String(): JO,
			interface{}(JP).(fmt.Stringer).String(): JP,
			interface{}(KE).(fmt.Stringer).String(): KE,
			interface{}(KG).(fmt.Stringer).String(): KG,
			interface{}(KH).(fmt.Stringer).String(): KH,
			interface{}(KI).(fmt.Stringer).String(): KI,
			interface{}(KM).(fmt.Stringer).String(): KM,
			interface{}(KN).(fmt.Stringer).String(): KN,
			interface{}(KP).(fmt.Stringer).String(): KP,
			interface{}(KR).(fmt.Stringer).String(): KR,
			interface{}(KW).(fmt.Stringer).String(): KW,
			interface{}(KY).(fmt.Stringer).String(): KY,
			interface{}(KZ).(fmt.Stringer).String(): KZ,
			interface{}(LA).(fmt.Stringer).String(): LA,
			interface{}(LB).(fmt.Stringer).String(): LB,
			interface{}(LC).(fmt.Stringer).String(): LC,
			interface{}(LI).(fmt.Stringer).String(): LI,
			interface{}(LK).(fmt.Stringer).String(): LK,
			interface{}(LR).(fmt.Stringer).String(): LR,
			interface{}(LS).(fmt.Stringer).String(): LS,
			interface{}(LT).(fmt.Stringer).String(): LT,
			interface{}(LU).(fmt.Stringer).String(): LU,
			interface{}(LV).(fmt.Stringer).String(): LV,
			interface{}(LY).(fmt.Stringer).String(): LY,
			interface{}(MA).(fmt.Stringer).String(): MA,
			interface{}(MC).(fmt.Stringer).String(): MC,
			interface{}(MD).(fmt.Stringer).String(): MD,
			interface{}(ME).(fmt.Stringer).String(): ME,
			interface{}(MF).(fmt.Stringer).String(): MF,
			interface{}(MG).(fmt.Stringer).String(): MG,
			interface{}(MH).(fmt.Stringer).String(): MH,
			interface{}(MK).(fmt.Stringer).String(): MK,
			interface{}(ML).(fmt.Stringer).String(): ML,
			interface{}(MM).(fmt.Stringer).String(): MM,
			interface{}(MN).(fmt.Stringer).String(): MN,
			interface{}(MO).(fmt.Stringer).String(): MO,
			interface{}(MP).(fmt.Stringer).String(): MP,
			interface{}(MQ).(fmt.Stringer).String(): MQ,
			interface{}(MR).(fmt.Stringer).String(): MR,
			interface{}(MS).(fmt.Stringer).String(): MS,
			interface{}(MT).(fmt.Stringer).String(): MT,
			interface{}(MU).(fmt.Stringer).String(): MU,
			interface{}(MV).(fmt.Stringer).String(): MV,
			interface{}(MW).(fmt.Stringer).String(): MW,
			interface{}(MX).(fmt.Stringer).String(): MX,
			interface{}(MY).(fmt.Stringer).String(): MY,
			interface{}(MZ).(fmt.Stringer).String(): MZ,
			interface{}(NA).(fmt.Stringer).String(): NA,
			interface{}(NC).(fmt.Stringer).String(): NC,
			interface{}(NE).(fmt.Stringer).String(): NE,
			interface{}(NF).(fmt.Stringer).String(): NF,
			interface{}(NG).(fmt.Stringer).String(): NG,
			interface{}(NI).(fmt.Stringer).String(): NI,
			interface{}(NL).(fmt.Stringer).String(): NL,
			interface{}(NO).(fmt.Stringer).String(): NO,
			interface{}(NP).(fmt.Stringer).String(): NP,
			interface{}(NR).(fmt.Stringer).String(): NR,
			interface{}(NT).(fmt.Stringer).String(): NT,
			interface{}(NU).(fmt.Stringer).String(): NU,
			interface{}(NZ).(fmt.Stringer).String(): NZ,
			interface{}(OM).(fmt.Stringer).String(): OM,
			interface{}(PA).(fmt.Stringer).String(): PA,
			interface{}(PE).(fmt.Stringer).String(): PE,
			interface{}(PF).(fmt.Stringer).String(): PF,
			interface{}(PG).(fmt.Stringer).String(): PG,
			interface{}(PH).(fmt.Stringer).String(): PH,
			interface{}(PK).(fmt.Stringer).String(): PK,
			interface{}(PL).(fmt.Stringer).String(): PL,
			interface{}(PM).(fmt.Stringer).String(): PM,
			interface{}(PN).(fmt.Stringer).String(): PN,
			interface{}(PR).(fmt.Stringer).String(): PR,
			interface{}(PS).(fmt.Stringer).String(): PS,
			interface{}(PT).(fmt.Stringer).String(): PT,
			interface{}(PW).(fmt.Stringer).String(): PW,
			interface{}(PY).(fmt.Stringer).String(): PY,
			interface{}(QA).(fmt.Stringer).String(): QA,
			interface{}(RE).(fmt.Stringer).String(): RE,
			interface{}(RO).(fmt.Stringer).String(): RO,
			interface{}(RS).(fmt.Stringer).String(): RS,
			interface{}(RU).(fmt.Stringer).String(): RU,
			interface{}(RW).(fmt.Stringer).String(): RW,
			interface{}(SA).(fmt.Stringer).String(): SA,
			interface{}(SB).(fmt.Stringer).String(): SB,
			interface{}(SC).(fmt.Stringer).String(): SC,
			interface{}(SD).(fmt.Stringer).String(): SD,
			interface{}(SE).(fmt.Stringer).String(): SE,
			interface{}(SF).(fmt.Stringer).String(): SF,
			interface{}(SG).(fmt.Stringer).String(): SG,
			interface{}(SH).(fmt.Stringer).String(): SH,
			interface{}(SI).(fmt.Stringer).String(): SI,
			interface{}(SJ).(fmt.Stringer).String(): SJ,
			interface{}(SK).(fmt.Stringer).String(): SK,
			interface{}(SL).(fmt.Stringer).String(): SL,
			interface{}(SM).(fmt.Stringer).String(): SM,
			interface{}(SN).(fmt.Stringer).String(): SN,
			interface{}(SO).(fmt.Stringer).String(): SO,
			interface{}(SR).(fmt.Stringer).String(): SR,
			interface{}(SS).(fmt.Stringer).String(): SS,
			interface{}(ST).(fmt.Stringer).String(): ST,
			interface{}(SU).(fmt.Stringer).String(): SU,
			interface{}(SV).(fmt.Stringer).String(): SV,
			interface{}(SX).(fmt.Stringer).String(): SX,
			interface{}(SY).(fmt.Stringer).String(): SY,
			interface{}(SZ).(fmt.Stringer).String(): SZ,
			interface{}(TA).(fmt.Stringer).String(): TA,
			interface{}(TC).(fmt.Stringer).String(): TC,
			interface{}(TD).(fmt.Stringer).String(): TD,
			interface{}(TF).(fmt.Stringer).String(): TF,
			interface{}(TG).(fmt.Stringer).String(): TG,
			interface{}(TJ).(fmt.Stringer).String(): TJ,
			interface{}(TK).(fmt.Stringer).String(): TK,
			interface{}(TL).(fmt.Stringer).String(): TL,
			interface{}(TM).(fmt.Stringer).String(): TM,
			interface{}(TN).(fmt.Stringer).String(): TN,
			interface{}(TO).(fmt.Stringer).String(): TO,
			interface{}(TP).(fmt.Stringer).String(): TP,
			interface{}(TR).(fmt.Stringer).String(): TR,
			interface{}(TT).(fmt.Stringer).String(): TT,
			interface{}(TV).(fmt.Stringer).String(): TV,
			interface{}(TW).(fmt.Stringer).String(): TW,
			interface{}(TZ).(fmt.Stringer).String(): TZ,
			interface{}(UA).(fmt.Stringer).String(): UA,
			interface{}(UG).(fmt.Stringer).String(): UG,
			interface{}(UK).(fmt.Stringer).String(): UK,
			interface{}(UM).(fmt.Stringer).String(): UM,
			interface{}(US).(fmt.Stringer).String(): US,
			interface{}(UY).(fmt.Stringer).String(): UY,
			interface{}(UZ).(fmt.Stringer).String(): UZ,
			interface{}(VA).(fmt.Stringer).String(): VA,
			interface{}(VC).(fmt.Stringer).String(): VC,
			interface{}(VE).(fmt.Stringer).String(): VE,
			interface{}(VG).(fmt.Stringer).String(): VG,
			interface{}(VI).(fmt.Stringer).String(): VI,
			interface{}(VN).(fmt.Stringer).String(): VN,
			interface{}(VU).(fmt.Stringer).String(): VU,
			interface{}(WF).(fmt.Stringer).String(): WF,
			interface{}(WS).(fmt.Stringer).String(): WS,
			interface{}(XK).(fmt.Stringer).String(): XK,
			interface{}(YE).(fmt.Stringer).String(): YE,
			interface{}(YT).(fmt.Stringer).String(): YT,
			interface{}(YU).(fmt.Stringer).String(): YU,
			interface{}(ZA).(fmt.Stringer).String(): ZA,
			interface{}(ZM).(fmt.Stringer).String(): ZM,
			interface{}(ZR).(fmt.Stringer).String(): ZR,
			interface{}(ZW).(fmt.Stringer).String(): ZW,
		}
	}
}

// MarshalJSON is generated so Landescode satisfies json.Marshaler.
func (r Landescode) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _LandescodeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Landescode: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Landescode satisfies json.Unmarshaler.
func (r *Landescode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Landescode should be a string, got %s", data)
	}
	v, ok := _LandescodeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Landescode %q", s)
	}
	*r = v
	return nil
}

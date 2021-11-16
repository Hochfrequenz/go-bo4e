package zeitreihentyp

type Zeitreihentyp int

const(
	EGS Zeitreihentyp = iota+1 // EGS ist eine Einspeisegangsumme
	LGS // LGS ist eine Lastgangsumme
	NZR // NZR ist eine Netzzeitreihe
	SES // SES ist eine Standardeinspeiseprofilsumme
	SLS // SLS ist eine Standardlastsumme
	TES //
)
/// <summary>Einspeisegangsumme</summary>
EGS,

/// <summary>Lastgangsumme</summary>
LGS,

/// <summary>Netzzeitreihe</summary>
NZR,

/// <summary>Standardeinspeiseprofilsumme</summary>
SES,

/// <summary>Standardlastsumme</summary>
SLS,

/// <summary>Tagesparameterabhängige Einspeiseprofilsumme</summary>
TES,

/// <summary>Tagesparameterabhängige Lastprofilsumme</summary>
TLS

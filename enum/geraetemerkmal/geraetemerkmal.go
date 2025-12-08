package geraetemerkmal

// Auflistung möglicher abzurechnender Gerätetypen.

//go:generate stringer --type Geraetemerkmal
//go:generate jsonenums --type Geraetemerkmal
type Geraetemerkmal int

const (
	EINTARIF             Geraetemerkmal = iota + 1 //     Eintarifzähler
	ZWEITARIF                                      //     Zweitarifzähler
	MEHRTARIF                                      //     Mehrtarifzähler
	GAS_G2P5                                       //     Gaszähler Größe G2.5
	GAS_G4                                         //     Gaszähler Größe G4
	GAS_G6                                         //     Gaszähler Größe G6
	GAS_G10                                        //     Gaszähler Größe G10
	GAS_G16                                        //     Gaszähler Größe G16
	GAS_G25                                        //     Gaszähler Größe G25
	GAS_G40                                        //     Gaszähler Größe G40
	GAS_G65                                        //     Gaszähler Größe G65
	GAS_G100                                       //     Gaszähler Größe G100
	GAS_G160                                       //     Gaszähler Größe G160
	GAS_G250                                       //     Gaszähler Größe G250
	GAS_G400                                       //     Gaszähler Größe G400
	GAS_G650                                       //     Gaszähler Größe G650
	GAS_G1000                                      //     Gaszähler Größe G1000
	GAS_G1600                                      //     Gaszähler Größe G1600
	GAS_G2500                                      //     Gaszähler Größe G2500
	IMPULSGEBER_G4_G100                            //     Impulsgeber für Zähler G4 - G100
	IMPULSGEBER_G100                               //     Impulsgeber für Zähler größer G100
	MODEM_GSM                                      //     Modem-GSM
	MODEM_GPRS                                     //     Modem-GPRS
	MODEM_FUNK                                     //     Modem-FUNK
	MODEM_GSM_O_LG                                 //     vom Messstellenbetreiber beigestelltes GSM-Modem ohne Lastgangmessung
	MODEM_GSM_M_LG                                 //     vom Messstellenbetreiber beigestelltes GSM-Modem mit Lastgangmessung
	MODEM_FESTNETZ                                 //     vom Messstellenbetreiber beigestelltes Festnetz-Modem
	MODEM_GPRS_M_LG                                //     vom Messstellenbetreiber bereitgestelltes GPRS-Modem Lastgangmessung
	PLC_COM                                        //     PLC-Kom.-Einrichtung (Powerline Communication)
	ETHERNET_KOM                                   //     Ethernet-Kom.-Einricht. LAN/WLAN
	DSL_KOM                                        //     Ethernet-Kom.-Einricht. DSL
	LTE_KOM                                        //     Ethernet-Kom.-Einricht. LTE
	RUNDSTEUEREMPFAENGER                           //     Rundsteuerempfänger
	TARIFSCHALTGERAET                              //     Tarifschaltgerät
	ZUSTANDS_MU                                    //     Zustandsmengenumwerter
	TEMPERATUR_MU                                  //     Temperaturmengenumwerter
	KOMPAKT_MU                                     //     Kompaktmengenumwerter
	SYSTEM_MU                                      //     Systemmengenumwerter
	UNBESTIMMT                                     //     Unbestimmtes Merkmal
	WASSER_MWZW                                    //  Wasserzähler Größe MWZW Meßkapsel Wohnungswasserzähler
	WASSER_WZWW                                    //  Wasserzähler Größe WZWW Wohnungswasserzähler
	WASSER_WZ01                                    //  Wasserzähler Größe WZ01 Wasserzähler W01 5 m³/h
	WASSER_WZ02                                    //  Wasserzähler Größe WZ02 Wasserzähler W02 10 m³/h
	WASSER_WZ03                                    //  Wasserzähler Größe WZ03 Wasserzähler W03 20 m³/h
	WASSER_WZ04                                    //  Wasserzähler Größe WZ04 Wasserzähler W04 30 m³/h
	WASSER_WZ05                                    //  Wasserzähler Größe WZ05 Wasserzähler W05 80 m³/h
	WASSER_WZ06                                    //  Wasserzähler Größe WZ06 Wasserzähler W06 120 m³/h
	WASSER_WZ07                                    //  Wasserzähler Größe WZ07 Wasserzähler W07 300 m³/h
	WASSER_WZ08                                    //  Wasserzähler Größe WZ08 Wasserzähler W08 180 m³/h
	WASSER_WZ09                                    //  Wasserzähler Größe WZ09 Wasserzähler W09 140 m³/h
	WASSER_WZ10                                    //  Wasserzähler Größe WZ10 Wasserzähler W10 600 m³/h
	WASSER_VWZ04                                   //  Wasserzähler Größe VWZ04 Verbundwasserzähler 30m³/h
	WASSER_VWZ05                                   //  Wasserzähler Größe VWZ05 Verbundwasserzähler 80m³/h
	WASSER_VWZ06                                   //  Wasserzähler Größe VWZ06 Verbundwasserzähler 120m³/h
	WASSER_VWZ07                                   //  Wasserzähler Größe VWZ07 Verbundwasserzähler 300m³/h
	WASSER_VWZ10                                   //  Wasserzähler Größe VWZ10 Verbundwasserzähler 600m³/h
	GAS_G350                                       //  Gaszähler Größe G350
	GAS_G4000                                      //  Gaszähler Größe G4000
	GAS_G6500                                      //  Gaszähler Größe G6500
	GAS_G10000                                     //  Gaszähler Größe G10000
	GAS_G12500                                     //  Gaszähler Größe G12500
	GAS_G16000                                     //  Gaszähler Größe G16000
)

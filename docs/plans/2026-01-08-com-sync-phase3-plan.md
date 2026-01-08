# Phase 3: Add Missing COM Structs

## Overview
Add COM structs that exist in C# BO4E.net but are entirely missing in Go.

## Reference Commit
`05dad017754da47d901ddb37c43aa609ec661465`

## Rules (from Phase 2 learnings)
1. Match nullability: C# `?` → Go `*pointer`, C# non-nullable → Go value type
2. Copy docstrings from C# library (German)
3. Run `gofmt` before committing
4. Run `go build ./...` to verify compilation
5. Use existing enum types where available
6. **Verify permalinks** using `curl -s <raw-url> | grep -n "JsonPropertyName"`
7. Branch naming: `add/com-<lowercase-name>`

## Missing COM Structs (30 total)

### Priority 1: Simple structs (few fields, no complex dependencies)
| Struct | Fields | Dependencies |
|--------|--------|--------------|
| Notiz | 2 | None |
| PhysikalischerWert | 2 | Mengeneinheit enum |
| GenericStringStringInfo | 2 | None |
| Aufgabe | 3 | Aufgabentyp enum |
| Rechenschritt | 3 | Enums |

### Priority 2: Medium complexity
| Struct | Fields | Dependencies |
|--------|--------|--------------|
| Geraet | ~5 | Geraetetyp, Geraetemerkmal enums |
| Energieherkunft | ~4 | Erzeugungsart enum |
| Energiemix | ~5 | Sparte, Oekozertifikat enums |
| AufAbschlag | ~5 | AufAbschlagstyp, Waehrungseinheit enums |
| Preisgarantie | ~4 | Preisgarantietyp enum |
| Kriteriumswert | 2 | Tarifregionskriterium enum |
| Kostenposition | ~6 | Kostenklasse enum, Betrag COM |
| Kostenblock | ~4 | Kostenart enum, Kostenposition COM |
| LokationsTypZuordnung | ~3 | Lokationstyp enum |
| Netzlokationsabrechnungsdaten | ~5 | Enums |

### Priority 3: Offer/Tender related
| Struct | Fields | Dependencies |
|--------|--------|--------------|
| Angebotsposition | ~5 | Enums |
| Angebotsteil | ~5 | Angebotsposition COM |
| Ausschreibungsdetail | ~8 | Multiple enums |
| Ausschreibungslos | ~10 | Ausschreibungsdetail COM |

### Priority 4: Tariff/Regional related
| Struct | Fields | Dependencies |
|--------|--------|--------------|
| Tarifpreisposition | ~6 | Multiple enums, Preisstaffel COM |
| Tarifeinschraenkung | ~4 | Enums |
| Tarifberechnungsparameter | ~8 | Complex |
| RegionaleGueltigkeit | ~3 | Regionskriterium COM |
| Regionskriterium | ~3 | Enums |
| RegionalePreisstaffel | ~4 | Preisstaffel, RegionaleGueltigkeit COMs |
| RegionalePreisgarantie | ~4 | Preisgarantie, RegionaleGueltigkeit COMs |
| RegionalerAufAbschlag | ~5 | AufAbschlag, RegionaleGueltigkeit COMs |
| RegionaleTarifpreisposition | ~5 | Tarifpreisposition, RegionaleGueltigkeit COMs |

### Priority 5: Misc
| Struct | Fields | Dependencies |
|--------|--------|--------------|
| RechnungspositionFlat | ~15 | Flattened version of Rechnungsposition |
| Zeitreihenprodukt | ~3 | Zeitreihenwertkompakt COM |

## Workflow per Struct

1. Fetch C# source: `curl -s "https://raw.githubusercontent.com/Hochfrequenz/BO4E-dotnet/05dad017754da47d901ddb37c43aa609ec661465/BO4E/COM/<Name>.cs"`
2. Get line numbers: `curl -s <url> | grep -n "JsonPropertyName"`
3. Check enum dependencies exist in `enum/` folder
4. Create branch: `git checkout -b add/com-<name>`
5. Create Go file: `com/<name>.go`
6. Run `gofmt -w com/<name>.go`
7. Run `go build ./...`
8. Commit with verified permalinks
9. Push and note PR link

## Progress Tracking

| Struct | Status | Branch | PR |
|--------|--------|--------|-----|
| Sigmoidparameter | ✅ Done (Phase 2) | add/com-sigmoidparameter | pending |
| Notiz | ✅ Done | add/com-notiz | pending |
| PhysikalischerWert | ✅ Done | add/com-physikalischerwert | pending |
| GenericStringStringInfo | ✅ Done | add/com-genericstringstringinfo | pending |
| Aufgabe | ✅ Done | add/com-aufgabe | pending |
| Geraet | ✅ Done | add/com-geraet | pending |
| Energieherkunft | ⏳ Pending | | |
| Energiemix | ⏳ Pending | | |
| AufAbschlag | ⏳ Pending | | |
| ... | | | |

## Verified Permalinks (curl -s ... | sed -n 'Np')
All permalinks verified to point to correct `[JsonPropertyName("...")]` lines.

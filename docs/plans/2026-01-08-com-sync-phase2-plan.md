# Phase 2: Add Missing Fields to COM Structs

## Status: COMPLETED

## Overview
Add fields that exist in C# BO4E.net COM classes but are missing in Go.

## Rules
1. Match nullability: C# `?` → Go `*pointer`
2. Copy docstrings from C# library
3. Run `gofmt` before committing
4. Run tests before pushing (at least roundtrip tests)
5. Use existing enum types where available

## Learnings from Phase 2

### Permalink Verification is Critical
- **Never trust WebFetch summaries** for line numbers - they hallucinate
- **Always verify** using `curl -s <raw-url> | grep -n "JsonPropertyName"` to get exact line numbers
- Permalinks must point to the `[JsonPropertyName("...")]` attribute line, not the property or summary

### Correct Permalink Format
```
https://github.com/Hochfrequenz/BO4E-dotnet/blob/<commit-sha>/BO4E/COM/<File>.cs#L<line>
```
- Use specific commit SHA (not `main`) for stable references
- Reference commit: `05dad017754da47d901ddb37c43aa609ec661465`

### Branch Naming Convention
- Phase 1 (fixes): `fix/com-<name>-serialization`
- Phase 2 (additions): `add/com-<name>-fields` or `feat/com-<name>-missing-fields`
- Phase 3 (new structs): `add/com-<name>`

### Commit Message Template
```
feat(com): add missing fields to <StructName>

Add fields from BO4E.net that were missing in Go:
- <Field1>: <description>
- <Field2>: <description>

C# reference (commit <sha>):
- <Field1>: <permalink>#L<verified-line>
- <Field2>: <permalink>#L<verified-line>

Co-Authored-By: Claude Opus 4.5 <noreply@anthropic.com>
```

## Structs with Missing Fields

### 1. Verbrauch (has test file)
Branch: `feat/com-verbrauch-missing-fields` from `main`
Missing fields:
- `Type`: `Verbrauchsmengetyp?` → `*verbrauchsmengetyp.Verbrauchsmengetyp` (json: "type")
- `Zeitfenster`: `Zeitfenster?` → `*Zeitfenster` (json: "zeitfenster")
- `Temperaturmasszahl`: `string?` → `*string` (json: "temperaturmasszahl")

### 2. Zaehlwerk (has test file)
Branch: `feat/com-zaehlwerk-missing-fields` from `main`
Missing fields:
- `Kennzahl`: `string?` → `*string` (json: "kennzahl")
- `Steuerbefreit`: `bool?` → `*bool` (json: "steuerbefreit")
- `Abrechnungsrelevant`: `bool?` → `*bool` (json: "abrechnungsrelevant")
- `EMobilitaetsart`: `EMobilitaetsart?` → `*emobilitaetsart.EMobilitaetsart` (json: "emobilitaetsart")
- `Verbrauchsarten`: `List<Verbrauchsart>?` → `[]verbrauchsart.Verbrauchsart` (json: "Verbrauchsarten")

### 3. Rechnungsposition (has test file)
Branch: `feat/com-rechnungsposition-missing-fields` from `main`
Missing fields:
- `VertragskontoId`: `string?` → `*string` (json: "vertragskontoId")
- `VertragsId`: `string?` → `*string` (json: "vertragsId")
- `Status`: `RechnungspositionsStatus?` → `*rechnungspositionsstatus.RechnungspositionsStatus` (json: "status")

### 4. Geokoordinaten (no test file - need to create)
Branch: `feat/com-geokoordinaten-missing-fields` from `main`
Missing fields:
- `OestlicheLaenge`: `decimal?` → `*decimal.Decimal` (json: "oestlicheLaenge")
- `NoerdlicheBreite`: `decimal?` → `*decimal.Decimal` (json: "noerdlicheBreite")
- `Zone`: `string?` → `*string` (json: "zone")
- `NordWert`: `string?` → `*string` (json: "nordWert")
- `OstWert`: `string?` → `*string` (json: "ostWert")

### 5. Lastprofil (has test file)
Branch: `feat/com-lastprofil-missing-fields` from `main`
Missing fields:
- `Profilschar`: `string?` → `*string` (json: "profilschar")
- `Profiltyp`: `Profilklasse?` → `*profilklasse.Profilklasse` (json: "profiltyp")
- `Normierungsfaktor`: `Normierungsfaktor?` → `*normierungsfaktor.Normierungsfaktor` (json: "normierungsfaktor")

### 6. Preisstaffel (no test file - need to create)
Branch: `feat/com-preisstaffel-missing-fields` from `main`
Missing fields:
- `Sigmoidparameter`: `Sigmoidparameter?` → `*Sigmoidparameter` (json: "sigmoidparameter")
Note: Sigmoidparameter COM struct may not exist in Go yet - check Phase 3

### 7. Hardware (has Phase 1 branch)
Branch: `feat/com-hardware-missing-fields` from `fix/com-hardware-serialization`
Missing fields:
- `Geraeteeigenschaften`: `Geraeteeigenschaften?` → `*Geraeteeigenschaften` (json: "geraeteeigenschaften")
- `Geraetenummer`: `string?` → `*string` (json: "geraetenummer")
- `Geraetereferenz`: `string?` → `*string` (json: "geraetereferenz")

### 8. Vertragskonditionen (has test file)
Branch: `feat/com-vertragskonditionen-missing-fields` from `main`
Missing fields:
- `NetznutzungsabrechnungIntervall`: `int?` → `*int` (json: "netznutzungsabrechnungIntervall")
- `BeinhaltetSingulaerGenutzteBetriebsmittel`: `bool?` → `*bool` (json: "beinhaltetSingulaerGenutzteBetriebsmittel")

## Workflow per Struct

1. Create branch from appropriate base
2. Fetch C# source to get exact docstrings and line numbers
3. Add missing fields with:
   - Correct Go type (matching nullability)
   - Correct json tag
   - Docstring from C#
4. Run `gofmt -w`
5. Update/create roundtrip test
6. Run `go test ./com/...`
7. Commit with permalink to C# source
8. Push and create PR

## Dependencies to Check

- `verbrauchsmengetyp` enum - verify exists
- `emobilitaetsart` enum - verify exists
- `rechnungspositionsstatus` enum - verify exists
- `normierungsfaktor` enum - verify exists
- `Sigmoidparameter` COM struct - may need Phase 3
- `Geraeteeigenschaften` COM struct - verify exists

<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes.
"State Machine Breaking" for breaking the AppState

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [v1.0.0] - 2023-04-24

This version base on [finschia-sdk v0.47.0](https://github.com/Finschia/finschia-sdk/releases/tag/v0.47.0), [Ostracon v1.1.0](https://github.com/Finschia/ostracon/tree/v1.1.0), [finschia/wasmd v0.1.3](https://github.com/Finschia/wasmd/releases/tag/v0.1.3) and [finschia/ibc-go v3.3.3](https://github.com/Finschia/ibc-go/releases/tag/v3.3.3).

### Features
* (build) [\#126](https://github.com/Finschia/finschia/pull/126) Automatically generates release note and binaries
* (x/wasmd) [\#129](https://github.com/Finschia/finschia/pull/129) chore: apply detached x/wasmd
* (build) [\#130](https://github.com/Finschia/finschia/pull/130) Add a release build for the linux/arm64, darwin/amd64, and darwin/arm64 platform
* (finschia-sdk) [\#137](https://github.com/Finschia/finschia/pull/137) Bump finschia-sdk to 6c84a4cffa
* (x/collection,token) [\#138](https://github.com/Finschia/finschia/pull/138) Add x/token and x/collection
* (ibc-go) [\#140](https://github.com/Finschia/finschia/pull/140) apply ibc-go
* (x/wasmplus) [\#141](https://github.com/Finschia/finschia/pull/141) change wasm module to wrapped `x/wasmplus`
* (finschia-sdk) [\#144](https://github.com/Finschia/finschia/pull/144) bump finschia-sdk v0.47.0-alpha1 (11966d1234155ebef20b64f2ae7a905beffdb33f)
* (build) [\#150](https://github.com/Finschia/finschia/pull/150) Modify the Makefile to build release bundles
* (build) [\#153](https://github.com/Finschia/finschia/pull/153) rename cli name to `fnsad`
* (finschia-sdk) [\#154](https://github.com/Finschia/finschia/pull/154) Bump finschia-sdk from v0.47.0-alpha1.0.20230214070148-11966d123415 to v0.47.0-rc1

### Improvements
* (x/wasmd) [\#147](https://github.com/Finschia/finschia/pull/147) update wasmd version
* (x/wasmd) [\#158](https://github.com/Finschia/finschia/pull/158) bump up wasmd version to v0.1.0
* (finschia-sdk) [\#159](https://github.com/Finschia/finschia/pull/159) Bump finschia-sdk from v0.47.0-rc1 to v0.47.0-rc2
* (wasmd) [\#171](https://github.com/Finschia/finschia/pull/171) bump up wasmd from v0.1.2-0.20230403061848-514953c0b244 to v0.1.2

### Bug Fixes
* (finschia-sdk) [\#162](https://github.com/Finschia/finschia/pull/162) Bump finschia-sdk from v0.47.0-rc2 to v0.47.0-rc3
* (finschia-sdk) [\#167](https://github.com/Finschia/finschia/pull/167) Bump finschia-sdk from v0.47.0-rc3 to v0.47.0-rc4
* (finschia-sdk) [\#178](https://github.com/Finschia/finschia/pull/178) Bump github.com/Finschia/finschia-sdk from v0.47.0-rc6 to v0.47.0-rc7
* (finschia-sdk) [\#168](https://github.com/Finschia/finschia/pull/168) Bump finschia-sdk from v0.47.0-rc4 to v0.47.0-rc4.0.20230410115856-b8421116b3f2
* (finschia-sdk) [\#172](https://github.com/Finschia/finschia/pull/172) Bump finschia-sdk from v0.47.0-rc4.0.20230410115856-b8421116b3f2 to v0.47.0-rc5
* (finschia-sdk) [\#174](https://github.com/Finschia/finschia/pull/174) bump up finschia-sdk from v0.47.0-rc5 to v0.47.0-rc5.0.20230414034539-489c442416cd

### Breaking Changes
* (api) [\#123](https://github.com/Finschia/finschia/pull/123) remove legacy REST API routes
* (ibc-go)[\#164](https://github.com/Finschia/finschia/pull/164) bump up ibc-go v3.3.2 for change ibc light client of Ostracon to Tendermint

### Build, CI
* (ci) [\#145](https://github.com/Finschia/finschia/pull/145) add github action to push docker image to docker.io
* (build) [\#157](https://github.com/Finschia/finschia/pull/157) add build args
* (ci)[\#163](https://github.com/Finschia/finschia/pull/163) fix `release-build` ci error occurred when adding assets after tagging
* (finschia-sdk,ostracon,wasmd,ibc-go) [\#176](https://github.com/Finschia/finschia/pull/176) Rename and bump up dependencies
* (ci) [\#180](https://github.com/Finschia/finschia/pull/180) add a CI to build darwin version when add tag

### Docs
* (doc) [\#156](https://github.com/Finschia/finschia/pull/156) modify broken link or typo doc file and add issue and pr template
* (doc)[\#165](https://github.com/Finschia/finschia/pull/165) update license notice and code_of_conduct
* (license) [\#170](https://github.com/Finschia/finschia/pull/170) fix license copyright holder and typo


## [v0.7.0] - 2022-11-29

This version base on [finschia-sdk v0.46.0](https://github.com/Finschia/finschia-sdk/releases/tag/v0.46.0).

### Features
* [\#108](https://github.com/Finschia/finschia/pull/108) Bump github.com/line/lbm-sdk from e19f863a8 to a389b6330
* [\#110](https://github.com/Finschia/finschia/pull/110) apply GovMint on x/foundation
* [\#111](https://github.com/Finschia/finschia/pull/111) Bump github.com/line/lbm-sdk from 66988a235 to 0.46.0-rc9
* (build) [\#113](https://github.com/Finschia/finschia/pull/113) enable to use libsodium version ostracon

### Improvements

### Bug Fixes
* (global) [\#103](https://github.com/Finschia/finschia/pull/103) replace deprecated functions of ioutil package
* (app) [\#107](https://github.com/Finschia/finschia/pull/107) change module order in `init genesis`
* (ci) [\#115](https://github.com/Finschia/finschia/pull/115) fix test flow to install libsodium
* (build) [\#118](https://github.com/Finschia/finschia/pull/118) fix docker build in Mac M1 device

### Improvements
* (app) [\#114](https://github.com/Finschia/finschia/pull/114) change the default compile setting to support ledger


## [v0.6.0]

This version based on [finschia-sdk v0.46.0-rc8](https://github.com/Finschia/finschia-sdk/releases/tag/v0.46.0-rc8)

### Bug Fixes
* (app) [\#96](https://github.com/Finschia/finschia/pull/96) fix the bug not setting `iavl-cache-size` value of the `app.toml`
* (baseapp) [\#97](https://github.com/Finschia/finschia/pull/97) fix max gas validation bug of lbm-sdk

### Breaking Changes
* (app) [\#87](https://github.com/Finschia/finschia/pull/87) remove unused modules from app
* (finschia-sdk) [\#95](https://github.com/Finschia/finschia/pull/95) apply the changes of lbm-sdk v0.46.0-rc8

### Build, CI
* (ci) [\#80](https://github.com/Finschia/finschia/pull/80) remove stale github action
* (ci) [\#83](https://github.com/Finschia/finschia/pull/83) enable tests on CI
* (build) [\#89](https://github.com/Finschia/finschia/pull/89) upgrade golang to 1.18


## [v0.5.0]

This version based on [finschia-sdk v0.46.0-rc6](https://github.com/Finschia/finschia-sdk/releases/tag/v0.46.0-rc6)

### Features
* (x/collection) [\#72](https://github.com/Finschia/finschia/pull/72) add x/collection
* (x/wasm) [\#79](https://github.com/Finschia/finschia/pull/79) chore: add iterator feature for wasm module

### Improvements
* (ci) [\#76](https://github.com/Finschia/finschia/pull/76) fix Dockerfile.static to build lbm instead of building wasmvm in the Dockerfile

### Bug Fixes
* (command) [\#81](https://github.com/Finschia/finschia/pull/81) add wrong address to genesis file in add-genesis-account command
* (x/collection) [\#86](https://github.com/Finschia/finschia/pull/86) add omitted cli commands on x/collection and fix Query/Balance
* (x/collection) [\#90](https://github.com/Finschia/finschia/pull/90) fix bugs in x/collection MsgModify


## [v0.4.0]

This version based on finschia-sdk v0.46.0-rc2

### Features
* (cosmos-sdk) [\#56](https://github.com/Finschia/finschia/pull/56) bump up cosmos-sdk v0.45.1
* (x/foundation) [\#62](https://github.com/Finschia/finschia/pull/62) add `x/foundation` module of lbm-sdk

### Improvements

### Bug Fixes
* (app) [\#60](https://github.com/Finschia/finschia/pull/60) register authz module store key


## [v0.3.0]

### Bug Fixes
* (build) [\#47](https://github.com/Finschia/finschia/pull/47) fix Docker build error

### Features
* (x/wasm) [\#41](https://github.com/Finschia/finschia/pull/41) upgrade x/wasm (merged original 0.19.0)
* (x/upgrade) [\#42](https://github.com/Finschia/finschia/pull/42) add token module and bump cosmos-sdk v0.42.11


## [v0.2.0]

### Features
* (x/upgrade) [\#33](https://github.com/Finschia/finschia/pull/33) To smoothen the update to the latest stable release, the SDK includes version map for managing migrations between SDK versions.
* (x/consortium) [\#34](https://github.com/Finschia/finschia/pull/34) add feegrant, consortium and stakingplus module
* (x/bank) [\#36](https://github.com/Finschia/finschia/pull/36) apply a feature that preventing sending coins to inactive contract (related to [finschia-sdk #400](https://github.com/Finschia/finschia-sdk/pull/400))

### Improvements
* (slashing) [\#31] (https://github.com/Finschia/finschia/pull/31) Apply VoterSetCounter

## [v0.1.0]

### Features
* (app) Revise bech32 prefix cosmos to link and tlink

### Improvements
* (sdk) Use fastcache for inter block cache and iavl cache
* (sdk) Enable signature verification cache
* (ostracon) Apply asynchronous receiving reactor
* (sdk) [\#21](https://github.com/Finschia/finschia/pull/21) Use lbm-sdk v0.43.0

### Bug Fixes

### Breaking Changes
* (sdk) (auth) [\#16](https://github.com/Finschia/lfb/pull/16) Introduce sig block height for the new replay protection
* (ostracon/sdk) [\#26](https://github.com/Finschia/lfb/pull/26) Use vrf-based consensus, address string treatment
* (global) [\#10](https://github.com/Finschia/finschia/pull/10) Re-brand lfb to lbm

### Build, CI
* (build) [\#25](https://github.com/Finschia/finschia/pull/25) Fix localnet-start

## [gaia v4.0.4] - 2021-03-15
Initial lbm is based on the tendermint v0.34.9+, cosmos-sdk v0.42.0+, gaia v4.0.4

* (tendermint) [v0.34.9](https://github.com/tendermint/tendermint/releases/tag/v0.34.9).
* (cosmos-sdk) [v0.42.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.42.0).
* (gaia) [v4.0.4](https://github.com/cosmos/gaia/releases/tag/v4.0.4).

Please refer [CHANGELOG_OF_GAIA_v4.0.4](https://github.com/cosmos/gaia/blob/v4.0.4/CHANGELOG.md)


<!-- Release links -->
[v1.0.0]: https://github.com/Finschia/finschia/releases/tag/v1.0.0
[v0.7.0]: https://github.com/Finschia/finschia/releases/tag/v0.7.0
[v0.6.0]: https://github.com/Finschia/finschia/releases/tag/v0.6.0
[v0.5.0]: https://github.com/Finschia/finschia/releases/tag/v0.5.0
[v0.4.0]: https://github.com/Finschia/finschia/releases/tag/v0.4.0
[v0.3.0]: https://github.com/Finschia/finschia/releases/tag/v0.3.0
[v0.2.0]: https://github.com/Finschia/finschia/releases/tag/v0.2.0
[v0.1.0]: https://github.com/Finschia/finschia/releases/tag/v0.1.0
[gaia v4.0.4]: https://github.com/cosmos/gaia/releases/tag/v4.0.4

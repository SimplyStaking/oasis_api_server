## 1.0.4

Released on 17th August 2020

#### Staking

~ Changed Response type from `[]staking_api.Event` to `[]*staking_api.Event`

#### Registry

~ Changed Response type from `[]registry_api.Event` to `[]*registry_api.Event`
~ Changed GetRuntimes to have a query parameter and it now takes suspended boolean value in the URL.

#### Other

* Added to docs the exact query parameters that are needed to return data from the server.
* Updated tendermint version to 0.34

## 1.0.3

Released on 19th June 2020

### Added

#### Staking

* GetLastBlockFees Handler at /api/staking/lastblockfees
* GetAddressFromPublicKey Handler at /api/staking/publickeytoaddress

#### Consensus:

* GetStatus Handler at /api/consensus/status
* GetGenesisDocument Handler at /api/consensus/genesisdocument

#### Registry:

* GetNodeStatus Handler at /api/registry/nodestatus
* GetRegistryEvents Handler at /api/registry/events

### Updated

#### Staking

* GetAccounts at /api/staking/accounts has been changed to GetAddresses at /api/staking/addresses
* GetAccountInfo at /api/staking/accountinfo has been changed to GetAccount at /api/staking/account, and "ownerKey" query parameter has been changed to "address"
* GetDelegations "ownerKey" query parameter has been changed to "address"
* GetDebondingDelegations "ownerKey" query parameter has been changed to "address"


## 1.0.2

Released on May 2020

### Changed

*  Changed endpoints by removing trailing `/` E.G `/api/ping/` is now `/api/ping`


## 1.0.1

Released on May 2020

### Changed

* Updated Tendermint Version together with Oasis-Core version.
* This update changes blocklast commits from Precommits to Signatures

## 1.0.0

Released on May 2020

### Added

* First version of the Oasis API Server
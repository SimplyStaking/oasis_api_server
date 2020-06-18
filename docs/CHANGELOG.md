## 1.0.3
Released on June 2020

ADDED:
    Staking:
	    GetLastBlockFees Handler at /api/staking/lastblockfees
	
	Consensus:
		GetStatus Handler at /api/consensus/status
		GetGenesisDocument Handler at /api/consensus/genesisdocument

	Registry:
		GetNodeStatus Handler at /api/registry/nodestatus
		GetRegistryEvents Handler at /api/registry/events

REMOVED
	

UPDATED
	Staking:
		GetAccounts:
			GetAccounts at /api/staking/accounts has been changed to GetAddresses at /api/staking/addresses
		GetAccountInfo:
			GetAccountInfo at /api/staking/accountinfo has been changed to GetAccount at /api/staking/account
			ownerKey query has been changed to address
		GetDelegations:
			ownerKey query has been changed to address
		GetDebondingDelegations:
			ownerKey query has been changed to address


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
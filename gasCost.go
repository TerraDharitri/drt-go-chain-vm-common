package vmcommon

// BaseOperationCost defines cost for base operation cost
type BaseOperationCost struct {
	StorePerByte      uint64
	ReleasePerByte    uint64
	DataCopyPerByte   uint64
	PersistPerByte    uint64
	CompilePerByte    uint64
	AoTPreparePerByte uint64
}

// BuiltInCost defines cost for built-in methods
type BuiltInCost struct {
	ChangeOwnerAddress       uint64
	ClaimDeveloperRewards    uint64
	SaveUserName             uint64
	SaveKeyValue             uint64
	DCDTTransfer             uint64
	DCDTBurn                 uint64
	DCDTLocalMint            uint64
	DCDTLocalBurn            uint64
	DCDTModifyRoyalties      uint64
	DCDTModifyCreator        uint64
	DCDTNFTCreate            uint64
	DCDTNFTRecreate          uint64
	DCDTNFTUpdate            uint64
	DCDTNFTAddQuantity       uint64
	DCDTNFTBurn              uint64
	DCDTNFTTransfer          uint64
	DCDTNFTChangeCreateOwner uint64
	DCDTNFTMultiTransfer     uint64
	DCDTNFTAddURI            uint64
	DCDTNFTSetNewURIs        uint64
	DCDTNFTUpdateAttributes  uint64
	SetGuardian              uint64
	GuardAccount             uint64
	TrieLoadPerNode          uint64
	TrieStorePerNode         uint64
}

// GasCost holds all the needed gas costs for system smart contracts
type GasCost struct {
	BaseOperationCost BaseOperationCost
	BuiltInCost       BuiltInCost
}

// SafeSubUint64 performs subtraction on uint64 and returns an error if it overflows
func SafeSubUint64(a, b uint64) (uint64, error) {
	if a < b {
		return 0, ErrSubtractionOverflow
	}
	return a - b, nil
}

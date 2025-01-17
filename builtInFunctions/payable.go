package builtInFunctions

import (
	"bytes"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data/vm"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
)

type payableCheck struct {
	payableHandler      vmcommon.PayableHandler
	enableEpochsHandler vmcommon.EnableEpochsHandler
}

// NewPayableCheckFunc returns a new component which checks if destination is payableCheck when needed
func NewPayableCheckFunc(
	payable vmcommon.PayableHandler,
	enableEpochsHandler vmcommon.EnableEpochsHandler,
) (*payableCheck, error) {
	if check.IfNil(payable) {
		return nil, ErrNilPayableHandler
	}
	if check.IfNil(enableEpochsHandler) {
		return nil, ErrNilEnableEpochsHandler
	}

	return &payableCheck{
		payableHandler:      payable,
		enableEpochsHandler: enableEpochsHandler,
	}, nil
}

func (p *payableCheck) mustVerifyPayable(vmInput *vmcommon.ContractCallInput, minLenArguments int) bool {
	typeToVerify := vm.AsynchronousCall
	if p.enableEpochsHandler.IsFlagEnabled(FixAsyncCallbackCheckFlag) {
		typeToVerify = vm.AsynchronousCallBack
		if vmInput.ReturnCallAfterError {
			return false
		}
	}
	if vmInput.CallType == typeToVerify || vmInput.CallType == vm.DCDTTransferAndExecute {
		return false
	}
	if bytes.Equal(vmInput.CallerAddr, core.DCDTSCAddress) {
		return false
	}
	if len(vmInput.Arguments) > minLenArguments {
		if p.enableEpochsHandler.IsFlagEnabled(CheckFunctionArgumentFlag) {
			if len(vmInput.Arguments[minLenArguments]) > 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// CheckPayable returns error if the destination account a non-payable smart contract and there is no sc call after transfer
func (p *payableCheck) CheckPayable(vmInput *vmcommon.ContractCallInput, dstAddress []byte, minLenArguments int) error {
	if !p.mustVerifyPayable(vmInput, minLenArguments) {
		return nil
	}

	isPayable, errIsPayable := p.payableHandler.IsPayable(vmInput.CallerAddr, dstAddress)
	if errIsPayable != nil {
		return errIsPayable
	}
	if !isPayable {
		return ErrAccountNotPayable
	}

	return nil
}

// DetermineIsSCCallAfter returns true if there is a smart contract call after execution
func (p *payableCheck) DetermineIsSCCallAfter(vmInput *vmcommon.ContractCallInput, destAddress []byte, minLenArguments int) bool {
	if len(vmInput.Arguments) <= minLenArguments {
		return false
	}
	if vmInput.ReturnCallAfterError && vmInput.CallType != vm.AsynchronousCallBack {
		return false
	}
	if !vmcommon.IsSmartContractAddress(destAddress) {
		return false
	}
	if p.enableEpochsHandler.IsFlagEnabled(CheckFunctionArgumentFlag) {
		if len(vmInput.Arguments[minLenArguments]) == 0 {
			return false
		}
	}

	return true
}

// IsInterfaceNil returns true if underlying object is nil
func (p *payableCheck) IsInterfaceNil() bool {
	return p == nil
}

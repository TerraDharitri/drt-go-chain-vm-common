package builtInFunctions

import (
	"testing"

	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/stretchr/testify/require"
)

func createGuardAccountArgs() GuardAccountArgs {
	return GuardAccountArgs{
		BaseAccountGuarderArgs: createBaseAccountGuarderArgs(),
	}
}

func TestBaseGuardAccount_SetNewGasConfig(t *testing.T) {
	t.Parallel()

	args := createGuardAccountArgs()
	baseGuardAccount, _ := newBaseGuardAccount(args)
	require.Equal(t, args.FuncGasCost, baseGuardAccount.funcGasCost)

	newGuardAccountCost := args.FuncGasCost + 1
	newGasCost := &vmcommon.GasCost{BuiltInCost: vmcommon.BuiltInCost{GuardAccount: newGuardAccountCost}}

	baseGuardAccount.SetNewGasConfig(newGasCost)
	require.Equal(t, newGuardAccountCost, baseGuardAccount.funcGasCost)
}

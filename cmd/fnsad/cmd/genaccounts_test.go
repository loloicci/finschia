package cmd_test

import (
	"context"
	"fmt"

	"testing"

	"github.com/Finschia/finschia-sdk/crypto/hd"
	"github.com/Finschia/finschia-sdk/crypto/keyring"
	sdk "github.com/Finschia/finschia-sdk/types"

	"github.com/Finschia/ostracon/libs/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/Finschia/finschia-sdk/client"
	"github.com/Finschia/finschia-sdk/client/flags"
	"github.com/Finschia/finschia-sdk/server"
	"github.com/Finschia/finschia-sdk/simapp"
	simcmd "github.com/Finschia/finschia-sdk/simapp/simd/cmd"
	"github.com/Finschia/finschia-sdk/testutil/testdata"
	"github.com/Finschia/finschia-sdk/types/module"
	"github.com/Finschia/finschia-sdk/x/genutil"
	genutiltest "github.com/Finschia/finschia-sdk/x/genutil/client/testutil"
)

var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})

func TestAddGenesisAccountCmd(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	tests := []struct {
		name        string
		addr        string
		denom       string
		withKeyring bool
		expectErr   bool
	}{
		{
			name:        "invalid address",
			addr:        "",
			denom:       "1000atom",
			withKeyring: false,
			expectErr:   true,
		},
		{
			name:        "valid address",
			addr:        addr1.String(),
			denom:       "1000atom",
			withKeyring: false,
			expectErr:   false,
		},
		{
			name:        "multiple denoms",
			addr:        addr1.String(),
			denom:       "1000atom, 2000stake",
			withKeyring: false,
			expectErr:   false,
		},
		{
			name:        "with keyring",
			addr:        "ser",
			denom:       "1000atom",
			withKeyring: true,
			expectErr:   false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			logger := log.NewNopLogger()
			cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
			require.NoError(t, err)

			appCodec := simapp.MakeTestEncodingConfig().Marshaler
			err = genutiltest.ExecInitCmd(testMbm, home, appCodec)
			require.NoError(t, err)

			serverCtx := server.NewContext(viper.New(), cfg, logger)
			clientCtx := client.Context{}.WithCodec(appCodec).WithHomeDir(home)

			if tc.withKeyring {
				path := hd.CreateHDPath(118, 0, 0).String()
				kr, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendMemory, home, nil)
				require.NoError(t, err)
				_, _, err = kr.NewMnemonic(tc.addr, keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
				require.NoError(t, err)
				clientCtx = clientCtx.WithKeyring(kr)
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
			ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

			cmd := simcmd.AddGenesisAccountCmd(home)
			cmd.SetArgs([]string{
				tc.addr,
				tc.denom,
				fmt.Sprintf("--%s=home", flags.FlagHome)})

			if tc.expectErr {
				require.Error(t, cmd.ExecuteContext(ctx))
			} else {
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}
}

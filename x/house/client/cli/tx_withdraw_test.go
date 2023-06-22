package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sge-network/sge/testutil/network"
	"github.com/sge-network/sge/x/house/client/cli"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
)

func TestTXWithdrawCLI(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	commonArgs := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String(),
		),
	}

	t.Run("Withdraw", func(t *testing.T) {
		for _, tc := range []struct {
			desc               string
			marketUid          string
			participationIndex uint64
			mode               int32
			amount             string
			ticket             string

			err  error
			code uint32
		}{
			{
				marketUid:          "6e31c60f-2025-48ce-ae79-1dc110f16355",
				amount:             "555",
				participationIndex: 1,
				mode:               1,
				ticket:             "ticket",

				desc: "valid",
			},
			{
				marketUid: "invalidUID",
				amount:    "555",
				ticket:    "ticket",

				desc: "validation failed",
				err:  fmt.Errorf("any error"),
			},
			{
				marketUid: "6e31c60f-2025-48ce-ae79-1dc110f16355",
				amount:    "invalidAmount",
				ticket:    "ticket",

				desc: "invalid amuont",
				err:  fmt.Errorf("any error"),
			},
		} {
			tc := tc
			t.Run(tc.desc, func(t *testing.T) {
				args := []string{
					tc.marketUid,
					cast.ToString(tc.participationIndex),
					tc.ticket,
					cast.ToString(tc.mode),
					tc.amount,
				}
				args = append(args, commonArgs...)
				out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdWithdraw(), args)
				if tc.err != nil {
					require.NotNil(t, err)
				} else {
					require.NoError(t, err)
					var resp sdk.TxResponse
					require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				}
			})
		}
	})
}

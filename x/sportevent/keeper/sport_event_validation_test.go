package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/sge-network/sge/testutil/sample"
	"github.com/sge-network/sge/x/sportevent/types"
	"github.com/stretchr/testify/require"
)

func Test_ValidateCreationEvent(t *testing.T) {
	k, _, wctx, _ := setupMsgServerAndKeeper(t)
	t1 := time.Now()
	params := k.GetParams(wctx)

	tests := []struct {
		name string
		msg  types.SportEvent
		err  error
	}{
		{
			name: "valid request",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				Details: "Winner of x:y",
			},
		}, {
			name: "same timestamp",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute).Unix()),
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "end timestamp before current timestamp",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				EndTS:   uint64(t1.Add(-time.Minute).Unix()),
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid uid",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     "invalid uuid",
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "few odds than required",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid odd id",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: "invalid id", Details: "invalid odds"},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "duplicate odds id",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: "8779cf93-925c-4818-bc81-13c359e0deb8", Details: "Odds 1"},
					{UID: "8779cf93-925c-4818-bc81-13c359e0deb8", Details: "invalid odds"},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid min amount, negative",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				BetConstraints: &types.EventBetConstraints{MinAmount: sdk.NewInt(-5)},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid min amount, less than required",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				BetConstraints: &types.EventBetConstraints{MinAmount: params.EventMinBetAmount.Sub(sdk.NewInt(5))},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid request, with bet constraint",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				BetConstraints: &types.EventBetConstraints{
					MinAmount: params.EventMinBetAmount,
					BetFee:    params.EventMinBetFee,
				},
				Details: "Winner of x:y",
			},
		},
		{
			name: "large details",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				Details: `Winner of x:y is the final winner of the game, 
				it is obvious the winner is not the champion yet but if it happens, 
				the winning users will reward 1M dollars each plus a furnished villa in the Beverley hills as a gift. 
				attention! this detail will not be stored in the chain because it's definitely a scam.`,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.MsgServerValidateCreationEvent(wctx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func Test_ValidateResolveEvent(t *testing.T) {
	k, _, _, _ := setupMsgServerAndKeeper(t)
	t1 := time.Now()

	tests := []struct {
		name string
		msg  types.ResolutionEvent
		err  error
	}{
		{
			name: "valid request",
			msg: types.ResolutionEvent{
				UID:            uuid.NewString(),
				ResolutionTS:   uint64(t1.Unix()),
				WinnerOddsUIDs: []string{uuid.NewString()},
				Status:         4,
			},
		}, {
			name: "invalid resolution ts",
			msg: types.ResolutionEvent{
				UID:            uuid.NewString(),
				ResolutionTS:   0,
				WinnerOddsUIDs: []string{uuid.NewString()},
				Status:         4,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid uid",
			msg: types.ResolutionEvent{
				UID:            "invalid uid",
				ResolutionTS:   uint64(t1.Unix()),
				WinnerOddsUIDs: []string{uuid.NewString()},
				Status:         4,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "empty winner odds",
			msg: types.ResolutionEvent{
				UID:          uuid.NewString(),
				ResolutionTS: uint64(t1.Unix()),
				Status:       4,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid winner odds",
			msg: types.ResolutionEvent{
				UID:            uuid.NewString(),
				ResolutionTS:   uint64(t1.Unix()),
				WinnerOddsUIDs: []string{"invalid winner odds"},
				Status:         4,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "msg status pending",
			msg: types.ResolutionEvent{
				UID:            uuid.NewString(),
				ResolutionTS:   uint64(t1.Unix()),
				WinnerOddsUIDs: []string{uuid.NewString()},
				Status:         0,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "msg invalid enum status",
			msg: types.ResolutionEvent{
				UID:            uuid.NewString(),
				ResolutionTS:   uint64(t1.Unix()),
				WinnerOddsUIDs: []string{uuid.NewString()},
				Status:         5,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "msg invalid enum status, pending",
			msg: types.ResolutionEvent{
				UID:            uuid.NewString(),
				ResolutionTS:   uint64(t1.Unix()),
				WinnerOddsUIDs: []string{uuid.NewString()},
				Status:         1,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.ValidateResolutionEvent(tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}

}

func Test_UpdateEvent(t *testing.T) {
	k, _, wctx, _ := setupMsgServerAndKeeper(t)
	params := k.GetParams(wctx)

	t1 := time.Now()
	tests := []struct {
		name string
		msg  types.SportEvent
		prev types.SportEvent
		err  error
	}{
		{
			name: "valid request",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
			},
		},
		{
			name: "same timestamp",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute).Unix()),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "end timestamp before current timestamp",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				EndTS:   uint64(t1.Add(-time.Minute).Unix()),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid min amount, negative",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				BetConstraints: &types.EventBetConstraints{
					MinAmount: sdk.NewInt(-5),
					BetFee:    params.EventMinBetFee,
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid min amount, less than required",
			msg: types.SportEvent{
				Creator: sample.AccAddress(),
				StartTS: uint64(t1.Add(time.Minute).Unix()),
				EndTS:   uint64(t1.Add(time.Minute * 2).Unix()),
				UID:     uuid.NewString(),
				Odds: []*types.Odds{
					{UID: uuid.NewString(), Details: "Odds 1"},
					{UID: uuid.NewString(), Details: "Odds 2"},
				},
				BetConstraints: &types.EventBetConstraints{
					MinAmount: params.EventMinBetAmount.Sub(sdk.NewInt(5)),
					BetFee:    params.EventMinBetFee,
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.MsgServerValidateUpdateEvent(wctx, tt.msg, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

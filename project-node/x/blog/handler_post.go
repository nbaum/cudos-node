package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cudos.org/cudos-node/x/blog/keeper"
	"cudos.org/cudos-node/x/blog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
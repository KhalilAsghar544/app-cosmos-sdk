package keeper

import (
	"context"
	"fmt"

	"hello/x/hello/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Creating a map
	// Using make() function
	var My_map = make(map[int]string)

	// As we already know that make() function
	// always returns a map which is initialized
	// So, we can add values in it
	My_map[1] = "Rohit"
	My_map[2] = "Sumit"

	// Checking if the map is nil or not
	if My_map == nil {

		fmt.Println("True")
	} else {

		fmt.Println("False")
	}

	fmt.Println(My_map)

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	return &types.QuerySayHelloResponse{Name: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

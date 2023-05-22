package types_test

import (
	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/testutil"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
		Black: alice,
		Red:   bob,
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	game := GetStoredGame1()
	black, err2 := game.GetBlackAddress()
	require.Equal(t, aliceAddress.String(), black.String())
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongBlack(t *testing.T) {
	game := GetStoredGame1()
	game.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4" // Bad last digit
	black, err := game.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t, err, "black address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, game.Validate(), err.Error())
}

func TestCanGetAddressRed(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	game := GetStoredGame1()
	red, err2 := game.GetRedAddress()
	require.Equal(t, bobAddress.String(), red.String())
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongRed(t *testing.T) {
	game := GetStoredGame1()
	game.Red = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4" // Bad last digit
	red, err := game.GetRedAddress()
	require.Nil(t, red)
	require.EqualError(t, err, "red address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, game.Validate(), err.Error())
}

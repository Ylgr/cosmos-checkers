package types

import (
	"fmt"
	"github.com/alice/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return game, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return game, sdkerrors.Wrapf(fmt.Errorf("turn %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) Validate() error {
	_, errBlack := storedGame.GetBlackAddress()
	if errBlack != nil {
		return errBlack
	}
	_, errRed := storedGame.GetRedAddress()
	if errRed != nil {
		return errRed
	}
	_, errGame := storedGame.ParseGame()
	if errGame != nil {
		return errGame
	}
	return nil
}

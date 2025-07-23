package types

import (
    
)

func NewMsgCreateCoin(creator string, name string) *MsgCreateCoin {
  return &MsgCreateCoin{
		Creator: creator,
    Name: name,
	}
}

func NewMsgUpdateCoin(creator string, id uint64, name string) *MsgUpdateCoin {
  return &MsgUpdateCoin{
        Id: id,
		Creator: creator,
    Name: name,
	}
}

func NewMsgDeleteCoin(creator string, id uint64) *MsgDeleteCoin {
  return &MsgDeleteCoin{
        Id: id,
		Creator: creator,
	}
}
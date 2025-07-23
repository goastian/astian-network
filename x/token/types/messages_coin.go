package types

func NewMsgCreateCoin(creator string, name string, amount string) *MsgCreateCoin {
	return &MsgCreateCoin{
		Creator: creator,
		Name:    name,
		Amount:  amount,
	}
}

func NewMsgUpdateCoin(creator string, id uint64, name string, amount string) *MsgUpdateCoin {
	return &MsgUpdateCoin{
		Id:      id,
		Creator: creator,
		Name:    name,
		Amount:  amount,
	}
}

func NewMsgDeleteCoin(creator string, id uint64) *MsgDeleteCoin {
	return &MsgDeleteCoin{
		Id:      id,
		Creator: creator,
	}
}

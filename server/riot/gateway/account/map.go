package account

import (
	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/protos/scylla/riot/account/v1"
)

func SerializeRiotAccount(account *RiotAccount) *ScyllaAccount {
	return &ScyllaAccount{
		Puuid:    account.Puuid,
		GameName: account.GameName,
		TagLine:  account.TagLine,
	}
}

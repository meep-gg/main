package account

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/scylla"
)

func CreateAccountTable(region string) {
	// Create the account table by puuid
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v_riot.account (
					puuid text PRIMARY KEY,
					gameName text,
					tagLine text
			)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating account_puuid table: %v", Highlight("r", err))
	}
}

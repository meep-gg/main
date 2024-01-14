package match

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolAPI_Match(t *testing.T) {
	t.Log(Title("Lol API Match"))
	t.Run("Id", TestId)
	t.Run("Stress", TestStress)
	t.Run("Match", TestMatch)
}

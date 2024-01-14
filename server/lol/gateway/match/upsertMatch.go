package match

import (
	"context"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"

	. "meep.gg/log"
)

func UpsertMatch(ctx context.Context, region string, gameId int64, match *RiotMatch) error {
	var bans []int32
	var participants []*ScyllaMatchParticipant
	var season, patch int64
	for _, teams := range match.Info.Teams {
		for _, ban := range teams.Bans {
			bans = append(bans, int32(ban.ChampionId))
		}
	}

	gameVersion := strings.Split(match.Info.GameVersion, ".")
	if len(gameVersion) > 1 {
		season64, _ := strconv.ParseUint(gameVersion[0], 10, 32)
		season = int64(season64)
		patch64, _ := strconv.ParseUint(gameVersion[1], 10, 32)
		patch = int64(patch64)
	}
	for _, participant := range match.Info.Participants {
		var p = &ScyllaMatchParticipant{
			Puuid:      participant.Puuid,
			GameName:   participant.RiotIdGameName,
			TagLine:    participant.RiotIdTagline,
			ChampionId: int32(participant.ChampionId),
		}
		participants = append(participants, p)
	}
	var new bool = false
	gameEndTimestamp := time.UnixMilli(int64(match.Info.GameEndTimestamp))
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	if gameEndTimestamp.After(thirtyDaysAgo) {
		new = true
	}
	avgRank, err := UpsertParticipants(ctx, region, gameId, new, match.Info.QueueId, match.Info.Participants)
	if err != nil {
		st := status.Convert(err)
		return status.Error(codes.Internal, st.Message())
	}
	_, err = ScyllaMatchClient.UpsertMatch(ctx, &ScyllaUpsertMatchReq{
		Region: region,
		Match: &ScyllaMatch{
			GameId:           gameId,
			Bans:             bans,
			GameDuration:     match.Info.GameDuration,
			GameEndTimestamp: match.Info.GameEndTimestamp,
			Participants:     participants,
			Patch:            int32(patch),
			QueueId:          match.Info.QueueId,
			Rank:             avgRank,
			Season:           int32(season),
		},
	})
	if err != nil {
		st := status.Convert(err)
		return status.Error(codes.Internal, st.Message())
	}

	return nil
}

func UpsertParticipants(ctx context.Context, region string, gameId int64, new bool, queueId int32, participants []*Participant) (int32, error) {
	var totalRank int32 = 0
	var ranked int32 = 0
	var wg sync.WaitGroup
	var errChan chan error
	if new {
		errChan = make(chan error, len(participants)*2)
	} else {
		errChan = make(chan error, len(participants))
	}

	for _, participant := range participants {
		if new {
			wg.Add(2)
		} else {
			wg.Add(1)
		}
		// participant
		go func(p *Participant) {
			defer wg.Done()
			var rank int32 = 0
			var queue = "RANKED_SOLO_5x5"
			if queueId == 440 {
				queue = "RANKED_FLEX_SR"
			}
			entry, err := ScyllaLeagueClient.GetLeague(ctx, &ScyllaLeagueReq{
				Region:     region,
				QueueType:  queue,
				SummonerId: p.SummonerId,
			})
			if err != nil {
				st := status.Convert(err)
				if st.Code() != codes.NotFound {
					errChan <- err
					return
				}
				log.Printf(Warn("League not found in scylladb"))
			}
			if entry != nil {
				rank = entry.Current
				totalRank += rank
				ranked++
			}
			_, err = ScyllaParticipantClient.UpsertParticipant(ctx, &ScyllaUpsertParticipantReq{
				Region: region,
				Participant: &ScyllaParticipant{
					Puuid:                       p.Puuid,
					GameId:                      gameId,
					ChampionId:                  p.ChampionId,
					ChampLevel:                  p.ChampLevel,
					Kda:                         round(p.Challenges.Kda),
					Kills:                       p.Kills,
					Deaths:                      p.Deaths,
					Assists:                     p.Assists,
					GoldEarned:                  p.GoldEarned,
					GoldPerMinute:               round(p.Challenges.GoldPerMinute),
					TotalMinionsKilled:          p.TotalMinionsKilled,
					NeutralMinionsKilled:        p.NeutralMinionsKilled,
					TotalDamageDealtToChampions: p.TotalDamageDealtToChampions,
					DamagePerMinute:             round(p.Challenges.DamagePerMinute),
					TotalDamageTaken:            p.TotalDamageTaken,
					TimeCCingOthers:             p.TimeCCingOthers,
					IndividualPosition:          p.IndividualPosition,
					Win:                         p.Win,
					Rank:                        rank,
					ParticipantId:               p.ParticipantId,
					Item0:                       p.Item0,
					Item1:                       p.Item1,
					Item2:                       p.Item2,
					Item3:                       p.Item3,
					Item4:                       p.Item4,
					Item5:                       p.Item5,
					Item6:                       p.Item6,
					Summoner1Id:                 p.Summoner1Id,
					Summoner2Id:                 p.Summoner2Id,
					PrimaryStyle:                p.Perks.Styles[0].Selections[0].Perk,
					SubStyle:                    p.Perks.Styles[1].Style,
					TeamId:                      p.TeamId,
					Placement:                   p.Placement,
					PlayerAugment1:              p.PlayerAugment1,
					PlayerAugment2:              p.PlayerAugment2,
					PlayerAugment3:              p.PlayerAugment3,
					PlayerAugment4:              p.PlayerAugment4,
					PlayerSubteamId:             p.PlayerSubteamId,
					TeamDamagePercentage:        round(p.Challenges.TeamDamagePercentage),
					VisionScorePerMinute:        round(p.Challenges.VisionScorePerMinute),
					KillParticipation:           round(p.Challenges.KillParticipation),
				},
			})
			if err != nil {
				errChan <- err
				return
			}
		}(participant)

		// Participant Details
		if new {
			go func(p *Participant) {
				defer wg.Done()
				_, err := ScyllaParticipantdetailClient.UpsertParticipantdetail(ctx, &ScyllaUpsertParticipantdetailReq{
					Region: region,
					Detail: &ScyllaParticipantdetail{
						GameId:                         gameId,
						ParticipantId:                  p.ParticipantId,
						DamageDealtToBuildings:         p.DamageDealtToBuildings,
						DamageDealtToObjectives:        p.DamageDealtToObjectives,
						DamageDealtToTurrets:           p.DamageDealtToTurrets,
						DamageSelfMitigated:            p.DamageSelfMitigated,
						MagicDamageDealt:               p.MagicDamageDealt,
						MagicDamageDealtToChampions:    p.MagicDamageDealtToChampions,
						MagicDamageTaken:               p.MagicDamageTaken,
						MaxCsAdvantageOnLaneOpponent:   p.Challenges.MaxCsAdvantageOnLaneOpponent,
						MaxLevelLeadLaneOpponent:       p.Challenges.MaxLevelLeadLaneOpponent,
						NeutralMinionsKilled:           p.NeutralMinionsKilled,
						Perks:                          SerialPerks(p.Perks),
						PhysicalDamageDealt:            p.PhysicalDamageDealt,
						PhysicalDamageDealtToChampions: p.PhysicalDamageDealtToChampions,
						PhysicalDamageTaken:            p.PhysicalDamageTaken,
						Spell1Casts:                    p.Spell1Casts,
						Spell2Casts:                    p.Spell2Casts,
						Spell3Casts:                    p.Spell3Casts,
						Spell4Casts:                    p.Spell4Casts,
						Summoner1Casts:                 p.Summoner1Casts,
						Summoner2Casts:                 p.Summoner2Casts,
						DamageTakenOnTeamPercentage:    round(p.Challenges.DamageTakenOnTeamPercentage),
					},
				})
				if err != nil {
					errChan <- err
					return
				}
			}(participant)
		}
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		st := status.Convert(err)
		return 0, status.Error(codes.Internal, st.Message())
	}

	if ranked == 0 {
		return 0, nil
	}
	return (totalRank / ranked), nil
}

func round(n float32) float32 {
	return float32(math.Round(float64(n)*100) / 100)
}

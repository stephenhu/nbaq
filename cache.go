package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/stephenhu/stats"
)


type Stats struct {
	Minutes               string        `json:"minutes"`
	Points								int           `json:"points"`
	Oreb									int           `json:"oreb"`
	Dreb									int           `json:"dreb"`
	Treb									int           `json:"treb"`
	Fta										int           `json:"fta"`
	Ftm										int           `json:"ftm"`
	Fg2a									int           `json:"fg2a"`
	Fg2m									int           `json:"fg2m"`
	Fg3a									int           `json:"fg3a"`
	Fg3m									int           `json:"fg3m"`
	Steals								int           `json:"steals"`
	Assists								int           `json:"assists"`
	Blocks								int           `json:"blocks"`
	Blocked								int           `json:"blocked"`
	Turnovers							int           `json:"turnovers"`
	Fouls									int           `json:"fouls"`
	Fouled								int           `json:"fouled"`
	FoulsOffensive				int           `json:"foulsOffensive"`
	Technicals						int           `json:"technicals"`
	Paint									int           `json:"paint"`
	Fastbreak							int           `json:"fastbreak"`
	SecondChance					int           `json:"secondChance"`
	PlusMinus							float64       `json:"plusMinus"`
}


type PlayerStats struct {
	ID                    int        		`json:"id"`
	First									string        `json:"first"`
	Last									string        `json:"last"`
	Full									string        `json:"full"`
	Abv										string        `json:"abv"`
	Jersey                string        `json:"jerseyNum"`
	Position              string        `json:"position"`
	Starter								string        `json:"starter"`
	Active                string        `json:"active"`
	Order                 int           `json:"order"`
	Stats			`json:"stats"`
}


type Period struct {
	Number								int           `json:"number"`
	PeriodType						string        `json:"periodType"`
	Away                  int           `json:"away"`
	Home                  int           `json:"home"`
}


type TeamStats struct {
	ID                		int           `json:"id"`
	Abv                   string        `json:"abv"`
	Score                 int           `json:"score"`
	Players               map[int]PlayerStats 	`json:"players"`
}


type Play struct {
	ID										string				`json:"id"`
	Detail								string				`json:"detail"`
}


type Game struct {
	ID                    string        `json:"id"`
	Date                  string        `json:"date"`
	Away									TeamStats			`json:"away"`
	Home									TeamStats			`json:"home"`
	Periods								map[int]Period			`json:"periods"`
	Plays                 map[string]Play     `json:"plays"`
}


type Standings struct {

}


type GameStats struct {
  GameID					string								`json:"gameId"`
	TeamID					int										`json:"teamId"`
	PlayerID				int										`json:"playerId"`
	Date            string                `json:"date"`
	Position				string								`json:"position"`
	Full						string								`json:"full"`
	Abv							string								`json:"abv"`
	Stats           Stats									`json:"stats"`
}


type SeasonData struct {


}


type Leaders struct {
  Players				map[int]PlayerStats					`json:"players"`
}


type Season struct {
	ID										string				`json:"id"`
	Games									map[string]Game	`json:"games"`
	Standings							Standings			`json:"standings"`
	Leaders								Leaders				`json:"leaders"`
	Players               map[int]map[string]GameStats		`json:"players"`
	Teams               	map[int]map[string]GameStats		`json:"teams"`
}


type Player struct {
	ID                    int           `json:"id"` 
	TeamID                int           `json:"teamId"`
	First									string				`json:"first"`
	Last									string				`json:"last"`
	Full									string				`json:"full"`
	Abv										string				`json:"abv"`
	Position							string				`json:"position"`
	Active								string				`json:"active"`
	Height								int						`json:"height"`
	Weight								int						`json:"weight"`
}


type Team struct {
	ID                		int           `json:"id"`
	Name									string				`json:"name"`
	Code									string				`json:"last"`
	City									string				`json:"full"`
	Mascot								string				`json:"abv"`
	Conf									string				`json:"position"`
	Div										string				`json:"active"`
}


type NbaCache struct {
  Seasons   						map[string]Season	`json:"seasons"`
	Players               map[int]Player		`json:"players"`
	Teams               	map[int]Team			`json:"teams"`
}


func TPlayerStats(p []stats.NbaPlayer) map[int]PlayerStats {

	ret := map[int]PlayerStats{}

	for _, player := range p {

		ret[player.ID] = PlayerStats {
			ID: player.ID,
			First: player.First,
			Last: player.Last,
			Full: player.Name,
			Abv: player.NameShort,
			Position: player.Position,
			Jersey: player.Jersey,
			Starter: player.Starter,
			Order: player.Order,
			Stats: Stats{
				Minutes: player.Statistics.Minutes,
				Points: player.Statistics.Points,
				Oreb: player.Statistics.Oreb,
				Dreb: player.Statistics.Dreb,
				Treb: player.Statistics.Treb,
				Fta: player.Statistics.Fta,
				Ftm: player.Statistics.Ftm,
				Fg2a: player.Statistics.Fg2a,
				Fg2m: player.Statistics.Fg2m,
				Fg3a: player.Statistics.Fg3a,
				Fg3m: player.Statistics.Fg3m,
				Assists: player.Statistics.Assists,
				Blocks: player.Statistics.Blocks,
				Steals: player.Statistics.Steals,
				Blocked: player.Statistics.Blocked,
				Turnovers: player.Statistics.Turnovers,
				Fouls: player.Statistics.Fouls,
				Fouled: player.Statistics.FoulsDrawn,
				Technicals: player.Statistics.Technicals,
				FoulsOffensive: player.Statistics.FoulsOff,
				Fastbreak: player.Statistics.PointsFast,
				Paint: player.Statistics.PointsPaint,
				SecondChance: player.Statistics.PointsSecond,
				PlusMinus: player.Statistics.PlusMinus,
			},
		}
	}

	return ret

} // TPlayerStats


func TPeriods(a []stats.NbaScoreData,
	h []stats.NbaScoreData) map[int]Period {

	ret := map[int]Period{}

	if len(h) != len(a) {
		
		log.Println("Error: periods number mismatch between home and away")
		return nil

	} else {

		for i, _ := range a {

			np := Period{
				Number: a[i].Period,
				Away: a[i].Score,
				Home: h[i].Score,
			}

			ret[a[i].Period] = np

		}

		return ret

	}

} // TPeriods


func TBoxscore(b stats.NbaBoxscore) {

	s := Season{
		ID: src,
		Games: map[string]Game{},
		Players: map[int]map[string]GameStats{},
	}

	_, ok := cache.Seasons[src]

	if !ok {
		cache.Seasons[src] = s
	}

	g := Game{
		ID: b.Game.ID,
		Date: b.Game.GameTime,
		Periods: TPeriods(b.Game.Away.Periods, b.Game.Home.Periods),
		Away: TeamStats{
			ID: b.Game.Away.ID,
			Abv: b.Game.Away.ShortName,
			Score: b.Game.Away.Score,
			Players: TPlayerStats(b.Game.Away.Players),
		},
		Home: TeamStats{
			ID: b.Game.Home.ID,
			Abv: b.Game.Home.ShortName,
			Score: b.Game.Home.Score,
			Players: TPlayerStats(b.Game.Home.Players),
		},
	}

	cache.Seasons[src].Games[b.Game.ID] = g

} // TBoxscore


func loadBoxscore(p string) {

	buf, err := os.ReadFile(p)

	if err != nil {
		log.Println(err)
	} else {

		b := stats.NbaBoxscore{}

		err := json.Unmarshal(buf, &b)

		if err != nil {
			log.Println(err)
		} else {
			TBoxscore(b)
		}

	}

} // loadBoxscore


func TPlayerGameStats() {

} // TPLayerGameStats


func initPlayerStats() {
	
	players := cache.Seasons[src].Players

	for _, g := range cache.Seasons[src].Games {

		for _, p := range g.Away.Players {

			_, ok := players[p.ID]

			if !ok {
				players[p.ID] = make(map[string]GameStats)
			}

			players[p.ID][g.ID] = GameStats{
				GameID: g.ID,
				TeamID: g.Away.ID,
				PlayerID: p.ID,
				Date: g.Date,
				Position: p.Position,
				Full: p.Full,
				Abv: p.Abv,
				Stats: Stats{
					Minutes: p.Stats.Minutes,
					Points: p.Stats.Points,
					Oreb: p.Stats.Oreb,
					Dreb: p.Stats.Dreb,
					Treb: p.Stats.Treb,
					Assists: p.Stats.Assists,
					Blocks: p.Stats.Blocks,
					Steals: p.Stats.Steals,
					Turnovers: p.Stats.Turnovers,
					Fouls: p.Stats.Fouls,
					Fouled: p.Stats.Fouled,
					FoulsOffensive: p.Stats.FoulsOffensive,
					Technicals: p.Stats.Technicals,
					Blocked: p.Stats.Blocked,
					Fta: p.Stats.Fta,
					Ftm: p.Stats.Ftm,
					Fg2a: p.Stats.Fg2a,
					Fg2m: p.Stats.Fg2m,
					Fg3a: p.Stats.Fg3a,
					Fg3m: p.Stats.Fg3m,
					Paint: p.Stats.Paint,
					Fastbreak: p.Stats.Fastbreak,
					SecondChance: p.Stats.SecondChance,
					PlusMinus: p.Stats.PlusMinus,
				},
			}

		}

	}

} // initPlayerStats


func initTeamStats() {

} // initTeamStats


func initCache() {

	cache = NbaCache{
		Seasons: map[string]Season{},
	}

	dirs := getFiles(src)

	for _, d := range dirs {

		if d.IsDir() {

			p := fmt.Sprintf("%s/%s", src, d.Name())

			files := getFiles(p)

			if files != nil {

				for _, f := range files {

					if filepath.Ext(f.Name()) == JSON_EXT &&
					  !strings.Contains(f.Name(), PLAY_BY_PLAY) {
						loadBoxscore(filepath.Join(p, f.Name()))
					}
						
				}

			}

		}

	}

	//log.Printf("%v\n", cache)

	initPlayerStats()

	initTeamStats()

} // initCache

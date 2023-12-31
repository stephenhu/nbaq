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


type PlayerData struct {
	ID                    string        `json:"id"`
	First									string        `json:"first"`
	Last									string        `json:"last"`
	Full									string        `json:"full"`
	Abv										string        `json:"abv"`
	Starter								string        `json:"starter"`
	Active                string        `json:"active"`
	Minutes               int           `json:"minutes"`
	Points								int           `json:"points"`
	Oreb									int           `json:"oreb"`
	Dreb									int           `json:"dreb"`
	Treb									int           `json:"treb"`
	Fga										int           `json:"fga"`
	Fgm										int           `json:"fgm"`
	Fta										int           `json:"fta"`
	Ftm										int           `json:"ftm"`
	Fg3a									int           `json:"fg3a"`
	Fg3m									int           `json:"fg3m"`
	Steals								int           `json:"steals"`
	Assists								int           `json:"assists"`
	Blocks								int           `json:"blocks"`
	Blocked								int           `json:"blocked"`
	Turnovers							int           `json:"turnovers"`
	Fouls									int           `json:"fouls"`
	Fouled								int           `json:"fouled"`
	Technicals						int           `json:"technicals"`
	Fragrants							int           `json:"fragrants"`
	Paint									int           `json:"paint"`
	Fastbreak							int           `json:"fastbreak"`
	SecondChance					int           `json:"secondChance"`
	PlusMinus							int           `json:"plusMinus"`
}


type Period struct {
	Number								int           `json:"number"`
	PeriodType						string        `json:"periodType"`
	Away                  int           `json:"away"`
	Home                  int           `json:"home"`
}


type TeamData struct {
	ID                		int           `json:"id"`
	Abv                   string        `json:"abv"`
	Score                 int           `json:"score"`
	Players               map[int]PlayerData 	`json:"players"`
}


type Play struct {
	ID										string				`json:"id"`
	Detail								string				`json:"detail"`
}


type Game struct {
	ID                    string        `json:"id"`
	Date                  string        `json:"date"`
	Away									TeamData			`json:"away"`
	Home									TeamData			`json:"home"`
	Periods								map[int]Period			`json:"periods"`
	Plays                 map[string]Play     `json:"plays"`
}


type Rankings struct {

}


type Leaders struct {

}


type Season struct {
	ID										string				`json:"id"`
	Games									map[string]Game	`json:"games"`
	Rankings							Rankings			`json:"rankings"`
	Leaders								Leaders				`json:"leaders"`
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


func AddPlayer() {

} // AddPlayer


func AddTeam() {

} // AddTeam


func TPlayerData(p []stats.NbaPlayer) map[int]PlayerData {

	ret := map[int]PlayerData{}

	return ret

} // TPlayerData


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
	}

	_, ok := cache.Seasons[src]

	if !ok {
		cache.Seasons[src] = s
	}

	g := Game{
		ID: b.Game.ID,
		Date: b.Game.GameTime,
		Periods: TPeriods(b.Game.Away.Periods, b.Game.Home.Periods),
		Away: TeamData{
			ID: b.Game.Away.ID,
			Abv: b.Game.Away.ShortName,
			Score: b.Game.Away.Score,
			Players: TPlayerData(b.Game.Away.Players),
		},
		Home: TeamData{
			ID: b.Game.Home.ID,
			Abv: b.Game.Home.ShortName,
			Score: b.Game.Home.Score,
			Players: TPlayerData(b.Game.Home.Players),
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

	log.Printf("%v\n", cache)

} // initCache

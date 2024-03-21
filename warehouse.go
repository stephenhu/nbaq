package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/stephenhu/stats"
	_ "github.com/marcboeker/go-duckdb"
	
)


const (
	DUCKDB_CONNECTOR				= "duckdb"
)


const (
  PLAYERS_PREFIX					= "players"
	GAMES_PREFIX						= "games"
)


const (
	SELECT                  = "SELECT"
	FROM                    = "FROM"
  WHERE                   = "WHERE"
)


const (
	QUERY_PLAYERS					= "SELECT " +
	                        "* " +
	                        "FROM '%s' " +
													"WHERE playerId = %s"
	QUERY_PLAYER_PPG			= "SELECT SUM(points) FROM '%s' WHERE playerId = %s"
)


var (
	db 				*sql.DB
)


func initWarehouse() {

	conn, err := sql.Open(DUCKDB_CONNECTOR, "")

	if err != nil {
		log.Println(err)
	} else {

		db = conn
		
	}
	

} // initWarehouse


func warehousePlayers() string {

	l := getLatest(PLAYERS_PREFIX)

	log.Println(l)
	
  return filepath.Join(dir, WAREHOUSE_DIR, getLatest(
		PLAYERS_PREFIX))
} // warehousePlayers


func getPlayerSeason(id string) *stats.PlayerSeason {

	ps := stats.PlayerSeason{
		Games: make(map[string]stats.PlayerGame),
		Info: stats.PlayerInfo{},
	}

	q := fmt.Sprintf(QUERY_PLAYERS,
		warehousePlayers(), id)

	rows, err := db.Query(q)

	if err != nil {
		log.Println(err)
	} else {

		defer rows.Close()

		for rows.Next() {

			pg := stats.PlayerGame{}

			err := rows.Scan(
				&pg.GameDate, &ps.Info.ID, &pg.TeamID, &pg.HomeTeamID,
				&pg.HomeCode, &pg.AwayTeamID, &pg.AwayCode, &ps.Info.First,
				&ps.Info.Last, &ps.Info.Name,
				&ps.Info.NameShort, &pg.GameID,
				&pg.Base.Points, &pg.Base.Oreb,
				&pg.Base.Dreb, &pg.Base.Treb, &pg.Base.Assists, &pg.Base.Steals,
				&pg.Base.Turnovers, &pg.Base.Blocks, &pg.Base.Blocked, &pg.Base.Fouls,
				&pg.Base.FoulsOffensive, &pg.Base.Technicals, &pg.Base.Fouled,
				&pg.Base.Fta, &pg.Base.Ftm, &pg.Base.Ftp,
				&pg.Base.Fg2a, &pg.Base.Fg2m, &pg.Base.Fg2p,
				&pg.Base.Fg3a, &pg.Base.Fg3m, &pg.Base.Fg3p,
				&pg.Base.Fgta, &pg.Base.Fgtm, &pg.Base.Fgtp,
				&pg.PlusMinus, &ps.Info.Position, &pg.Minutes, 
				&pg.Base.Fastbreak, &pg.Base.Paint, &pg.Base.SecondChance,
				&pg.GameType,
			)

			if err != nil {
				log.Println(err)
			} else {
				ps.Games[pg.GameID] = pg
			}

		}

	}

	return &ps

} // getPlayerSeason

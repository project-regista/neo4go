package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const matchQuery = `
// Merge match
MERGE (m:Match{id: {ID} })
ON CREATE SET m.temperature = {Temperature}, m.htScore = {HtScore}, m.ftScore = {FtScore}, m.etScore = {EtScore}, m.homeScorePenalties = {HomeScorePenalties}, m.awayScorePenalties = {AwayScorePenalties}, m.dateTimeTba = {DateTimeTba}, m.spectators = {Spectators}, m.startingDate = {StartingDate}, m.startingTime = {StartingTime}, m.status = {Status}, m.minute = {Minute}, m.injuryTime = {InjuryTime}, m.extraTime = {ExtraTime}, m.roundID = {RoundID}, m.stageID = {StageID}, m.aggregate = {Aggregate}, m.weather_code = {Weather_Code}, m.weather_Type = {Weather_Type}, m.weather_temperature_temp = {Weather_Temperature_Temp}, m.weather_temperature_unit = {Weather_Temperature_Unit}, m.weather_clouds = {Weather_Clouds}, m.weather_humidity = {Weather_Humidity}, m.weather_wind_speed = {Weather_Wind_Speed}, m.formation_home = {Formation_Home}, m.formation_away = {Formation_Away}
ON MATCH SET m.temperature = {Temperature}, m.htScore = {HtScore}, m.ftScore = {FtScore}, m.etScore = {EtScore}, m.homeScorePenalties = {HomeScorePenalties}, m.awayScorePenalties = {AwayScorePenalties}, m.dateTimeTba = {DateTimeTba}, m.spectators = {Spectators}, m.startingDate = {StartingDate}, m.startingTime = {StartingTime}, m.status = {Status}, m.minute = {Minute}, m.injuryTime = {InjuryTime}, m.extraTime = {ExtraTime}, m.roundID = {RoundID}, m.stageID = {StageID}, m.aggregate = {Aggregate}, m.weather_code = {Weather_Code}, m.weather_Type = {Weather_Type}, m.weather_temperature_temp = {Weather_Temperature_Temp}, m.weather_temperature_unit = {Weather_Temperature_Unit}, m.weather_clouds = {Weather_Clouds}, m.weather_humidity = {Weather_Humidity}, m.weather_wind_speed = {Weather_Wind_Speed}, m.formation_home = {Formation_Home}, m.formation_away = {Formation_Away}
// Merge home team
MERGE (ht:Team{id: {HomeTeam_ID} })
ON CREATE SET ht.name = {HomeTeam_Name}, ht.logo = {HomeTeam_Logo}, ht.twitter = {HomeTeam_Twitter}, ht.venue = {HomeTeam_VenueID}, ht.coach = {HomeTeam_CoachID}, ht.chairman = {HomeTeam_ChairmanID}
ON MATCH SET ht.name = {HomeTeam_Name}, ht.logo = {HomeTeam_Logo}, ht.twitter = {HomeTeam_Twitter}, ht.venue = {HomeTeam_VenueID}, ht.coach = {HomeTeam_CoachID}, ht.chairman = {HomeTeam_ChairmanID}
// Merge away team
MERGE (at:Team{id: {AwayTeam_ID} })
ON CREATE SET at.name = {AwayTeam_Name}, at.logo = {AwayTeam_Logo}, at.twitter = {AwayTeam_Twitter}, at.venue = {AwayTeam_VenueID}, at.coach = {AwayTeam_CoachID}, at.chairman = {AwayTeam_ChairmanID}
ON MATCH SET at.name = {AwayTeam_Name}, at.logo = {AwayTeam_Logo}, at.twitter = {AwayTeam_Twitter}, at.venue = {AwayTeam_VenueID}, at.coach = {AwayTeam_CoachID}, at.chairman = {AwayTeam_ChairmanID}
// Merge venue
MERGE (v:Venue{id: {Venue_ID} })
ON CREATE SET v.name = {Venue_Name}, v.city = {Venue_City}, v.address = {Venue_Address}, v.phone = {Venue_Phone}, v.fax = {Venue_Fax}, v.yearOfManufacture = {Venue_YearOfManufacture}, v.seats = {Venue_Seats}, v.fieldType  = {Venue_FieldType}
ON MATCH SET v.name = {Venue_Name}, v.city = {Venue_City}, v.address = {Venue_Address}, v.phone = {Venue_Phone}, v.fax = {Venue_Fax}, v.yearOfManufacture = {Venue_YearOfManufacture}, v.seats = {Venue_Seats}, v.fieldType  = {Venue_FieldType}
// Merge Home stats
MERGE (hs:Stats{id: {HomeStats_ID} })
ON CREATE SET hs.teamID = {HomeStats_TeamID}, hs.shotsOfftarget = {HomeStats_ShotsOffTarget}, hs.shotsOnTarget = {HomeStats_ShotsOnTarget}, hs.shotsTotal = {HomeStats_ShotsTotal}, hs.foulsTotal = {HomeStats_FoulsTotal}, hs.cornersTotal = {HomeStats_CornersTotal}, hs.offsidesTotal = {HomeStats_OffsidesTotal}, hs.possession = {HomeStats_Possession}, hs.yellowCards = {HomeStats_YellowCards}, hs.redCards = {HomeStats_RedCards}, hs.saves = {HomeStats_Saves}, hs.throwInTotal = {HomeStats_ThrowInTotal}
ON MATCH SET hs.teamID = {HomeStats_TeamID}, hs.shotsOfftarget = {HomeStats_ShotsOffTarget}, hs.shotsOnTarget = {HomeStats_ShotsOnTarget}, hs.shotsTotal = {HomeStats_ShotsTotal}, hs.foulsTotal = {HomeStats_FoulsTotal}, hs.cornersTotal = {HomeStats_CornersTotal}, hs.offsidesTotal = {HomeStats_OffsidesTotal}, hs.possession = {HomeStats_Possession}, hs.yellowCards = {HomeStats_YellowCards}, hs.redCards = {HomeStats_RedCards}, hs.saves = {HomeStats_Saves}, hs.throwInTotal = {HomeStats_ThrowInTotal}
//Merge away stats
MERGE (as:Stats{id: {AwayStats_ID} })
ON CREATE SET as.teamID = {AwayStats_TeamID}, as.shotsOfftarget = {AwayStats_ShotsOffTarget}, as.shotsOnTarget = {AwayStats_ShotsOnTarget}, as.shotsTotal = {AwayStats_ShotsTotal}, as.foulsTotal = {AwayStats_FoulsTotal}, as.cornersTotal = {AwayStats_CornersTotal}, as.offsidesTotal = {AwayStats_OffsidesTotal}, as.possession = {AwayStats_Possession}, as.yellowCards = {AwayStats_YellowCards}, as.redCards = {AwayStats_RedCards}, as.saves = {AwayStats_Saves}, as.throwInTotal = {AwayStats_ThrowInTotal}
ON MATCH SET as.teamID = {AwayStats_TeamID}, as.shotsOfftarget = {AwayStats_ShotsOffTarget}, as.shotsOnTarget = {AwayStats_ShotsOnTarget}, as.shotsTotal = {AwayStats_ShotsTotal}, as.foulsTotal = {AwayStats_FoulsTotal}, as.cornersTotal = {AwayStats_CornersTotal}, as.offsidesTotal = {AwayStats_OffsidesTotal}, as.possession = {AwayStats_Possession}, as.yellowCards = {AwayStats_YellowCards}, as.redCards = {AwayStats_RedCards}, as.saves = {AwayStats_Saves}, as.throwInTotal = {AwayStats_ThrowInTotal}
// Merge season
MERGE (s:Season{id: {Season_ID} })
ON CREATE SET s.name = {Season_Name}, s.active = {Season_Active}
ON MATCH SET s.name = {Season_Name}, s.active = {Season_Active}
// Merge referee
MERGE (r:Referee{id: {Referee_ID} })
ON CREATE SET r.name = {Referee_Name}
ON MATCH SET r.name = {Referee_Name}
// Links
MERGE (m)-[:HAS_HOME_TEAM]->(ht)
MERGE (m)-[:HAS_AWAY_TEAM]->(at)
MERGE (m)-[:AT_VENUE]->(v)
MERGE (m)-[:HAS_HOME_STATS]->(hs)
MERGE (m)-[:HAS_AWAY_STATS]->(as)
MERGE (m)-[:HAS_HOME_TEAM]->(ht)
MERGE (m)-[:HAS_AWAY_TEAM]->(at)
MERGE (m)-[:OCCURS_ON_SEASON]->(s)
MERGE (m)-[:HAS_REFEREE]->(r)
`

// MatchStmt prepare a Neo4j statement for a single match
func MatchStmt(conn bolt.Conn) (bolt.Stmt, error) {
	stmt, err := conn.PrepareNeo(matchQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create season statement: %s", err)
	}
	return stmt, nil
}

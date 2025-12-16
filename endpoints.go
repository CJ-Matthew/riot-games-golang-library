package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// **************
// User Interface
// **************

// Create new riot client that holds api key and region for querying
func NewRiotClient(apiKey string, cluster string, region string) *RiotClient {
	clusterURL := "https://" + cluster + ".api.riotgames.com"
	regionURL := "https://" + region + ".api.riotgames.com"

	
	return &RiotClient{
		apiKey:  apiKey,
		cluster:  cluster,
		region: region,
		clusterURL: clusterURL,
		regionURL: regionURL,
	}
}

// Riot client data 
type RiotClient struct {
	apiKey string
	cluster string
	region string
	clusterURL string
	regionURL string
}

// *****************
// ACCOUNT ENDPOINTS
// *****************

// Get the PUUID for an riot account using their riot name and tag
func (riotClient RiotClient) GetAccountByRiotID(name string, tag string) (Account, error) {

	url := riotClient.clusterURL + "/riot/account/v1/accounts/by-riot-id/" + name + "/" + tag

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return Account{}, err
	}

	// Unmarshall the response bodys
	var account Account
	err = json.Unmarshal(body, &account)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return Account{}, err
	}
	
	return account, nil
}

// Get the active region for an account
func (riotClient RiotClient) GetAccountRegion(puuid string, game string) (AccountRegion , error) {
	url := riotClient.clusterURL + "/riot/account/v1/region/by-game/" + game + "/by-puuid/" + puuid

	body, err := riotClient.sendGetRequest(url) 

	if err != nil {
		return AccountRegion{}, err
	}

	// Unmarshall the response bodys
	var accountRegion AccountRegion
	err = json.Unmarshal(body, &accountRegion)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return AccountRegion{}, err
	} 

	return accountRegion, nil
}

// **************************
// CHAMPION MASTERY ENDPOINTS
// **************************

// Returns the champion masteries in decending order. if count is -1 then all masteries returned otherwise the top count are returned
func (riotClient RiotClient) GetAllChampionMasteries(puuid string, count int) ([]ChampionMastery, error) {
	var url string
	if count == -1 {
		url = riotClient.regionURL + "/lol/champion-mastery/v4/champion-masteries/by-puuid/" + puuid
	} else {
		url = riotClient.regionURL + "/lol/champion-mastery/v4/champion-masteries/by-puuid/" + puuid + "/top?count=" + strconv.Itoa(count)
	}
	
	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return nil, err
	}

	// Unmarshall the response bodys
	var championMasteries []ChampionMastery
	err = json.Unmarshal(body, &championMasteries)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return nil, err
	}
	
	return championMasteries, nil
	
}

// Get mastery details for a specific champion
func (riotClient RiotClient) GetChampionMastery(championID int, puuid string) (ChampionMastery, error) {
	url := riotClient.regionURL + "/lol/champion-mastery/v4/champion-masteries/by-puuid/" + puuid + "/by-champion/" + strconv.Itoa(championID)
	
	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return ChampionMastery{}, err
	}

	// Unmarshall the response bodys
	var champMastery ChampionMastery
	err = json.Unmarshal(body, &champMastery)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return ChampionMastery{}, err
	}
	
	return champMastery, nil
}

// Get mastery score for a user
func (riotClient RiotClient) GetMasteryScore(puuid string) (int, error) {
	url := riotClient.regionURL + "/lol/champion-mastery/v4/scores/by-puuid/" + puuid

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(string(body))
	if err != nil {
		log.Fatal(err)
	}



	return count, nil
}

// ******************
// CHAMPION ENDPOINTS
// ******************

// Get the champions on free rotation, low level accounts have different champion rotations
func (riotClient RiotClient) GetChampionRotation() (ChampionRotation, error) {
	url := riotClient.regionURL + "/lol/platform/v3/champion-rotations"

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return ChampionRotation{}, err
	}

	// Unmarshall the response bodys
	var champRotation ChampionRotation
	err = json.Unmarshal(body, &champRotation)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return ChampionRotation{}, err
	}
	
	return champRotation, nil

}

// ***************
// CLASH ENDPOINTS
// ***************

// Get the registered clash tournaments for a player (saturday and sunday are different tournaments)
func (riotClient RiotClient) GetClashPlayer(puuid string) ([]ClashPlayer, error) {
	url := riotClient.regionURL + "/lol/clash/v1/players/by-puuid/" + puuid

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return nil, err
	}

	// Unmarshall the response bodys
	var clashPlayer []ClashPlayer
	err = json.Unmarshal(body, &clashPlayer)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return nil, err
	}
	
	return clashPlayer, nil
}

// Get the clash team details 
func (riotClient RiotClient) GetClashTeam(teamId int) (ClashTeam, error) {
	url := riotClient.regionURL + "/lol/clash/v1/teams/" + strconv.Itoa(teamId)

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return ClashTeam{}, err
	}

	// Unmarshall the response bodys
	var clashTeam ClashTeam
	err = json.Unmarshal(body, &clashTeam)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return ClashTeam{}, err
	}
	
	return clashTeam, nil
}

// Get upcoming tournaments for a region
func (riotClient RiotClient) GetClashTournaments() ([]ClashTournament, error) {
	url := riotClient.regionURL + "/lol/clash/v1/tournaments"

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return nil, err
	}

	// Unmarshall the response bodys
	var clashTournaments []ClashTournament
	err = json.Unmarshal(body, &clashTournaments)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return nil, err
	}
	
	return clashTournaments, nil
}

// Get tournament details for a team
func (riotClient RiotClient) GetClashTeamTournamentDetails(teamId int) (ClashTournament, error) {
	url := riotClient.regionURL + "/lol/clash/v1/tournaments/by-team/" + strconv.Itoa(teamId)

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return ClashTournament{}, err
	}

	// Unmarshall the response bodys
	var clashTeamTournament ClashTournament
	err = json.Unmarshal(body, &clashTeamTournament)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return ClashTournament{}, err
	}
	
	return clashTeamTournament, nil
}

// Get tournament deatils
func (riotClient RiotClient) GetClashTournamentDetails(tournamentId int) (ClashTournament, error) {
	url := riotClient.regionURL + "/lol/clash/v1/tournaments/" + strconv.Itoa(tournamentId)

	body, err := riotClient.sendGetRequest(url)

	if err != nil {
		return ClashTournament{}, err
	}

	// Unmarshall the response bodys
	var clashTournament ClashTournament
	err = json.Unmarshal(body, &clashTournament)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return ClashTournament{}, err
	}
	
	return clashTournament, nil
}

// ****************
// Helper Functions
// ****************

func (riotClient RiotClient) sendGetRequest(url string) ([]byte, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	// Add headers to the request
	req.Header.Add("X-Riot-Token", riotClient.apiKey)
	req.Header.Set("Accept", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to " + url + ":", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Status code error " + resp.Status) 
		return nil, fmt.Errorf("200 response not given")
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	return body, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	asiaOC1Client := NewRiotClient(os.Getenv("RIOT_KEY"), "asia", "oc1")
	// account, _ := asiaOC1Client.GetAccountByRiotID("CJM", "00000")
	// accountRegion, _ := asiaClient.GetAccountRegion(account.PUUID, "lol")\
	// champMasteries, _ := asiaClient.GetAllChampionMasteries(account.PUUID, 3)
	// cm, _ := asiaClient.GetChampionMastery("412", account.PUUID)
	fmt.Println(asiaOC1Client.GetClashTournaments())
	

}
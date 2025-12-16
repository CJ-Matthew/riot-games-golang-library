package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Create new riot client that holds api key and region for querying
func NewRiotClient(apiKey string, cluster string) *RiotClient {
	clusterURL := "https://" + cluster + ".api.riotgames.com"

	
	return &RiotClient{
		apiKey:  apiKey,
		region:  cluster,
		clusterURL: clusterURL,
		regionURL: "",
	}
}

// Riot client data 
type RiotClient struct {
	apiKey string
	region string
	clusterURL string
	regionURL string
}

// ACCOUNT ENDPOINTS

// Get the PUUID for an riot account using their riot name and tag
func (riotClient RiotClient) GetPUUID(name string, tag string) (string, error) {

	url := riotClient.clusterURL + "/riot/account/v1/accounts/by-riot-id/" + name + "/" + tag

	body, err := riotClient.sendRequest(url)

	if err != nil {
		return "", err
	}

	type Response struct {
		PUUID    string `json:"puuid"`
		GameName string `json:"gameName"`
		TagLine  string `json:"tagLine"`
	}

	// Unmarshall the response bodys
	var response Response
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return "", err
	}
	
	return response.PUUID, nil
}

// Get the active region for an account
func (riotClient RiotClient) GetActiveRegion(puuid string, game string) (string , error) {
	url := riotClient.clusterURL + "/riot/account/v1/region/by-game/" + game + "/by-puuid/" + puuid

	body, err := riotClient.sendRequest(url) 

	if err != nil {
		return "", err
	}

	type Response struct {
		PUUID    string `json:"puuid"`
		Game string `json:"game"`
		Region  string `json:"region"`
	}

	// Unmarshall the response bodys
	var response Response
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Println("Error unmarshalling response body", err)
		return "", err
	} 

	return response.Region, nil
}


func (riotClient RiotClient) sendRequest(url string) ([]byte, error) {
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
	asiaClient := NewRiotClient(os.Getenv("RIOT_KEY"), "asia")
	puuid, _ := asiaClient.GetPUUID("CJM", "00000")
	fmt.Println(asiaClient.GetActiveRegion(puuid, "lol"))
}
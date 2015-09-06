package steamgauggo

import (
	"encoding/json"
	"net/http"
)

//This package to check if steam api is online or not
//Please refer to https://steamgaug.es/doc for more details
//440 TF2
//570 Dota2
//730 CSGO

type SteamInfoJson struct {
	IsSteamClient struct {
		Online int `json:"online"`
	} `json:"ISteamClient"`

	SteamCommunity struct {
		Online int    `json:"online"`
		Time   int    `json:"time"`
		Error  string `json:"error"`
	} `json:"SteamCommunity"`

	SteamStore struct {
		Online int    `json:"online"`
		Time   int    `json:"time"`
		Error  string `json:"error"`
	} `json:"SteamStore"`

	ISteamUser struct {
		Online int    `json:"online"`
		Time   int    `json:"time"`
		Error  string `json:"error"`
	} `json:"ISteamUser"`

	IEconItems struct {
		TF2 struct {
			Online int    `json:"online"`
			Time   int    `json:"time"`
			Error  string `json:"error"`
		} `json:"440"`

		DOTA2 struct {
			Online int    `json:"online"`
			Time   int    `json:"time"`
			Error  string `json:"error"`
		} `json:"570"`

		CSGO struct {
			Online int    `json:"online"`
			Time   int    `json:"time"`
			Error  string `json:"error"`
		} `json:"730"`
	} `json:"IEconItems"`

	ISteamGameCoordinator struct {
		TF2 struct {
			Online int    `json:"online"`
			Schema string `json:"schema"`
			Error  int    `json:"error"`
			Stats  struct {
				SpyScore  int `json:"spyScore"`
				EngiScore int `json:"engiScore"`
			} `json:"stats"`
		} `json:"440"`

		DOTA2 struct {
			Online int    `json:"online"`
			Error  string `json:"error"`
			Stats  struct {
				PlayerSearching int `json:"players_searching"`
			} `json:"stats"`
		} `json:"570"`

		CSGO struct {
			Online int    `json:"online"`
			Error  string `json:"error"`
			Stats  struct {
				PlayerSearching  int    `json:"players_searching"`
				AverageWait      int    `json:"average_wait"`
				OngoingMatches   int    `json:"ongoing_matches"`
				ServersAvailable int    `json:"servers_available"`
				ServersOnline    int    `json:"servers_online"`
				MenuURL          string `json:"menu_url"`
				PlayerOnline     int    `json:"players_online"`
			} `json:"stats"`
		} `json:"730"`
	} `json:"ISteamGameCoordinator"`
}

func GetSteamInfo() (*SteamInfoJson, error) {
	resp, err := http.Get("https://steamgaug.es/api/v2")
	if err != nil {
		return nil, err
	}
	var info *SteamInfoJson

	err = json.NewDecoder(resp.Body).Decode(&info)

	if err != nil {
		return nil, err
	}

	return info, nil
}

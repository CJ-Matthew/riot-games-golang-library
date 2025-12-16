package main

type Account struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type AccountRegion struct {
	PUUID    string `json:"puuid"`
	Game string `json:"game"`
	Region  string `json:"region"`
}

type ChampionMastery struct {
    Puuid                        string                    `json:"puuid"`
    ChampionID                   int64                     `json:"championId"`
    ChampionLevel                int                       `json:"championLevel"`
    ChampionPoints               int                       `json:"championPoints"`
    LastPlayTime                 int64                     `json:"lastPlayTime"`
    ChampionPointsSinceLastLevel int64                     `json:"championPointsSinceLastLevel"`
    ChampionPointsUntilNextLevel int64                     `json:"championPointsUntilNextLevel"`
    MarkRequiredForNextLevel     int                       `json:"markRequiredForNextLevel"`
    TokensEarned                 int                       `json:"tokensEarned"`
    ChampionSeasonMilestone      int                       `json:"championSeasonMilestone"`
    ChestGranted                 bool                      `json:"chestGranted,omitempty"`
    MilestoneGrades              []string                  `json:"milestoneGrades"`
    NextSeasonMilestone          *NextSeasonMilestone   `json:"nextSeasonMilestone"`
}

type NextSeasonMilestone struct {
    RequireGradeCounts   map[string]int `json:"requireGradeCounts"`
    RewardMarks          int            `json:"rewardMarks"`
    Bonus                bool           `json:"bonus"`
    TotalGamesRequires   int            `json:"totalGamesRequires"`
    RewardConfig         *RewardConfig `json:"rewardConfig,omitempty"`
}

type RewardConfig struct {
    RewardValue   string `json:"rewardValue"`
    RewardType    string `json:"rewardType"`
    MaximumReward int    `json:"maximumReward"`
}
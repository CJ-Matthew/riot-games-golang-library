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

type ChampionRotation struct {
    MaxNewPlayerLevel int `json:"maxNewPlayerLevel"`
    FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
    FreeChampionIds []int `json:"freeChampionIds"`
}

type ClashPlayer struct {
    PUUID    string `json:"puuid"`
    TeamID    string `json:"teamId"`
    Position    string `json:"position"`
    Role    string `json:"role"`
}

type ClashTeam struct {
    ID    string `json:"id"`
    TournamentID    int `json:"tournamentId"`
    Name    string `json:"name"`
    IconId    int `json:"iconId"`
    Tier int `json:"tier"`
    Captain string `json:"captain"`
    Abbreviation string `json:"abbreviation"`
    Players []ClashPlayer `json:"players"`
}

type ClashTournament struct {
    ID    int `json:"id"`
    ThemeID    int `json:"themeId"`
    NameKey    string `json:"nameKey"`
    NameKeySecondary    string `json:"nameKeySecondary"`
    Schedule []ClashTournamentPhase `json:"schedule"`
}

type ClashTournamentPhase struct {
    ID    int `json:"id"`
    RegistrationTime    uint64 `json:"registrationTime"`
    StartTime uint64 `json:"startTime"`
    Cancelled bool `json:"cancelled"`
}
package main

import "time"

type World struct {
	VersionInteger     int32
	SavefileType       uint8
	Revision           uint32
	IsFavorite         uint64
	TileFrameImportant []bool
	WorldName          string
	GeneratorSeed      string
	GeneratorVersion   uint64
	UUID               string
	ID                 int32
	BoundsVec          []int32
	WorldHeight        int32
	WorldWidth         int32
	DifficultyValue    int32

	// World modifiers
	IsDrunkWorld       bool
	IsForTheWorthy     bool
	IsTenthAnniversary bool
	IsTheConstant      bool
	IsBeeWorld         bool
	IsUpsideDown       bool
	IsTrapWorld        bool
	IsZenithWorld      bool
	CreatedOn          string

	// Grouped data (to be implemented)
	// GameProgression    GameProgression
	// SavedNPCs          SavedNPCs
	// Environment        WorldEnvironment
	// WeatherEvents      WeatherAndEvents
	// Invasions          InvasionData

	// Spawn and level data
	SpawnPointX      int32
	SpawnPointY      int32
	UndergroundLevel float64
	CavernLevel      float64
	DungeonPointX    int32
	DungeonPointY    int32
	WorldEvilType    bool

	// Hardmode ores and misc
	HardmodeOre1 int32
	HardmodeOre2 int32
	HardmodeOre3 int32
	Ore1         int32
	Ore2         int32
	Ore3         int32
	Ore4         int32

	// Pets and items
	HasCat             bool
	HasDog             bool
	HasBunny           bool
	CombatBookUsed     bool
	CombatBook2Used    bool
	PeddlerSatchelUsed bool

	// Angler quest data
	AnglerTodayQuestCompletedBy []string
	AnglerDailyQuestTarget      int32

	// Mob data
	MobKills []int32

	// Sundial and moondial
	SundialCooldown   uint8
	SundialIsRunning  bool
	MoondialIsRunning bool
	MoondialCooldown  uint8

	// World content (to be implemented)
	// Tiles              TileMatrix
	// ChestsMaxItems     int16
	// Chests             []Chest
	// Signs              []Sign
	// NPCs               []NPC
	// Mobs               []Mob
	// ShimmeredNPCs      []int32
	// TileEntities       []TileEntity
	// WeighedPressurePlates []WeighedPressurePlate
	// Rooms              []Room
	// Bestiary           Bestiary
	// JourneyPowers      JourneyPowers
}

func NewWorld(worldName, worldSize, difficulty, corruptionType string) *World {
	var worldWidth, worldHeight int32
	switch worldSize {
	case "small":
		worldWidth, worldHeight = 4200, 1200
	case "medium":
		worldWidth, worldHeight = 6400, 1800
	case "large":
		worldWidth, worldHeight = 8400, 2400
	default:
		worldWidth, worldHeight = 4200, 1200
	}

	var difficultyValue int32
	switch difficulty {
	case "journey":
		difficultyValue = 3
	case "classic":
		difficultyValue = 0
	case "expert":
		difficultyValue = 1
	case "master":
		difficultyValue = 2
	default:
		difficultyValue = 0
	}

	worldEvilType := false
	if corruptionType == "crimson" {
		worldEvilType = true
	}

	return &World{
		VersionInteger:              279,
		SavefileType:                2,
		Revision:                    1,
		IsFavorite:                  0,
		TileFrameImportant:          make([]bool, 0),
		WorldName:                   worldName,
		GeneratorSeed:               "go/terraria-wld-parser",
		GeneratorVersion:            0,
		UUID:                        uuid.New().String(),
		ID:                          int32(time.Now().UnixNano()),
		BoundsVec:                   []int32{0, worldWidth * 16, 0, worldHeight * 16},
		WorldHeight:                 worldHeight,
		WorldWidth:                  worldWidth,
		DifficultyValue:             difficultyValue,
		IsDrunkWorld:                false,
		IsForTheWorthy:              false,
		IsTenthAnniversary:          false,
		IsTheConstant:               false,
		IsBeeWorld:                  false,
		IsUpsideDown:                false,
		IsTrapWorld:                 false,
		IsZenithWorld:               false,
		CreatedOn:                   time.Now().Format("2006-01-02 15:04:05.0000000"),
		SpawnPointX:                 worldWidth / 2,
		SpawnPointY:                 worldHeight / 2,
		UndergroundLevel:            0,
		CavernLevel:                 0,
		DungeonPointX:               0,
		DungeonPointY:               0,
		WorldEvilType:               worldEvilType,
		HardmodeOre1:                -1,
		HardmodeOre2:                -1,
		HardmodeOre3:                -1,
		Ore1:                        7,
		Ore2:                        6,
		Ore3:                        9,
		Ore4:                        169,
		HasCat:                      false,
		HasDog:                      false,
		HasBunny:                    false,
		CombatBookUsed:              false,
		CombatBook2Used:             false,
		PeddlerSatchelUsed:          false,
		AnglerTodayQuestCompletedBy: []string{},
		AnglerDailyQuestTarget:      0,
		MobKills:                    make([]int32, 688),
		SundialCooldown:             0,
		SundialIsRunning:            false,
		MoondialIsRunning:           false,
		MoondialCooldown:            0,
		// Diğer alanlar ve içerikler varsayılan olarak boş bırakıldı
	}
}

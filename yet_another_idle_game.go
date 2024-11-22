package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
	sloggin "github.com/samber/slog-gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
	"yet_another_idle_game/battle"
	"yet_another_idle_game/creation"
	"yet_another_idle_game/monolith"
)

var playerCreationId creation.Id
var playerMonolithId monolith.Id

func main() {
	// Get the working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Print the working directory
	fmt.Println("Working directory:", wd)

	// Open the SQLite database file
	db, err := sql.Open("sqlite", wd+"/database.db")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err != nil {
		log.Fatal(err)
	}

	// Create the Gin router
	r := gin.Default()

	creationService := creation.NewCreationService()

	battleService := battle.NewBattleService()
	battleInitializeService := battle.NewBattleInitializeService(battleService, creationService)

	monolithService := monolith.NewMonolithService()
	monolithPriceService := monolith.NewPriceService(monolithService)
	monolithUpgradeService := monolith.NewUpgradeService(monolithService, monolithPriceService, creationService)

	createPlayer(creationService, monolithService)

	runSoulEnergyIncreasing(monolithService)

	runBattles(creationService, battleInitializeService)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r.Use(sloggin.New(logger))

	r.GET("/character/stats", func(c *gin.Context) { getCharacter(c, creationService, monolithService) })
	r.POST("/buy/soulEnergyPerSecond", func(c *gin.Context) { upgradeSoulEnergyPerSecond(c, monolithUpgradeService) })
	r.POST("/buy/damage", func(c *gin.Context) { upgradeDamage(c, monolithUpgradeService) })
	r.POST("/buy/hp", func(c *gin.Context) { upgradeHP(c, monolithUpgradeService) })
	r.GET("/character/battles", func(c *gin.Context) { getBattles(c, battleService, creationService) })

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func runBattles(creationService *creation.CreationService, initializeService *battle.BattleInitializeService) {
	go func() {
		for {
			enemy := &creation.Creation{
				Name:         "Скелет",
				HP:           10,
				MaxHP:        10,
				DamagePerHit: 1,
				AttackSpeed:  1,
			}
			enemyId, _ := creationService.Save(enemy)

			_, _ = initializeService.InitiateBattle(map[battle.FractionId][]creation.Id{
				battle.PLAYER: {
					playerCreationId,
				},
				battle.ENEMY: {
					enemyId,
				},
			})
			time.Sleep(time.Second * 10)
		}
	}()
}

func runSoulEnergyIncreasing(monolithService *monolith.MonolithService) {
	go func() {
		for {
			playerMonolith := monolithService.Get(playerMonolithId)
			playerMonolith.IncreaseSoulEnergyByTick()
			_, _ = monolithService.Save(playerMonolith)
			time.Sleep(time.Second * 1)
		}
	}()
}

func createPlayer(creationService *creation.CreationService, monolithService *monolith.MonolithService) {
	createPlayerCreation(creationService)
	createPlayerMonolith(monolithService)
}

func createPlayerMonolith(monolithService *monolith.MonolithService) {
	playerMonolith := &monolith.Monolith{
		SoulEnergyPerTick: 1,
		SoulEnergy:        0,
		CreationId:        playerCreationId,
	}
	playerMonolithId, _ = monolithService.Save(playerMonolith)
}

func createPlayerCreation(creationService *creation.CreationService) {
	playerCreation := &creation.Creation{
		Name:         "Vpopudanets",
		HP:           10,
		MaxHP:        10,
		DamagePerHit: 1,
		AttackSpeed:  1,
	}
	playerCreationId, _ = creationService.Save(playerCreation)
}

func getCharacter(c *gin.Context, creationService *creation.CreationService, monolithService *monolith.MonolithService) {
	playerCreation := creationService.Get(playerCreationId)
	playerMonolith := monolithService.Get(playerMonolithId)

	characterDto := MainCharacterDto{
		Name:                playerCreation.Name,
		HP:                  playerCreation.HP,
		DamagePerHit:        playerCreation.DamagePerHit,
		DamagePerSecond:     playerCreation.DamagePerHit * playerCreation.AttackSpeed,
		AttackSpeed:         playerCreation.AttackSpeed,
		SoulEnergyPerSecond: playerMonolith.SoulEnergyPerTick,
		SoulEnergy:          playerMonolith.SoulEnergy,
	}

	c.JSON(http.StatusOK, characterDto)
}

type MainCharacterDto struct {
	Name                string `json:"name,omitempty"`
	HP                  int    `json:"HP,omitempty"`
	DamagePerHit        int    `json:"damagePerHit,omitempty"`
	DamagePerSecond     int    `json:"damagePerSecond,omitempty"`
	AttackSpeed         int    `json:"attackSpeed,omitempty"`
	SoulEnergyPerSecond int    `json:"soulEnergyPerSecond,omitempty"`
	SoulEnergy          int    `json:"soulEnergy,omitempty"`
}

func upgradeSoulEnergyPerSecond(c *gin.Context, upgradeService *monolith.UpgradeService) {
	upgradeService.UpgradeSoulEnergyPerTick(playerMonolithId)
}

func upgradeDamage(c *gin.Context, upgradeService *monolith.UpgradeService) {
	upgradeService.UpgradeDamagePerHit(playerMonolithId)
}

func upgradeHP(c *gin.Context, upgradeService *monolith.UpgradeService) {
	upgradeService.UpgradeHP(playerMonolithId)
}

func getBattles(c *gin.Context, battleService *battle.BattleService, creationService *creation.CreationService) {
	battles, err := battleService.GetBattles(playerCreationId)

	if err != nil {
		c.JSON(500, err)
		return
	}

	battleDtos := make([]*BattleDto, 0)

	for _, b := range battles {
		player := creationService.Get(b.Members[battle.PLAYER][0])
		enemy := creationService.Get(b.Members[battle.ENEMY][0])

		battleDtos = append(battleDtos, &BattleDto{
			BattleStatus: b.BattleStatus,
			CharacterInfo: CreationBattleInfoDto{
				Name: player.Name,
			},
			EnemyInfo: CreationBattleInfoDto{
				Name: enemy.Name,
			},
		})
	}

	c.JSON(http.StatusOK, battleDtos)
}

type BattleDto struct {
	BattleStatus  battle.BattleStatus   `json:"battleStatus,omitempty"`
	CharacterInfo CreationBattleInfoDto `json:"characterInfo"`
	EnemyInfo     CreationBattleInfoDto `json:"enemyInfo"`
}

type CreationBattleInfoDto struct {
	Name string `json:"name,omitempty"`
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ks-redis/model"
	"ks-redis/postgres"
	sRedis "ks-redis/redis"
	"ks-redis/setup"
	"net/http"
	"strconv"
	"time"
)

var ctx = context.Background()

func main() {
	// Init Database
	postgres.InitDB("postgres://user:password@localhost:5432/db")
	db := postgres.GetDB()
	defer postgres.CloseDB()

	// Init Redis
	rdb := sRedis.InitRedis()
	defer rdb.Close()

	// Setup Database
	err := setup.Setup(db)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	// Echo server
	e := echo.New()

	// Demo 1 - Load test
	e.GET("/load-test", func(c echo.Context) error {
		sleep10ms()
		//sleep1000ms()
		c.String(http.StatusOK, "ok")
		return nil
	})

	// Demo 1 - Call DB
	e.GET("/latest-members-db", func(c echo.Context) error {

		members, _ := queryLatestMembersFromDatabase(db)
		count, _ := queryCountAllMembersFromDatabase(db)

		resp := map[string]interface{}{
			"status": "ok",
			"count":  count,
			"items":  members,
		}

		c.JSON(http.StatusOK, resp)
		return nil
	})

	// Demo 1 - Call Redis
	e.GET("/latest-members-redis", func(c echo.Context) error {
		cacheTimeout := 60 * 2 * time.Second

		memberCacheKey := "members::latest"
		members := []*model.Member{}
		membersJSON, _ := rdb.Get(ctx, memberCacheKey).Result()
		// Member cache miss
		if membersJSON == "" {
			fmt.Println("Member cache miss")
			members, err = queryLatestMembersFromDatabase(db)

			membersJSON, _ := json.Marshal(members)
			rdb.Set(ctx, memberCacheKey, membersJSON, cacheTimeout)
		}
		// Member cache hit
		if len(membersJSON) > 0 {
			fmt.Println("Member cache hit")
			json.Unmarshal([]byte(membersJSON), &members)
		}

		countCacheKey := "members::count"
		count := -1
		countStr, _ := rdb.Get(ctx, countCacheKey).Result()
		// Count cache miss
		if countStr == "" {
			fmt.Println("Count cache miss")
			count, err = queryCountAllMembersFromDatabase(db)
			rdb.Set(ctx, countCacheKey, count, cacheTimeout)
		}
		// Count cache hit
		if len(countStr) > 0 {
			fmt.Println("Count cache hit")
			count, _ = strconv.Atoi(countStr)
		}

		resp := map[string]interface{}{
			"status": "ok",
			"count":  count,
			"items":  members,
		}

		c.JSON(http.StatusOK, resp)
		return nil
	})

	// Demo 2 - Call Redis using MGET,MSET
	e.GET("/latest-members-redis-v2", func(c echo.Context) error {
		cacheTimeout := 60 * 2 * time.Second

		memberCacheKey := "members::latest"
		countCacheKey := "members::count"

		members := []*model.Member{}
		count := -1

		// Get cache using MGET
		cacheItems, _ := rdb.MGet(ctx, memberCacheKey, countCacheKey).Result()
		membersJSON := cacheItems[0]
		countStr := cacheItems[1]

		itemToCaches := map[string]interface{}{}

		// Member cache miss
		if membersJSON == nil {
			fmt.Println("Member cache v2 miss")
			members, err = queryLatestMembersFromDatabase(db)

			membersJSON, _ := json.Marshal(members)
			itemToCaches[memberCacheKey] = membersJSON
		}
		// Member cache hit
		if membersJSON != nil {
			fmt.Println("Member cache v2 hit")
			json.Unmarshal([]byte(membersJSON.(string)), &members)
		}

		// Count cache miss
		if countStr == nil {
			fmt.Println("Count cache v2 miss")
			count, err = queryCountAllMembersFromDatabase(db)
			itemToCaches[countCacheKey] = count
		}
		// Count cache hit
		if countStr != nil {
			fmt.Println("Count cache v2 hit")
			count, _ = strconv.Atoi(countStr.(string))
		}

		// Set cache using MSET
		if len(itemToCaches) > 0 {
			fmt.Println("Set cache using MSET")
			rdb.MSet(ctx, itemToCaches)
			for key := range itemToCaches {
				rdb.Expire(ctx, key, cacheTimeout)
				fmt.Println("Set cache expire : ", key)
			}
		}

		resp := map[string]interface{}{
			"status": "ok",
			"count":  count,
			"items":  members,
		}

		c.JSON(http.StatusOK, resp)
		return nil
	})

	// Demo 2 - Save to Redis

	e.Start(":8085")
}
func sleep10ms() {
	time.Sleep(10 * time.Millisecond)
}

func sleep1000ms() {
	time.Sleep(1000 * time.Millisecond)
}

func queryLatestMembersFromDatabase(db *gorm.DB) ([]*model.Member, error) {
	var members []*model.Member
	err := db.Where("is_active = ?", 1).Order("register_order DESC").Limit(100).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func queryCountAllMembersFromDatabase(db *gorm.DB) (int, error) {
	var count int64

	err := db.Model(&model.Member{}).Where("is_active = ?", 1).Count(&count).Error
	if err != nil {
		return -1, err
	}

	return int(count), nil
}

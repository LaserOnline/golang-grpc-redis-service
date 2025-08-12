// app.config.go
package config

import (
	"log"
	"strings"
	"time"
)

type RedisConfig struct {
	MasterName    string
	SentinelAddrs []string
	Password      string
	DB            int
	MaxRetries    int
	DialTimeout   time.Duration
	ReadTimeout   time.Duration
	RewriteAddr   func(addr string) string
}

func LoadRedisConfig() RedisConfig {
	LoadDotenv()
	profile := ActiveProfile()

	s1 := strings.TrimSpace(getp(profile, "SENTINEL_1_ADDRESS", ""))
	s2 := strings.TrimSpace(getp(profile, "SENTINEL_2_ADDRESS", ""))
	s3 := strings.TrimSpace(getp(profile, "SENTINEL_3_ADDRESS", ""))

	sentinels := make([]string, 0, 3)
	for _, s := range []string{s1, s2, s3} {
		if s != "" {
			sentinels = append(sentinels, s)
		}
	}
	if len(sentinels) == 0 {
		log.Printf("[redis] WARN: empty SentinelAddrs for profile=%s", profile)
	}

	pw := getp(profile, "REDIS_PASSWORD", "")

	return RedisConfig{
		MasterName:    "mymaster",
		SentinelAddrs: sentinels,
		Password:      pw,
		DB:            0,
		MaxRetries:    3,
		DialTimeout:   2 * time.Second,
		ReadTimeout:   2 * time.Second,
		RewriteAddr:   nil,
	}
}

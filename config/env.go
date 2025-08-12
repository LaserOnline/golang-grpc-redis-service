// env.go
package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var loadOnce sync.Once

func LoadDotenv() {
	loadOnce.Do(func() {
		paths := []string{
			".env",
			"/app/.env",
			"../.env",
			"../../.env",
		}
		for _, p := range paths {
			if _, err := os.Stat(p); err == nil {
				m, err := godotenv.Read(p)
				if err != nil {
					log.Printf("dotenv: cannot read %s: %v", p, err)
					continue
				}
				// set env เฉพาะ key ที่ยังไม่ถูก set และ value ไม่ว่าง
				for k, v := range m {
					if _, exists := os.LookupEnv(k); !exists && strings.TrimSpace(v) != "" {
						_ = os.Setenv(k, v)
					}
				}
				log.Printf("dotenv loaded from %s (%d keys)", p, len(m))
				break // เจอไฟล์แล้วหยุด เพื่อความเสถียร
			}
		}
	})
}

type Profile string

const (
	ProfileUAT Profile = "uat"
	ProfilePRO Profile = "pro"
)

func ActiveProfile() Profile {
	p := strings.ToLower(strings.TrimSpace(os.Getenv("APP_PROFILE")))
	if p == "" {
		return ProfileUAT
	}
	switch p {
	case "uat":
		return ProfileUAT
	case "pro", "prod", "production":
		return ProfilePRO
	default:
		log.Printf("unknown APP_PROFILE=%s, fallback to uat", p)
		return ProfileUAT
	}
}

func prefixFor(p Profile) string {
	if p == ProfilePRO {
		return "PRO_"
	}
	return "UAT_"
}

func getp(p Profile, key, def string) string {
	if v := strings.TrimSpace(os.Getenv(prefixFor(p) + key)); v != "" {
		return v
	}
	return def
}

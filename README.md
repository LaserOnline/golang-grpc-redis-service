# Doc Golang MicroService Redis GRPC

[GO Pckage](https://pkg.go.dev/)

- create package project
```
go mod init github.com/auth-service
```

- update dependency
```
go mod tidy 
```

- package all project
```
go get github.com/joho/godotenv
go get github.com/redis/go-redis/v9
go get google.golang.org/grpc
go get google.golang.org/protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- cmd go controler version
```
gvm help แสดงคำสั่งช่วยเหลือทั้งหมดของ gvm
gvm listall แสดง version ติดตั้ง
gvm list แสดง Go เวอร์ชันที่ติดตั้งแล้ว
gvm use go1.22.0 --default ตั้ง Go เวอร์ชัน 1.22.0 ให้เป็นค่าเริ่มต้นถาวร
gvm use go1.22.0 ใช้งาน Go เวอร์ชัน 1.22.0 (เฉพาะใน shell ปัจจุบัน)
gvm uninstall go1.22.0 ลบ Go เวอร์ชัน 1.22.0 ออกจากระบบ
```

- generate stub go root project
```
protoc -I=.proto \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  .proto/request.proto
```


```
go mod tidy
go clean -cache
go build ./...   # ให้แน่ใจว่าคอมไพล์ผ่านด้วย pb.go ใหม่
go run ./cmd/server   # หรือวิธีที่คุณสตาร์ต
```

---
- golang build project
```
go fmt ./...   # จัดรูปแบบโค้ด
go vet ./...   # ตรวจ warning เบื้องต้น
go test ./...  # ถ้ามี unit test
```

```
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o bin/auth-service-linux cmd/main.go

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o bin/auth-service-arm cmd/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/auth-service.exe cmd/main.go

# macOS ARM
GOOS=darwin GOARCH=arm64 go build -o bin/auth-service-mac cmd/main.go

```

---
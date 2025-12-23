# Website Portfolio (Golang)

Dokumentasi singkat untuk menjalankan dan mengembangkan project portfolio ini.

## Ringkasan singkat
- Backend: Go
- Router: chi
- DB: PostgreSQL (pgx)
- Template: html/template

## Persyaratan
- Go 1.20+
- PostgreSQL (opsional: Docker untuk dev)

## Cara cepat (Quick Start)

1) Jalankan PostgreSQL (opsional pake Docker):

	docker run --name portfolio-db -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=portfolio -p 5432:5432 -d postgres:15

2) Jalankan migrasi (semua tabel sudah ada di `migrations/init.sql`):

	psql -d portfolio -f migrations/init.sql

3) Set environment (contoh):

	export DATABASE_URL="postgres://postgres:pass@localhost:5432/portfolio"
	export SERVER_PORT=8080

4) Jalankan aplikasi:

	go run ./cmd/server

5) Buka: http://localhost:8080

Catatan: folder `migrations` sudah menggabungkan semua perintah ke `init.sql` — file `001/002/003` yang lama sudah dihapus.

## Perintah berguna (Makefile)

- Jalankan: `make run` (atau `go run ./cmd/server`)
- Jalankan test: `make test` atau `go test ./... -v`
- Jalankan migrasi: `make migrate` (menjalankan `migrations/init.sql`)

## Testing

Jalankan semua unit tests:

```
go test ./... -v
```

Untuk melihat coverage:

```
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## CI
Project sudah punya GitHub Actions CI yang menjalankan test dan memeriksa coverage (threshold 50%).

Butuh bantuan lagi? Saya bisa tambahkan `docker-compose.yml` untuk setup dev yang lebih mudah, atau menambahkan badge coverage ke README.
```

```makefile file="Makefile"
.PHONY: run build test clean migrate

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./tests/...

clean:
	rm -rf bin/

migrate:
	psql -d portfolio_db -f migrations/001_create_users.sql
	psql -d portfolio_db -f migrations/002_create_projects.sql
	psql -d portfolio_db -f migrations/003_create_contacts.sql

deps:
	go mod download
	go mod tidy

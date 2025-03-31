# ⚓️ Harbor Health Go Server

This is the backend server for the Harbor Health platform, built with **Go**, **Gin**, **PostgreSQL**, and **SQLC**.  
It handles API requests, business logic, database interactions, authentication, and third-party service integrations.

---

## ⚙️ Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) – HTTP router
- [PostgreSQL](https://www.postgresql.org/) – Relational database
- [SQLC](https://sqlc.dev/) – Type-safe SQL code generation
- [golang-migrate](https://github.com/golang-migrate/migrate) – DB migrations
- [Docker](https://www.docker.com/) – Local Postgres environment
- [Viper](https://github.com/spf13/viper) – Config management
- [Twilio](https://www.twilio.com/), [Brevo](https://www.brevo.com/), [Stream](https://getstream.io/) integrations

---

## 📁 Project Structure

```
.
├── cmd/harbor             # Main entry point
├── internal/
│   ├── db/                # SQLC and migration logic
│   │   ├── migration/     # SQL migration files
│   │   └── sqlc/          # Generated DB code
│   ├── token/             # Auth token logic
│   └── util/              # Email integrations and config
├── app.env.example        # Env variable sample
├── Makefile               # Dev scripts
└── ...
```

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/IamDushu/Harbor-Health-Server.git
cd Harbor-Health-Server
go mod tidy
```

---

## 🐘 Local PostgreSQL Setup

### Start PostgreSQL in Docker

```bash
make postgres
```

### Create Database

```bash
make createdb
```

### Run Migrations

```bash
make migrateup
```

### (Optional) Roll Back Migrations

```bash
make migratedown
```

---

## 🛠️ Configuration with `.env`

Create a file named `.env` in the root directory based on the provided example:

### ✅ `.env` file example

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/harbordb?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080

AUTH_TOKEN_EXPIRY=30m

TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=43200m # 30 days

# Twilio
Twillio_Account_SID=xyz
Twillio_Auth_Token=xyz

# Brevo
Brevo_API_Key=xyz
Template_ID=xyz

# Stream
STREAM_API_KEY=xyz
STREAM_SECRET_KEY=xyz
```

> ⚠️ Never commit your real `.env` file — use `.env.example` for sharing configuration structure.

---

## 🧬 Code Generation & Server Start

### Generate SQLC Code

```bash
make sqlc
```

### Start the Server

```bash
make server
```

The server should now be running at: [http://localhost:8080](http://localhost:8080)

---

## 📦 Makefile Commands

| Command            | Description                          |
| ------------------ | ------------------------------------ |
| `make postgres`    | Start PostgreSQL using Docker        |
| `make createdb`    | Create the `harbordb` database       |
| `make dropdb`      | Drop the `harbordb` database         |
| `make migrateup`   | Apply all up migrations              |
| `make migratedown` | Roll back the latest migration       |
| `make sqlc`        | Generate type-safe SQL code via SQLC |
| `make server`      | Run the server                       |

---

## 🛡️ Security & Auth

- PASETO-based authentication system (Better than JWT)
- Access/Refresh token support
- Encrypted token generation using a 32-byte symmetric key

---

## ☁️ Third-Party Integrations

- 📞 **Twilio** – for phone number verification
- 📧 **Brevo (Sendinblue)** – for transactional emails
- 📹 **Stream** – for video chat and real-time interactions

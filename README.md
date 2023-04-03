# course-forum-forum-servcie
## Table of Contents
- [Configuration Manage](#configuration-manage)
  - [ENV Manage](#env-manage)
  - [Server Configuration](#server-configuration)
  - [Database Configuration](#database-configuration)
- [Installation](#installation)
  - [Local Setup Instruction](#local-setup-instruction)
- [Middlewares](#middlewares)
- [Folder Structure](#folder-structure)
- [Use Packages](#use-packages)

### Configuration Manage

#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`
```text
# Server Configuration
SECRET=
DEBUG=True # `False` in Production
ALLOWED_HOSTS=0.0.0.0
SERVER_HOST=0.0.0.0
SERVER_PORT=8000

# Database Configuration
DB_NAME=course_forum
DB_USER=root
DB_PASSWORD=password
DB_HOST=db
DB_PORT=5432
DB_LOG_MODE=True # `False` in Production
SSL_MODE=disable
```
- Server `DEBUG` set `False` in Production
- Database Logger `DB_LOG_MODE`  set `False` in production
- If ENV Manage from YAML file add a config.yml file and configuration [db.go](config/db.go) and [server.go](config/server.go). See More [ENV YAML Configure](#env-yaml-configure)

#### Server Configuration
- Use [Gin](https://github.com/gin-gonic/gin) Web Framework

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM
- Use database `DB_HOST` value set as `localhost` for local development, and use `db` for docker development 

### Installation
#### Local Setup Instruction
Follow these steps:
- Copy [.env.example](.env.example) as `.env` and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- Check Application health available on [localhost:8000](http://localhost:8000/)

### Middlewares
- Use Gin CORSMiddleware
```go
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(middleware.CORSMiddleware())
```

### Folder Structure
<pre>├── <font color="#3465A4"><b>config</b></font>
│   ├── config.go
│   ├── db.go
│   └── server.go
├── <font color="#3465A4"><b>controllers</b></font>
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── <font color="#3465A4"><b>infra</b></font>
│   ├── <font color="#3465A4"><b>database</b></font>
│   │   └── database.go
│   └── <font color="#3465A4"><b>logger</b></font>
│       └── logger.go
├── LICENSE
├── main.go
├── <font color="#3465A4"><b>migrations</b></font>
│   └── migration.go
├── <font color="#3465A4"><b>models</b></font>
├── README.md
├── <font color="#3465A4"><b>repository</b></font>
├── <font color="#3465A4"><b>routers</b></font>
│   ├── index.go
│   ├── <font color="#3465A4"><b>middleware</b></font>
│   │   └── cors.go
│   └── router.go
</pre>

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.


# eFishery BE Task

Simple service dengan bahasa pemograman go (auth app) & nodejs (fetch app) dalam satu repository (monorepo)

## Doc API (Swagger)

Auth App  : -

Fetch App : -



## Initial Project Setup

Clone Project

```bash
  git clone https://github.com/RindangRamadhan/efishery-be-task.git
```

Go to the project directory

```bash
  cd efishery-be-task
```

1. Create a file `.env` based on file `.env.example`
2. Run in terminal `docker volume create --name=efishery_go-modules` to create a volume **efishery_go-modules**
2. Run in terminal `docker network create efishery-network` to create a network **efishery-network**

Start the server

```bash
  docker-compose up -d
```


## Documentation

1. [Markdown Documentation](https://github.com/didintri196/microservice/blob/master/API.md)
2. [Auth App (Swagger)](https://documenter.getpostman.com/view/3419442/UVeMKQFS)
3. [Fetch App (Swagger)](https://documenter.getpostman.com/view/3419442/UVeMKQFS)

## Context & Deployment

### Context
![Logo](https://raw.githubusercontent.com/RindangRamadhan/efishery-be-task/master/database/services_system.png)

### Deployment
![Logo](https://raw.githubusercontent.com/RindangRamadhan/efishery-be-task/master/database/deployment_system.png)


## Tech Stack

**Server:** Go, Node, Docker, SQLite

## Goals

- [x]  Servers bisa dinyalakan di port berbeda
- [x]  Semua endpoint berfungsi dengan semestinya (3 endpoint auth, 3 endpoint fetch)
- [x]  Dokumentasi endpoint dengan format OpenAPI (API.md)
- [x]  Dokumentasi system diagram-nya dalam format C4 Model (Context dan Deployment)
- [x]  Pergunakan satu repo git untuk semua apps (mono repo)
- [x]  Dockerfile untuk masing-masing app
- [x]  Petunjuk penggunaan dan instalasi di README.md yang memudahkan

## Additional Goals

- [x]  Deployed ke Host/Penyedia Layanan (semacam surge, heroku, vercel, firebase, glitch,
host anda pribadi)
- [x]  Docker Compose
- [ ]  Unit Testing

## Authors

- [@rindangramadhan](https://www.github.com/rindangramadhan)
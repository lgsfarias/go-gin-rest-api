<p align="center">
  <a href="https://github.com/lgsfarias/go-gin-rest-api">
    <img src="./assets/icon.png" alt="readme-logo" width="100">
  </a>

  <h3 align="center">
    GO Gin Rest API
  </h3>
  <p align="center">
    API RESTful com Gin e GORM
    <br />
    <a href="https://github.com/lgsfarias/go-gin-rest-api"><strong>Explore the docs »</strong></a>
    <br />
</p>

## 🎯 Goals

This project was built during a Golang course on the [Alura](https://www.alura.com.br/) platform. The course was taught by Professor [Guilherme Lima](https://www.linkedin.com/in/guilherme-lima-458925178/) and the objective of the course was:

- [x] Create an API RESTful with Go and Gin
- [x] Integrate Go API with a database running on Docker
- [x] Connect to a database using GORM

<br/>

## ⛏️ Built With

![Golang](https://img.shields.io/badge/-Golang-00ADD8?style=flat-square&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/-Gin-00ADD8?style=flat-square&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/-GORM-00ADD8?style=flat-square&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/-Docker-2496ED?style=flat-square&logo=docker&logoColor=white)
![DockerCompose](https://img.shields.io/badge/-DockerCompose-2496ED?style=flat-square&logo=docker&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-336791?style=flat-square&logo=postgresql&logoColor=white)
![git](https://img.shields.io/badge/-git-F05032?style=flat-square&logo=git&logoColor=white)

<br/>

## 🏁 Getting Started

To run this project, you will need to have [Go](https://golang.org/), [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed on your machine.

<br/>

## 🚀 Installation

Clone the repository

```sh
git clone https://github.com/lgsfarias/go-gin-rest-api
```

Acssess the project folder

```sh
cd go-gin-rest-api
```

Fill .env file acording to .env.example

```bash
cp .env.example .env
```

<br/>

## 🏃🏽 Running the project

Up the database

```sh
docker-compose up
```

build the project

```sh
go build
```

run the project

```sh
./go-gin-rest-api
```

<br/>

## 🚀 Routes

```yml
GET /students
- Get all students
```

```yml
GET /students/:id
- Get student by id
```

```yml
POST /students
- Create a new student
- Body: {
  "name": "string",
  "email": "string",
  "cpf": "string"
}
```

```yml
DELETE /students/:id
- Delete student by id
```

```yml
PATCH /students/:id
- Update student by id
- Body: {
  "name": "string",
  "email": "string",
  "cpf": "string"
}
```

```yml
GET /students/cpf/:cpf
- Get student by cpf
```

<br/>

## Contact

<div>
  <a href="https://www.linkedin.com/in/lgsfarias" target="_blank"><img src="https://img.shields.io/badge/-LinkedIn-%230077B5?style=for-the-badge&logo=linkedin&logoColor=white" target="_blank"></a>
  <a href = "mailto:lgsfarias.dev@gmail.com"><img src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white" target="_blank"></a>
</div>

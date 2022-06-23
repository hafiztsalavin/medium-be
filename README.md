# Backend for Article 

## Overview
API to manage Article and Tag to provide basic features of Article (like medium) and implemented in Golang Language. For pass the auth, you can use [login-gate](https://github.com/hafiztsalavin/login-go).

## Model
This for schema db.
![model-db](https://github.com/hafiztsalavin/medium-be/blob/master/docs/images/schema-medium-article-tags.png)

## Todo

- [x] Creates the migrations/seeds for the database.
- [x] CRUD operation.
- [x] Routes guarded and secure by role.
- [x] Docker
- [ ] Documentation API.
- [ ] Unit testing.

## Application Dependency

* [**Go**](https://golang.org/), here using Go v1.17.
* [**PostgreSQL**](https://www.postgresql.org/), and here I'm using postgresql for RDMS.
* [**Docker**](https://www.docker.com/) (optional), if you want to containerize the app and the db.
* [**Postman**](https://www.postman.com/) (optional), to use or test the API.


## Library/Module

* `https://echo.labstack.com` as Go web application framework.
* `https://gorm.io` is ORM (Object Relational Mapping) to help access database also `gorm.io/driver/postgres`as a driver to help GORM access Postgre DB.
* `github.com/golang-jwt/jwt` as a helper to auth JWT.
* `github.com/joeshaw/envdecode` to help fetch environment variable from file to configuration.
* `github.com/go-playground/validator/v10` to help validate request input.

## Running locally
### Using docker
1. Add `.env` file and copy `env.sample` to `.env`. and adjust the environment variable to your own environment.
```
cp env.sample .env
```
2. Pull images GO and Postgresql or you just run in terterminal
   ```
      sudo docker pull postgres
      sudo docker pull golang:1.17-alpine 
   ```
3. Install Go library that was listed in go.mod. 

```
go mod tidy
```
4. Migrate model before start API using.
```
make migrate
```

Actually, you can skip step 2 - 3 using docker-compose
```
   docker-compose up
```
This comment will be run file in `docker-compose.yaml`


### Running local

1. Add `.env` file and copy `env.sample` to `.env`. and adjust the environment variable to your own environment.
```
cp env.sample .env
```
2. Install Go library that was listed in go.mod. 
```
go mod tidy
```
3. Make sure postgres is already based on your `.env` file
4. Create empty database to migrate database if you haven't. 
```
make migrate
```
5. Make a seed data. 
```
make seed
```
6. Run api.
```
make start
```
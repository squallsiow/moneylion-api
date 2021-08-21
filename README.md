## ⭐️ Installation and Run

### Run locally with Golang installed
1. Install GO https://golang.org/doc/install
```bash
export GO111MODULE=on

go run main.go
```
or if you familiar with [Make](https://makefiletutorial.com/)
```bash
make build-local
```
To terminate :
Ctrl+C

### Run with Docker & Docker compose 
1. Install Docker https://docs.docker.com/get-docker/
2. Install Docker compose https://docs.docker.com/compose/install/
```bash
docker build -t moneylion .

docker-compose up -d
```
To terminate :
```
docker-compose down
```

### Infrastructure 
1. [Echo Labstack](https://echo.labstack.com/) - Base web server framework
2. [Mongo Atlas](https://www.mongodb.com/cloud/atlas/lp/try2?utm_source=google&utm_campaign=gs_apac_malaysia_search_core_brand_atlas_desktop&utm_term=mongo%20atlas&utm_medium=cpc_paid_search&utm_ad=e&utm_ad_campaign_id=12212624356&gclid=CjwKCAjwyIKJBhBPEiwAu7zll6wsnieymPwG-ImrkSvJ5j3P1yNDH1-KDe_w0ml0_MtH64hk7tuo1hoC48wQAvD_BwE) - Database 

### Tools

Tools used for this development :

1. [Golang](https://golang.org/doc/install)
2. [Goland](jetbrains.com/go/download/)
3. [POSTMAN](https://www.getpostman.com/)
4. [Mongo Atlas](https://www.mongodb.com/cloud/atlas/lp/try2?utm_source=google&utm_campaign=gs_apac_malaysia_search_core_brand_atlas_desktop&utm_term=mongo%20atlas&utm_medium=cpc_paid_search&utm_ad=e&utm_ad_campaign_id=12212624356&gclid=CjwKCAjwyIKJBhBPEiwAu7zll6wsnieymPwG-ImrkSvJ5j3P1yNDH1-KDe_w0ml0_MtH64hk7tuo1hoC48wQAvD_BwE)

## How to Test
Embedded POSTMAN collection [.json file](https://github.com/squallsiow/moneylion-api/blob/master/moneylion.postman_collection.json), you may use POSTMAN to try out the API directly

## 📚 Available API
### Create User
API [POST]       : http://localhost:8080/user
```json
{
  "name" : "Admin",
  "email" : "admin@email.com"
}
```

### Create Feature
API [POST]       : http://localhost:8080/feature-new
```json
{
  "name" : "Global Settings"
}
```

### Test 1 : Get Feature Access
API [GET]       : localhost:8080/feature?featureName=Global%20Settings&email=admin@email.com

### Test 2: Grant Feature Access
API [POST]       : http://localhost:8080/feature
```json
{
  "featureName" : "Global Settings",
  "email" : "admin@email.com",
  "enable" : true
}
```

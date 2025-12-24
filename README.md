# contacts

## local

### depencencies
* docker 25+
* nvm 0.39+
* golang 1.25

open one terminal
```bash
make compose-up
```

open another terminal
```bash
cd apps/api
go mod download
go run .
```
api running on http://localhost:8010/contacts


and another terminal
```bash
cd apps/web
nvm use
npm i
npm start
```
web running on http://localhost:1234



## validate docker images

for api
```bash
cd apps/api
docker build --tag contacts-api:local .
```

for web
```bash
cd apps/web
docker build --tag contacts-web:local .
```
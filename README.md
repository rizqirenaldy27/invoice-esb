# invoice-esb

## Step Running application
1. Rubah .env.example menjadi .env
2. Lakukan go mod download
```
go mod download
```
3. Lakukan create database dan table dari file invoice.sql
4. Build docker
```
docker build -t invoice-esb .
```
5. Run image docker
```
docker run -d -p 4000:4000 invoice-esb
```
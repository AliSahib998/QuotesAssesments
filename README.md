# Quotes API

This is a simple crud operations which return random quote to user

## Build
```groovy
go build
```

## Run
```groovy
docker compose up -d
```

## Endpoints

### Authentication and Registration

- POST `/user/register`
    - Description: for user register
- POST `/user/login`
    - Description: for user authentication

### Quotes Api

- GET `/quote?priority={low, high}`
    - Description: get quote random from client or high rated or low rated from elastic search engine
- POST `/quote/{id}`
    - Description: Like given quote
- POST `/quote/search`
    - Description: search quote based criteria

# How to run:
1. Setup postgres database: Make sure docker is installed. Run `docker compose up -d`
2. Start server: Run `go run main.go`. Migration for database table will be executed when starting application
3. Access `localhost:3005/homepage` on browser to start testing

# Endpoint List:
Frontend:
- GET /homepage -> Show homepage
- GET /input -> Show input form
- GET /output -> Show output form

Backend:
- GET /handphone -> Get list of handphone & provider data
- POST /handphone -> Create handphone & provider data
- POST /handphone/auto -> Generate 25 handphone & provider data
- PATCH /handphone -> Edit specific no_handphone by id
- DELETE /handphone/:id -> Delete specific handphone data by id
- GET /ws -> Accept websocket connection

<hr>
Notes: 

- Database port by default is set to 5450 instead of 5432 to avoid port already in use

- Front end apiUrl is hardcoded to `localhost:3005`, so make sure HOST env variable is set to `localhost:3005`

<br>
Tech stack used:

- Go, with gin 

- Postgres, with GORM ORM

- Websocket, with gorilla/websocket
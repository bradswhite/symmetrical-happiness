# Rest API

### API Route Docs
 - (get) /login
    -> { id, token, username }
 - (post) /sign-up/{username}
    -> { id }
 - (delete) /delete-account/{username} (auth)

 - (get) /software
    -> { id, name, title, description, image, url, user_id, username, created_at }[]
 - (get) /software/id/{software-id}
    -> { id, name, title, description, image, url, user_id, username, created_at }
 - [group together in same handler: /software/{software-id}]
 - (post) (auth)
    -> { id, name, title, description, image, url, user_id, username, created_at }
 - (put) (auth)
    -> { id, name, title, description, image, url, user_id, username, created_at }
 - (delete) (auth)

 - (get) /software-likes/{software-id}
 - [group together in same handler: /software-likes/{software-id}/user/{user-id}]
 - (post) /software-likes/{software-id}/user/{user-id} (auth)
 - (delete) /software-likes/{software-id}/user/{user-id} (auth)


### Command to set up database (locally):
`docker run -d --name exit_db -e POSTGRES_PASSWORD=password -p 5432:5432 postgres`

### Command to run (locally):
`go run cmd/main.go --seed=true` (to seed the database and start)
`go run cmd/main.go`
build for prod...


### Notes:
- ADD CRUD UPDATE FEATURES (WHEN APPLICATABLE) THROUGHTOUT THE APP


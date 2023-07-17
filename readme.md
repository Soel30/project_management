# Build & Run
1. Install go if not installed on your machine.
2. Install MongoDB if not installed on your machine.
3. Important: Change the DB_HOST to localhost (DB_HOST=localhost) in .env and change DB_NAME to your databa
4. Install neccessary Golang packages, `cd cmd` then `go install`
4. Run go `run cmd/main.go`.
Access API using http://localhost:8080

# Note
1. Just Change DB_Name to create new database, because it automatically creates a new database if it doesn't exist
2. Delete auto migrate inside file cmd/main.go, line 33 if if you don't need anymore
```sh
 dbBase.Debug().Migrator().AutoMigrate(&domain.Post{})
```
to
```sh
 dbBase.Debug().Migrator().AutoMigrate()
```
3. I also created a db file if you want to do it manually

4. Postman
   <https://www.postman.com/cloudy-meteor-109482/workspace/sharing-vision/collection/16254877-27031389-9a2b-4388-80c2-35a8a5987133?action=share&creator=16254877>

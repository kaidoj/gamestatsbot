## Discord bot built with Go, MongoDb, Docker

Installation:  
- copy .env.example to .env and fill out the blank fields
- copy discord-bot/config.yml.example and fill out blank fields

Run:  
- ```docker-compose up -d``` or ```docker-compose up``` to see output

Commands:
- ```!mk name``` displays name with little twist
- ```!mk top``` displays top users by message count
- ```!mk admin messages``` updates user messages count from discord api to local db
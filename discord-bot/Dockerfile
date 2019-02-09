FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/kaidoj/gamestatsbot/discord-bot
COPY . .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get github.com/bwmarrin/discordgo && \
    go get -u -v github.com/spf13/viper && \
    go get github.com/globalsign/mgo && \
    go get github.com/Henry-Sarabia/igdb

RUN CGO_ENABLED=0 go build -o main

# final stage
FROM alpine
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build-env /go/src/github.com/kaidoj/gamestatsbot/discord-bot/config.yml .
COPY --from=build-env /go/src/github.com/kaidoj/gamestatsbot/discord-bot/main . 
ENTRYPOINT ./main


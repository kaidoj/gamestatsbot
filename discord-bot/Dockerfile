FROM golang:alpine
WORKDIR /go/src/github.com/kaidoj/gamestatsbot/discord-bot
COPY . .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get github.com/bwmarrin/discordgo && \
    go get -u -v github.com/spf13/viper && \
    go get github.com/globalsign/mgo

RUN CGO_ENABLED=0 go build -o main . && \
    mv main /usr/local/bin/main

CMD /usr/local/bin/main


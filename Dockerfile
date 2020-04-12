FROM alpine
ADD TwitchChat-service /TwitchChat-service
ENTRYPOINT [ "/TwitchChat-service" ]

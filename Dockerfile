FROM alpine:3.9

RUN apk add --no-cache ca-certificates

EXPOSE 8080

ADD main /bin/main
ADD youtube-data-235117-firebase-adminsdk-646xx-6b0333d15c.json /etc/TRBackend/youtube-data-235117-firebase-adminsdk-646xx-6b0333d15c.json
ENV GOOGLE_APPLICATION_CREDENTIALS="/etc/TRBackend/youtube-data-235117-firebase-adminsdk-646xx-6b0333d15c.json"

ENTRYPOINT ["./bin/main"]
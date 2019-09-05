FROM balenalib/%%BALENA_MACHINE_NAME%%-alpine-golang:latest as build

WORKDIR /src
COPY . .
RUN make build

FROM scratch
ENTRYPOINT ["/disglair"]

COPY files/slack_com.crt /etc/ssl/certs/slack_com.crt
COPY --from=build /src/bin/disglair /disglair


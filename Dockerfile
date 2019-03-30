FROM alpine:latest

ADD build/docker/one-oauth2-server /

EXPOSE 8080

ENTRYPOINT [ "/one-oauth2-server"]

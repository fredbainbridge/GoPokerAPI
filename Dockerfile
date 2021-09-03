FROM ubuntu:latest
ENV SECRET_PATH=/tmp
COPY PokerAPI /tmp
COPY secrets.json /tmp/secrets.json
WORKDIR /tmp
#RUN ./PokerAPI
EXPOSE 8080
ENTRYPOINT ./PokerAPI

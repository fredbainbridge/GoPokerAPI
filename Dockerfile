FROM ubuntu:latest
COPY PokerAPI /tmp
WORKDIR /tmp
#RUN ./PokerAPI
EXPOSE 8080
ENTRYPOINT ./PokerAPI

FROM ubuntu:latest
LABEL authors="ibanrama"

ENTRYPOINT ["top", "-b"]
FROM ubuntu:xenial

RUN mkdir /app
WORKDIR /app

COPY ./aihealth ./aihealth
EXPOSE 10086

CMD ["/app/aihealth","-config /app/config.yaml"]

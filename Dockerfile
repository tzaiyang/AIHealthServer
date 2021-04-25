FROM alpine

RUN mkdir /app
WORKDIR /app

COPY ./aihealth ./aihealth
EXPOSE 8080

CMD ["/app/aihealth"]

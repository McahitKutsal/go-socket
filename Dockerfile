FROM golang:1.16.2-alpine3.9

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8098

CMD [ "/app/main" ]
FROM golang:latest

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8098

CMD [ "/app/main" ]

#   docker build -t socket-app .
#   docker run -it -p 8098:8098 socket-app
FROM golang:1.21

WORKDIR /server

# Copies your source code into the app directory
COPY . .

RUN go mod download

WORKDIR ./app

RUN go build -o /serverapp

EXPOSE 8080

CMD [ "/serverapp" ]
FROM golang:1.23-alpine

# Install necessary packages
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app .

CMD [ "./app" ]
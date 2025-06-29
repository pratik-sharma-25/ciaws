FROM golang:1.24-alpine

# Install necessary packages
WORKDIR /app

COPY . .

RUN go build -o app .

CMD [ "./app" ]
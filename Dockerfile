FROM golang:1.14

LABEL base.name="demoapiclientwithgo"

WORKDIR /app

COPY . .

RUN go build -o application .

EXPOSE 8080

ENTRYPOINT ["./application"]

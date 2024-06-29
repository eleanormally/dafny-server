FROM golang:1.22 as builder

WORKDIR /dafny-server
COPY . . 
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o main


FROM tchajed/dafny

WORKDIR /dafny-server
COPY --from=builder /dafny-server/main .
COPY --from=builder /dafny-server/allowedOrigins.txt .
# RUN chmod +x ./main
EXPOSE 80
CMD ["./main"]

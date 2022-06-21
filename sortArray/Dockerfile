FROM golang:1.18-alpine AS builder

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sortarray .

FROM alpine:3.15
EXPOSE 8080
COPY --from=builder /build/sortarray .

CMD [ "./sortarray" ]

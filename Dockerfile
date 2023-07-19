FROM golang:1.20.0

WORKDIR /app
ADD . .
RUN go build -o /label-artist-fetcher  ./cmd

EXPOSE 3000
CMD ["/label-artist-fetcher"]
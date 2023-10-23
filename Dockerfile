FROM golang:1.20
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
RUN git config --global --add safe.directory /app # git config for air
COPY go.mod go.sum ./
RUN go mod download
CMD ["air", "-c", ".air.toml"]

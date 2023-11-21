FROM golang:1.20
LABEL authors="jiangtao"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY cmd/*.go ./cmd/
# Build
RUN CGO_ENABLED=0 GOOS=linux go build

CMD ["bash"]
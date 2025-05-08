
FROM golang:1.24.0-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app .

FROM alpine:latest
WORKDIR /home/appuser/

# Install CA certificates and create non-root user
RUN apk add --no-cache ca-certificates && \
    adduser -D appuser && \
    mkdir -p /home/appuser/ && \
    chown appuser:appuser /home/appuser/

# Copy artifacts
COPY --from=build /app/app /home/appuser/app
COPY --chown=appuser:appuser firebase/serviceAccountKey.json /home/appuser/serviceAccountKey.json

RUN chmod +x /home/appuser/app

RUN ls -la /home/appuser/
USER appuser

EXPOSE 8080
ENTRYPOINT ["./app"]

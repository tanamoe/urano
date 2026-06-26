dev.up:
	docker compose up -d

dev.start: 
	OTEL_METRICS_EXPORTER=console \
	OTEL_LOGS_EXPORTER=console \
	OTEL_TRACES_EXPORTER=console \
	go run cmd/urano/main.go

run-dev:
	clear && \
	. scripts/env.sh && \
	go build && \
	./gin-gorm
	
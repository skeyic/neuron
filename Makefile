.PHONY: test
test:
	swag init && \
	go run main.go -v 10 -logtostderr
build:
	go build -gcflags="all=-N -l" -o ./bin/inventory_tracker_app ./cmd/inventory_tracker_app

run: 
	./bin/inventory_tracker_app

all: build run  
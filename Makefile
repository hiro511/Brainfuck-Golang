GO_FILES = $(shell find . -name "*.go")
LDFLAGS = ""
BIN_DIR = bin
TARGET = bf

all: bin/bf

bin/bf: bin $(GO_FILES)
	go build -o $(BIN_DIR)/$(TARGET) $(GO_FILES)

bin:
	mkdir -p $(BIN_DIR)

clean:
	rm -r $(BIN_DIR)

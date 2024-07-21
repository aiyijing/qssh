BINARY_NAME=qssh
BUILD_DIR=bin

ARCHS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64

ARCH_TARGETS=$(ARCHS:%=$(BUILD_DIR)/qssh-%)

build: $(ARCH_TARGETS)

$(BUILD_DIR)/qssh-%:
	@mkdir -p $(BUILD_DIR)
	GOOS=$(word 1,$(subst -, ,$(subst /,-,$*))) GOARCH=$(word 2,$(subst -, ,$(subst /,-,$*))) go build -o $(BUILD_DIR)/qssh-$(subst /,-,$*) main.go

install: build
	install ./$(BUILD_DIR)/qssh-$(GOOS)-$(GOARCH) /usr/local/bin/qssh

clean:
	rm -rf $(BUILD_DIR)

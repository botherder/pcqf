BUILD_FOLDER  = $(shell pwd)/build
ASSETS_FOLDER = $(shell pwd)/assets

FLAGS_LINUX   = GOOS=linux GOARCH=amd64
FLAGS_DARWIN  = GOOS=darwin
FLAGS_WINDOWS = GOOS=windows GOARCH=amd64 CC=i686-w64-mingw32-gcc CGO_ENABLED=1

WINPMEM_URL		= https://github.com/Velocidex/WinPmem/releases/download/v4.0.rc1/winpmem_mini_x86.exe
# OSXPMEM_URL		= https://github.com/google/rekall/releases/download/v1.5.1/osxpmem-2.1.post4.zip

LOCAL_WINPMEM_FILE = "winpmem.exe"
LOCAL_OSXPMEM_FILE = "osxpmem.zip"

lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...

clean:
	rm -f ./cmd/bindata.go
	rm -f rsrc.syso
	rm -rf $(ASSETS_FOLDER)
	rm -rf $(BUILD_FOLDER)

mkdirs:
	@echo "[mkdirs] Creating build folders..."
	@mkdir -p $(BUILD_FOLDER)
	@mkdir -p $(ASSETS_FOLDER)

deps:
	@echo "[deps] Installing dependencies..."
	go mod download
	go get -u github.com/akavel/rsrc/...
	go get -u github.com/go-bindata/go-bindata/v3/...
	go mod tidy
	@echo "[deps] Dependencies installed."

windows: deps mkdirs
	@find $(ASSETS_FOLDER) -type f ! -name $(LOCAL_WINPMEM_FILE) -exec rm -f {} \;

	@if [ ! -f $(ASSETS_FOLDER)/$(LOCAL_WINPMEM_FILE) ]; then          \
		echo "[builder] Downloading WinPmem";                          \
		wget $(WINPMEM_URL) -O $(ASSETS_FOLDER)/$(LOCAL_WINPMEM_FILE); \
	fi

	@echo "[builder] Preparing assets"
	@go-bindata -pkg main -o ./cmd/bindata.go -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)/
	@rsrc -manifest pcqf.manifest -o rsrc.syso

	@echo "[builder] Building Windows executable"
	@$(FLAGS_WINDOWS) go build --ldflags '-s -w -extldflags "-static"' -o $(BUILD_FOLDER)/pcqf_windows_amd64.exe ./cmd/

	@echo "[builder] Done!"

darwin: deps mkdirs
# 	@find $(ASSETS_FOLDER) -type f ! -name $(LOCAL_OSXPMEM_FILE) -exec rm -f {} \;

# 	@if [ ! -f $(ASSETS_FOLDER)/$(LOCAL_OSXPMEM_FILE) ]; then          \
# 		echo "[builder] Downloading OSXPmem";                          \
# 		wget $(OSXPMEM_URL) -O $(ASSETS_FOLDER)/$(LOCAL_OSXPMEM_FILE); \
# 	fi

# 	@echo "[builder] Preparing assets"
# 	@go-bindata -pkg main -o ./cmd/bindata.go -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)/

	@echo "[builder] Building Darwin amd64 executable"
	@$(FLAGS_DARWIN) GOARCH=amd64 go build --ldflags '-s -w' -o $(BUILD_FOLDER)/pcqf_darwin_amd64 ./cmd/

	@echo "[builder] Building Darwin arm64 executable"
	@$(FLAGS_DARWIN) GOARCH=arm64 go build --ldflags '-s -w' -o $(BUILD_FOLDER)/pcqf_darwin_arm64 ./cmd/

	@echo "[builder] Done!"

linux: deps mkdirs
	@echo "[builder] Building Linux executable"
	@$(FLAGS_LINUX) go build --ldflags '-s -w' -o $(BUILD_FOLDER)/pcqf_linux_amd64 ./cmd/

	@echo "[builder] Done!"


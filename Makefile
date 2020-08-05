BUILD_FOLDER	= $(shell pwd)/build
ASSETS_FOLDER	= $(shell pwd)/assets

FLAGS_DARWIN	= GOOS=darwin
FLAGS_WINDOWS	= GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CGO_ENABLED=1

WINPMEM_URL		= https://github.com/google/rekall/releases/download/v1.5.1/winpmem-2.1.post4.exe
OSXPMEM_URL		= https://github.com/google/rekall/releases/download/v1.5.1/osxpmem-2.1.post4.zip

lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...

deps:
	@echo "[deps] Installing dependencies..."
	go mod download
	go get github.com/akavel/rsrc
	go get -u github.com/jteeuwen/go-bindata/...
	@echo "[deps] Dependencies installed."

darwin:
	@mkdir -p $(BUILD_FOLDER)/darwin
	@mkdir -p $(ASSETS_FOLDER)

	@if [ ! -f $(ASSETS_FOLDER)/osxpmem.zip ]; then          \
		echo "[builder] Downloading OSXPmem";                \
		wget $(OSXPMEM_URL) -O $(ASSETS_FOLDER)/osxpmem.zip; \
	fi

	@echo "[builder] Preparing assets"
	@go-bindata -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)/

	@echo "[builder] Building Darwin executable"
	@$(FLAGS_DARWIN) go build --ldflags '-s -w' -o $(BUILD_FOLDER)/darwin/snoopdigg

	@echo "[builder] Done!"

windows:
	@mkdir -p $(BUILD_FOLDER)/windows
	@mkdir -p $(ASSETS_FOLDER)

	@if [ ! -f $(ASSETS_FOLDER)/winpmem.exe ]; then          \
		echo "[builder] Downloading WinPmem";                \
		wget $(WINPMEM_URL) -O $(ASSETS_FOLDER)/winpmem.exe; \
	fi

	@echo "[builder] Preparing assets"
	@go-bindata -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)/
	@rsrc -manifest snoopdigg.manifest -ico snoopdigg.ico -o rsrc.syso

	@echo "[builder] Building Windows executable"
	@$(FLAGS_WINDOWS) go build --ldflags '-s -w -extldflags "-static"' -o $(BUILD_FOLDER)/windows/snoopdigg.exe

	@echo "[builder] Done!"

clean:
	rm -f bindata.go
	rm -f rsrc.syso
	rm -rf $(ASSETS_FOLDER)
	rm -rf $(BUILD_FOLDER)

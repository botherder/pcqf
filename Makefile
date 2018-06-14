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
	go get -u github.com/Sirupsen/logrus
	go get -u github.com/mattn/go-colorable
	go get -u github.com/botherder/go-files
	go get -u github.com/botherder/go-autoruns
	go get -u github.com/matishsiao/goInfo
	go get -u github.com/satori/go.uuid
	go get -u github.com/shirou/gopsutil/mem
	go get -u github.com/StackExchange/wmi
	go get -u github.com/mattn/go-shellwords
	go get -u github.com/shirou/w32
	@echo "[deps] Dependencies installed."

darwin:
	@mkdir -p $(BUILD_FOLDER)/darwin
	@mkdir -p $(ASSETS_FOLDER)

	@echo "[builder] Downloading OSXPmem"
	@wget $(OSXPMEM_URL) -O assets/osxpmem.zip

	@echo "[builder] Preparing assets"
	@go-bindata -prefix $(ASSETS_FOLDER) -pkg main $(ASSETS_FOLDER)/osxpmem.zip

	@echo "[builder] Building Darwin executable"
	@go build --ldflags '-s -w' -o $(BUILD_FOLDER)/darwin/snoopdigg

	@echo "[builder] Done!"

windows:
	@mkdir -p $(BUILD_FOLDER)/windows
	@mkdir -p $(ASSETS_FOLDER)

	@echo "[builder] Downloading WinPmem"
	@wget $(WINPMEM_URL) -O assets/winpmem.exe

	@echo "[builder] Preparing assets"
	@go-bindata -prefix $(ASSETS_FOLDER) -pkg main $(ASSETS_FOLDER)/winpmem.exe
	@rsrc -manifest snoopdigg.manifest -ico snoopdigg.ico -o rsrc.syso

	@echo "[builder] Building Windows executable"
	@$(FLAGS_WINDOWS) go build --ldflags '-s -w -extldflags "-static"' -o $(BUILD_FOLDER)/windows/snoopdigg.exe

	@echo "[builder] Done!"

clean:
	rm -f bindata.go
	rm -f rsrc.syso
	rm -rf $(ASSETS_FOLDER)
	rm -rf $(BUILD_FOLDER)

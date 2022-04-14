
arm: 
	wails build --clean --platform darwin/arm64

amd: 
	wails build --clean --platform darwin

win:
	wails build --clean --platform windows/amd64

wails:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
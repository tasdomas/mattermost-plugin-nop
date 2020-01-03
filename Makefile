BUILDS = dist/plugin-linux-amd64 dist/plugin-darwin-amd64 dist/plugin-windows-amd64.exe

package: $(BUILDS) plugin.yaml
	tar -czvf plugin.tar.gz dist plugin.yaml

build: $(BUILDS)

clean:
	rm -rf dist

dist/plugin-linux-amd64: go.mod plugin.go hooks/hooks.go
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o $@ plugin.go

dist/plugin-darwin-amd64: go.mod plugin.go hooks/hooks.go
	mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -o $@ plugin.go

dist/plugin-windows-amd64.exe: go.mod plugin.go hooks/hooks.go
	mkdir -p dist
	GOOS=windows GOARCH=amd64 go build -o $@ plugin.go

generate;
	go generate ./hooks

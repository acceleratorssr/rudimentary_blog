.PHONY: docker

export GOOS = linux
export GOARCH = arm

docker:
	@del blog 2>nul || true
	@go build -o blog .
	@docker rmi -f acc/blog:V0.5 2>nul || true
	@docker build -t acc/blog:V0.5 .
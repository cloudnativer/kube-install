NAME=kube-install
VERSION=v0.5.0

all:kube-install

kube-install:
	@echo Start building kube-install.
	go build -o $(NAME) kube-install.go
	@echo Finished building.


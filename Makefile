NAME=kube-install
VERSION=v0.7.0

all:kube-install

kube-install:
	@echo Start building kube-install.
	go build -o $(NAME)
	@echo Finished building.


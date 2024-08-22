run:
	gow run ./cmd

build:
	go build  -o ~/.gopm/bin/gopm ./cmd

alias:
	echo "alias gopm='~/.gopm/bin/gopm'" >> ~/.profile
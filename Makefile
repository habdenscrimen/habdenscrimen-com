build_bot:
	mkdir -p functions
	go get ./...
	go build -o functions/deploy-succeeded ./tg_bot.go

build_site:
	vuepress build .

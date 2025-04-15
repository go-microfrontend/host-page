build:
	templ generate
	./tailwindcss \
		--input assets/tailwind/tailwind.css \
		--output assets/stylesheets/index.css \
		--minify
	docker compose build

up-app:
	docker compose up -d app

up: up-app

down:
	docker compose down

restart: down up

app-shell:
	docker compose exec app sh

init-tailwind:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

tailwind:
	./tailwindcss \
		--input assets/tailwind/tailwind.css \
		--output assets/stylesheets/index.css

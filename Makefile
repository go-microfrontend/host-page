build:
	templ generate
	# ./tailwindcss \
	# 	--input tailwind/tailwind.css \
	# 	--output assets/css/main.css \
	# 	--minify
	docker compose build

up-app:
	docker compose up -d app

up: up-app

down:
	docker compose down

restart: down up

app-shell:
	docker compose exec app sh

# init-tailwind:
# 	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
# 	chmod +x tailwindcss-linux-x64
# 	mv tailwindcss-linux-x64 tailwindcss
#
# generate-tailwind:
# 	./tailwindcss \
# 		--input tailwind/tailwind.css \
# 		--output assets/css/main.css

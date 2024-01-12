# clone twitter app

# How to setup
1. `git clone ${this repository}`
1. `cd ${this repository}`
1. `cp ./frontend/.env.template ./frontend/.env.local`
1. `docker compose --profile dev up -d --build`

# Warning
1. backendよりdbが先に立ち上がる場合があるため、その時は`docker compose restart backend`をする。


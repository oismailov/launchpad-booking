# launchpad-booking
Book your launchpad and be ready for the journey on time.

## How it works
 Application has multiple commands to parse SpaceX API and populate local database:
- `make refresh-upcoming-launches` - loads upcoming launches
- `make refresh-launchpad-list` - loads available launchpads

These commands can be run as a cron job scripts to periodically get fresh data from external API.  
This is done in that way to split booking creation and SpaceX API parser.

## Project Setup
- Init docker volumes: `make init`
- Start docker containers: `make start`
- Create database schemas: `make migrate`
- Load SpaceX launches and launchpad data: `make load-spase-x-data`
- Open `http://127.0.0.1:9000` in browser to get list of available endpoints (swagger)
- Run tests: `make tests`

## Project Stack
- Golang
- Postgresql
- Docker
- Swagger

## Configuration
- Config files are localed in `config` directory;
- `config/config.json` - local config
- `config/config_test.json` - local for tests

## Development
- If you'd like to apply any changes to the project, please re-build it by running `make re-build` command.

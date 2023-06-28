# DelegationZ - Get latest and historical delegations events on Tezos Chain

## Introduction

Lorem Ipsum Dolor Sic Amet ..

### Specifications

[LINK HERE]()

### Architecture

As per the specifications above, we will need at least three distinct parts :

- [Importer](cmd/importer/main.go) \* :
  First brick to run in order to sync historical datas from reference [tzkt.io](https://api.tzkt.io) api.

- [Watcher](cmd/watcher/main.go) \* :
  A "realtime" or "cron like" job to keep a synchronized state with the reference [tzkt.io](https://api.tzkt.io) api.

- [Api](cmd/api/main.go) :
  A "RESTish" API to expose gathered data.

**\*** Those parts could be combined in a "start sync job if realtime is not up to date or no data in store" binary.

### Implementation

Langage and design : 

Go is the lang because of clarity, consistency and coolest mascot.

To first fetch reference api data's, the need to develop an isolated, tested and reusable tzkt api client wrapper comes pretty quickly* (regarding the requirements only [Delegations](pkg/services/tzkt/delegations.go) is actually implemented).


\* As of 28/06/28 [DipDup Go SDK](https://github.com/dipdup-io/go-lib) seems to be a promising start but currently not used in this project.

Persistence and communication : 

The three parts detailed above should be able to communicate together or at least referring themselves to a single source of truth. In a will to be "simple", this implementation is using Postgres to store and retrieve delegations datas. For realtime need or added observability, a "pubsub style" messaging queue could be later introduced for example to share sync progress or to receive filtered notifications for a particular "baker" or else...



### Limits

- Service is highly dependent to tzkt api for its synchronization speed and overall functionning.
- Relation with the data provider could go wrong and lead to corrupted or no datas.
- Maybe delegation events could be acquired from a node [rpc source](https://tezos.gitlab.io/active/rpc.html) directly ?

## Setup

**Tldr**: With your favorite (UN*X) distribution on hand, fire up your shell and run according to your package manager :
```bash
yay -S go nvm jq bash docker
nvm install 20
nvm use 20
docker run --env POSTGRES_PASSWORD=supersecret --env POSTGRES_DB=dev -d --rm -p 5432:5432 docker.io/postgres:alpine
# Check docker logs , once pg startup is done you should be able to psql "postgres://postgres:supersecret@127.0.0.1:5432/dev" your way in
cd migrations && npm i && npx run prisma migrate deploy
go mod download
go run cmd/api/main.go &
curl "http://localhost:8080/xtz/delegations" | jq
## 2 latest delegation events order by latest first
```

### Go

Install and setup a modern (>=1.20) golang env

#### Sqlboiler

```
go install github.com/volatiletech/sqlboiler/v4@latest 
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
export PATH=$PATH:/$HOME/go/bin
```


### Node , Npm , Npx ...

Install [nvm](https://github.com/nvm-sh/nvm) and use latest nodejs version

### Docker

To help during the development, a postgres db is used, install docker or podman and use 

```
make testdb
```

to test the success of your installation.

## Development

Lorem Ipsum

### Git

#### Conventions

If not already mandatory in your institution, [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/#summary) is a clear and concise way to write your commit messages and easily understand them later.

```
<type>(scope): <description>

[optional body]

[optional footer(s)]
```

Eg :

```
chore(repo): init repo
...
feat(api): pagination filters for /tzc/delegations endpoint
...
```

## Deployment

Regarding the deployment strategy, it should be carefully planned based on required business and engineering requirements such as availability, latency , usage volume / related costs, existing pipelines and infrastucture ...

As far as this project is limited by human and financials means, this repository project is **not**
actually deployed somewhere, sorry about that.

## Misc

During the [importer](cmd/importer/main.go) development, a strange reference api behavior appear, reproduction and developement is summarised in [tzkt_pagination_test.sh](utils/scripts/tzkt_pagination_test.sh) script.

[Makefile](Makefile) usage and pagination test script run : 

[![asciicast](https://asciinema.org/a/9xHrFGtIHTFUAV7lWeUkolxyd.svg)](https://asciinema.org/a/9xHrFGtIHTFUAV7lWeUkolxyd)
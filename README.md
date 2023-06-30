# [DelegationZ](https://delegationz.fly.dev/) - Get latest and historical delegations events on Tezos Chain
[![forthebadge](https://forthebadge.com/images/badges/built-with-grammas-recipe.svg)](https://api.tzkt.io/) [![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://stats.uptimerobot.com/6pOz8UVrqA) 

## Introduction

A Golang service to fetch delegations events from [TZkt.io](https://api.tzkt.io/#operation/Operations_GetDelegations) api, save and present them.

### Specifications

[LINK HERE]()

### Architecture

As per the specifications above, we will need at least three distinct parts :

- [Importer / Watcher](cmd/importer/main.go) \* :
  First brick to run in order to sync historical datas from reference [tzkt.io](https://api.tzkt.io) api. Once sync is completed, a "realtime" or "cron like" job is kept running to keep it.

- [Api](cmd/api/main.go) \*:
  A "RESTish" API to expose gathered data.

**\*** The whole parts are combined [in one binary](cmd/delegationz/main.go) or could be built and deployed separately according to requirements.

### Implementation

Data fetching :

To first fetch reference api data's, the need to develop an isolated, tested and reusable tzkt api client wrapper comes pretty quickly\* (regarding the requirements only [Delegations](pkg/services/tzkt/delegations.go) is actually implemented).

\* As of 28/06/28 [DipDup Go SDK](https://github.com/dipdup-io/go-lib) open source project with the same goals seems to be a promising start but currently not used in this project.

Persistence and communication :

The three parts detailed above should be able to communicate together or at least referring themselves to a single source of truth. In a will to be "simple", this implementation is using Postgres to store and retrieve delegations datas. For realtime need or added observability, a "pubsub style" messaging queue could be later introduced for example to share sync progress or to receive filtered notifications for a particular "baker" or else...

API :

| Path | Params | Response |
|---|:---:|---|
| /* | NA | [HTML](https://delegationz.fly.dev)<br> |
| /xtz/delegations | QueryParams:<br> **year** : YYYY<br>**limit** : <= 100<br>**page** : int  | JSON <br><br>200 :<br> ```[{timestamp: string, amount: string, delegator: string, block: string}, ...]```<br>500 :<br>```{message: string}```<br>400 :<br>```[{timestamp: string, amount: string, delegator: string, block: string}, ...]```<br>500 :<br>```{message: string}``` |
| /xtz/sync | NA | JSON<br><br>200 :<br>```{id: number, timestamp: string, amount: number, delegator: string, block_hash: string, block_level: number}```<br>500 :<br>```{message: string}``` |
| /health | NA | HTTP NO CONTENT CODE 200 |

### Limits

- Service is highly dependent to tzkt api for its synchronization speed and overall functionning.
- Relation with the data provider could go wrong and lead to corrupted or no datas.
- Maybe delegation events could be acquired from a node [rpc source](https://tezos.gitlab.io/active/rpc.html) directly ?

- This implementation lacks a proper caching strategy, a request rate limit had been put in place to avoid database load but it's far from being sufficient if usage increase.

## Setup

**Tldr**: With your favorite (UN\*X) distribution on hand, fire up your shell and run according to your package manager :

```bash
apt-get install golang nvm jq docker
nvm install 20
nvm use 20
make testdb DOCKER=docker
go mod download
go run cmd/api/main.go &
curl "http://localhost:8080/xtz/delegations" | jq
## 2 latest delegation events order by latest first
```

### Go

Install and setup a modern (>=1.20) golang env

#### Sqlboiler

To fetch delegations datas stored in postgres sql without having to worry about golang models drifting or incompactibility with the database scheme, runnning `make gen` will let sqlboiler tool introspect the running database scheme and generate a repository / orm\* style object to use on api resolvers. To use it run :

```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
export PATH=$PATH:/$HOME/go/bin
```

\* Of course, such an orm is only used for the "READ" database operations, the data import processes use "custom" sql queries to avoid too much overhead and optimise the sensible "WRITE" operations.

### Node , Npm , Npx ...

Install [nvm](https://github.com/nvm-sh/nvm) and use latest nodejs version

### Docker

To help during the development, a postgres db is used, install docker or podman and use

```
make testdb
```

to test the success of your installation.

## Development

### Database scheme and migrations

In order to ease, control and track the database scheme(s) definitions, the [prisma](https://www.prisma.io/docs/getting-started/quickstart) tool is used. To udpate the db scheme, update [schema.prisma](./migrations/prisma/schema.prisma) and create a new mutation with
`cd migrations && npx --yes prisma@latest migrate dev`.

### Unit Testing

To illustrate a robust and maintenable project setup, some unit tests have been written for tzkt client and automatically generated for the database repository. Run them with `make test`

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

~~This repository project is **not** actually deployed somewhere, sorry about that.~~ 
**UPDATE** : It is ! 🚀🚀🚀🚀

## Misc

- ~~During the [importer](cmd/importer/main.go) development, a strange reference api behavior appear, reproduction and developement is summarised in [tzkt_pagination_test.sh](utils/scripts/tzkt_pagination_test.sh) script.~~ (Tss .. those developers who can't read a documentation ...)
- In the specifications requirement , the `block` field asked for the api response to looks like a stringified bigint (as the block level) (`"block": "2338084"`) but the reference api is returning the hash like : `"block": "BLwRUPupdhP8TyWp9J6TbjLSCxPPW6tyhVPF2KmNAbLPt7thjPw",` . Current implementation will return the `level` as the `block` response field. It could been suggested that each api features a more descriptive name for their fields as `block_level` or `block_hash`.

- In the specifiction requirements, the `delegator` field has been mapped from `"newDelegate": { "address":"..."}` with no confidence in the fact that `newDelegate` is the asked `delegator`.

- Regarding the delegator field, it could sometimes be empty .. 🤷

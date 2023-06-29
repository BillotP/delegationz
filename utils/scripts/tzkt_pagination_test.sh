#!/usr/bin/env bash
###
##   To run, install  bash and  [jq](https://jqlang.github.io/jq/download/)
###
###
## This script goal is to illustrate "strange" api paginated responses returned by 
## [api.tzkt.io/v1/operations/delegations](https://api.tzkt.io/#operation/Operations_GetDelegations) 
##  endpoint.
###
mkdir out
## Max results in page
limit=600
## Page number
offset=0
## First delegation date RFC3339 format ...
start="2018-06-30T19:30:27Z"
## ... and 48H later 
end="2018-07-02T19:30:27Z"
## delegation status filter 
##  (eg: applied, failed, backtracked, skipped)
dstatus="applied"
echo "[INFO] Delegations items from $start to $end results count aggregate using api /delegations path :"
curl -s "https://api.tzkt.io/v1/operations/delegations?status.eq=$dstatus&limit=$limit&offset=$offset&timestamp.ge=$start&timestamp.lt=$end" | jq '. | length'
echo "[INFO] Delegations items from $start to $end results using api /delegations/count path :"
curl -s "https://api.tzkt.io/v1/operations/delegations/count?timestamp.ge=$start&timestamp.lt=$end"
echo ''
echo "[INFO] Expected 545 items, found 545 || behavior : 👌"
end="2018-07-03T19:30:27Z"
echo "[INFO] Setting end date to 72H later to $end"
fname="out/delegation_${start}_${end}_${limit}_${offset}.json"
curl -s "https://api.tzkt.io/v1/operations/delegations?status.eq=$dstatus&limit=$limit&offset=$offset&timestamp.ge=$start&timestamp.lt=$end" | jq > $fname
echo "[INFO] Delegations items from $start to $end results count aggregate using api /delegations path with page $offset : $(jq '. | length' $fname)"
offset=1
fname="out/delegation_${start}_${end}_${limit}_${offset}.json"
curl -s "https://api.tzkt.io/v1/operations/delegations?status.eq=$dstatus&limit=$limit&offset=$offset&timestamp.ge=$start&timestamp.lt=$end" | jq > $fname
echo "[INFO] Delegations items from $start to $end results count aggregate using api /delegations path with page $offset : $(jq '. | length' $fname)"
count=$(curl -s "https://api.tzkt.io/v1/operations/delegations/count?timestamp.ge=$start&timestamp.lt=$end")
echo "[INFO] Delegations items from $start to $end results using api /delegations/count path : $count"
echo -e "[WARN] /delegations api call on page 0 + page 1 items count doesn't seem to match /delegations/count results : want $count got 1 200 (600*2) || behavior :  ⚠️"

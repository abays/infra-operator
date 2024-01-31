#!/bin/bash
set -eux

case "$1" in
    readiness)
        # ready if we're the master or if we're a slave connected to the current master
        output=$(redis-cli info replication | tr -d '\r')
        declare -A state
        while IFS=: read -r key value; do state[$key]=$value; done < <(echo "$output")
        [[ "${state[role]}" == "master" ]] || [[ "${state[role]}" == "slave" && "${state[master_link_status]}" == "up" ]]
        ;;
    liveness)
        redis-cli -e ping >/dev/null;;
    *)
        echo "Invalid probe option '$1'"
        exit 1;;
esac

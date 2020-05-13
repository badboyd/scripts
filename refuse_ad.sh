#!/bin/bash
DBNAME="blocketdb"
TOKEN=$1
REASON=$2
USERNAME="postgres"
DBHOST="10.50.10.170"
CONNECT_STR="-h $DBHOST -U $USERNAME $DBNAME"
QUERY="SELECT ad_id, action_id FROM ad_queues"
psql $CONNECT_STR -c "$QUERY" | tail -n +3 | while read -a Record; do
	echo "clear ad with id=${Record[0]} action_id=${Record[2]}"
	curl -X "POST" "http://10.50.10.252:32755/v1/ads/refuse" \
	-H "Content-Type: application/json" \
	-d "{\"token\": $TOKEN, \"remote_addr\": \"10.50.10.129\", \"do_not_send_mail\": 1, \
	\"reason\": $2, \"ad_id\": ${Record[0]}, \"action_id\": ${Record[2]}}"
done

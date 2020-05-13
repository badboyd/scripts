#!/bin/bash
DBNAME="blocketdb"

USERNAME="postgres"
DBHOST="10.60.7.10"
CONNECT_STR="-h $DBHOST -U $USERNAME $DBNAME"
QUERY="SELECT ad_id, action_id FROM ad_actions where queue = 'normal' and state = 'unverified'"
psql $CONNECT_STR -c "$QUERY" | tail -n +3 | while read -a Record; do
	echo "clear ad with id=${Record[0]} action_id=${Record[2]}"
	curl -X "POST" "http://10.60.3.47:5657/v1/ads/clear" -H "Content-Type: application/json" -d "{\"ad_id\": ${Record[0]},\"action_id\": ${Record[2]}}"
done





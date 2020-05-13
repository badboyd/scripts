#!/bin/sh

loop_time=$1

echo "Run ${loop_time} times"

for i in `seq 1 $loop_time`
do
	echo "Run $i time"

	curl -X POST "https://cloudgw.chotot.org/v1/private/images/upload" \
             -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2OTY0NTAsImlzcyI6ImNob3RvdCIsInN1YiI6IjEwMDAyMzMxIn0.2jkO_PoudWqgthjqOd8F_i-xZJxI0USCmC1iGHVbVT0" \
             -H "content-type: multipart/form-data" \
             -F "image=@/Users/trandat/Downloads/IMG_2762 copy.jpg"
done

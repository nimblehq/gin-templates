#!/bin/sh

set -x

until psql $DATABASE_URL -c "select 1" > /dev/null 2>&1; do
  sleep 1
done

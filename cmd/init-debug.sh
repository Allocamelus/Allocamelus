#!/bin/sh
echo "-- Debug Start --"
echo "$CONFIG_PATH"
/bin/allocamelus-setup
if [ $? -eq 0 ]; then
  touch $NEW_CONTAINER
  /go/bin/dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec /bin/allocamelus
  else
  echo "Debug Start failed. Please check the logs."
fi
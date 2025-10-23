#!/bin/sh
NEW_CONTAINER="RAN_CONTAINER_SETUP"
if [ ! -e $NEW_CONTAINER ]; then
    echo "-- First Start --"
    /bin/allocamelus-setup
    if [ $? -eq 0 ]; then
      touch $NEW_CONTAINER
      /bin/allocamelus
      else
      echo "Setup failed. Please check the logs."
    fi
else
/bin/allocamelus
fi
#!/bin/bash

attempt_counter=0
max_attempts=60

# this will wait until 60 seconds for Jenkins to be online

until [ "$(curl --head --location --connect-timeout 5 --write-out '%{http_code}' --silent --output /dev/null jenkins:8080)" != 000 ]; do
    if [ ${attempt_counter} -eq ${max_attempts} ];then
      echo "Max attempts reached"
      break
    fi

    echo 'Waiting Jenkins be online.'
    sleep 5
    attempt_counter=$(($attempt_counter+1))
done

sleep 5

# unset the variable to avoid conflicts in the shell
unset attempt_counter
unset max_attempts

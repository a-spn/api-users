#!/bin/bash

if [ -z "$1" ]; then
  echo "[ERROR] Output folder must be given in parameter."
  exit 1
elif [ ! -d "$1" ]; then
  echo "[ERROR]  : \"$1\" is not a valid folder."
  exit 1
fi

openssl genpkey -algorithm RSA -out "$1/access_jwt.key" -pkeyopt rsa_keygen_bits:2048
openssl rsa -in "$1/access_jwt.key" -pubout -outform PEM -out "$1/access_jwt.key.pub"
openssl genpkey -algorithm RSA -out "$1/refresh_jwt.key" -pkeyopt rsa_keygen_bits:2048
openssl rsa -in "$1/refresh_jwt.key" -pubout -outform PEM -out "$1/refresh_jwt.key.pub"
#!/usr/bin/env bash
docker build -t xushikuan/remind-backend .
docker tag xushikuan/remind-backend:latest xushikuan/remind-backend:1.2
docker push xushikuan/remind-backend:1.2
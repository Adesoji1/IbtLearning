#!/bin/sh
kubectl set image deployment/users-api users-container=php:7.4-apache-v1

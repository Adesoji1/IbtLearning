#We can perform a rolling update by running the following command below:
kubectl set image deployment/users-api users-container=php:7.4-apache-v2 

#To perform a rollback, we can run the same command but specify the previous image:
kubectl set image deployment/users-api users-container=php:7.4-apache-v1

# IbtLearning
Assume you’re running on your favorite cloud (Azure, AWS, GCP) - you don’t have to make this work specifically for GCP GKE. Create a README.md that outlines your line of thinking for the solution.

Can you crack this?

Let's see what you know.

___________________
Preamble
Assume you’re running on your favorite cloud (Azure, AWS, GCP) - you don’t have to make this work specifically for GCP GKE.
Create a README.md that outlines your line of thinking for the solution.
Create plain Kubernetes resources (yaml or json). Please return this file in your response with any other materials you want to share with us.
You can make the following assumptions:
Each system you’re deploying has its own isolated database. You don’t have to worry about the type. You can assume the database is in the same region as your k8s.
You can use any docker image you’d like for your containers. It’s just an example and does not have to work. Any, say, default php docker you can deploy on a pod. What the container it is, does not matter - but we’ll be talking about two different containers in the exercise, one for users, and one for shifts.
Assume daily bell-curve scaling. High traffic during the day, low traffic during the night.

Exercise
1.) We want to deploy two containers that scale independently from one another.
Container 1: This container runs code that runs a small API that returns users from a database.
Container 2: This container runs code that runs a small API that returns shifts from a database.
2.) For the best user experience auto scale this service when the average CPU reaches 70%.
3.) Ensure the deployment can handle rolling deployments and rollbacks.
4.) Your development team should not be able to run certain commands on your k8s cluster, but you want them to be able to deploy and roll back. What types of IAM controls do you put in place?

Bonus
·    How would you apply the configs to multiple environments (staging vs production)?
·    How would you auto-scale the deployment based on network latency instead of CPU?

# Solution Below

1. Deploying Two Independent Containers:
To deploy two independent containers, we can create two separate Deployments, each having its own set of replicas, pods, and services. We can also create a horizontal pod autoscaler (HPA) for each Deployment to automatically scale the pods based on CPU utilization. Deploying Two Independent Containers: We can create two Deployments, one for users( view 1.yaml in the Ibt Folder) and one for shifts (view 2.yaml in the Ibt Folder, each with its own set of replicas, pods, and services:



2. Autoscaling the Service:
To auto scale the service, we can use the Kubernetes Horizontal Pod Autoscaler (HPA) with the "CPU utilization" metric. The HPA monitors the average CPU utilization of the pods and adjusts the number of replicas based on the target CPU utilization. The HPA can be configured to scale up or down the replicas to maintain the target CPU utilization percentage. We can set the HPA to scale up when the average CPU utilization reaches 70% and scale down when it goes below 50%.kindly view the autoscale1.yaml manifest file for users and autoscale2.yaml for shift users

```
https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
```

3. Rolling Deployments and Rollbacks:
To handle rolling deployments and rollbacks, we can use Kubernetes Deployments with Rolling Updates. Rolling Updates allows us to deploy new versions of the application without downtime. When a new version is deployed, Kubernetes replaces the old pods with new ones gradually, ensuring that the service remains available throughout the process. If any issues are detected during the deployment, we can perform a rollback by simply rolling back to the previous version.
Rolling Deployments and Rollbacks:
We can enable Rolling Updates for the Deployments by adding the following lines below which has been appended to autoscale1.yaml and autoscale2.yaml:
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0



```
https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/
```
You can perform a rolling update by running the following command below in your bsh terminal :

```
kubectl set image deployment/users-api users-container=php:7.4-apache-v2
```
To perform a rollback, we can run the same command below but specify the previous image :

```
kubectl set image deployment/users-api users-container=php:7.4-apache-v1
```
###### What is Kubectl?  kubectl, allows you to run commands against Kubernetes cluster. Furthermore, The kubectl set command is used to overwrite or set the given cluster. It allows the user to overwrite the property while working similarly to the kubectl run command

4. IAM(Identity and Access Management) Control:
To restrict certain commands from the development team, we can use Kubernetes Role-Based Access Control (RBAC). We can create two roles: one with full access for the admin team and another with limited access for the development team. The limited access role can be configured to allow the team to deploy and roll back, but not to run certain commands.To restrict certain commands on the Kubernetes cluster, we can use RBAC (Role-Based Access Control) and create a Role with the appropriate permissions as seen in restrict.yaml file. Furthermore, This will give the dev-team user permissions to create, update, and delete Deployments, but not other resources such as Pods, Services, or ConfigMaps.


```
https://kubernetes.io/docs/reference/access-authn-authz/rbac/
```

5. Bonus:
To apply the configs to multiple environments (staging vs. production), we can use Kubernetes ConfigMaps and Secrets. We can create separate ConfigMaps and Secrets for each environment and use Kubernetes labels to associate the appropriate resources with each environment. This will allow us to easily manage the configurations for each environment and ensure that the correct configurations are applied to the correct resources. The implementation is found in the config.yaml file

To auto-scale the deployment based on network latency instead of CPU, we can use the Kubernetes Horizontal Pod Autoscaler with the "network latency" metric. This metric can be obtained using Kubernetes custom metrics, which can be gathered using Prometheus or another monitoring tool. We can create a custom metric that measures the network latency between the pods and use it to scale up or down the replicas based on the target network latency.kindly view networklatency.yaml

```

https://cloud.google.com/kubernetes-engine/docs/how-to/horizontal-pod-autoscaling
```

This HPA in networklatency.yaml will scale the users Deployment based on the average network latency between the Pods.

6. To enable network latency metrics, we need to install Kubernetes Metrics Server and create a ServiceMonitor for our application. Here's an example of Installing Metrics Server below :

```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

```

7. Create a ServiceMonitor for users and shifts Deployments is found in servicemonitor.yaml file.

8.Add latency metric to our application code could be written in python or go language. or the latency metric,a `go` script was used and it is named addlatency.go file

Other visible files in the Ibt folder are rolling update, rollback and install metric files


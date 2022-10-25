# Kubernetes 101

## Begriffe

# Kubernetes

OpenSource Software um Container zu orchestrieren (deployment, scaling, management, service discovery, distributed storage).

Notes: Deklarativ, Tags, Resourcen

# POD

Ein oder mehrere Container die in einem Linux Namespace zusammen laufen (teilen sich Resourcen, Netzwerk etc.)

# SideCar Pattern

Ein Container der neben einem anderen Container läuft und "unterstützende" Funktionen bereitstellt.

# Deployment

Beschreibt einen POD meist in YML, welche Tags, wie viele Instanzen, welche Volumes, welche.

# Service

Beschreibung wie ein POD von der Außenwelt (andere PODs, Umgebung) erreichbar ist. (ClusterIP, NodePort, LoadBalancer)

# Node

Bare metal das als Resource im Cluster zur Verfügung steht (kubelet, kube-proxy, container runtime).

# Control Plane

Überwacht den Cluster und stellt sicher dass der gewünschte Zustand eingehalten wird (etcd, kube-scheduler, kube-controller-manager, cloud-controller-manager).

# Namespace

Segmentiert den Cluster in Teilbereiche die "keinen/kontrollierten" Zugriff zwischen anderen Namespaces haben.

# Other stuff

## CMDs

```
gcloud container clusters get-credentials kubernetesbasics-101 --region europe-west1 --project powerpuffgirls

kubectl apply  -f deployment.yml 

```

## Liveness, Readiness and Startup Probes

#### Liveness - If the app is in an broken state and needs a restart.
#### Startup - For slow starting containers. Can it receive traffic yet?
#### Readiness - Temporary not able to receive traffic, does not need a restart, will recover.

## Links

* https://k8syaml.com/ Generate deployment descriptors.
* https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
* https://cloud.google.com/kubernetes-engine/docs/concepts/ingress
* https://cloud.google.com/kubernetes-engine/docs/concepts/container-native-load-balancing
* https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
* https://www.telepresence.io/

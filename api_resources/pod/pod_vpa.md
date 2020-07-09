# 一 概述:
## (1)vertical pod autoscaler(VPA):
- VPA: frees the users from necessity of setting up-to-date resource limits and requests for the containers in their pods. 
- When configured, it will set the requests automatically based on usage and thus allow proper scheduling onto nodes so that appropriate resource amount is available for each pod. 

## (2)备注:
- https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler

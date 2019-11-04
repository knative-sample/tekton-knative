# tekton-knative
tekton knative auto deploy demo

```
kubectl apply -f tasks/source-to-image.yaml -f tasks/deploy-using-kubectl.yaml  -f resources/picalc-git.yaml -f image-secret.yaml -f pipeline-account.yaml -f pipeline/build-and-deploy-pipeline.yaml
```

```
kubectl create -f run/picalc-pipeline-run.yaml
```
# download kubebuilder and install locally.
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/

kubebuilder init --domain kubequant.com --repo github.com/kubequant/KubeQuant
kubebuilder create api --group controller --version v1 --kind KubeQuantRecommendation

make manifests

make install / make uninstall
make run


### Create a new api
```bash
kubebuilder create api --group webapp --version v1 --kind MyResourceName

```


```bash
make docker-build docker-push IMG=<some-registry>/<project-name>:tag
make deploy IMG=<some-registry>/<project-name>:tag

# loading image into the kind
kind load docker-image <your-image-name>:tag --name <your-kind-cluster-name>


```

kubectl apply -k config/samples/

# download kubebuilder and install locally.
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/

kubebuilder init --domain kubequant.com --repo github.com/kubequant/KubeQuant
kubebuilder create api --group controller --version v1 --kind KubeQuantPodResource

make manifests

make install
make run


kubectl apply -k config/samples/

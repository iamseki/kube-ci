FROM golang:alpine as builder-installer

WORKDIR /app
COPY . .

# Building golang app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o kube-ci .

# Download kubectl bin and make it executable
RUN  apk add curl && curl -LO -o usr/bin/kubectl "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" && \
chmod +x kubectl

# Setting kubeconfig properly
# if this envs was provided the kubeconfig will be mounted based on a config downloaded
# from some cloud proider
ARG CONFIG_FILE_URL
ARG CONFIG_PROVIDER
# otherwise will be mounted with cat /app/kubeconfig >> /.kube/config file 
RUN ./kube-ci --config

# Final image layer
FROM alpine

WORKDIR /root

COPY --from=builder-installer /app/kube-ci /app/kubectl /usr/bin/ 
COPY --from=builder-installer /app/kubeconfig .kube/config

#ENTRYPOINT ["kube-ci"]
CMD ["kube-ci","--help"]
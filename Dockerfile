FROM ubuntu
RUN apt update && \
      apt install -y curl && \
      curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && \
      chmod +x ./kubectl && \
      mv ./kubectl /usr/local/bin/kubectl
ENV KUBERNETES_SERVICE_HOST=kubernetes
ENV KUBERNETES_SERVICE_PORT=443
COPY . /app
CMD kubectl cluster-info
CMD ["./app/main"]

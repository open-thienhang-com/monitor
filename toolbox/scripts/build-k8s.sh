#!/bin/bash

# Description: Build K8S
# Author: thienhang.com
# Date: Feb 1, 2024

# Check if the script is run as root
if [[ $EUID -ne 0 ]]; then
    echo "This script must be run as root"
    exit 1
fi

# Install necessary packages
echo "Installing Docker and dependencies..."
apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update && apt-get install -y docker-ce docker-ce-cli containerd.io

# Add Kubernetes apt repository
echo "Adding Kubernetes apt repository..."
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" >/etc/apt/sources.list.d/kubernetes.list
apt-get update

# Install kubeadm, kubelet, and kubectl
echo "Installing Kubernetes components..."
apt-get install -y kubelet kubeadm kubectl

# Enable and start Docker and kubelet services
echo "Enabling and starting Docker and kubelet services..."
systemctl enable docker && systemctl start docker
systemctl enable kubelet && systemctl start kubelet

# Initialize Kubernetes cluster
echo "Initializing Kubernetes cluster with kubeadm..."
kubeadm init

# Configure kubectl for the current user
echo "Configuring kubectl for the current user..."
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config

# (Optional) Deploy pod network (e.g., Calico)
echo "Deploying pod network (Calico)..."
kubectl apply -f https://docs.projectcalico.org/v3.14/manifests/calico.yaml

# (Optional) Allow scheduling pods on the master node
echo "Allowing scheduling pods on the master node..."
kubectl taint nodes --all node-role.kubernetes.io/master-

echo "Kubernetes installation completed successfully!"

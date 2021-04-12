FROM quay.io/podman/stable:latest

ARG OPERATOR_SDK_VERSION=v1.5.0
ARG GOLANG_VERSION=1.15.2
ARG OCCLI_VERSION=4.5.9
ARG OPM_VERSION=1.15.3

# operator-sdk
RUN curl -OJL https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/operator-sdk_linux_amd64 && \
    chmod +x operator-sdk_linux_amd64 && \
    cp operator-sdk_linux_amd64 /usr/local/bin/operator-sdk && \
    rm operator-sdk_linux_amd64

# golang
RUN curl -OJL https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/* /usr/local/bin

# oc cli
RUN curl -OJL https://mirror.openshift.com/pub/openshift-v4/clients/ocp/${OCCLI_VERSION}/openshift-client-linux.tar.gz && \
    tar -C /usr/local/bin -xzf openshift-client-linux.tar.gz && \
    rm openshift-client-linux.tar.gz

# development tools (make, gcc, etc)
RUN yum -y groupinstall "Development Tools" && yum clean all

# opm cli
RUN curl -L https://github.com/operator-framework/operator-registry/releases/download/v${OPM_VERSION}/linux-amd64-opm -o opm && \
    chmod +x opm && \
    mv opm /usr/local/bin

# kustomize
RUN yum -y install findutils && yum clean all && \
    curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | \
    bash && \
    chmod +x kustomize && \
    mv ./kustomize /usr/local/bin/

# ansible-operator
RUN export ARCHOPER=$(uname -m); \
    export OSOPER=$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/darwin/apple-darwin/' | sed 's/linux/linux-gnu/'); \
    curl -L -o ansible-operator https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/ansible-operator_linux_amd64 && \
    chmod +x ansible-operator && \
    mv ansible-operator /usr/local/bin/ansible-operator

# helm-operator
RUN export ARCHOPER=$(uname -m); \
    export OSOPER=$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/darwin/apple-darwin/' | sed 's/linux/linux-gnu/'); \
    curl -L -o helm-operator https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/helm-operator_linux_amd64 && \
    chmod +x helm-operator && \
    mv helm-operator /usr/local/bin/helm-operator

RUN yum -y install which && yum -y install openssl && yum clean all

RUN yum -y install nano

ENV GOPATH=/go
RUN mkdir -p ${GOPATH} && chmod +w ${GOPATH}

RUN yum -y install docker && \
    systemctl enable docker
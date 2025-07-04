FROM eelectronn7998/dev-containers:fedora

RUN dnf update -y
RUN dnf upgrade -y

RUN sudo dnf install make automake gcc gcc-c++ kernel-devel wget -y
RUN git config --global --add safe.directory '*'


RUN wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/workspaces"
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v golang.org/x/tools/gopls@latest

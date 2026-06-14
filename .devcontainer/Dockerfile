ARG UBUNTU_VERSION=24.04
FROM mcr.microsoft.com/devcontainers/base:ubuntu${UBUNTU_VERSION}

# install devtool
RUN apt-get update \
    && apt-get install -y \
        curl \
        sudo \
        unzip \
    && rm -rf /var/lib/apt/lists/*

# install lefhook
ARG LEFTHOOK_VERSION=2.1.9
RUN curl -fsSLO https://github.com/evilmartians/lefthook/releases/download/v${LEFTHOOK_VERSION}/lefthook_${LEFTHOOK_VERSION}_amd64.deb \
    && dpkg -i lefthook_${LEFTHOOK_VERSION}_amd64.deb

# install gitleaks
ARG GITLEAKS_VERSION=8.30.1
RUN curl -fsSLO https://github.com/gitleaks/gitleaks/releases/download/v${GITLEAKS_VERSION}/gitleaks_${GITLEAKS_VERSION}_linux_x64.tar.gz \
    && tar -xf gitleaks_${GITLEAKS_VERSION}_linux_x64.tar.gz -C /tmp \
    && mv /tmp/gitleaks /usr/bin/gitleaks

# install golang
ARG GOLANG_VERSION=1.26.4
RUN curl -fsSLO https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz \
    && tar -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

ARG USERNAME=vscode
USER ${USERNAME}

# persistence bash history
RUN mkdir -p /home/${USERNAME}/.bash \
    && touch /home/${USERNAME}/.bash/.bash_history \
    && chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/.bash \
    && echo 'export HISTFILE="$HOME/.bash/.bash_history"' >> /home/${USERNAME}/.bashrc \
    && echo 'shopt -s histappend' >> /home/${USERNAME}/.bashrc \
    && echo 'PROMPT_COMMAND="history -a;history -n;$PROMPT_COMMAND"' >> /home/${USERNAME}/.bashrc

# install cobra-cli
RUN go install github.com/spf13/cobra-cli@latest \
    && echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> /home/${USERNAME}/.bashrc

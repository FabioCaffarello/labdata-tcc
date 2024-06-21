ARG VARIANT="jammy"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

# Install necessary packages for SSH agent forwarding
RUN apt-get update && \
    apt-get install -y openssh-client

# Create SSH config to use the SSH agent
RUN mkdir -p /home/vscode/.ssh && \
    touch /home/vscode/.ssh/config && \
    echo "Host *\n  ForwardAgent yes" > /home/vscode/.ssh/config && \
    chown -R vscode:vscode /home/vscode/.ssh

EXPOSE 8000
EXPOSE 9000
EXPOSE 27017
EXPOSE 27018
EXPOSE 5672
EXPOSE 15672
EXPOSE 8082
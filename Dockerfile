FROM golang:latest

RUN apt-get update && apt-get install -y \
    iputils-ping \
    net-tools \
    curl \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

CMD ["/bin/bash"]
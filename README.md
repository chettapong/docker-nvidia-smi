# Docker Nvidia SMI Exporter

Dockerized exporter for monitoring of Nvidia GPUs using nvidia-smi written in Go.

Supports multiple GPUs.

# Deps

NVIDIA/nvidia-docker

https://github.com/NVIDIA/nvidia-docker

# Setup

    $ cd docker-nvidia-smi
    $ docker build -t gardenminer/nvidia-smi .
    $ nvidia-docker run -d \
          --name=nvidia-smi \
          -p 9202:9202 \
          gardenminer/nvidia-smi

You can check the output of ccminer with:

    $ docker logs -f nvidia-smi
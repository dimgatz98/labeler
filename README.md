# Installing

```sh
make all
```

# Building docker image
```sh
docker build .
```

# Running locally
```sh
./bin/main
```

# Running on kubernetes
```sh
cd kubernetes && kubectl apply -f . && cd .. 
```

# Using the client 
Example for kubernetes:
```sh
# Node
./bin/client node -i 172.20.0.3 -p 30000 -l "test:123" -n k3d-gputest-server-0
# Pod
./bin/client pod -i 172.20.0.3 -p 30000 -l "test:123" -o dcgm-exporter-1667664784-95x75 
```

Example for local deployment:
```sh
# Node
./bin/client node -l "test:123" -n k3d-gputest-server-0 -c "/home/dimitris/.kube/config"
# Pod
./bin/client pod -l "test:123" -o dcgm-exporter-1667664784-95x75 -c "/home/dimitris/.kube/config"
# Note: Replace "/home/dimitris/.kube/config" with your own kube config path
```

Client Usage:
```sh
./bin/client --help
```

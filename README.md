# metrics-processor
Nats subscriber to OpenFaaS function metrics

## Usage

Expose Nats as a kubernetes service:
```bash
kubectl apply -f kubernetes/nats-external.yml
```

Retrieve Nats hostname and nodeport:
```bash
HOSTNAME=<cluster ip>
PORT=$(kubectl get svc -n openfaas nats-external -o jsonpath="{.spec.ports[0].nodePort}")
```

Run metrics-processor:
```bash
go build -o bin/metrics-processor
bin/metrics-processor -s http://$HOSTNAME:$PORT ${SUBJECT}
```
## Jaeger Go demo

# Build docker image 

cd hello_go_http
docker build . -t venugopalhegde/go-helm-demo:v3
docker push venugopalhegde/go-helm-demo:v3

# Run helm install 
helm install hell-go-helm

# Install Jaeger (Only for development purpose, for production use check https://github.com/jaegertracing/jaeger-kubernetes)
kubectl apply -f jaeger-deployment.yaml

# access your app by port forwarding, and make request to http://localhost:8080/helloworld
kubectl port-forward go-chart-5d458c7bb9-cfdb5 8080:8080


# Access the Jaeger ui with the following steps 
kubectl proxy
http://localhost:8001/api/v1/namespaces/default/services/jaeger-query:80/proxy/search?end=1660237909992000&limit=20&lookback=1h&maxDuration&minDuration&service=my-first-tracer&start=1660234309992000
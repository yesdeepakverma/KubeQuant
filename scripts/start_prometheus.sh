docker network rm prometheus
docker network create prometheus

docker rm prometheus
docker run --name prometheus \
    -d \
    -p 9090:9090 \
    --network prometheus \
    -v ./scripts/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus


docker rm pushgateway
docker run --name pushgateway -d -p 9091:9091 --network prometheus prom/pushgateway


docker ps -a --filter "name=pushgateway|prometheus"
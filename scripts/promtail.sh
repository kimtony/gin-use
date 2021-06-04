
readonly tag=latest

docker pull grafana/promtail:${tag}
docker stop promtail
docker rm promtail

docker run -d  --name promtail \
-v $(pwd):/etc/promtail --restart=always \
-v /data/:/data/ grafana/promtail:${tag} \
-config.file=/data/promtail/promtail.yaml \
grafana/promtail:${tag}
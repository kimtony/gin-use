readonly docker_addr=docker.test.com    #nexus的docker镜像仓库地址
readonly slu_name=go
readonly app_name=gin-use
readonly env=master
readonly tag=
readonly port=8081
readonly file_path=/home/linkai/data/workspace/goWork/src/gin-use

#docker pull ${docker_addr}/${slu_name}/${app_name}/${env}
docker stop ${slu_name}-${app_name}
docker rm ${slu_name}-${app_name}
docker run --restart=always  -d \
--network host \
--name=${slu_name}-${app_name} \
-e CONSUL_HOST=192.168.1.7 \
-e CONSUL_PORT=8500 \
-e CONSUL_CONFIG_PATH=test/test_config \
-v ${file_path}/configs:/data/app/configs \
-v ${file_path}/logs:/data/app/logs \
${docker_addr}/${slu_name}/${app_name}/${env}
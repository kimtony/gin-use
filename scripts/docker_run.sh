readonly docker_addr=docker.dev.isecsp.com
readonly slu_name=go
readonly app_name=gin-use
readonly env=master
readonly tag=
readonly port=8081
readonly file_path=/home/linkai/data/workspace/goWork/src/gin-use

docker pull ${docker_addr}/${slu_name}/${app_name}-${env}
docker stop ${slu_name}-${app_name}
docker rm ${slu_name}-${app_name}
docker run --restart=on-failure:3 \
-p ${port}:${port} -d \
--name=${slu_name}-${app_name} \
-v ${file_path}/docs:/data/app/docs \
-v ${file_path}/configs:/data/app/configs \
-v ${file_path}/logs:/data/app/logs \
${docker_addr}/${slu_name}/${app_name}-${env}
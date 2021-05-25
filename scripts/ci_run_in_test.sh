#!/usr/bin/env bash

echo "开始执行部署......."
slu_name=$1
app_name=$2
branch_name=$3
commit_sha=$4
test_docker_svr=$5
container_port=$6
docker_repo=$7

echo "slu_name:${slu_name}"
echo "app_name:${app_name}"
echo "branch_name:${branch_name}"
echo "commit_sha:${commit_sha}"

echo "test_docker_svr:${test_docker_svr}"
echo "container_port:${container_port}"
echo "docker_repo:${docker_repo}"

echo "docker -H ${test_docker_svr} inspect ${slu_name}-${app_name}"
docker -H ${test_docker_svr} inspect ${slu_name}-${app_name} > /dev/null
isExist=$?

echo "${slu_name}-${app_name} isExist: ${isExist}"
if [ $isExist -eq 0 ]
then
    echo "docker -H ${test_docker_svr} stop ${slu_name}-${app_name}"
    docker -H ${test_docker_svr} stop ${slu_name}-${app_name}

    echo "docker -H ${test_docker_svr} rm ${slu_name}-${app_name}"
    docker -H ${test_docker_svr} rm ${slu_name}-${app_name}
fi

echo "docker -H ${test_docker_svr} run  \\
-v /data/logs/${slu_name}/${app_name}:/root/logs  \\
-p ${container_port}:${container_port}  \\
-d --name=${slu_name}-${app_name}  \\
${docker_repo}/${slu_name}/${app_name}-${branch_name}:${commit_sha} "

docker -H ${test_docker_svr} run  \
-v /data/logs/${slu_name}/${app_name}:/root/logs  \
-p ${container_port}:${container_port}  \
-d --name=${slu_name}-${app_name}  \
${docker_repo}/${slu_name}/${app_name}-${branch_name}:${commit_sha} 
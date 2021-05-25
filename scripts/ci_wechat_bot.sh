#!/usr/bin/env bash
# msg_type =[ docker_img_build ,  push_to_repo , deploy ] 
msg_type=$1
slu_name=$2
name=$3
wechat_hook_key=$4
if [[ ${msg_type} == "docker_img_build" ]];
then
  img_build_state=$5
  user_name=$6
  commit_branch=$7
  commit_sha=$8
  commit_desc=$9
  commit_message=$10

  msg_content="\"${user_name} 对 ${commit_branch} 进行提交镜像打包 . \n  commit sha:${commit_sha}.\n  commit_desc:${commit_desc}.\n  commit_message:${commit_message}\n  docker镜像打包结果: ${img_build_state}！\n  [点击查看](https://git.isecsp.com/${slu_name}/${name}/pipelines)\""
elif [[ ${msg_type} == "notify" ]];
then
  img_build_state=$5
  user_name=$6
  commit_branch=$7
  commit_sha=$8
  commit_desc=$9
  commit_message=$10

  msg_content="\"${user_name} 对 ${commit_branch} 进行提交 测试-打包-推送-部署 . \n  commit sha:${commit_sha}.\n  commit_desc:${commit_desc}.\n  commit_message:${commit_message}\n  进行提交测试-打包-部署后的结果: ${img_build_state}！\n  [点击查看](https://git.isecsp.com/${slu_name}/${name}/pipelines)\""  
fi

echo "
{
  \"msgtype\": \"markdown\",
  \"markdown\": {
    \"content\": ${msg_content}
  }
}"
#https://httpbin.isecsp.com/post
curl "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=${wechat_hook_key}" \
-H 'Content-Type: application/json' \
-d "
{
  \"msgtype\": \"markdown\",
  \"markdown\": {
    \"content\": ${msg_content}
  }
}"

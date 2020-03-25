#!/usr/bin/env bash

registry=10.4.47.129:5000
#registry=192.168.64.1:5000

pull_image() {
    docker pull $registry/$1
    docker tag $registry/$1 $1
    docker rmi $registry/$1
}

push_image() {
    docker tag $1 $registry/$1
    docker push $registry/$1
    docker rmi $registry/$1
}

if [ $# -ge 1 ]; then
    push_image cloustone/pandas-$1
else
    push_image cloustone/pandas-apimachinery
    push_image cloustone/pandas-dmms
    push_image cloustone/pandas-pms
    push_image cloustone/pandas-headmast
    push_image cloustone/pandas-rulechain
    push_image cloustone/pandas-shiro
    push_image cloustone/pandas-lbs
fi
#push_image redis:alpine
#push_image bitnami/rabbitmq
#push_image postgres:latest
#push_image elcolio/etcd

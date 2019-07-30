# ipns

## 产生背景

由于ipfs的特性是按照hash访问内容，但是当内容更新后，用户需要使用新的hash去访问新内容。这就导致用户获取更新困难的问题。故产生了类似于dns的ipns解析服务。

## 原理

将节点id（可比作是域名）绑定到目标文件夹的hash(可比作ip地址)，以后可通过节点id访问该文件夹。

## 过程

### 查看节点id

    ipfs id

### 发布（绑定）

    ipfs name publish 文件夹hash
    响应结果：（响应时间略长，耐心等待）
    Publish to 节点id

### 访问

    http://localhost:8080/ipns/节点id
    注意：是ipns

### 更新（重新部署）

    注意，当内容更新后，解析并不会自动升级，需要

### 反向解析（resolve）

    ipfs name resolve
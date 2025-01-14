<img src="https://grstatic.oss-cn-shanghai.aliyuncs.com/images/rainbond%20log_full.png" width="80%">

[![Go Report Card](https://goreportcard.com/badge/github.com/goodrain/rainbond)](https://goreportcard.com/report/github.com/goodrain/rainbond)
[![GitHub stars](https://img.shields.io/github/stars/goodrain/rainbond.svg?style=flat-square)](https://github.com/goodrain/rainbond/stargazers)
![Rainbond version](https://img.shields.io/badge/version-v5.1-brightgreen.svg)
[![Build Status](https://travis-ci.org/goodrain/rainbond.svg?branch=master)](https://travis-ci.org/goodrain/rainbond)
[![GoDoc](https://godoc.org/github.com/goodrain/rainbond?status.svg)](https://godoc.org/github.com/goodrain/rainbond)

[项目官网](http://www.rainbond.com) • [文档](https://www.rainbond.com/docs/)

**Rainbond** 企业应用云操作系统

Rainbond（云帮）是云原生下企业应用操作系统。 Rainbond支撑企业应用的开发、架构、交付和运维的全流程，通过“无侵入”架构，无缝衔接各类企业应用，底层资源可以对接和管理IaaS、虚拟机和物理服务器。

```
企业应用包括：
各类信息系统、OA、CRM、ERP、数据库、大数据、物联网、互联网平台、微服务架构等运行在企业内部的各种系统。
```

<img src="https://grstatic.oss-cn-shanghai.aliyuncs.com/images/Rainbond%E4%BA%A7%E5%93%81%E6%9E%B6%E6%9E%84.png" width="100%">

## 应用场景

* 企业应用开发

DevOps开发流水线、微服务架构、服务治理及各类技术工具“开箱即用”，不改变开发习惯，让企业专注核心业务，提升10倍效率。

* 企业应用交付

支持持续交付、企业应用市场交付、SaaS化、企业应用销售、二次开发等交付流程，客户统一管理，兼顾标准化交付和个性化交付。

* 企业应用运维

建设私有云平台，透明对接管理多种计算资源，天然实现多云和混合云，企业应用自动化运维，提高2倍资源利用率。业务、资源多维度集成监控系统开箱即用。

## 主要功能特性

| 特性                       | 描述                                                         |
| -------------------------- | ------------------------------------------------------------ |
| Kubernetes             | 平台底层基于Kubernetes，但用户无需学习和编辑复杂的yaml文件，通过应用级图形界面操作使用，实现业务流程开箱即用。 |
| Service Mesh微服务架构 | 内置跨语言、跨协议、代码无侵入的Service Mesh微服务架构原生支持，传统应用直接变成微服务架构。同时支持常见微服务架构Spring Cloud、Dubbo等，通过插件扩展架构能力及治理功能。 |
| 一体化DevOps               | 衔接需求、开发、测试、构建、上线、运维的一体化DevOps。支持对接第三方软件（Jira、Sonar、Jenkins、Gitlab等） |
| 企业级应用市场             | 非镜像市场和服务目录，支持各类企业级应用，像手机应用一样即点即用，全流程管理（应用开发、应用发布、应用展示、应用离线导入/导出、应用安装/升级、应用运维） |
| 自动化运维                 | 应用自动化运维。节点自动安装、扩容、监控、容错。平台支持高可用、多数据中心管理、多租户管理。 |
| Serverless PaaS            | 以应用为核心，使用过程不需要了解服务器相关概念，简单灵活。通过对接行业应用，快速构建行业专有PaaS。 |
| 应用网关                   | 基于HTTP、HTTPs、TCP、UDP等协议应用访问控制策略，轻松操作应用灰度发布、A/B测试。 |
| 异构服务统一管理            | 支持集群内外不同架构服务统一管理、监控和通信治理。                              |

更多功能特性详见： [Rainbond功能特性说明](https://www.rainbond.com/docs/quick-start/edition/)
## 快速开始

1.  [快速安装 Rainbond 集群](https://www.rainbond.com/docs/quick-start/rainbond_install/)
2.  [创建第一个应用（服务）](https://www.rainbond.com/docs/user-manual/app-creation/)
3.  [观看教程视频，快速学习Rainbond](https://www.rainbond.com/video.html)
4.  [搭建 ServiceMesh 微服务架构](https://www.rainbond.com/docs/advanced-scenarios/micro/)

## 社区

[Rainbond 开源社区](https://t.goodrain.com)        

[Rainbond 项目官网](https://www.rainbond.com)

<center><img width="200px" src="https://grstatic.oss-cn-shanghai.aliyuncs.com/images/12141565594759_.pic_hd.jpg"/></center>
<center> 添加微信，申请加入微信群，了解Rainbond更多资讯 </center>

## 参与贡献

你可以参与Rainbond社区关于平台、应用、插件等领域的贡献和分享。

[参与Rainbond项目](https://www.rainbond.com/docs/contribution/)

[Rainbond 贡献者社区](https://t.goodrain.com/c/contribution)

## 相关项目

当前仓库为Rainbond数据中心端核心服务实现代码，项目还包括以下子项目：

   * [Rainbond-Console](https://github.com/goodrain/rainbond-console) Rainbond控制台服务端
   * [Rainbond-Console-UI](https://github.com/goodrain/rainbond-ui) Rainbond控制台前端
   * [Rainbond-Ansible](https://github.com/goodrain/rainbond-ansible) Rainbond安装工具
   * [Rainbond-Builder](https://github.com/goodrain/builder) Rainbond源码构建工具集
   * [Rainbond-Docs](https://github.com/goodrain/rainbond-docs) Rainbond文档
   * Rainbond-Resource/UI (企业版)
   * Rainbond-APP-Store (企业版)

## License

Rainbond 遵循 LGPL-3.0 license 协议发布，详情查看[LICENSE](https://github.com/goodrain/rainbond/blob/master/LICENSE)及[Licensing](https://github.com/goodrain/rainbond/blob/master/Licensing.md)。

## 鸣谢

感谢以下开源项目

- [Kubernetes](https://github.com/kubernetes/kubernetes)
- [Docker/Moby](https://github.com/moby/moby)
- [Heroku Buildpacks](https://github.com/heroku?utf8=%E2%9C%93&q=buildpack&type=&language=)
- [OpenResty](https://github.com/openresty/)
- [Calico](https://github.com/projectcalico)
- [Midonet](https://github.com/midonet/midonet)
- [Etcd](https://github.com/coreos/etcd)
- [Prometheus](https://github.com/prometheus/prometheus)
- [GlusterFS](https://github.com/gluster/glusterfs)
- [Ceph](https://github.com/ceph/ceph)
- [CockroachDB](https://github.com/cockroachdb/cockroach)
- [MySQL](https://github.com/mysql/mysql-server)
- [Weave Scope](https://github.com/weaveworks/scope)
- [Ant Design](https://github.com/ant-design/ant-design)

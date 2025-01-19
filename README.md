# Muxi-Micro-CLI
木犀微服务框架CLI工具

## 命令结构

```shell
muxi
├── init        # 初始化项目
└── help        # 查看帮助
```

## 安装方法
```shell
go install github.com/muxi-Infra/muxi@latest
```

## 使用方法
```shell
muxi init [project-name]
```

## 更新方法
1. 通过muxi命令更新(暂不支持，之后实现)
    ```shell
    muxi update
    ```
2. 再次执行安装命令(建议先清除缓存)
    ```shell 
    go clean -modcache // 清除缓存
    go install github.com/muxi-Infra/muxi@latest
    ```
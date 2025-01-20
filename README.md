# Muxi-Micro-CLI
木犀微服务框架CLI工具

## Notice
如果遇到 `🚫 Failed to remove .git directory ❗️ Please remove .git directory manually` 的报错提示请手动删除 `.git` 文件夹

## 命令结构
```shell
muximicro
├── init        # 初始化项目
└── help        # 查看帮助
```

## 安装方法
```shell
go install github.com/muxi-Infra/muximicro@latest
```

## 使用方法
```shell
muximicro init [project-name]
```

## 更新方法
1. 通过muxi命令更新(暂不支持，之后实现)
    ```shell
    muximicro update
    ```
2. 再次执行安装命令(如果无效先清除缓存)
    ```shell
    go install github.com/muxi-Infra/muximicro@latest
    ```
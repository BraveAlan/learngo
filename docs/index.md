# 电子书制作流程

mkdocs常用于文档写作，官方文档（英文）[传送门](https://mkdocs.readthedocs.io/en/stable/)

## 1. 安装依赖
```sh
pip install mkdocs
```
常用命令如下：

* `mkdocs new [dir-name]` - 创建一个新项目
* `mkdocs serve` - 启动热加载服务器，监测文档变更并在浏览器上实时预览
* `mkdocs build` - 在当前目录下生成site文件夹，将编写的文档转变成HTML文件
* `mkdocs -h` - 命令帮助

## 2. 创建项目

使用`mkdocs new [dir-name]`创建**dir-name**文件夹，其结构如下所示：

    mkdocs.yml    # 配置文件
    docs/
        index.md  # 主页
        ...       # 其他页面

通常是为一个已经存在的git项目编写文档。因此一般的做法是将**dir-name**文件夹放置在项目的根目录下

## 3. 启动mkdocs的内置dev-server

MkDocs带有内置的开发服务器，可用于处理文档时预览文档。在与mkdocs.yml配置文件位于同一目录中，


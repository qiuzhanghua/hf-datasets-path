# 查找Huggingface的数据集路径

## prepare
```bash
go install github.com/qiuzhanghua/autotag@latest
```
来生成go.sum和下载依赖。

## dependency
```bash
go get github.com/qiuzhanghua/common@latest
```

## build
```bash
make all
```

## Usage
```bash
bin/hf-datasets-path rotten_tomatoes
```
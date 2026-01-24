# macOS 安装问题修复指南

## 问题描述

在 macOS 上通过 npm 安装 BrowserWing 后，运行时可能会被系统终止（killed），系统日志显示：

```
load code signature error 2 for file "browserwing"
Security policy would not allow process
```

这是因为二进制文件没有经过 Apple 代码签名，macOS Gatekeeper 阻止了它运行。

## 快速修复（3种方法）

### 方法 1：移除 Quarantine 属性（推荐）

```bash
# 找到 browserwing 安装位置
which browserwing

# 移除 quarantine 属性
xattr -d com.apple.quarantine $(which browserwing)

# 验证
browserwing --version
```

### 方法 2：Ad-hoc 签名

```bash
# 使用 codesign 进行本地签名
codesign -s - $(which browserwing)

# 验证
browserwing --version
```

### 方法 3：临时允许运行

如果上述方法不起作用，可以尝试：

```bash
# 手动运行一次（会弹出安全提示）
$(which browserwing) --version

# 此时打开"系统偏好设置" > "安全性与隐私"
# 点击"仍要打开"按钮
```

## 对于 nvm 用户

如果你使用 nvm 管理 Node.js 版本，路径通常在：

```bash
# 查找实际路径
ls ~/.nvm/versions/node/*/lib/node_modules/browserwing/bin/

# 移除 quarantine（替换实际路径）
xattr -d com.apple.quarantine ~/.nvm/versions/node/v24.13.0/lib/node_modules/browserwing/bin/browserwing*
```

## 验证修复

运行以下命令确认问题已解决：

```bash
browserwing --version
browserwing --help
```

## 为什么会出现这个问题？

1. BrowserWing 的二进制文件是在 Linux CI 环境中构建的
2. 这些二进制文件没有 Apple Developer 证书签名
3. macOS Gatekeeper 默认阻止运行未签名的可执行文件
4. 从互联网下载的文件会被标记 quarantine 属性

## 长期解决方案

我们正在努力实现：
- 使用 Apple Developer 证书对 macOS 二进制进行签名
- 通过 Apple Notarization 服务验证

在此之前，请使用上述临时解决方案。

## 相关链接

- [Apple Gatekeeper 文档](https://support.apple.com/en-us/HT202491)
- [codesign 手册](https://www.unix.com/man-page/osx/1/codesign/)
- [GitHub Issue #XXX](https://github.com/browserwing/browserwing/issues/XXX)

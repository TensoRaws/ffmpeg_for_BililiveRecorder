# 录播姬自动压制插件
## 使用教程
1.  安装 FFmpeg

在使用之前，必须先安装 FFmpeg 并通过 `$PATH` 环境变量访问。

有多种方法可以安装 FFmpeg，例如 [官方下载链接](https://ffmpeg.org/download.html)，或使用您选择的包管理器（例如 Debian/Ubuntu 上的 `sudo apt install ffmpeg`、OS X 上的 `brew install ffmpeg` 等）。

2. 修改配置文件文件  

配置文件说明如下
|  参数   | 内容  |
|  ----  | ----  |
| workdir  | 录播姬工作目录 |
| outputdir | 输出目录 |
| ffmpegtype | 配置序号 |
| ffmpegcom | 自定义参数 |

配置序号表
|  配置序号   | 内容  | 备注 |
|  ----  | ----  | ----  |
|0|执行`ffmpegcom`中的参数||

3. 运行程序

命令行运行`./ffmpeg_for_BililiveRecorder`


## 鸣谢
[wintbiit/brec-sdk-go](https://github.com/wintbiit/brec-sdk-go)
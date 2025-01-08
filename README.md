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
|10|`-i {fullworkpath} -c:v libsvtav1 -c:a copy -crf 40 -preset 6 {fulloutpath}`|av1压制,建议在ffmpeg7.x版本中使用|
|11|`-i {fullworkpath} -c:v libx265 -c:a copy -preset slow {fulloutpath}`|h265压制，该方法有概率压制失败|

3. 运行程序

命令行运行`./ffmpeg_for_BililiveRecorder`

4. 添加webhook  

在录播姬中的`设置`->`Webhook V2`填写`http://127.0.0.1:8080/ffmpeg`


## 鸣谢
[wintbiit/brec-sdk-go](https://github.com/wintbiit/brec-sdk-go)
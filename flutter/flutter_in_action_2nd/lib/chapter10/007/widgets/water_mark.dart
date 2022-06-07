import 'package:flutter/material.dart';
import 'dart:ui' as ui;

class WaterMark extends StatefulWidget {
  const WaterMark({
    Key? key,
    this.repeat = ImageRepeat.repeat,
    required this.painter,
  }) : super(key: key);

  /// 单元水印画笔
  final WaterMarkPainter painter;

  /// 单元水印的重复方式
  final ImageRepeat repeat;

  @override
  _WaterMarkState createState() => _WaterMarkState();
}

class _WaterMarkState extends State<WaterMark> {
  late Future<MemoryImage> _memoryImageFuture;

  @override
  void initState() {
    // 缓存提promise
    _memoryImageFuture = _getWaterMarkImage();
    super.initState();
  }

  /// 离屏绘制单元水印并将绘制结果保存为图片缓存起来
  Future<MemoryImage> _getWaterMarkImage() async {
    // 创建一个canvas 进行离屏绘制
    final recorder = ui.PictureRecorder();
    final canvas = Canvas(recorder);
    // 绘制单元水印并获取其大小
    final Size size = widget.painter.paintUnit(
      canvas,
      MediaQueryData.fromWindow(ui.window).devicePixelRatio,
    );
    final picture = recorder.endRecording();
    // 将单元水印导为图片并缓存起来
    final img = await picture.toImage(size.width.ceil(), size.height.ceil());
    final byteData = await img.toByteData(format: ui.ImageByteFormat.png);
    final pngBytes = byteData!.buffer.asUint8List();

    return MemoryImage(pngBytes);
  }

  @override
  void didUpdateWidget(WaterMark oldWidget) {
    // 如果画笔发生了变化（类型或者配置）则重新绘制水印
    if (widget.painter.runtimeType != oldWidget.painter.runtimeType ||
        widget.painter.shouldRepaint(oldWidget.painter)) {
      // 先释放之前的缓存
      _memoryImageFuture.then((value) => value.evict());
      // 重新绘制并缓存
      _memoryImageFuture = _getWaterMarkImage();
    }

    super.didUpdateWidget(oldWidget);
  }

  @override
  void dispose() {
    // 释放图片缓存
    _memoryImageFuture.then((value) => value.evict());
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox.expand(
      child: FutureBuilder(
          future: _memoryImageFuture,
          builder: (BuildContext context, AsyncSnapshot snapshot) {
            if (snapshot.connectionState != ConnectionState.done) {
              return Container();
            } else {
              // 如果单元水印已经绘制好，则渲染水印
              return DecoratedBox(
                decoration: BoxDecoration(
                  image: DecorationImage(
                    image: snapshot.data, // 背景图，绘制的单元水印图片
                    repeat: widget.repeat, // 重复方式
                    alignment: Alignment.topLeft,
                  ),
                ),
              );
            }
          }),
    );
  }
}

/// 定义水印画笔
abstract class WaterMarkPainter {
  /// 绘制单元水印，完整的水印是由单元水印重复平铺组成，返回值为单元水印占用空间的大小。
  /// [devicePixelRatio]: 因为最终要将绘制内容保存图片，所以在绘制时需要根据屏幕的DPR
  /// 来放大，以防止失真
  Size paintUnit(Canvas canvas, double devicePixelRatio);

  /// 是否需要重绘
  bool shouldRepaint(covariant WaterMarkPainter oldPainter) => true;
}

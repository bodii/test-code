import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

class CustomTapGestureRecognizer extends TapGestureRecognizer {
  @override
  void rejectGesture(int pointer) {
    // 失败
    // super.rejectGesture(pointer);
    // 宣布成功
    super.acceptGesture(pointer);
  }
}

/// 创建一个新GestureDetector,用自定义的CustomTapGestureRecognizer替换默认的
RawGestureDetector customGestureDetector({
  GestureTapCallback? onTap,
  GestureTapDownCallback? onTapDown,
  Widget? child,
}) {
  return RawGestureDetector(
    child: child,
    gestures: {
      CustomTapGestureRecognizer:
          GestureRecognizerFactoryWithHandlers<CustomTapGestureRecognizer>(
        () => CustomTapGestureRecognizer(),
        (detector) => detector.onTap = onTap,
      ),
    },
  );
}

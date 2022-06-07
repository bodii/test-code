import 'package:flutter/material.dart';

/// CustomWidget 包含子组件
class CustomWidget extends LeafRenderObjectWidget {
  const CustomWidget({Key? key}) : super(key: key);

  @override
  RenderObject createRenderObject(BuildContext context) {
    // 创建RenderObject
    return RenderCustomObject();
  }

  @override
  void updateRenderObject(
      BuildContext context, RenderCustomObject renderObject) {
    // 更新 RenderObject
    super.updateRenderObject(context, renderObject);
  }
}

class RenderCustomObject extends RenderBox {
  @override
  void performLayout() {
    // 实现布局逻辑
  }

  @override
  void paint(PaintingContext context, Offset offset) {
    // 实现绘制
  }
}

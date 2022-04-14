import 'package:flutter/material.dart';

class CustomAppBar extends StatelessWidget {
  final Widget leftIcon, rightIcon;
  final Function? leftCallBack, rightCall;

  const CustomAppBar(
      {required this.leftIcon,
      required this.rightIcon,
      this.leftCallBack,
      this.rightCall,
      Key? key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 25),
      padding: EdgeInsets.only(
        top: MediaQuery.of(context).padding.top,
        left: 25,
        right: 25,
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          GestureDetector(
            child: _buildIcon(leftIcon),
            onTap: leftCallBack != null ? () => leftCallBack!() : null,
          ),
          GestureDetector(
            child: _buildIcon(rightIcon),
            onTap: rightCall != null ? () => rightCall!() : null,
          ),
        ],
      ),
    );
  }

  Container _buildIcon(Widget icon) {
    return Container(
      padding: const EdgeInsets.all(8),
      decoration: const BoxDecoration(
        shape: BoxShape.circle,
        color: Colors.white,
      ),
      child: icon,
    );
  }
}

import 'package:flutter/material.dart';

import 'widgets/use_alert_dialog_widget.dart';
import 'widgets/use_futrue_show_dialog_widget.dart';
import 'widgets/use_simple_dialog_widget.dart';
import 'widgets/use_dialog_list_view_widget.dart';
import 'widgets/use_constrained_box_list_view_widget.dart';
import 'widgets/use_custom_dialog_widget.dart';
import 'widgets/use_dialog_state_widget.dart';
import 'widgets/use_dialog_state_builder_widget.dart';
import 'widgets/use_show_modal_bottom_sheet_widget.dart';
import 'widgets/use_show_loading_dialog_widget.dart';
import 'widgets/use_show_loading_dialog_custom_widget.dart';
import 'widgets/use_show_date_picker_01_widget.dart';
import 'widgets/use_show_date_picker_02_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('use Dialog')),
      body: const UseShowDatePicker02Widget(),
    );
  }
}

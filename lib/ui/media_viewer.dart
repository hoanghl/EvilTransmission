import 'package:eviltrans_fe/utils.dart';
import 'package:flutter/material.dart';

class MediaViewer extends StatelessWidget {
  final MediaType mediaType;
  final String mediaPath;
  const MediaViewer(
      {super.key, required this.mediaType, required this.mediaPath});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onPanUpdate: (details) {
        if (details.delta.dx < 0) {
          Navigator.pop(context);
        }
      },
      child: Expanded(
        child: Image.asset(mediaPath),
      ),
    );
  }
}

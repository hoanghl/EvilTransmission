import 'dart:core';

import 'package:eviltrans_fe/data/protocol.dart';
import 'package:eviltrans_fe/ui/media_viewer.dart';
import 'package:eviltrans_fe/utils.dart';
import 'package:flutter/material.dart';

class Viewer extends StatefulWidget {
  const Viewer({super.key});

  @override
  State<Viewer> createState() => _ViewerState();
}

class _ViewerState extends State<Viewer> {
  final List<String> tabsName = ["Image", "Video"];
  List<String>? videoPaths;
  List<String>? imagePaths;
  Future<List>? imagesInfo;

  @override
  void initState() {
    imagePaths = [
      "assets/image.jpg",
      "assets/image.jpg",
      "assets/image.jpg",
      "assets/image.jpg",
      "assets/image.jpg",
      "assets/image.jpg",
      "assets/image.jpg",
    ];
    imagesInfo = getResInfo();

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: tabsName.length,
      child: Scaffold(
        appBar: AppBar(
          leading: IconButton(
            icon: const Icon(Icons.menu),
            onPressed: () {
              // TODO: HoangLe [Mar-07]: Implement this
            },
          ),
          title: const Text(
            "Viewer",
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 20,
            ),
          ),
          actions: [
            IconButton(
              onPressed: () {},
              icon: const Icon(Icons.settings),
            )
          ],
          bottom: TabBar(
            tabs: tabsName.map((e) => Tab(text: e)).toList(),
          ),
        ),
        body: Container(
          padding: const EdgeInsets.all(10),
          child: TabBarView(
            children: [
              FutureBuilder(
                future: imagesInfo,
                builder: (context, snapshot) => snapshot.hasData
                    ? GridView.count(
                        crossAxisCount: 2,
                        children: snapshot.data!.map((e) {
                          var record = e as Map<String, dynamic>;
                          return Column(
                            children: [
                              Text(record['id'].toString()),
                              Text(record['name']),
                              Text(record['res_type']),
                            ],
                          );
                        }).toList(),
                      )
                    : const SizedBox(
                        width: 60,
                        height: 60,
                        child: CircularProgressIndicator(),
                      ),
              ),
              // GridView.count(
              //   crossAxisCount: 2,
              //   children: imagePaths!
              //       .map(
              //         (e) => Container(
              //           padding: const EdgeInsets.all(10),
              //           margin: const EdgeInsets.all(10),
              //           decoration: BoxDecoration(
              //             color: Colors.amberAccent,
              //             borderRadius: BorderRadius.circular(10),
              //           ),
              //           child: Image(
              //             image: AssetImage(e),
              //           ),
              //         ),
              //       )
              //       .toList(),
              // ),
              GridView.count(
                crossAxisCount: 2,
                children: imagePaths!
                    .map(
                      (e) => GestureDetector(
                        onTap: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => MediaViewer(
                                  mediaType: MediaType.image, mediaPath: e),
                            ),
                          );
                        },
                        child: Container(
                          padding: const EdgeInsets.all(10),
                          margin: const EdgeInsets.all(10),
                          decoration: BoxDecoration(
                            color: Colors.amberAccent,
                            borderRadius: BorderRadius.circular(10),
                          ),
                          child: Image(
                            image: AssetImage(e),
                          ),
                        ),
                      ),
                    )
                    .toList(),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

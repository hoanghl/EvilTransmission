import 'dart:convert';

import 'package:flutter/material.dart';
import 'dart:io';
import 'dart:typed_data';

import 'protocol.dart';
import 'utils.dart';
import 'client.dart';

import 'package:file_picker/file_picker.dart';
import 'package:crypto/crypto.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
      ),
      home: HomePage(),
    );
  }
}

class HomePage extends StatelessWidget {
  bool isSetup = false;
  late final Protocol3Vil protocolProcessor;
  HomePage() {
    protocolProcessor = Protocol3Vil();
  }

  void readData(Socket socket) async {
    FilePickerResult? result = await FilePicker.platform.pickFiles();

    if ((result != null) && (result.files.single.path != null)) {
      var path = result.files.single.path;

      var file = File(path!);
      var futureBytes = file.readAsBytes();
      futureBytes.then((Uint8List data) {
        var buffer = protocolProcessor.createPacket(
            PacketTypes.hash, Uint8List.fromList(sha256.convert(data).bytes));
        sendData(buffer);

        // buffer = protocolProcessor.createPacket(PacketTypes.data, data);
        // sendData(buffer);
      });
    }
  }

  void triggerReading() async {
    if (!isSetup) {
      await setup();
      isSetup = !isSetup;
    }
    // connect to the socket server
    // final socket = await Socket.connect(HOST, PORT);
    // print('Connected to: ${socket.remoteAddress.address}:${socket.remotePort}');

    // // Set event Listen for responses from the server
    // socket.listen(
    //   // handle data from the server
    //   (Uint8List data) {
    //     final serverResponse = String.fromCharCodes(data);
    //     print('Server: $serverResponse');
    //   },

    //   // handle errors
    //   onError: (error) {
    //     print(error);
    //     socket.destroy();
    //   },

    //   // handle server ending connection
    //   onDone: () {
    //     print('Server left.');
    //     socket.destroy();
    //   },
    // );

    // Pick file and send packets
    // NOTE: Enable the following after finish testing
    // readData(socket);

    FilePickerResult? result = await FilePicker.platform.pickFiles();

    if ((result != null) && (result.files.single.path != null)) {
      var path = result.files.single.path;

      var file = File(path!);
      var data = await file.readAsBytes();

      var buffer = protocolProcessor.createPacket(
          PacketTypes.hash, Uint8List.fromList(sha256.convert(data).bytes));
      sendData(buffer);

      await Future.delayed(Duration(seconds: 2));

      buffer = protocolProcessor.createPacket(PacketTypes.data, data);
      sendData(buffer);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        child: Center(
          child: ElevatedButton(
            child: Text(
              "Press to send",
              style: TextStyle(fontSize: 20),
            ),
            onPressed: triggerReading,
          ),
        ),
      ),
    );
  }
}

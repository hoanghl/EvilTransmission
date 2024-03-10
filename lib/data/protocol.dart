import 'dart:convert';

import 'package:eviltrans_fe/utils.dart';
import 'package:http/http.dart' as http;

import 'dart:io';

import 'dart:typed_data';

const String host = "192.168.0.163";
const int port = 30556;

Future<List> getResInfo() async {
  var uri = Uri(scheme: "http", host: host, port: port, path: "res");
  var resp = await http.get(uri);
  logger.i("Sent GET to ${uri.toString()}");

  var output = jsonDecode(resp.body) as List;

  return output;
}

// Future<[]byte> getRes(int resId) async {
//   return
// }

void sendData(Uint8List data) async {
  // connect to the socket server
  final socket = await Socket.connect(host, port);
  print('Connected to: ${socket.remoteAddress.address}:${socket.remotePort}');

  // listen for responses from the server
  socket.listen(
    // handle data from the server
    (Uint8List data) {
      final serverResponse = String.fromCharCodes(data);
      print('Server: $serverResponse');
    },

    // handle errors
    onError: (error) {
      print(error);
      socket.destroy();
    },

    // handle server ending connection
    onDone: () {
      print('Server left.');
      socket.destroy();
    },
  );

  socket.add(data);

  await socket.close();
}

void readData(Socket socket, path) async {
  File file;

  if (path != null) {
    file = File(path);

    var futureBytes = file.readAsBytes();
    futureBytes.then((Uint8List value) {
      print("Total bytes: ${value.length}");
      // var tmp = Uint8List(8)
      //   ..buffer.asByteData().setInt64(0, value.length, Endian.big);
      // socket.add(tmp);

      socket.add(value);
    });
  }
}

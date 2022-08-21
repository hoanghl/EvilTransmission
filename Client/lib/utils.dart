import 'dart:typed_data';
import 'dart:io';

import 'package:yaml/yaml.dart';
import 'package:flutter/services.dart' show rootBundle;

List<int> convInt2Byte(int num, int nBytes) {
  var tmp = Uint8List(8)..buffer.asByteData().setInt64(0, num);

  return tmp.getRange(tmp.length - nBytes, tmp.length).toList();
}

int convBytes2Int(List<int> bytes) =>
    Uint8List.fromList(bytes).buffer.asByteData().getInt16(0);

const PACKET_TYPE = 0;
const RESPONSE_TYPE = 1;

class Type {
  late String name;
  late int code;
  late int type;
  Type(this.name, this.code, this.type);
}

class Types {
  late final Set<String> names;
  late final Set<int> codes;
  late final Map mapCode2Name;
  late final Map mapName2Code;

  Types.setup(List types) {
    var nameSet = <String>{}, codeSet = <int>{};
    for (var x in types) {
      nameSet.add(x.name!);
      codeSet.add(x.code!);
    }

    names = nameSet;
    codes = codeSet;
  }
}

class PacketTypes extends Types {
  static final data = Type('data', 0, PACKET_TYPE);
  static final hash = Type('hash', 1, PACKET_TYPE);
  static final response = Type('response', 2, PACKET_TYPE);

  PacketTypes() : super.setup([data, hash, response]);
}

class ResponseTypes extends Types {
  static final internalErr = Type("internal_err", 0, RESPONSE_TYPE);
  static final permitted = Type("permitted", 1, RESPONSE_TYPE);
  static final notPermitted = Type("not_permitted", 2, RESPONSE_TYPE);

  ResponseTypes() : super.setup([internalErr, permitted, notPermitted]);
}

final packetTypes = PacketTypes();
final responseTypes = ResponseTypes();

//==========================================================
// Define filters used for decomposing flag
int filterPacketType = int.parse("00000110", radix: 2);
int filterRepsType = int.parse("00011000", radix: 2);
int filterSucessive = int.parse("00000001", radix: 2);

//==========================================================
// Define shifts

class Shifts {
  static const shifts = {PACKET_TYPE: 1, RESPONSE_TYPE: 3};

  static int getShift(Type type) {
    return shifts[type.type]!;
  }
}

late final int MAX_SIZE_PACKET;
late final String HOST;
late final int PORT;
late final int LEN_FLAG;
late final int LEN_SIZE;

Future setup() async {
  var configs = loadYaml(await rootBundle.loadString('assets/configs.yaml'));

  MAX_SIZE_PACKET = configs['MAX_SIZE_PACKET'];
  HOST = configs['SERVER']['HOST'];
  PORT = configs['SERVER']['PORT'];
  LEN_FLAG = configs['LEN_FLAG'];
  LEN_SIZE = configs['LEN_SIZE'];

  print("Done setting up things...");
}

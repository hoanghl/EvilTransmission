import 'package:eviltrans_fe/data/protocol.dart';
import 'package:eviltrans_fe/utils.dart';

Future<Map<MediaType, List<String>>> getResources() async {
  // Get resource info
  List resInfo = await getResInfo();

  // Get resource by id
  List<Future<PathResource>> list = [];
  list.addAll(resInfo.map((e) {
    var record = e as Map<String, dynamic>;
    return getRes(
      record['id'],
      record['name'],
      MediaType.values.byName(record['res_type']),
    );
  }));

  List<PathResource> pathRes = await list.wait;

  // Establish map of media type
  Map<MediaType, List<String>> resMap = {
    MediaType.image: [],
    MediaType.video: [],
    MediaType.thumbnail: [],
  };
  for (final r in pathRes) {
    logger.i("Path to resource: ${r.resPath}");

    resMap[r.resType]!.add(r.resPath);
  }

  return resMap;
}

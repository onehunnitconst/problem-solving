import 'dart:io';

class Point {
  int x;
  int y;

  Point(this.x, this.y);
}

void main() {
  int n = int.parse(stdin.readLineSync()!);

  List<Point> points = [];

  for (int i = 0; i < n; i++) {
    List<int> input =
        stdin.readLineSync()!.split(' ').map((e) => int.parse(e)).toList();
    points.add(Point(input[0], input[1]));
  }

  points.sort((a, b) {
    if (a.y == b.y) {
      return a.x - b.x;
    } else {
      return a.y - b.y;
    }
  });

  for (final point in points) {
    print('${point.x} ${point.y}');
  }
}

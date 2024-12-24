import java.util.Scanner

data class Point (val x: Int, val y: Int)

fun main(args: Array<String>) {
  val sc: Scanner = Scanner(System.`in`)
  var n = sc.nextInt()

  var points: MutableList<Point> = mutableListOf();

  for (i in 0 until n) {
    var x = sc.nextInt()
    var y = sc.nextInt()
    points.add(Point(x, y))
  }

  points.sortWith(compareBy({it.y}, {it.x}))

  for (point in points) {
    println("${point.x} ${point.y}")
  }
}
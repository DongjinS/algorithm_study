fun main() {
    val x = readLine()!!.toInt()
    val y = readLine()!!.toInt()

    if (x > 0 && y > 0) {
        println(1)
    } else if (x < 0 && y > 0) {
        println(2)
    } else if (x < 0 && y < 0) {
        println(3)
    } else if (x > 0 && y < 0) {
        println(4)
    }
}
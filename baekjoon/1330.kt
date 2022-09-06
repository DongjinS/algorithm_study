import java.util.Scanner

fun main() {
    val sc : Scanner = Scanner(System.`in`)
    var a = sc.nextInt()
    var b = sc.nextInt()

    if (a>b) {
        println(">")
    } else if (a<b) {
        println("<")
    } else println("==")
}
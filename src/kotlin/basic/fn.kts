val n = listOf(1,2,3,4,5)
n.fold(0, {
    // param参数
    acc: Int, i: Int ->
    println("hello")
    val res = acc + i
    println(res)
    // return返回值
    res
})
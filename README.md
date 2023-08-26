# the JOJO programming language

JOJO 是一个使用 Golang 实现的类 TypeScript 语言。

JOJO 语法简洁，无隐式类型转换，没有 JavaScript 的种种历史包袱和迷惑行为。

(const|var) identifier[: type] = value

```jojo
var a: number = 1
const b = 'hello'

a.toString() // "1"

// 可直接在常量上调用方法
true.toString() // "true"

// 使用fn作为函数关键字
fn main() {
  print('hello jojo') // 使用 print 输出控制台
}

// 指定长度数组，数组只支持一种数据类型
var arr: []number = [1,2,3]

arr.push(4) // [1, 2, 3, 4]

// map 只支持特定类型的映射
const map: [string]number = {
  "a": 1
  "b": 2
}

map.keys()
map.values()

for (key in map.keys()) {
  print(key) // a, b
}

```


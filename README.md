# the JOJO programming language

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


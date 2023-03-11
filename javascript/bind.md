## JS Bind 函数

`bind` 是函数原型对象上的方法，用于绑定函数执行时的 `this`。并返回一个新的函数。

### 手写 bind 函数

```javascript
Function.myBind = function(thisArg, ...args) {
  // 保存原函数
  const fn = this;

  return function(...otherArgs) {
    return fn.apply(thisArg, [...args, ...otherArgs])
  }
}
```

## JS Apply 函数

`apply` 方法是函数原型对象上的方法，用于指定 `this` 的值，同时将参数以数组的形式传递给函数。


### 手写 Apply 函数
```javascript

Fcuntion.prototype.myApply = function(thisArg, args) {
  // 保存原函数
  const fn = this;

  // 如果没有指定 this
  if (!thisArg) {
    // 设置成全局对象，node 为 global，浏览器为 window
    thisArg = typeof window === 'undefined' ? global : window
  }

  if (!args || !Array.isArray(args)) {
    args = [];
  }

  thisArg.__fn = fn;
  const result = thisArg.__fn(...args);
  delete thisArg.__fn;
  return result;
}
```

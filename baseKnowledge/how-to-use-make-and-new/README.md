### 1、new:
  - 用于各种类型的内存分配。
  - 需要初始化后才能使用。
  - new返回指针：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。

### 2、make:
  - 用于内建类型（map、slice、channel）的内存分配。
    ```go
        slice1 := make([]int, 5, 10) // [0 0 0 0 0]
        
        map1 := make(map[string]int)
    ```
  - 可以直接使用。
  - 返回一个有初始值(非零)的T类型。

### 3、总结：


### 4、传值、传引用
> 观点：go函数参数一律传值，
  - 预声明类型如int，string、struct类型没什么好说的， 无论是传递该类型的值还是指针作为函数参数，本质上都是传值。
PS：函数内部的修改不会反馈到外部，除非是传递引用(&student1)。
  - slice，map和chan作为参数传递到函数中时是传的引用，其实这个说法不准确。 
    - PS：主要是因为函数内部的修改可以反馈到外部。 
    - slice、map、channel看上去像「传引用」只是因为 
      - slice其实是一个含有指针的结构体，注：由于slice会扩容，使用新的地址指针，要注意⚠️
      ```go
      type slice struct {
        array unsafe.Pointer // 连续分配的内存块，此指针指向内存的首地址，即该切片。
        len int // 切片长度。
        cap int // 切片容量。
      }
      ```
      - map和channel本身就是一个指针。
      ```go
        func makemap(t *maptype, hint int, h *hmap) *hmap
      ```

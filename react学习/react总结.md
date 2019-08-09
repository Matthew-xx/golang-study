# React总结

## react的特点

将js与html，css通过自有语法，融合到一起了，有利于封装组件。

## 总结

1. jsx对象，包含一组模板标签，类似于模板语法，使用（）包裹，
    
    - 内部只能有1个基标签
    - 内部的语句都用{}包裹，可包裹变量，表达式，被调用的函数，注释
    - 内部的注释只能使用{/* */}
    - 标签内的class属性需要写成className

2. 组件的创建可以使用两种方式

    - 方法创建
    - 类创建:
        
        定义类：一个类继承自Reac.Component
        类内：类内部定义render(),return一个jsx对象。
        调用：直接在渲染器中把类名当做标签用。

    一般使用类创建，因为可以绑定方法。

3. 组件可以嵌套

4. 组件绑定事件

    - onclick需要写成onClick
    - 组件的方法内部this容易与标签本身的this冲突，需要使用this.fnName.bind(this,prams)进行绑定。
    fnName为自定义的方法名称，prams为此方法的参数。

5. 组件状态属性

    - react中所有自定义属性，都保存着props对象中
    - 通过构造方法初始化自定义属性。需要传递props对象，并使用父类的构造函数初始化。
    - 通过state保存所有属性值得键值对{key:value}，但在构造方法中初始化之后，只能读，不能写
    - 通过setState()修改键值对。
    - 默认变量prevState指代当前的value

6. 数组渲染成列表

    - 数组对象调用map()进行遍历，跟range很像,map()接收1个方法，方法的参数为遍历的value和index,返回1个标签组。 
    - 将遍历的value和index分别写入<li>中，推荐自定义属性key存储index.<li>与</li>之间放value

7. onChange方法自带一个默认对象ev，传递当前事件对象(ev)的目标标签（target）的值（value）

8. 数组对象的splice(i,N,item)方法，从第i个元素开始剪切，减掉N个,在此位置追加元素item

9. 
# css

## 布局原理

根据设计图，先划分行，在把每行划分若干列

## 布局优化

### 语义化

能用语义化标签的就不用div或者span
如：
    <p>
    <a>

### 结构优化

如果一个容器内只有一个标签，那这个容器就不应该使用、

## css基本语法

    选择器{属性：值；属性：值；}

## css引入方式

### 内联式

使用标签的style属性。优先级最高，但乱

### 嵌入式

使用<style></style>专门开辟一块区域于<head></head>
需使用选择器

### 外链式

使用<link>引用专门的文件于<head></head>

## 选择器

### 标签选择器

按照某一类标签进行选择

特点：选择范围大，一般用来做一些通用设置。或用在层级选择器之中

    div{color:red}
    <div></div>

### 类选择器

通过给元素起class名称，归属到一类。从而选择一类。

特点：应用灵活，可复用，用的比较多

    .name{color:blue}
    <div class="name"></div>

### 层级选择器

在以上两种选择器中，增加了层级关系

    .con span{}

层级选择器可以越级选择

## css盒子模型

1. 宽与高 width：20px;height:300px；
2. 外边距 margin
3. 内边距 padding
4. 边框 border

注意点：

1. border、padding都会增加宽和高
2. 合并设置：
    padding:20px 80px 160px 40px；上右下左
    padding:20px 80px  40px；上左右下
    padding:20px 80px；上下左右
    padding:20px；所有
3. 很多标签有自己的固有属性，需要特别注意，如<h></h><body>都是有默认margin的

### 盒子真实尺寸计算

    盒子宽度=width+左右pading+左右border
    盒子高度=height+上下pading+上下border
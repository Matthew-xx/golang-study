# jquery

## 两种写法的区别

    //这个是原生写法，很慢
    window.onload() = function(){

    }
    //这个写法,运行更快
    $(document).ready(function(){

    })

## juery选择器

类似于css的选择器，很容易获取到对应标签

    $('#myid')选择id
    $('.myclass') 选择类
    $('li')选择所有li标签
    $('#uli li span')子孙选择器
    $('input[属性名=属性值]')选择属性名=属性值的input标签

## 选择器过滤

缩小筛选范围

    $('A').has('B'); 包含，选择有B元素的所有**A元素**
    $('A').not('B'); 不包含，选择不符合B的所有A元素
    $('A').eq(3);选择第3个某元素

## 选择集转移

    .pre() 向上转移1个兄弟元素
    .preALL 向上转移所有兄弟元素
    .next()
    .nextAll()
    .parent()
    .children()
    .siblings() 反选所有兄弟元素
    .find() 选择某元素内部包含的某一类元素，可使用子孙选择器代替

## 操作css

获取标签，与操作标签属性使用同一个函数
    $('').css('');
例如：
    //获取一个div对象
    var $div = $("#my_div");
    //读出这个div的字体大小
    var size = $div.css("font-size");
    //设置文字颜色（单属性操作）
    $div.css("color","red");
    //设置字体大小和粗细样式（多属性操作）
    $div.css({"font-size":"30","font-weignt":"bold"});

注意操作css时，如果目标选择器结果是一个集合，则只操作第一个元素。


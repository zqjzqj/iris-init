layui.use(['tree'], function(){
    var tree = layui.tree

    data = JSON.parse($("#rolePerms").attr("data-json"))

    //基本演示
    tree.render({
        elem: '#rolePerms'
        ,data: data
        ,showCheckbox: true  //是否显示复选框
        ,id: 'rolePerms'
        ,isJump: true //是否允许点击节点时弹出新窗口跳转
    });
})
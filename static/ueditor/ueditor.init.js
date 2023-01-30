$(function(){
    $("[data-ueditor]").each(function(k, v){
        var initialFrameHeight = $(v).attr("data-height")
        if (initialFrameHeight === "" || initialFrameHeight === 0 || initialFrameHeight === undefined) {
            initialFrameHeight = 400
        }
        UE.getEditor($(v).attr("id"), {
            initialFrameHeight:initialFrameHeight,
            initialFrameWidth:'100%',
        });
    })
})
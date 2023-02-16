var ReqStatusSuccess = 0
var ReqStatusErr = 1
layui.use(['form', 'laydate', 'upload'], function(){
    $("[data-action=upload]").change(function(){
        file = this.files[0]
        _this = $(this)
        var uploadView = _this.attr("data-view")
        if (uploadView) {
            window.GUtils.fileReadAsDataURL(file, function(r) {
                $(uploadView).attr('src', r)
            })
        }
        var uploadInput = _this.attr("data-input")
        if (uploadInput) {
            loadingShow()
            window.GUtils.cosPutObject({file: file}, function(url){
                var uploadInput = _this.attr("data-input")
                if (uploadInput) {
                    $(uploadInput).val(url)
                }
                loadingClose()
            })
        }
    })

    $("[data-action=tips]").hover(function(){
        layer.tips($(this).attr("data-content"), $(this), {
            tips: [1, '#3595CC'],
            time: 0,
        });
    }, function(){
        layer.tips($(this).attr("data-content"), $(this), {
            tips: [1, '#3595CC'],
            time: 1,
        });
    })

    //ifrIndex ifr层index
    $("[data-action=form]").submit(function() {
        var funcCallback = $(this).attr("data-callback");
        var callbackUrl = $(this).attr("data-callback-url")
        var ifrIndex = $(this).attr("data-ifr-index")
        if (callbackUrl) {
            funcSuccess = 'successCallbackUrl(callbackUrl,response.Msg, response.Data, ifrIndex)'
        } else {
            funcSuccess = $(this).attr("data-success") || 'successCallback(response.Msg, response.Data, ifrIndex)';
        }
        funcError = $(this).attr("data-error") || 'errorCallback(response.Msg, response.Data, ifrIndex)';
        var beforeFun = $(this).attr("data-before") || "";
        if ( beforeFun ) {
            var beforeFunRet = eval(beforeFun)();//console.log(beforeFunRet)
            if ( typeof  beforeFunRet === "object") {
                if ( !beforeFunRet.status ) {
                    loadingClose();
                    alertError(beforeFunRet.message);
                    return false;
                }
            } else {
                if ( !beforeFunRet ) {
                    loadingClose();
                    return false;
                }
            }
        }
        $("[data-mark]").text("")
        loadingShow();
        $(this).ajaxSubmit({
            dataType: 'json',
            success: function (response) {
                loadingClose();
                if ( funcCallback ) {
                    eval(funcCallback)
                }
                if ( response.Code == ReqStatusSuccess ) {
                    eval(funcSuccess);
                } else {
                    eval(funcError);
                }
            },
            error: function(response){
                loadingClose();
                alertError("系统出错：" + response.statusText)
            }
        });
        return false;
    });


    $("[data-action=prompt]").on("click", function(){
        var url = $(this).attr("href");
        var title = $(this).attr("data-title") || "默认对话框";
        var msg = $(this).attr("data-msg") || "";
        var id = $(this).attr("data-id") || 0;
        var ifrIndex = $(this).attr("data-ifr-index")
        var funcCallback = $(this).attr("data-callback") || "";
        var funcSuccess = $(this).attr("data-success") || 'successCallback(response.Msg, response.Data, ifrIndex)';
        var funcError = $(this).attr("data-error") || 'errorCallback(response.Msg, response.Data, ifrIndex)';
        var requestData = $(this).attr('data-params') || {};
        var beforeFun = $(this).attr("data-before") || "";

        if ( beforeFun ) {
            eval(beforeFun);
        }

        alertWarning(title, msg, function(){
            loadingShow();
            if ( typeof requestData !== "object" ) {
                /*console.log(requestData)
                console.log(typeof requestData)*/
                requestData = JSON.parse(requestData)
            }

            if ( id ) {
                requestData['ID'] = id;
            }
            $.ajax({
                type:"POST",
                url: url,
                data: requestData,
                dataType:'json',
                success: function (response) {
                    loadingClose();
                    if ( funcCallback ) {
                        eval(funcCallback)
                    }
                    if ( response.Code == ReqStatusSuccess ) {
                        eval(funcSuccess)
                    } else {
                        eval(funcError)
                    }
                }
            })
        });

        return false;
    });

    $("[data-action=del]").on("click", function(){
        var url = $(this).attr("href");
        var title = $(this).attr("data-title") || "你确定要删除吗?";
        var msg = $(this).attr("data-msg") || "";
        var method = $(this).attr('data-method') || "POST";
        var id = $(this).attr("data-id") || 0;
        var ifrIndex = $(this).attr("data-ifr-index")
        var funcSuccess = $(this).attr("data-success") || 'successCallback(response.Msg, response.Data, ifrIndex)';
        var funcError = $(this).attr("data-error") || 'errorCallback(response.Msg, response.Data, ifrIndex)';
        var requestData = $(this).attr('data-params') || {};
        if ( typeof requestData !== "object" ) {
            requestData = JSON.parse(requestData)
        }

        if ( id ) {
            requestData['ID'] = id;
        }
        alertWarning(title, msg, function(){
            loadingShow();

            $.ajax({
                type:method,
                url: url,
                data:requestData,
                dataType:'json',
                success: function (response) {
                    loadingClose();
                    if ( response.Code == ReqStatusSuccess ) {
                        eval(funcSuccess);
                    } else {
                        eval(funcError);
                    }
                }
            })
        });

        return false;
    });

    $("[data-action=open]").on("click", function () {
        var width = $(this).attr('data-width') || "85%"
        var height = $(this).attr('data-height') || "95%"
        var title = $(this).attr('data-title');
        var type = $(this).attr('data-type') || 2;
        var before = $(this).attr("data-before") || null;
        var callbackEnd = $(this).attr('data-callback-end') || "";
        var callbackSuccess = $(this).attr('data-callback-success') || "";

        if ( before ) {
            var beforeFunRet = eval(before)($(this));//console.log(beforeFunRet)
            if ( typeof  beforeFunRet === "object") {
                if ( !beforeFunRet.status ) {
                    loadingClose();
                    alertError(beforeFunRet.message);
                    return false;
                }
            } else {
                if ( !beforeFunRet ) {
                    loadingClose();
                    return false;
                }
            }
        }
        close=1;
        if(title=="false") {
            title=false;
            close=0;
        }
        var init = {
            type: parseInt(type),
            title: title,
            shadeClose: false,
            shade: 0.3,
            maxmin: true, //开启最大化最小化按钮
            area: [width, height],
            closeBtn:close,
            content: $(this).attr("href"),
            success : function (layero, index) {
            }
        };

        if ( $(this).attr("data-offset") ) {
            init['offset'] = $(this).attr("data-offset");
        }

        if ( $(this).attr('data-iframe-id') ) {
            init['id'] = $(this).attr('data-iframe-id');
        }

        if ( callbackEnd ) {
            init["end"] = function () {
                eval(callbackEnd + "()");
            };
        }

        if ( callbackSuccess ) {
            init["success"] = function (index, layero) {
                eval(callbackSuccess + "(index, layero)");
            };
        }

        top.layer.open(init);

        return false;
    });


    $("[data-action=input]").on("click", function(){
        var title = $(this).attr("data-title");
        var formType = $(this).attr("data-type") || 3;
        var method = $(this).attr("data-method") || "post";
        var url = $(this).attr("href");
        var name = $(this).attr("data-name");
        var params = $(this).attr("data-params");
        var value = $(this).attr("data-val")
        var ifrIndex = $(this).attr("data-ifr-index")
        if ( params ) {
            data = JSON.parse(params);
        } else {
            var data = {};
        }
        data[name] = "";
        layer.prompt({title: title, formType: formType, value:value}, function(pass, index){
            layer.close(index);
            data[name] = pass;
            loadingShow();
            $.ajax({
                type : method,
                url : url,
                data : data,
                dataType:'json',
                success: function (response) {
                    loadingClose();
                    if ( response.Code == ReqStatusSuccess ) {
                        successCallback(response.Msg, response.Data, ifrIndex)
                    } else {
                        errorCallback(response.Msg, response.Data, ifrIndex)
                    }
                }
            })
        });
        return false;
    });
    var dataUrlEl = $("[data-action=url]")
    dataUrlEl.on("click", function(){
        var url = $(this).data("data-url");
        if ( !url ) {
            url = $(this).attr("href")
        }
        var requestData = $(this).attr('data-params') || {};
        var method = $(this).attr("data-method") || "GET";
        if( typeof(requestData) == 'string' ){
            requestData = JSON.parse(requestData);
        }
        var ifrIndex = $(this).attr("data-ifr-index")
        loadingShow();
        $.ajax({
            type:method,
            url: url,
            data:requestData,
            dataType:'json',
            success: function (response) {
                loadingClose();
                if ( response.Code == ReqStatusSuccess ) {
                    successCallback(response.Msg, response.Data, ifrIndex)
                } else {
                    errorCallback(response.Msg, response.Data, ifrIndex)
                }
            }
        });

        return false;
    });
    dataUrlEl.on("change", function(){
        var url = $(this).data("url");
        var requestData = $(this).attr('data-params') || {};
        var method = $(this).attr("data-method") || "GET";
        if(typeof(requestData)=='string'){
            requestData = JSON.parse(requestData);
        }
        var ifrIndex = $(this).attr("data-ifr-index")
        loadingShow();
        $.ajax({
            type:method,
            url: url,
            data:requestData,
            dataType:'json',
            success: function (response) {
                loadingClose();
                if ( response.Code == ReqStatusSuccess ) {
                    successCallback(response.Msg, response.Data, ifrIndex)
                } else {
                    errorCallback(response.Msg, response.Data, ifrIndex)
                }
            }
        });

        return false;
    });

    $("[data-action=checkboxAll]").click(function () {
        var mark = $(this).attr("data-main-mark");
        var items = $("[data-mark=" + mark + "]");

        if ( $("[data-mark=" + mark + "]:checked").length === items.length ) {
            items.removeAttr("checked");
            $(this).removeAttr("checked");;
        } else {
            $(this).prop("checked");
            items.prop("checked", "checked")
        }
    });

    var dataCheckbox = $("[data-action=checkbox]")
    dataCheckbox.click(function () {
        var mark = $(this).attr("data-mark");
        var main = $("[data-main-mark=" + mark + "]");
        if ( $("[data-mark=" + mark + "]:checked").length === $("[data-mark=" + mark + "]").length ) {
            main.prop("checked", "checked")
        } else {
            main.removeAttr("checked");
        }
    });
    if ( dataCheckbox.length > 0 ) {
        $.each($("[data-action=checkbox]"), function(){
            var mark = $(this).attr("data-mark");
            if ( mark ) {
                var main = $("[data-main-mark=" + mark + "]");
                if ( $("[data-mark=" + mark + "]:checked").length === $("[data-mark=" + mark + "]").length ) {
                    main.prop("checked", "checked")
                } else {
                    main.removeAttr("checked");
                }
            }
        })
    }

    $(".checked-with-tr").click(function () {
        $(this).find(".checked-td").click();
    });

    var laydate = layui.laydate;
    $('#date-range').find('input').each(function(){
        var type = $(this).attr("data-type") || "datetime";
        if ( $(this).attr("data-format") ) {
            var format = $(this).attr("data-format");
        } else {
            var format =  type == "datetime" ? "yyyy-MM-dd HH:mm:ss" : "yyyy-MM-dd";
        }
        if($(this).attr('name').indexOf('_ymd')!=-1) {
            type='date';
            format = 'yyyy-MM-dd';
        }
        laydate.render({
            elem: this,
            type: type,
            format:  format
        });
    });
    $(".dateInputs").each(function(){
        var type = $(this).attr("data-type") || "datetime";
        if ( $(this).attr("data-format") ) {
            var format = $(this).attr("data-format");
        } else {
            var format =  type == "datetime" ? "yyyy-MM-dd HH:mm:ss" : "yyyy-MM-dd";
        }
        if($(this).attr('name').indexOf('_ymd')!=-1) {
            type='date';
            format = 'yyyy-MM-dd';
        }
        laydate.render({
            elem: this,
            type: type,
            format:  format
        });
    })

    $("[data-action=checkboxChecked]").each(function(k, v){
        var dataVal = $(v).attr("data-value")
        _dataVal = dataVal.split(",")
        $(v).find("input:checkbox").each(function(k, _v){
            var _val = $(_v).val()
            if (_val !== "" && $.inArray(_val, _dataVal) !== -1) {
                $(_v).removeAttr("disabled")
                $(_v).next().removeClass("layui-checkbox-disbaled")
                $(_v).next().trigger("click")
                //这里在手动添加上选择 使得重置可以按钮生效
                $(_v).attr("checked", "checked")
            }
        })
    })

    setTimeout(function(){
        $("[data-action=upload-multiple-tbody]").find("._filename").each(function(k, v){
            var _this = $(this)
            _this.parent().parent().find('.upload-delete').on('click', function(){
                $(this).parent().parent().remove();
                return false
            })
            var _url = $(v).find("img")[0].src
            fetch(_url).then(function(res){
                return res.blob()
            }).then(function(data){
                if (data.type.indexOf("image") === -1) {
                    var lastItem = decodeURI(_url.substring(_url.lastIndexOf('/') + 1))
                    _this.find("img").remove()
                    _this.html("<span style='color: #00a0e9'>"+lastItem+" 【点击下载附件】</span>")
                    _this.attr("href", _url)
                    _this.attr("target", "_blank")
                    _this.unbind("click")
                }
                _this.parent('td').next('.file-size').text((data.size/1014).toFixed(1) + "kb")
            })
        })
    }, 100)
    $("[data-action=upload-multiple]").each(function(k ,v){
        var id = $(v).attr("id")
        var inputName = $(v).attr("data-input-name")
        var accept = $(v).attr("data-accept") || "file"
        var fileName = $(v).attr("data-file-name") || "File"
        var ListView = $('#list-view-'+id)
        var dataMore = $(v).attr("data-more")
        var more = false
        if (dataMore === "true"  || dataMore === "True"  || dataMore === true) {
            more = true
        }
        _uploadUrl = $(v).attr("upload-url")
        if (_uploadUrl === "" || _uploadUrl === undefined) {
            _uploadUrl = "/upload"
        }
        uploadListIns = layui.upload.render({
            elem: '#'+id
            ,url: _uploadUrl
            ,accept: accept
            ,multiple: more
            ,auto: true
            ,field: fileName
            ,bindAction: '#upFile-'+id
            ,choose: function(obj){
                var files = this.files = obj.pushFile(); //将每次选择的文件追加到文件队列
                //读取本地文件
                obj.preview(function(index, file, result){
                    var _td0 = ""
                    if (file.type.indexOf("image") !== -1) {
                        _td0 =  '<a href="javascript:void(0);" id="upload-view-file-'+index+'">' +
                                '<img width="80" height="80" src=""> '+ file.name +'' +
                                '</a>'
                    } else {
                        _td0 = '<a href="javascript:void(0);" id="upload-view-file-'+index+'">'+file.name+'</a>'
                    }
                    var tr = $(['<tr id="upload-'+ index +'">'
                        ,'<td>'+_td0+'</td>'
                        ,'<td>'+ (file.size/1014).toFixed(1) +'kb</td>'
                        ,'<td>等待上传</td>'
                        ,'<td>'
                        ,'<button class="layui-btn layui-btn-mini layui-btn-danger upload-delete">删除</button>'
                        ,'</td>'
                        ,'</tr>'].join(''));

                    //删除
                    tr.find('.upload-delete').on('click', function(){
                        delete files[index]; //删除对应的文件
                        tr.remove();
                        uploadListIns.config.elem.next()[0].value = ''; //清空 input file 值，以免删除后出现同名文件不可选
                    });
                    if (!more) {
                        ListView.empty()
                    }
                    ListView.append(tr);
                    if (file.type.indexOf("image") !== -1) {
                        var reader = new FileReader();
                        reader.readAsDataURL(file)
                        reader.onload = function(e){
                            var urlData = this.result;
                            var _upAImg = $('#upload-view-file-'+index)
                            _upAImg.find("img").attr("src", urlData)
                            _upAImg.attr("href", "<img src='"+urlData+"'>")
                            _upAImg.on('click', function(){
                                layer.open({
                                    type: 1,
                                    shade: 0.3,
                                    title: "查看图片",
                                    shadeClose: false,
                                    maxmin: true, //开启最大化最小化按钮
                                    area: ['auto', '95%'],
                                    closeBtn:1,
                                    content: "<img src='"+urlData+"'>",
                                })
                                return false;
                            })
                        }
                    }/*
                    if (!more) {
                       // obj.upload(index, file);
                    }*/
                });
            }
            ,before: function(){
                var i = 0
                for (k in this.files) {
                    i++
                }
                if (i === 0) {
                    return
                }
                loadingShow()
            }
            ,done: function(res, index, upload){
                var i = 0
                for (k in this.files) {
                    i++
                }
                if (i === 1) {
                    loadingClose()
                }
                if(res.Code === 0){ //上传成功
                    var tr = ListView.find('tr#upload-'+ index)
                        ,tds = tr.children();
                    tds.eq(2).html('<span style="color: #5FB878;">上传成功</span>');
                    $('#upload-view-file-'+index).append("<input type='hidden' name='"+inputName+"' value='"+res.Data.Url+"'>")
                    return delete this.files[index]; //删除文件队列已经上传成功的文件
                }
                this.error(index, upload);
            }
            ,error: function(index, upload){
                var i = 0
                for (k in this.files) {
                    i++
                }
                if (i === 1) {
                    loadingClose()
                }
                var tr = ListView.find('tr#upload-'+ index)
                    ,tds = tr.children();
                tds.eq(2).html('<span style="color: #FF5722;">上传失败</span>');
                delete this.files[index]; //删除文件队列已经上传成功的文件
            }
        });
    })

})
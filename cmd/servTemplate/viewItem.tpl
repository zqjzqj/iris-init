{% extends "layouts/main.html" %}
{% block content %}
<div class="layui-fluid">
    <div class="layui-card">
        <div class="layui-card-body" style="padding: 15px;">
            <form class="layui-form" action="/{{.Alias}}/edit" method="post" data-action="form" data-ifr-index="0" lay-filter="component-form-group">
                {{- range .ModelField}}
                <div class="layui-form-item">
                    <label class="layui-form-label">{{.Label}}</label>
                    <div class="layui-input-block">
                        <input type="text" name="{{.Name}}" value="{{printf "{{Data.Item"}}.{{.Name}}{{print "}}"}}" autocomplete="off" placeholder="ID" class="layui-input">
                    </div>
                </div>
                {{- end}}
                <div class="layui-form-item layui-layout-admin">
                    <div class="layui-input-block">
                        <div class="layui-footer" style="left: 0;">
                            <input type="hidden" name="ID" value="{{printf "{{Data.Item.ID}}"}}">
                            <button class="layui-btn">立即提交</button>
                            <button type="reset" class="layui-btn layui-btn-primary">重置</button>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>
{% endblock %}
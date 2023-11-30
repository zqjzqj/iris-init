{% extends "layouts/main.html" %}
{% block content %}
<div class="layui-fluid">
    <div class="layui-row layui-col-space15">
        <div class="layui-col-md12">
            <div class="layui-card">
                <div class="layui-card-header">列表</div>
                <div class="layui-card-body">
                    <div class="layui-form layui-card-header layuiadmin-card-header-auto" style="margin-bottom: 10px;">
                        <form action="" class="list-search">
                            搜索：
                            <div class="layui-inline">
                                <input class="layui-input" name="ID" placeholder="ID">
                            </div>
                         {{- range .ModelField}}
                            {{- if .Search}}
                            <div class="layui-inline">
                                {{- if eq .Type "string"}}
                                <input class="layui-input" name="{{.Name}}Like" placeholder="{{.Label}}">
                                {{- else}}
                                <input class="layui-input" name="{{.Name}}" placeholder="{{.Label}}">
                                {{- end}}
                            </div>
                            {{- end}}
                         {{- end}}
                            <button class="layui-btn" type="submit">搜索</button>
                            <button class="layui-btn layui-btn-primary" type="reset">重置</button>
                        </form>
                    </div>

                    <div style="padding-bottom: 10px;">
                        <button data-perm data-perm-val="POST@/{{.Alias}}/edit"
                         data-action="open" href="/{{.Alias}}/item" data-title="添加" class="layui-btn layuiadmin-btn-list" >
                            添加
                        </button>
                    </div>
                    <table class="layui-table">
                        <thead>
                        <tr>
                            {{- range .ModelField}}
                            <th>{{.Label}}</th>
                            {{- end}}
                            <th>操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {% for item in Data.List %}
                        <tr>
                            {{- range .ModelField}}
                            <td>{{print "{{item"}}.{{.Name}}{{print "}}"}}</td>
                            {{- end}}
                            <td>
                                <button data-perm data-perm-val="GET@/{{.Alias}}/item"
                                        data-action="open" href="/{{.Alias}}/item?ID={{print "{{item.ID}}"}}&form-disabled=1"
                                        data-title="详情"
                                        class="layui-btn layui-btn-primary layui-btn-sm">
                                    详情
                                </button>
                                <button data-perm data-perm-val="POST@/{{.Alias}}/edit"
                                        data-action="open" href="/{{.Alias}}/item?ID={{print "{{item.ID}}"}}"
                                        data-title="编辑"
                                        class="layui-btn layui-btn-primary layui-btn-sm">
                                    编辑
                                </button>
                                <button
                                        data-perm data-perm-val="POST@/{{.Alias}}/delete"
                                        data-action="del" href="/{{.Alias}}/delete"
                                        data-id="{{print "{{item.ID}}"}}" class="layui-btn layui-btn-primary layui-btn-sm">
                                    删除
                                </button>
                            </td>
                        </tr>
                        {% endfor %}
                        </tbody>
                    </table>
                    {% include "../partials/pager.html" %}
                </div>
            </div>
        </div>
    </div>
</div>
{% endblock %}
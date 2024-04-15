{{ define "blogs/index.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <div>
            <h1>Blogs</h1>

            <ul>
                {{ template "blogs/blog_list.tpl" . }}
            </ul>
        </div>
    {{ template "layouts/footer.tpl" .}}
{{end}}
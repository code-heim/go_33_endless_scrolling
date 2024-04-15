{{ define "blogs/blog_list.tpl"}}
    {{ with .blogs }}
        {{ range . }}
            <li>
                <div>
                    <a href="/blogs/{{.ID}}">
                        <h5>{{ .Title }} </h5>
                    </a>
                    <p>{{ .Content }}</p>
                </div>
                <br/>
            </li>
        {{ end }}
    {{ end }}

    <div hx-trigger="revealed" hx-target="this"
        hx-get="/blogs?page={{ .nextPage }}"
        hx-swap="afterend">
    </div>
{{ end }}
package global

{{- if .HasGlobal }}

import "github.com/dcncy/gin-vue-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}
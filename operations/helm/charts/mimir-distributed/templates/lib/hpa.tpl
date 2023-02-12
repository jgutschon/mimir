{{/*
Mimir common HorizontalPodAutoscaler definition
Params:
  ctx = . context
  component = name of the component
  kind = Kind of object to be scaled
  zoneName = name of the zone, e.g. zone-a (optional, StatefulSets only)
  rolloutZone = used to get replicas per zone when using zone-aware replication (optional, StatefulSets only)
  zoneAware = if the sts uses zone-aware replication (optional, StatefulSets only)
*/}}
{{- define "mimir.lib.horizontalPodAutoscaler" -}}
{{- $componentSection := include "mimir.componentSectionFromName" . | fromYaml -}}
{{- $autoscaling := $componentSection.autoscaling -}}
{{- $args := dict "ctx" $.ctx "component" $.component -}}
{{- $name := (and (eq ($.zoneAware | toString) "true") (ne ($.zoneName | toString) "")) | ternary
  (printf "%s-%s" (include "mimir.resourceName" $args) $.zoneName)
  (include "mimir.resourceName" $args)
-}}

{{- if $autoscaling.enabled -}}
apiVersion: {{ include "mimir.hpa.version" . }}
kind: HorizontalPodAutoscaler
metadata:
  name: {{ $name }}
  labels:
    {{- include "mimir.labels" $args | nindent 4 -}}
  namespace: {{ .Release.Namespace | quote }}
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: {{ $.kind }}
    name: {{ $name }}
  {{- if and $.zoneAware (ne ($.zoneName | toString) "") }}
  minReplicas: {{ $.rolloutZone.replicas }}
  {{- else }}
  minReplicas: {{ $autoscaling.minReplicas }}
  {{- end }}
  maxReplicas: {{ $autoscaling.maxReplicas }}
  {{- if eq (include "mimir.hpa.version" .) "autoscaling/v2" -}}
  metrics:
    {{- if $autoscaling.averageMemoryUtilization }}
    - type: Resource
      resource:
        name: memory
        target:
          type: Utilization
          averageUtilization: {{ $autoscaling.averageMemoryUtilization }}
    {{- end }}
    {{- if $autoscaling.averageCpuUtilization }}
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: {{ $autoscaling.averageCpuUtilization }}
    {{- end }}
  behavior:
    {{- toYaml $autoscaling.behavior | nindent 4 -}}
  {{- else -}}
  metrics:
    {{- if $autoscaling.averageMemoryUtilization }}
    - type: Resource
      resource:
        name: memory
        targetAverageUtilization: {{ $autoscaling.averageMemoryUtilization }}
    {{- end }}
    {{- if $autoscaling.averageCpuUtilization }}
    - type: Resource
      resource:
        name: cpu
        targetAverageUtilization: {{ $autoscaling.averageCpuUtilization }}
    {{- end }}
  {{- end -}}
{{- end -}}
{{- end -}}

{{/*
Returns the HorizontalPodAutoscaler API version for this verison of kubernetes.
*/}}
{{- define "mimir.hpa.version" -}}
{{- if semverCompare ">= 1.23-0" .Capabilities.KubeVersion.Version -}}
autoscaling/v2
{{- else -}}
autoscaling/v2beta1
{{- end -}}
{{- end -}}

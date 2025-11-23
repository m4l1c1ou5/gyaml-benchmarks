package benchmarks

import (
	"testing"

	"github.com/m4l1c1ou5/gyaml"
	"gopkg.in/yaml.v3"
)

const yamlDocument = `
widget:
  debug: "on"
  window:
    title: "Sample Konfabulator Widget"
    name: "main_window"
    width: 500
    height: 500
  image:
    src: "Images/Sun.png"
    hOffset: 250
    vOffset: 250
    alignment: "center"
  text:
    data: "Click Here"
    size: 36
    style: "bold"
    vOffset: 100
    alignment: "center"
    onMouseUp: "sun1.opacity = (sun1.opacity / 100) * 90;"
`

const complexYAML = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: production
  labels:
    app: nginx
    version: "1.14.2"
  annotations:
    description: "This is a sample nginx deployment"
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
        env:
        - name: ENVIRONMENT
          value: "production"
        - name: LOG_LEVEL
          value: "info"
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
      - name: sidecar
        image: busybox:latest
        command: ["sh", "-c", "while true; do echo hello; sleep 10; done"]
`

const largeArrayYAML = `
users:
  - id: 1
    name: "Alice Johnson"
    email: "alice@example.com"
    age: 28
    active: true
    roles: ["admin", "developer"]
  - id: 2
    name: "Bob Smith"
    email: "bob@example.com"
    age: 34
    active: true
    roles: ["developer"]
  - id: 3
    name: "Charlie Davis"
    email: "charlie@example.com"
    age: 42
    active: false
    roles: ["viewer"]
  - id: 4
    name: "Diana Prince"
    email: "diana@example.com"
    age: 31
    active: true
    roles: ["admin", "manager"]
  - id: 5
    name: "Eve Wilson"
    email: "eve@example.com"
    age: 25
    active: true
    roles: ["developer", "tester"]
`

var paths = []string{
	"widget.window.name",
	"widget.image.hOffset",
	"widget.text.onMouseUp",
}

var complexPaths = []string{
	"spec.template.spec.containers.0.image",
	"metadata.labels.version",
	"spec.replicas",
}

var arrayPaths = []string{
	"users.0.name",
	"users.2.email",
	"users.#.name",
}

// GYAML Benchmarks
func BenchmarkGYAMLGet(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(yamlDocument, paths[i%len(paths)])
	}
}

func BenchmarkGYAMLGetBytes(b *testing.B) {
	data := []byte(yamlDocument)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.GetBytes(data, paths[i%len(paths)])
	}
}

func BenchmarkGYAMLGetComplex(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(complexYAML, complexPaths[i%len(complexPaths)])
	}
}

func BenchmarkGYAMLGetArray(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(largeArrayYAML, arrayPaths[i%len(arrayPaths)])
	}
}

func BenchmarkGYAMLParse(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Parse(yamlDocument)
	}
}

func BenchmarkGYAMLValid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Valid(yamlDocument)
	}
}

// YAML v3 Benchmarks (for comparison)
func BenchmarkYAMLv3UnmarshalMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		yaml.Unmarshal([]byte(yamlDocument), &data)
		_ = data["widget"].(map[string]interface{})["window"].(map[string]interface{})["name"]
	}
}

func BenchmarkYAMLv3UnmarshalStruct(b *testing.B) {
	type Window struct {
		Title  string `yaml:"title"`
		Name   string `yaml:"name"`
		Width  int    `yaml:"width"`
		Height int    `yaml:"height"`
	}
	type Image struct {
		Src       string `yaml:"src"`
		HOffset   int    `yaml:"hOffset"`
		VOffset   int    `yaml:"vOffset"`
		Alignment string `yaml:"alignment"`
	}
	type Text struct {
		Data      string `yaml:"data"`
		Size      int    `yaml:"size"`
		Style     string `yaml:"style"`
		VOffset   int    `yaml:"vOffset"`
		Alignment string `yaml:"alignment"`
		OnMouseUp string `yaml:"onMouseUp"`
	}
	type Widget struct {
		Debug  string `yaml:"debug"`
		Window Window `yaml:"window"`
		Image  Image  `yaml:"image"`
		Text   Text   `yaml:"text"`
	}
	type Root struct {
		Widget Widget `yaml:"widget"`
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data Root
		yaml.Unmarshal([]byte(yamlDocument), &data)
		_ = data.Widget.Window.Name
	}
}

// Memory efficiency benchmarks
func BenchmarkGYAMLMultipleGets(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r1 := gyaml.Get(yamlDocument, "widget.window.name")
		r2 := gyaml.Get(yamlDocument, "widget.image.hOffset")
		r3 := gyaml.Get(yamlDocument, "widget.text.onMouseUp")
		_, _, _ = r1, r2, r3
	}
}

func BenchmarkYAMLv3MultipleGets(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		yaml.Unmarshal([]byte(yamlDocument), &data)
		widget := data["widget"].(map[string]interface{})
		r1 := widget["window"].(map[string]interface{})["name"]
		r2 := widget["image"].(map[string]interface{})["hOffset"]
		r3 := widget["text"].(map[string]interface{})["onMouseUp"]
		_, _, _ = r1, r2, r3
	}
}

// Deep path benchmarks
func BenchmarkGYAMLDeepPath(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(complexYAML, "spec.template.spec.containers.0.resources.limits.memory")
	}
}

func BenchmarkYAMLv3DeepPath(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		yaml.Unmarshal([]byte(complexYAML), &data)
		spec := data["spec"].(map[string]interface{})
		template := spec["template"].(map[string]interface{})
		tSpec := template["spec"].(map[string]interface{})
		containers := tSpec["containers"].([]interface{})
		container := containers[0].(map[string]interface{})
		resources := container["resources"].(map[string]interface{})
		limits := resources["limits"].(map[string]interface{})
		_ = limits["memory"]
	}
}

// Query benchmarks (GYAML-specific features)
func BenchmarkGYAMLQueryAll(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(largeArrayYAML, "users.#.name")
	}
}

func BenchmarkGYAMLQueryConditional(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(largeArrayYAML, "users.#(active==true)#.name")
	}
}

// Iteration benchmarks
func BenchmarkGYAMLForEach(b *testing.B) {
	result := gyaml.Parse(yamlDocument)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.ForEach(func(key, value gyaml.Result) bool {
			return true
		})
	}
}

func BenchmarkGYAMLArray(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := gyaml.Get(largeArrayYAML, "users")
		_ = result.Array()
	}
}

func BenchmarkGYAMLMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := gyaml.Get(yamlDocument, "widget.window")
		_ = result.Map()
	}
}

// Type conversion benchmarks
func BenchmarkGYAMLGetString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := gyaml.Get(yamlDocument, "widget.window.name")
		_ = result.String()
	}
}

func BenchmarkGYAMLGetInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := gyaml.Get(yamlDocument, "widget.window.width")
		_ = result.Int()
	}
}

func BenchmarkGYAMLGetBool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := gyaml.Get(largeArrayYAML, "users.0.active")
		_ = result.Bool()
	}
}

// Large document benchmarks
const largeDocument = `
database:
  connections:
  - host: "db1.example.com"
    port: 5432
    username: "admin"
    password: "secret123"
    database: "production"
    ssl: true
    maxConnections: 100
  - host: "db2.example.com"
    port: 5432
    username: "admin"
    password: "secret456"
    database: "production"
    ssl: true
    maxConnections: 100
  - host: "db3.example.com"
    port: 5432
    username: "admin"
    password: "secret789"
    database: "production"
    ssl: true
    maxConnections: 100
  settings:
    pool:
      minSize: 10
      maxSize: 50
      timeout: 30
    retry:
      maxAttempts: 3
      backoff: 1000
    monitoring:
      enabled: true
      interval: 60
      alerts:
      - type: "connection_error"
        threshold: 5
        action: "email"
      - type: "slow_query"
        threshold: 1000
        action: "log"
`

func BenchmarkGYAMLLargeDocument(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gyaml.Get(largeDocument, "database.settings.monitoring.alerts.1.threshold")
	}
}

func BenchmarkYAMLv3LargeDocument(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		yaml.Unmarshal([]byte(largeDocument), &data)
		db := data["database"].(map[string]interface{})
		settings := db["settings"].(map[string]interface{})
		monitoring := settings["monitoring"].(map[string]interface{})
		alerts := monitoring["alerts"].([]interface{})
		alert := alerts[1].(map[string]interface{})
		_ = alert["threshold"]
	}
}

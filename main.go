package main

import (
    "embed"
    "flag"
    "fmt"
    "log"
    "os"
    "text/template"
    "gopkg.in/yaml.v2"
)

//go:embed templates/swagger_ui.html
var swaggerTemplateFS embed.FS

func main() {
    // Параметры командной строки
    openAPIFilePath := flag.String("yaml", "api/openapi.yaml", "Путь к YAML файлу OpenAPI")
    outputFilePath := flag.String("output", "api/swagger_ui.html", "Путь к выходному HTML файлу")
    flag.Parse()

    // Чтение и обработка YAML файла
    openAPIYamlFile, err := os.ReadFile(*openAPIFilePath)
    if err != nil {
        log.Fatalf("Ошибка при чтении YAML-файла: %v", err)
    }

    var data map[string]interface{}
    err = yaml.Unmarshal(openAPIYamlFile, &data)
    if err != nil {
        log.Fatalf("Ошибка при разборе YAML-данных: %v", err)
    }

    // Чтение HTML шаблона из встроенной файловой системы
    swaggerUIHTMLTemplateFile, err := swaggerTemplateFS.ReadFile("templates/swagger_ui.html")
    if err != nil {
        log.Fatalf("Ошибка при чтении HTML-шаблона: %v", err)
    }

    tmpl, err := template.New("swagger-ui").Parse(string(swaggerUIHTMLTemplateFile))
    if err != nil {
        log.Fatalf("Ошибка при парсинге шаблона: %v", err)
    }

    // Создание выходного файла
    swaggerUIHTMLFile, err := os.Create(*outputFilePath)
    if err != nil {
        log.Fatalf("Ошибка при создании файла: %v", err)
    }
    defer swaggerUIHTMLFile.Close()

    // Заполнение шаблона
    err = tmpl.Execute(swaggerUIHTMLFile, string(openAPIYamlFile))
    if err != nil {
        log.Fatalf("Ошибка при выполнении шаблона: %v", err)
    }

    fmt.Println("Файл успешно сгенерирован:", *outputFilePath)
}

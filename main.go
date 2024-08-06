package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "text/template"

    "gopkg.in/yaml.v2"
)

func main() {
    // Параметры командной строки
    openAPIFilePath := flag.String("yaml", "api/openapi.yaml", "Путь к YAML файлу OpenAPI")
    swaggerTemplateFilePath := flag.String("template", "cmd/swagger/templates/swagger_ui.html", "Путь к HTML шаблону Swagger UI")
    outputFilePath := flag.String("output", "api/swagger_ui.html", "Путь к выходному HTML файлу")
    flag.Parse()

    // Чтение и обработка YAML файла
    openAPIYamlFile, err := os.ReadFile(*openAPIFilePath)
    if err != nil {
        log.Fatalf("Ошибка при чтении YAML-файла: %v", err)
    }

    var data map[string]any
    err = yaml.Unmarshal(openAPIYamlFile, &data)
    if err != nil {
        log.Fatalf("Ошибка при разборе YAML-данных: %v", err)
    }

    // Чтение и обработка HTML шаблона
    swaggerUIHTMLTemplateFile, err := os.ReadFile(*swaggerTemplateFilePath)
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
